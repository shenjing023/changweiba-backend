package repository

import (
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"time"

	"cw_account_service/conf"
	"cw_account_service/repository/ent"

	"cw_account_service/repository/ent/migrate"
	"cw_account_service/repository/ent/user"

	"log"

	"database/sql"

	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v5/stdlib"
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
		log.Fatal("connect to redis error: ", err)
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

// GetRandomAvatar 随机获取一个头像url
func GetRandomAvatar(ctx context.Context) (url string, err error) {
	avatars := []string{
		"https://thumbs2.imgbox.com/2b/7b/oSiwD7s7_t.jpg",
		"https://thumbs2.imgbox.com/10/b9/QczUmHlM_t.jpg",
		"https://thumbs2.imgbox.com/41/50/gky4xYjJ_t.jpg",
		"https://thumbs2.imgbox.com/fd/ea/4OD0kAmq_t.jpg",
		"https://thumbs2.imgbox.com/45/78/A105RLJe_t.jpg",
		"https://thumbs2.imgbox.com/12/d8/5jDVsWPr_t.jpg",
	}

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := seed.Intn(len(avatars))
	url = avatars[index]
	return
}

// InsertUser insert new user
func InsertUser(ctx context.Context, name, password, avatar string) (int64, error) {
	user, err := entClient.User.Create().
		SetName(name).
		SetPassword(password).
		SetAvatar(avatar).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return int64(user.ID), nil
}

// GetUserByID get user by user_id
func GetUserByID(ctx context.Context, id int64) (*ent.User, error) {
	user, err := entClient.User.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// CheckUserExistByName 检查user是否已存在
func CheckUserExistByName(ctx context.Context, userName string) (bool, error) {
	if count, err := entClient.User.Query().Where(user.Name(userName)).Count(ctx); err != nil {
		return false, err
	} else if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetUserByName get user by name
func GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	if user, err := entClient.User.Query().Where(user.Name(name)).Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
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

// BytesToInt64 []byte转int64
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
func GetUsers(ctx context.Context, ids []int64) ([]*ent.User, error) {
	// TODO redis cache
	if len(ids) == 0 {
		return nil, fmt.Errorf("no ids provided")
	}

	_ids := make([]int, len(ids))
	for i, id := range ids {
		_ids[i] = int(id)
	}
	users, err := entClient.User.Query().Where(user.IDIn(_ids...)).Order(func(s *entsql.Selector) {
		s.OrderBy(user.FieldID)
	}).All(ctx)
	if err != nil {
		return nil, err
	}

	//可能有的id不存在或重复,需要再排序
	var (
		results []*ent.User
		m       = make(map[int]*ent.User)
	)
	for _, v := range users {
		m[v.ID] = v
	}
	for _, id := range ids {
		if _, ok := m[int(id)]; ok {
			results = append(results, m[int(id)])
		} else {
			results = append(results, &ent.User{})
		}
	}
	return results, nil
}
