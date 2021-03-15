package repository

import (
	"cw_post_service/common"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"cw_post_service/conf"

	"github.com/go-redis/redis/v8"
	log "github.com/shenjing023/llog"
	"golang.org/x/net/context"
	"golang.org/x/sync/singleflight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	redisClient     *redis.Client
	dbOrm           *gorm.DB
	postsCountCache singleflight.Group
)

const (
	// POSTSCOUNTKEY redis 保存当前帖子总数
	POSTSCOUNTKEY = "posts_count_key"
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

	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Cfg.DB.User, conf.Cfg.DB.Password, conf.Cfg.DB.Host, conf.Cfg.DB.Port, conf.Cfg.DB.Dbname)
	c := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局禁用表名复数形式
		},
	}
	if conf.Cfg.Debug {
		// default log
		c.Logger = logger.Default
	}
	dbOrm, err = gorm.Open(mysql.Open(dsn), &c)
	if err != nil {
		log.Error("mysql connection error: ", err)
		os.Exit(1)
	}
	sqlDB, err := dbOrm.DB()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if conf.Cfg.DB.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(conf.Cfg.DB.MaxIdle)
	}
	if conf.Cfg.DB.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(conf.Cfg.DB.MaxOpen)
	}
}

// Close close db connection
func Close() {
	sqlDB, _ := dbOrm.DB()
	sqlDB.Close()
	redisClient.Close()
}

// InsertPost insert new post
func InsertPost(userID int64, topic, content string) (int64, error) {
	session := dbOrm.Begin()
	now := time.Now().Unix()
	post := Post{
		UserID:     userID,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	if err := session.Create(&post).Error; err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	}
	if err := redisClient.Incr(context.Background(), POSTSCOUNTKEY).Err(); err != nil {
		session.Rollback()
		return 0, common.NewServiceErr(common.Internal, err)
	}
	return post.ID, nil
}

// GetPost get post
func GetPost(postID int64) (*Post, error) {
	var post Post
	if err := dbOrm.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewServiceErr(common.NotFound, err)
		}
		return nil, common.NewServiceErr(common.Internal, err)
	}
	return &post, nil
}

// GetPosts get posts by page and page_size
func GetPosts(page, pageSize int) ([]*Post, error) {
	var posts []*Post
	// TODO 待优化
	rows, err := dbOrm.Raw(`SELECT t1.* FROM cw_post t1, 
		(SELECT id FROM cw_post WHERE status=? ORDER BY last_update DESC, id DESC LIMIT ?,?) t2 
		WHERE t1.id=t2.id`,
		0, pageSize*(page-1), pageSize).Rows()
	if err != nil {
		return nil, common.NewServiceErr(common.Internal, err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		dbOrm.ScanRows(rows, &p)
		posts = append(posts, &p)
	}
	return posts, nil
}

// GetPostsTotalCount get all post count
func GetPostsTotalCount() (int64, error) {
	total, err := redisClient.Get(context.Background(), POSTSCOUNTKEY).Result()
	if err == redis.Nil {
		// 不存在，防穿透
		value, err, _ := postsCountCache.Do("posts_count", func() (ret interface{}, err error) {
			var count int64
			if err := dbOrm.Raw("select count(*) from cw_post where status=0").Scan(&count).Error; err != nil {
				return 0, err
			}
			redisClient.Set(context.Background(), POSTSCOUNTKEY, count, 0)
			return count, nil
		})
		if err != nil {
			return 0, common.NewServiceErr(common.Internal, err)
		}
		return value.(int64), nil
	} else if err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	}
	return strconv.ParseInt(total, 10, 64)
}

// InsertComment add new comment
func InsertComment(userID int64, postID int64, content string) (int64, error) {
	var (
		floor int64
		ctx   = context.Background()
		key   = fmt.Sprintf("comment_count_post_%d", postID)
	)
	//先获取楼层数
	_, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// 不存在
		sql := "SELECT count(*) AS total FROM cw_comment WHERE post_id=?"
		if err := dbOrm.Raw(sql, postID).Scan(&floor).Error; err != nil {
			return 0, common.NewServiceErr(common.Internal, err)
		}
		floor += 1
		r, err := redisClient.SetNX(ctx, key, floor, 0).Result()
		if err != nil {
			return 0, common.NewServiceErr(common.Internal, err)
		}
		if !r {
			// 已存在
			floor, err = redisClient.Incr(ctx, key).Result()
			if err != nil {
				return 0, common.NewServiceErr(common.Internal, err)
			}
		}
	} else if err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	} else {
		floor, err = redisClient.Incr(ctx, key).Result()
		if err != nil {
			return 0, common.NewServiceErr(common.Internal, err)
		}
	}

	now := time.Now().Unix()
	comment := Comment{
		UserID:     userID,
		PostID:     postID,
		Content:    content,
		CreateTime: now,
		Floor:      floor,
		Status:     0,
	}
	if err := dbOrm.Create(&comment).Error; err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	}
	go increasePostReplyNum(postID)
	return comment.ID, nil
}

//帖子回复数+1
func increasePostReplyNum(postID int64) {
	sql := "UPDATE cw_post SET reply_num=reply_num+1,last_update=UNIX_TIMESTAMP() WHERE id=?"
	dbOrm.Exec(sql, postID)
}

// InsertReply add new reply
func InsertReply(userID, postID, commentID, parentID int64, content string) (int64, error) {
	session := dbOrm.Begin()
	//先获取楼层数
	var floor int64
	// 行锁
	sql := "SELECT count(*) AS total FROM reply WHERE comment_id=? FOR UPDATE"
	if err := session.Raw(sql, commentID).Scan(&floor).Error; err != nil {
		session.Rollback()
		return 0, common.NewServiceErr(common.Internal, err)
	}

	reply := Reply{
		UserID:     userID,
		PostID:     postID,
		Content:    content,
		ParentID:   parentID,
		CommentID:  commentID,
		CreateTime: time.Now().Unix(),
		Status:     0,
		Floor:      floor + 1,
	}
	if err := session.Create(&reply).Error; err != nil {
		session.Rollback()
		return 0, common.NewServiceErr(common.Internal, err)
	}

	session.Commit()
	go increasePostReplyNum(postID)
	return reply.ID, nil
}
