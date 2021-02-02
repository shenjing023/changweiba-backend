package repository

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"cw_account_service/common"
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

// Close close db connection
func Close() {
	sqlDB, _ := dbOrm.DB()
	sqlDB.Close()
	redisClient.Close()
}

// GetRandomAvatar 随机获取一个头像url
func GetRandomAvatar() (url string, err error) {
	var avatars []Avatar
	if err = dbOrm.Where("status=0").Select("url").Find(&avatars).Error; err != nil {
		return "", common.NewServiceErr(common.Internal, err)
	}
	if len(avatars) == 0 {
		return "", common.NewServiceErr(common.Internal, errors.New("there are no avatar data in db"))
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := seed.Intn(len(avatars))
	url = avatars[index].URL
	return
}

// InsertUser insert new user
func InsertUser(userName, password, ip, avatar string) (int64, error) {
	now := time.Now().Unix()
	user := User{
		Name:       userName,
		Password:   password,
		IP:         InetAtoi(ip),
		CreateTime: now,
		LastUpdate: now,
		Avatar:     avatar,
	}
	if err := dbOrm.Create(&user).Error; err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	}
	return user.ID, nil
}

// GetUserByID get user by user_id
func GetUserByID(id int64) (*User, error) {
	var user User
	if err := dbOrm.Where("id=?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewServiceErr(common.NotFound, err)
		}
		return nil, common.NewServiceErr(common.Internal, err)
	}
	return &user, nil
}

// CheckUserExistByName 检查user是否已存在
func CheckUserExistByName(userName string) (bool, error) {
	var count int64
	if err := dbOrm.Model(&User{}).Where("name=?", userName).Count(&count).Error; err != nil {
		return false, common.NewServiceErr(common.Internal, err)
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// GetUserByName get user by name
func GetUserByName(name string) (*User, error) {
	var user User
	if err := dbOrm.Where("name=?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewServiceErr(common.NotFound, err)
		}
		return nil, common.NewServiceErr(common.Internal, err)
	}
	return &user, nil
}

// InetAtoi ip地址string->int
func InetAtoi(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// InetItoa ip地址int->string
func InetItoa(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

//BytesToInt64 []byte转int64
func BytesToInt64(buf []byte) int64 {
	r, _ := strconv.ParseInt(string(buf), 10, 64)
	return r
}

// BytesToInt32 []byte转int32
func BytesToInt32(buf []byte) int32 {
	r, _ := strconv.ParseInt(string(buf), 10, 32)
	return int32(r)
}
