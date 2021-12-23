package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"stock_service/repository"

	"github.com/robfig/cron/v3"
	log "github.com/shenjing023/llog"
)

func UpdateTradeData() {
	// 先获取订阅数不为0的stock
	stocks, err := repository.GetSubscribedStocks()
	if err != nil {
		log.Errorf("get subscribed stocks error:%v", err)
		return
	}
	// 循环获取每个股票的最近拉取的时间
	for _, stock := range stocks {
		// 获取最近拉取的时间
		lastPulledTime, err := repository.GetStockLastPullTime(stock.ID)
		if err != nil {
			log.Errorf("get last pulled time error:%v", err)
			continue
		}
		tradeData := getTradeData(stock.Symbol, lastPulledTime)
		if tradeData == nil || tradeData.Data == nil {
			log.Errorf("get trade data error:%v", err)
			continue
		}
		log.Info("trade data:", tradeData)
		xqComment, err := GetXqCommentData(lastPulledTime, stock.Symbol)
		if err != nil {
			log.Errorf("get xq comment error:%v", err)
			continue
		}
		for _, data := range tradeData.Data[strings.ToLower(stock.Symbol)].Qfqday {
			// 插入数据库
			date := data[0]
			close, _ := strconv.ParseFloat(data[2], 64)
			volume, _ := strconv.ParseFloat(data[5], 64)
			xq := 0
			if _, ok := xqComment[date]; ok {
				xq = xqComment[date]
			}
			if err := repository.InsertStockTradeDate(stock.ID, date, close, volume, int64(xq)); err != nil {
				log.Errorf("insert stock trade date error:%v", err)
			}
		}
	}
	log.Info("update trade data done")
}

func getTradeData(symbol string, lastPulledTime int64) (data *TradeData) {
	// 获取当前时间
	now := time.Now()
	// 两个时间相差的天数
	days := int(now.Sub(time.Unix(lastPulledTime, 0)).Hours() / 24)
	if days <= 0 {
		return
	}
	if days > 600 {
		days = 600
	}

	url := fmt.Sprintf("http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=%s,day,,,%d,qfq&r=0.%d", strings.ToLower(symbol), days, now.UnixNano())
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("get trade data error:%v", err)
		return
	}
	defer resp.Body.Close()
	t, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read trade data error:%v", err)
		return
	}
	s := strings.Split(string(t), "=")
	if len(s) < 2 {
		log.Errorf("split trade data error:%v", err)
		return
	}
	data = &TradeData{}
	err = json.Unmarshal([]byte(s[1]), data)
	if err != nil {
		log.Errorf("unmarshal trade data error:%v", err)
		return
	}
	return
}

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
