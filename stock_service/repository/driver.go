package repository

import (
	"fmt"
	"os"
	"stock_service/common"
	"stock_service/repository/ent"
	"stock_service/repository/ent/stock"

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

// SubscribeStock subscribe stock
func SubscribeStock(stockID int64, userID int64) error {
	user, err := entClient.User.Get(context.Background(), uint64(userID))
	if err != nil {
		if ent.IsNotFound(err) {
			return common.NewServiceErr(common.NotFound, err)
		}
		return common.NewServiceErr(common.Internal, err)
	}
	if err = user.Update().AddSubscribeStockIDs(uint64(stockID)).Exec(context.Background()); err != nil {
		return common.NewServiceErr(common.Internal, err)
	}
	return nil
}

// UnSubscribeStock unsubscribe stock
func UnSubscribeStock(stockID int64, userID int64) error {
	user, err := entClient.User.Get(context.Background(), uint64(userID))
	if err != nil {
		if ent.IsNotFound(err) {
			return common.NewServiceErr(common.NotFound, err)
		}
		return common.NewServiceErr(common.Internal, err)
	}
	if err = user.Update().RemoveSubscribeStockIDs(uint64(stockID)).Exec(context.Background()); err != nil {
		return common.NewServiceErr(common.Internal, err)
	}
	return nil
}

func GetStockBySymbols(symbols ...string) ([]*ent.Stock, error) {
	stocks, err := entClient.Stock.
		Query().
		Where(stock.SymbolIn(symbols...)).
		All(context.Background())
	if err != nil {
		return nil, common.NewServiceErr(common.Internal, err)
	}
	var (
		results []*ent.Stock
		m       = make(map[string]*ent.Stock)
	)
	for _, stock := range stocks {
		m[stock.Symbol] = stock
	}
	for _, symbol := range symbols {
		if stock, ok := m[symbol]; ok {
			results = append(results, stock)
		} else {
			results = append(results, &ent.Stock{})
		}
	}
	return results, nil
}

func InsertStocks(symbols, names []string) ([]*ent.Stock, error) {
	if len(symbols) != len(names) {
		return nil, common.NewServiceErr(common.InvalidArgument, fmt.Errorf("symbols and names length not equal"))
	}
	bulk := make([]*ent.StockCreate, len(symbols))
	for i, symbol := range symbols {
		bulk[i] = entClient.Stock.Create().SetSymbol(symbol).SetName(names[i])
	}
	stocks, err := entClient.Stock.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return nil, common.NewServiceErr(common.Internal, err)
	}
	return stocks, nil
}
