package repository

import (
	"cw_post_service/repository/ent"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/reply"
	"fmt"
	"os"
	"strconv"
	"time"

	"cw_post_service/conf"

	"entgo.io/ent/dialect/sql"
	"github.com/cockroachdb/errors"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/shenjing023/llog"
	er "github.com/shenjing023/vivy-polaris/errors"
	"golang.org/x/net/context"
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

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Cfg.DB.User, conf.Cfg.DB.Password, conf.Cfg.DB.Host, conf.Cfg.DB.Port, conf.Cfg.DB.Dbname)

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error("mysql connection error: ", err)
		os.Exit(1)
	}
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	if conf.Cfg.DB.MaxIdle > 0 {
		db.SetMaxIdleConns(conf.Cfg.DB.MaxIdle)
	}
	if conf.Cfg.DB.MaxOpen > 0 {
		db.SetMaxOpenConns(conf.Cfg.DB.MaxOpen)
	}
	entClient = ent.NewClient(ent.Driver(drv))
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
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	now := time.Now().Unix()
	post, err := tx.Post.Create().
		SetUserID(uint64(userID)).
		SetTitle(title).
		SetStatus(0).
		SetCreateAt(now).
		SetUpdateAt(now).
		SetReplyNum(0).
		SetContent(content).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	if err := redisClient.Incr(ctx, POSTSCOUNTKEY).Err(); err != nil {
		tx.Rollback()
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	if err := tx.Commit(); err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return int64(post.ID), nil
}

// GetPostByID get post by postID
func GetPostByID(ctx context.Context, id int64) (*ent.Post, error) {
	post, err := entClient.Post.Get(ctx, uint64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, er.NewServiceErr(er.NotFound, errors.New("post not exist"))
		}
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
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
		return nil, er.NewServiceErr(er.Internal, err)
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
			count, err = entClient.Post.Query().Where(post.Status(0)).Count(ctx)
			if err != nil {
				return 0, err
			}
			redisClient.Set(ctx, POSTSCOUNTKEY, count, 0)
			return count, nil
		})
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
		}
		return value.(int64), nil
	} else if err != nil {
		return 0, er.NewServiceErr(er.Internal, err)
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
		t, err := entClient.Post.Query().Where(post.ID(uint64(postID))).QueryComments().Count(ctx)
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
		}

		floor = int64(t) + 1
		r, err := redisClient.SetNX(ctx, key, floor, 0).Result()
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
		}
		// 二次检查
		if !r {
			// 已存在
			floor, err = redisClient.Incr(ctx, key).Result()
			if err != nil {
				return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
			}
		}
	} else if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	} else {
		floor, err = redisClient.Incr(ctx, key).Result()
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
		}
	}

	now := time.Now().Unix()
	comment, err := entClient.Comment.Create().
		SetUserID(uint64(userID)).
		SetContent(content).
		SetFloor(uint64(floor)).
		SetCreateAt(now).
		SetStatus(0).
		SetOwnerID(uint64(postID)).
		Save(ctx)
	if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
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
	entClient.Post.UpdateOneID(uint64(postID)).
		AddReplyNum(1).
		SetUpdateAt(time.Now().Unix()).
		Save(context.Background())
}

// 帖子评论数+1
func increasePostCommentNum(postID int64) {
	var (
		ctx   = context.Background()
		key   = fmt.Sprintf("%s_%d", COMMENTCOUNTKEY, postID)
		count int
	)
	count, err := entClient.Post.Query().Where(post.ID(uint64(postID))).QueryComments().Where(comment.Status(0)).Count(ctx)
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
	count, err := entClient.Comment.Query().Where(comment.ID(uint64(commentID))).QueryReplies().Where(reply.Status(0)).Count(ctx)
	// count, err := entClient.Reply.Query().Where(reply.CommentID(commentID), reply.Status(0)).Count(ctx)
	if err != nil {
		return
	}
	redisClient.SetEX(ctx, key, count, time.Hour*24)
}

// InsertReply add new reply
func InsertReply(ctx context.Context, userID, postID, commentID, parentID int64, content string) (int64, error) {
	id, err := ent.InsertReply(ctx, entClient, uint64(userID), uint64(commentID), uint64(parentID), content)
	if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	go increasePostReplyNum(postID)
	go increaseCommentReplyNum(commentID)
	return id, nil
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
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
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
				ID:      uint64(t.ID),
				Content: t.Content,
				Status:  t.Status,
			}
		}
	}
	if len(ids) == 0 {
		return results, nil
	}

	var _ids []uint64
	for _, id := range ids {
		_ids = append(_ids, uint64(id))
	}
	tmp, err := entClient.Post.Query().Where(post.IDIn(_ids...)).QueryComments().Where(comment.Status(0), comment.Floor(1)).Order(func(s *sql.Selector) {
		s.OrderBy(comment.FieldID)
	}).All(ctx)
	// tmp, err := entClient.Comment.Query().Where(comment.PostIDIn(ids...), comment.Floor(1)).Order(func(s *sql.Selector) {
	// 	s.OrderBy(comment.FieldID)
	// }).All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}

	var m = make(map[uint64]*ent.Comment)
	for _, v := range tmp {
		m[v.ID] = v
	}
	for _, id := range ids {
		if _, ok := m[uint64(id)]; ok {
			results[idsIndex[id]] = m[uint64(id)]
			go SaveFirstComment(id, map[string]interface{}{
				"id":      m[uint64(id)].ID,
				"content": m[uint64(id)].Content,
				"status":  m[uint64(id)].Status,
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
	_, err := entClient.Post.UpdateOneID(uint64(postID)).
		SetStatus(1).
		SetUpdateAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	if err := redisClient.IncrBy(context.Background(), POSTSCOUNTKEY, -1).Err(); err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	return nil
}

// GetCommentsByPostID 获取帖子所属的评论
func GetCommentsByPostID(ctx context.Context, postID int64, page, pageSize int) (comments []*ent.Comment, err error) {
	comments, err = entClient.Post.Query().Where(post.ID(uint64(postID))).
		QueryComments().Where(comment.Status(0)).Offset(pageSize * (page - 1)).
		Limit(pageSize).All(ctx)
	// comments, err = ent.GetCommentsByPostID(context.Background(), entClient, postID, page, pageSize)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return
}

// GetPostCommentTotalCount get post all comment count
func GetPostCommentTotalCount(ctx context.Context, postID int64) (count int64, err error) {
	key := fmt.Sprintf("%s_%d", COMMENTCOUNTKEY, postID)
	total, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		t, err := entClient.Post.Query().Where(post.ID(uint64(postID))).
			QueryComments().Where(comment.Status(0)).Count(context.Background())
		// t, err := entClient.Comment.Query().Where(comment.PostID(postID), comment.Status(0)).Count(context.Background())
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
		}
		count = int64(t)
		redisClient.Set(ctx, key, count, time.Hour*24)
		return count, nil
	} else if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	return strconv.ParseInt(total, 10, 64)
}

// GetRepliesByCommentID 获取评论所属的回复
func GetRepliesByCommentID(ctx context.Context, commentID int64, page int, pageSize int) (replies []*ent.Reply, err error) {
	// TODO 前几个回复使用 redis list元素保存json格式的hash
	replies, err = entClient.Comment.Query().Where(comment.ID(uint64(commentID))).
		QueryReplies().Where(reply.Status(0)).Offset(pageSize * (page - 1)).
		Limit(pageSize).All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return
}

// GetCommentReplyTotalCount get comment all reply count
func GetCommentReplyTotalCount(ctx context.Context, commentID int64) (count int64, err error) {
	key := fmt.Sprintf("%s_%d", REPLYCOUNTKEY, commentID)
	total, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		t, err := entClient.Comment.Query().Where(comment.ID(uint64(commentID))).QueryReplies().Where(reply.Status(0)).Count(context.Background())
		if err != nil {
			return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
		}
		count = int64(t)
		redisClient.Set(ctx, key, count, time.Hour*24)
		return count, nil
	} else if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	return strconv.ParseInt(total, 10, 64)
}

func GetPostsByUserId(ctx context.Context, userID int64, page, pageSize int) (posts []*ent.Post, err error) {
	posts, err = entClient.Post.Query().Where(post.UserID(uint64(userID)), post.Status(0)).
		Offset(pageSize * (page - 1)).
		Limit(pageSize).Order(ent.Desc(post.FieldUpdateAt)).All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return
}

func GetUserPostCount(ctx context.Context, userID int64) (count int64, err error) {
	t, err := entClient.Post.Query().Where(post.UserID(uint64(userID)),
		post.Status(0)).Count(ctx)
	if err != nil {
		return 0, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
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
