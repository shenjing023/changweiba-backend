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
	"cw_account_service/repository/ent"

	"cw_account_service/repository/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	log "github.com/shenjing023/llog"
	"golang.org/x/net/context"
)

var (
	redisClient *redis.Client
	entClient   *ent.Client
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

// GetRandomAvatar 随机获取一个头像url
func GetRandomAvatar() (url string, err error) {
	var avatars []*ent.Avatar
	avatars, err = entClient.Avatar.Query().All(context.Background())
	if err != nil {
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
func InsertUser(userName, password, avatar string) (int64, error) {
	now := time.Now().Unix()
	user, err := entClient.User.Create().
		SetNickName(userName).
		SetPassword(password).
		SetCreateAt(now).
		SetUpdateAt(now).
		SetAvatar(avatar).
		Save(context.Background())
	if err != nil {
		return 0, common.NewServiceErr(common.Internal, err)
	}
	return user.ID, nil
}

// GetUserByID get user by user_id
func GetUserByID(id int64) (*ent.User, error) {
	user, err := entClient.User.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, common.NewServiceErr(common.NotFound, err)
		}
		return nil, common.NewServiceErr(common.Internal, err)
	}
	return user, nil
}

// CheckUserExistByName 检查user是否已存在
func CheckUserExistByName(userName string) (bool, error) {
	if count, err := entClient.User.Query().Where(user.NickName(userName)).Count(context.Background()); err != nil {
		return false, common.NewServiceErr(common.Internal, err)
	} else if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetUserByName get user by name
func GetUserByName(name string) (*ent.User, error) {
	if user, err := entClient.User.Query().Where(user.NickName(name)).Only(context.Background()); err != nil {
		if ent.IsNotFound(err) {
			return nil, common.NewServiceErr(common.NotFound, err)
		}
		return nil, common.NewServiceErr(common.Internal, err)
	} else {
		return user, nil
	}
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

// GetUsers 批量获取用户信息
func GetUsers(ids []int64) ([]*ent.User, error) {
	// TODO redis
	users, err := entClient.User.Query().Where(user.IDIn(ids...)).Order(func(s *sql.Selector) {
		s.OrderBy(user.FieldID)
	}).All(context.Background())
	if err != nil {
		return nil, common.NewServiceErr(common.Internal, err)
	}

	//可能有的id不存在或重复,需要再排序
	var (
		results []*ent.User
		m       = make(map[int64]*ent.User)
	)
	for _, v := range users {
		m[v.ID] = v
	}
	for _, id := range ids {
		if _, ok := m[id]; ok {
			results = append(results, m[id])
		} else {
			results = append(results, &ent.User{})
		}
	}
	return results, nil
}

// GetBannedReason 获取禁言原因
func GetBannedReason(bannedType int64) (string, error) {
	// TODO redis
	ban, err := entClient.BanType.Get(context.Background(), bannedType)
	if err != nil {
		return "", common.NewServiceErr(common.Internal, err)
	}
	return ban.Content, nil
}
