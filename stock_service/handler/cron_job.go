package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"stock_service/conf"
	"stock_service/repository"

	"github.com/robfig/cron/v3"
	log "github.com/shenjing023/llog"
)

// func UpdateTradeData() {
// 	// 先获取订阅数不为0的stock
// 	stocks, err := repository.GetSubscribedStocks(context.Background())
// 	if err != nil {
// 		log.Errorf("get subscribed stocks error:%+v", err)
// 		return
// 	}
// 	// 循环获取每个股票的最近拉取的时间
// 	for _, stock := range stocks {
// 		// 获取最近拉取的时间
// 		lastPulledTime, err := repository.GetStockLastPullTime(context.Background(), stock.ID)
// 		if err != nil {
// 			log.Errorf("get last pulled time error:%+v", err)
// 			continue
// 		}
// 		tradeData := getTradeData(stock.Symbol, lastPulledTime)
// 		if tradeData == nil || tradeData.Data == nil {
// 			log.Errorf("get trade data error:%v", err)
// 			continue
// 		}
// 		xqComment, err := GetXqCommentData(lastPulledTime, stock.Symbol)
// 		if err != nil {
// 			log.Errorf("get xq comment error:%v", err)
// 			continue
// 		}
// 		for _, data := range tradeData.Data[strings.ToLower(stock.Symbol)].Qfqday {
// 			// 插入数据库
// 			date := data[0]
// 			close, _ := strconv.ParseFloat(data[2], 64)
// 			volume, _ := strconv.ParseFloat(data[5], 64)
// 			xq := 0
// 			if _, ok := xqComment[date]; ok {
// 				xq = xqComment[date]
// 			}
// 			if err := repository.InsertStockTradeDate(context.Background(), stock.ID, date, close, volume, int64(xq)); err != nil {
// 				log.Errorf("insert stock trade date error:%+v", err)
// 			}
// 		}
// 	}
// 	log.Info("update trade data done")
// }

// func getTradeData(symbol string, lastPulledTime int64) (data *TradeData) {
// 	// 获取当前时间
// 	now := time.Now()
// 	// 两个时间相差的天数
// 	days := int(now.Sub(time.Unix(lastPulledTime, 0)).Hours() / 24)
// 	if days <= 0 {
// 		return
// 	}
// 	if days > 600 {
// 		days = 600
// 	}

// 	url := fmt.Sprintf("http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=%s,day,,,%d,qfq&r=0.%d", strings.ToLower(symbol), days, now.UnixNano())
// 	fmt.Printf("url:%s\n", url)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Errorf("get trade data error:%v", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	t, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Errorf("read trade data error:%v", err)
// 		return
// 	}
// 	s := strings.Split(string(t), "=")
// 	if len(s) < 2 {
// 		log.Errorf("split trade data error:%v", err)
// 		return
// 	}
// 	data = &TradeData{}
// 	err = json.Unmarshal([]byte(s[1]), data)
// 	if err != nil {
// 		log.Errorf("unmarshal trade data error:%v", err)
// 		return
// 	}
// 	return
// }

type TradeData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data map[string]struct {
		/*
			格式: [["2021-12-22","17.620","17.390","17.640","17.300","976928.000"]]
			日期，开盘价，收盘价，最高价，最低价，成交量
		*/
		Qfqday [][]string `json:"qfqday"`
	} `json:"data"`
}

func RunCronJob() {
	c := cron.New()
	// 每天凌晨1点更新股票交易数据
	c.AddFunc("0 1 * * *", UpdateTradeData)
	c.Start()
}

type updateStock struct {
	Symbol  string `json:"symbol"`
	StockId int64  `json:"stock_id"`
}

// 获取需要更新数据的stock
func getUpdateStocks(ctx context.Context) (uStocks []*updateStock) {
	stocks, _ := repository.GetAllStocks(ctx)
	for _, stock := range stocks {
		users, _ := repository.GetSubscribedUsersByStockId(ctx, stock.ID)
		// 没有订阅用户且订阅时间超过限制
		if len(users) == 0 &&
			time.Since(stock.LastSubscribeAt).Abs().Hours() > float64(conf.Cfg.SubscribeDurationLimit) {
			continue
		}
		uStocks = append(uStocks, &updateStock{
			Symbol:  stock.Symbol,
			StockId: int64(stock.ID),
		})
	}
	return
}

func UpdateTradeData() {
	ctx := context.Background()
	uStocks := getUpdateStocks(ctx)
	for _, uStock := range uStocks {
		// 获取日k数据
		tradeData, _ := getTradeData(uStock.Symbol)
		if tradeData == nil || tradeData.Data == nil {
			log.Errorf("get trade_data stock[%s] error", uStock.Symbol)
			continue
		}
		// 获取wencai数据
		wencaiData := getWencaiData(uStock.Symbol)

		fmt.Printf("data:%+v\n", tradeData.Data)

		for _, data := range tradeData.Data[strings.ToLower(uStock.Symbol)].Qfqday {
			// 插入数据
			date := data[0]
			close, _ := strconv.ParseFloat(data[2], 64)
			volume, _ := strconv.ParseFloat(data[5], 64)
			open, _ := strconv.ParseFloat(data[1], 64)
			max, _ := strconv.ParseFloat(data[3], 64)
			min, _ := strconv.ParseFloat(data[4], 64)

			if err := repository.InsertStockTradeDate(ctx, uint64(uStock.StockId),
				date, open, close, max, min, volume, 0,
				wencaiData.Bull, wencaiData.Short); err != nil {
				log.Errorf("insert stock[%s] trade_date error:%+v", uStock.Symbol, err)
			}
		}

		// 更新bull
		if err := repository.UpdateStockBullAndShort(ctx, uint64(uStock.StockId), wencaiData.Bull, wencaiData.Short); err != nil {
			log.Errorf("update stock[%s] bull error:%+v", uStock.Symbol, err)
		}
	}
}

func getTradeData(symbol string) (*TradeData, error) {
	now := time.Now()
	var data *TradeData
	url := fmt.Sprintf("http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=%s,day,,,1,qfq&r=0.%d",
		strings.ToLower(symbol), now.UnixNano())
	fmt.Printf("url:%s\n", url)
	for i := 0; i < 3; i++ {
		resp, err := http.Get(url)
		if err != nil {
			log.Infof("trade_data url[%s] request failed: %v", url, err)
			time.Sleep(time.Second)
			continue
		}
		defer resp.Body.Close()
		t, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("read trade data error:%v", err)
			return nil, err
		}
		s := strings.Split(string(t), "=")
		if len(s) < 2 {
			log.Errorf("split trade data error:%v", err)
			return nil, err
		}
		data = &TradeData{}
		err = json.Unmarshal([]byte(s[1]), data)
		if err != nil {
			log.Errorf("unmarshal trade data error:%v", err)
			return nil, err
		}
		if data.Code != 0 {
			log.Info("get trade data retry")
			time.Sleep(time.Second)
			continue
		}
		return data, nil
	}
	return nil, nil
}

func getWencaiData(symbol string) *WencaiStockData {
	data, err := getWencaiStock(context.Background(), symbol)
	if err != nil {
		log.Errorf("get wencai data error:%v", err)
		return nil
	}
	return data
}
