package repository

import (
	"fmt"
	"os"
	"stock_service/repository/ent"
	"stock_service/repository/ent/user"

	"stock_service/conf"

	"entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/shenjing023/llog"
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

// subscribeStock subscribe stock
func SubscribeStock(symbol string, userID int64) error {
	stockID := 0
	user, err := entClient.User.Query().Where(user.ID(uint64(userID))).Only(context.Background())
	if err != nil {
		return err
	}
	user.Update().AddSubscribeStockIDs(symbol).Exec(context.Background())
}
