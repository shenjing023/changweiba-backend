package repository

import (
	"fmt"
	"os"

	"cw_account_service/conf"

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
