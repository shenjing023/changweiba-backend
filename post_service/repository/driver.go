package repository

import (
	"cw_post_service/repository/ent"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/migrate"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/reply"
	"fmt"
	"os"
	"strconv"
	"time"

	"cw_post_service/conf"

	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/shenjing023/llog"
	"golang.org/x/sync/singleflight"
)

var (
	redisClient     *redis.Client
	entClient       *ent.Client
	postsCountCache singleflight.Group
)

const (
	// POSTSCOUNTKEY redis 保存当前帖子总数
	POSTSCOUNTKEY = "post:post:totalcount"
	// 帖子下共有多少楼
	COMMENTFLOORKEY = "post:comment:totalcount"
	// 帖子的一楼评论
	FIRSTCOMMENTKEY = "post:post:first_comment"
	// 帖子的总评论数
	COMMENTCOUNTKEY = "post:comments_allcount"
	// 评论的总回复数
	REPLYCOUNTKEY = "post:reply_count_comment"
)

// Init init mysql and redis orm
func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Cfg.Redis.Host, conf.Cfg.Redis.Port),
		Password: conf.Cfg.Redis.Password,
		DB:       0,
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Error("connect to redis error: ", err)
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.Cfg.DB.Host, conf.Cfg.DB.User, conf.Cfg.DB.Password, conf.Cfg.DB.Dbname, conf.Cfg.DB.Port)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	if conf.Cfg.DB.MaxIdle > 0 {
		db.SetMaxIdleConns(conf.Cfg.DB.MaxIdle)
	}
	if conf.Cfg.DB.MaxOpen > 0 {
		db.SetMaxOpenConns(conf.Cfg.DB.MaxOpen)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	entClient = ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := entClient.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("db connection success")
}

// Close close db connection
func Close() {
	entClient.Close()
	redisClient.Close()
}

// InsertPost insert new post
func InsertPost(ctx context.Context, userID int64, title, content string) (int64, error) {
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return 0, err
	}

	post, err := tx.Post.Create().
		SetAuthorID(int(userID)).
		SetTitle(title).
		SetReplyNum(0).
		SetContent(content).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := redisClient.Incr(ctx, POSTSCOUNTKEY).Err(); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return int64(post.ID), nil
}

// GetPostByID get post by postID
func GetPostByID(ctx context.Context, id int64) (*ent.Post, error) {
	post, err := entClient.Post.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return post, nil
}

// GetPosts get posts by page and page_size
func GetPosts(ctx context.Context, page, pageSize int) ([]*ent.Post, error) {
	// posts, err := entClient.Post.Query().Order(ent.Desc(post.FieldUpdateAt)).
	// 	Offset((page - 1) * pageSize).Limit(pageSize).All(ctx)
	// if err != nil {
	// 	return nil, er.NewServiceErr(er.Internal, err)
	// }
	// return posts, nil
	posts, err := ent.GetPosts(ctx, entClient, page, pageSize)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPostsTotalCount get all post count
func GetPostsTotalCount(ctx context.Context) (int64, error) {
	total, err := redisClient.Get(ctx, POSTSCOUNTKEY).Result()
	if err == redis.Nil {
		// 不存在，防穿透
		value, err, _ := postsCountCache.Do("posts_count", func() (ret interface{}, err error) {
			var count int
			count, err = entClient.Post.Query().Where(post.StatusEQ(post.StatusNORMAL)).Count(ctx)
			if err != nil {
				return 0, err
			}
			redisClient.Set(ctx, POSTSCOUNTKEY, count, 0)
			return count, nil
		})
		if err != nil {
			return 0, err
		}
		return value.(int64), nil
	} else if err != nil {
		return 0, err
	}
	return strconv.ParseInt(total, 10, 64)
}

// InsertComment add new comment
func InsertComment(ctx context.Context, userID int64, postID int64, content string) (int64, error) {
	var (
		floor int64
		key   = fmt.Sprintf("%s_%d", COMMENTFLOORKEY, postID)
	)
	//先获取楼层数
	_, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		t, err := entClient.Post.Query().Where(post.ID(int(postID))).QueryComments().Count(ctx)
		if err != nil {
			return 0, err
		}

		floor = int64(t) + 1
		r, err := redisClient.SetNX(ctx, key, floor, 0).Result()
		if err != nil {
			return 0, err
		}
		// 二次检查
		if !r {
			// 已存在
			floor, err = redisClient.Incr(ctx, key).Result()
			if err != nil {
				return 0, err
			}
		}
	} else if err != nil {
		return 0, err
	} else {
		floor, err = redisClient.Incr(ctx, key).Result()
		if err != nil {
			return 0, err
		}
	}

	comment, err := entClient.Comment.Create().
		SetAuthorID(int(userID)).
		SetContent(content).
		SetFloor(int(floor)).
		SetOwnerID(int(postID)).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	if floor == 1 {
		// 一楼，保存到redis
		go SaveFirstComment(postID, map[string]interface{}{
			"id":      comment.ID,
			"content": content,
			"status":  0,
		})
	}
	go increasePostReplyNum(postID)
	go increasePostCommentNum(postID)
	return int64(comment.ID), nil
}

// 帖子回复数+1
func increasePostReplyNum(postID int64) {
	entClient.Post.UpdateOneID(int(postID)).
		AddReplyNum(1).
		Save(context.Background())
}

// 帖子评论数+1
func increasePostCommentNum(postID int64) {
	var (
		ctx   = context.Background()
		key   = fmt.Sprintf("%s_%d", COMMENTCOUNTKEY, postID)
		count int
	)
	count, err := entClient.Post.Query().Where(post.ID(int(postID))).
		QueryComments().Where(comment.StatusEQ(comment.StatusNORMAL)).Count(ctx)
	// count, err := entClient.Comment.Query().Where(comment.PostID(postID), comment.Status(0)).Count(ctx)
	if err != nil {
		return
	}
	redisClient.SetEX(ctx, key, count, time.Hour*24)
}

// 评论回复数+1
func increaseCommentReplyNum(commentID int64) {
	var (
		ctx   = context.Background()
		key   = fmt.Sprintf("%s_%d", REPLYCOUNTKEY, commentID)
		count int
	)
	count, err := entClient.Comment.Query().Where(comment.ID(int(commentID))).
		QueryReplies().Where(reply.StatusEQ(reply.StatusNORMAL)).Count(ctx)
	// count, err := entClient.Reply.Query().Where(reply.CommentID(commentID), reply.Status(0)).Count(ctx)
	if err != nil {
		return
	}
	redisClient.SetEX(ctx, key, count, time.Hour*24)
}

// InsertReply add new reply
func InsertReply(ctx context.Context, userID, postID, commentID, parentID int64, content string) (int64, error) {
	// id, err := ent.InsertReply(ctx, entClient, int(userID), int(commentID),
	// 	int(parentID), content)
	// if err != nil {
	// 	return 0, err
	// }

	rc := entClient.Reply.Create().
		SetAuthorID(int(userID)).
		SetContent(content).
		SetOwnerID(int(commentID))
	if parentID > 0 {
		rc.SetParentID(int(parentID))
	}
	r, err := rc.Save(ctx)

	if err != nil {
		return 0, err
	}
	go increasePostReplyNum(postID)
	go increaseCommentReplyNum(commentID)
	return int64(r.ID), nil
}

// FirstComment
type FirstComment struct {
	ID      int64  `redis:"id"`
	Content string `redis:"content"`
	Status  int8   `redis:"status"`
}

// GetPostFirstComment 获取帖子的第一条评论
// 先从redis中查，记录redis中没有的id，然后再到mysql查，最后拼接结果
func GetPostFirstComment(ctx context.Context, postIDs []int64) ([]*ent.Comment, error) {
	var (
		pipe = redisClient.Pipeline()
	)
	// TODO redis集群时使用需谨慎
	for id := range postIDs {
		pipe.HMGet(ctx, fmt.Sprintf("%s_%d", FIRSTCOMMENTKEY, id), "id", "content", "status")
	}
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	var (
		// 保存redis中不存在的key的id
		ids     []int64
		results = make([]*ent.Comment, len(postIDs))
		// redis不存在的key的id对应的最后结果的索引
		idsIndex = make(map[int64]int)
	)
	for i, cmder := range cmders {
		cmd := cmder.(*redis.SliceCmd)
		var t FirstComment
		cmd.Scan(&t)
		if t.ID == 0 && t.Content == "" {
			// redis HMGet 返回的err不能判断key是否存在,所以用这个方法
			ids = append(ids, postIDs[i])
			idsIndex[postIDs[i]] = i
		} else {
			results[i] = &ent.Comment{
				ID:      int(t.ID),
				Content: t.Content,
				// Status:  t.Status,
			}
		}
	}
	if len(ids) == 0 {
		return results, nil
	}

	var _ids []int
	for _, id := range ids {
		_ids = append(_ids, int(id))
	}
	tmp, err := entClient.Post.Query().Where(post.IDIn(_ids...)).
		QueryComments().Where(comment.StatusEQ(comment.StatusNORMAL), comment.Floor(1)).
		Order(func(s *entsql.Selector) {
			s.OrderBy(comment.FieldID)
		}).All(ctx)
	// tmp, err := entClient.Comment.Query().Where(comment.PostIDIn(ids...), comment.Floor(1)).Order(func(s *sql.Selector) {
	// 	s.OrderBy(comment.FieldID)
	// }).All(ctx)
	if err != nil {
		return nil, err
	}

	var m = make(map[int]*ent.Comment)
	for _, v := range tmp {
		m[v.ID] = v
	}
	for _, id := range ids {
		if _, ok := m[int(id)]; ok {
			results[idsIndex[id]] = m[int(id)]
			go SaveFirstComment(id, map[string]interface{}{
				"id":      m[int(id)].ID,
				"content": m[int(id)].Content,
				"status":  m[int(id)].Status,
			})
		} else {
			results[idsIndex[id]] = &ent.Comment{}
		}
	}
	return results, nil
}

func SaveFirstComment(postID int64, data map[string]interface{}) error {
	var (
		key = fmt.Sprintf("%s_%d", FIRSTCOMMENTKEY, postID)
		ctx = context.Background()
	)
	if err := redisClient.HSet(ctx, key, data).Err(); err != nil {
		return err
	}
	return redisClient.Expire(ctx, key, time.Hour*24*7).Err()
}

func DeletePost(ctx context.Context, postID int64) error {
	_, err := entClient.Post.UpdateOneID(int(postID)).
		SetStatus(post.StatusDELETED).
		Save(ctx)
	if err != nil {
		return err
	}
	if err := redisClient.IncrBy(context.Background(), POSTSCOUNTKEY, -1).Err(); err != nil {
		return err
	}
	return nil
}

// GetCommentsByPostID 获取帖子所属的评论
func GetCommentsByPostID(ctx context.Context, postID int64, page, pageSize int) (comments []*ent.Comment, err error) {
	comments, err = entClient.Post.Query().Where(post.ID(int(postID))).
		QueryComments().Where(comment.StatusEQ(comment.StatusNORMAL)).Offset(pageSize * (page - 1)).
		Limit(pageSize).All(ctx)
	// comments, err = ent.GetCommentsByPostID(context.Background(), entClient, postID, page, pageSize)
	if err != nil {
		return nil, err
	}
	return
}

// GetPostCommentTotalCount get post all comment count
func GetPostCommentTotalCount(ctx context.Context, postID int64) (count int64, err error) {
	key := fmt.Sprintf("%s_%d", COMMENTCOUNTKEY, postID)
	total, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		t, err := entClient.Post.Query().Where(post.ID(int(postID))).
			QueryComments().Where(comment.StatusEQ(comment.StatusNORMAL)).Count(context.Background())
		// t, err := entClient.Comment.Query().Where(comment.PostID(postID), comment.Status(0)).Count(context.Background())
		if err != nil {
			return 0, err
		}
		count = int64(t)
		redisClient.Set(ctx, key, count, time.Hour*24)
		return count, nil
	} else if err != nil {
		return 0, err
	}
	return strconv.ParseInt(total, 10, 64)
}

// GetRepliesByCommentID 获取评论所属的回复
func GetRepliesByCommentID(ctx context.Context, commentID int64, page int, pageSize int) (replies []*ent.Reply, err error) {
	// TODO 前几个回复使用 redis list元素保存json格式的hash
	replies, err = entClient.Comment.Query().Where(comment.ID(int(commentID))).
		QueryReplies().Where(reply.StatusEQ(reply.StatusNORMAL)).Offset(pageSize * (page - 1)).
		Limit(pageSize).All(ctx)
	if err != nil {
		return nil, err
	}
	return
}

// GetCommentReplyTotalCount get comment all reply count
func GetCommentReplyTotalCount(ctx context.Context, commentID int64) (count int64, err error) {
	key := fmt.Sprintf("%s_%d", REPLYCOUNTKEY, commentID)
	total, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		t, err := entClient.Comment.Query().Where(comment.ID(int(commentID))).
			QueryReplies().Where(reply.StatusEQ(reply.StatusNORMAL)).Count(context.Background())
		if err != nil {
			return 0, err
		}
		count = int64(t)
		redisClient.Set(ctx, key, count, time.Hour*24)
		return count, nil
	} else if err != nil {
		return 0, err
	}
	return strconv.ParseInt(total, 10, 64)
}

func GetPostsByUserId(ctx context.Context, userID int64, page, pageSize int) (posts []*ent.Post, err error) {
	posts, err = entClient.Post.Query().Where(post.AuthorID(int(userID)),
		post.StatusEQ(post.StatusNORMAL), post.Pin(0)).
		Offset(pageSize * (page - 1)).
		Limit(pageSize).Order(ent.Desc(post.FieldUpdatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func GetUserPostCount(ctx context.Context, userID int64) (count int64, err error) {
	t, err := entClient.Post.Query().Where(post.AuthorID(int(userID)),
		post.StatusEQ(post.StatusNORMAL)).Count(ctx)
	if err != nil {
		return 0, err
	}
	count = int64(t)
	return count, nil
	// key := fmt.Sprintf("%s_%d", POSTSCOUNTKEY, userID)
	// total, err := redisClient.Get(ctx, key).Result()
	// if err == redis.Nil {
	// 	// 不存在
	// 	t, err := entClient.Post.Query().Where(post.UserID(uint64(userID)), post.Status(0)).Count(context.Background())
	// 	if err != nil {
	// 		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	// 	}
	// 	count = int64(t)
	// 	redisClient.Set(ctx, key, count, time.Hour*24)
	// 	return count, nil
	// } else if err != nil {
	// 	return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	// }
	// return strconv.ParseInt(total, 10, 64)
}

func PinPost(ctx context.Context, postID int64, pinStatus int) error {
	_, err := entClient.Post.UpdateOneID(int(postID)).
		SetPin(int8(pinStatus)).
		// SetUpdateAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetPinPostsByUserId(ctx context.Context, userID int64) (posts []*ent.Post, err error) {
	posts, err = entClient.Post.Query().Where(post.AuthorID(int(userID)),
		post.StatusEQ(post.StatusNORMAL), post.Pin(1)).
		Order(ent.Desc(post.FieldUpdatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return
}
