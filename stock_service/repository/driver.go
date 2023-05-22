package repository

import (
	"fmt"
	"os"
	"stock_service/repository/ent"
	"stock_service/repository/ent/stock"
	"stock_service/repository/ent/tradedate"
	"stock_service/repository/ent/user"
	"strings"
	"time"

	"stock_service/conf"

	"entgo.io/ent/dialect/sql"
	"github.com/cockroachdb/errors"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/shenjing023/llog"
	er "github.com/shenjing023/vivy-polaris/errors"
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

// SubscribeStock subscribe stock
func SubscribeStock(ctx context.Context, userID int64, symbol, name string) error {
	user, err := entClient.User.Get(ctx, uint64(userID))
	if err != nil {
		if ent.IsNotFound(err) {
			return er.NewServiceErr(er.NotFound, errors.New("user not exist"))
		}
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	// 先查看用户订阅的股票数量
	tmp, _ := GetSubscribedStocksByUserID(ctx, userID)
	if len(tmp) > conf.Cfg.SubscribeStockLimit {
		return er.NewServiceErr(er.Unavailable, errors.New("too many subscribe stocks"))
	}
	// 先查看股票是否存在
	stocks, err := GetStockBySymbols(ctx, symbol)
	if err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	stockID := uint64(0)
	if len(stocks) > 0 {
		if stocks[0].ID == 0 {
			// 不存在，插入新股票
			newStocks, err := InsertStocks(ctx, []string{symbol}, []string{name})
			if err != nil {
				return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
			}
			stockID = newStocks[0].ID
		} else {
			stockID = stocks[0].ID
			// update latest subscribe time
			entClient.Stock.UpdateOneID(stockID).SetLastSubscribeAt(time.Now()).Exec(ctx)
		}
	}
	if err = user.Update().AddSubscribeStockIDs(stockID).Exec(ctx); err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	} else if ent.IsConstraintError(err) {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return nil
}

// UnSubscribeStock unsubscribe stock
func UnSubscribeStock(ctx context.Context, symbol string, userID int64) error {
	user, err := entClient.User.Get(ctx, uint64(userID))
	if err != nil {
		if ent.IsNotFound(err) {
			return er.NewServiceErr(er.NotFound, errors.New("user not exist"))
		}
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	// 先查看股票是否存在
	stocks, err := GetStockBySymbols(ctx, symbol)
	if err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	stockID := uint64(0)
	if len(stocks) > 0 {
		if stocks[0].ID == 0 {
			return nil
		} else {
			stockID = stocks[0].ID
		}
	}
	if err = user.Update().RemoveSubscribeStockIDs(stockID).Exec(ctx); err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return nil
}

func GetStockBySymbols(ctx context.Context, symbols ...string) ([]*ent.Stock, error) {
	stocks, err := entClient.Stock.
		Query().
		Where(stock.SymbolIn(symbols...)).
		All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
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

func InsertStocks(ctx context.Context, symbols, names []string) ([]*ent.Stock, error) {
	if len(symbols) != len(names) {
		return nil, er.NewServiceErr(er.InvalidArgument, errors.New("symbols and names length not equal"))
	}
	bulk := make([]*ent.StockCreate, len(symbols))
	for i, symbol := range symbols {
		bulk[i] = entClient.Stock.Create().SetSymbol(strings.ToUpper(symbol)).SetName(names[i])
	}
	stocks, err := entClient.Stock.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return stocks, nil
}

func GetSubscribedStocksByUserID(ctx context.Context, userID int64) ([]*ent.Stock, error) {
	stocks, err := entClient.User.Query().Where(user.ID(uint64(userID))).QuerySubscribeStocks().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, er.NewServiceErr(er.NotFound, errors.New("user not exist"))
		}
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return stocks, nil
}

// 订阅数不为0的股票
func GetSubscribedStocks(ctx context.Context) ([]*ent.Stock, error) {
	stocks, err := entClient.Stock.Query().Where(stock.HasSubscribers()).All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return stocks, nil
}

// 获取股票交易数据最近拉取的时间
func GetStockLastPullTime(ctx context.Context, stockID uint64) (int64, error) {
	td, err := entClient.TradeDate.Query().Where(tradedate.StockID(stockID)).Order(ent.Desc(tradedate.FieldTDate)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}
	t, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", td.TDate, time.Local)
	return t.Unix(), nil
}

// 插入股票每日交易数据
func InsertStockTradeDate(ctx context.Context, stockID uint64, tradeDate string, open, close,
	max, min, volume float64, xq int64, bull int, short string) error {
	now := time.Now().Unix()
	_, err := entClient.TradeDate.Create().
		SetStockID(stockID).
		SetTDate(tradeDate).
		SetClose(close).
		SetVolume(volume).
		SetXueqiuCommentCount(xq).
		SetCreateAt(now).
		SetUpdateAt(now).
		SetOpen(open).
		SetMax(max).
		SetMin(min).
		SetBull(bull).
		SetShort(short).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil
		}
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return nil
}

// 获取股票交易数据
func GetStockTradeDate(ctx context.Context, stockID uint64) ([]*ent.TradeDate, error) {
	data, err := entClient.TradeDate.Query().Where(tradedate.StockID(stockID)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		d.TDate = d.TDate[:10]
	}
	return data, nil
}

// 获取所有stock
func GetAllStocks(ctx context.Context) ([]*ent.Stock, error) {
	stocks, err := entClient.Stock.Query().All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return stocks, nil
}

func GetStockById(ctx context.Context, id uint64) (*ent.Stock, error) {
	stock, err := entClient.Stock.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, er.NewServiceErr(er.NotFound, errors.New("stock not exist"))
		}
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return stock, nil
}

func GetSubscribedUsersByStockId(ctx context.Context, stockID uint64) ([]*ent.User, error) {
	users, err := entClient.Stock.Query().Where(stock.ID(stockID)).
		QuerySubscribers().All(ctx)
	if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "ent error"))
	}
	return users, nil
}
