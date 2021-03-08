package repository

import (
	"cw_post_service/common"
	"errors"
	"fmt"
	"os"
	"time"

	"cw_post_service/conf"

	"github.com/go-redis/redis/v8"
	log "github.com/shenjing023/llog"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	redisClient *redis.Client
	dbOrm       *gorm.DB
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
	now := time.Now().Unix()
	post := Post{
		UserID:     userID,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	if err := dbOrm.Create(&post).Error; err != nil {
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
		(SELECT id FROM cw_post WHERE status=? ORDER BY last_update,id DESC LIMIT ?,?) t2 
		WHERE t1.id=t2.id`,
		0, pageSize, pageSize*(page-1)).Rows()
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
