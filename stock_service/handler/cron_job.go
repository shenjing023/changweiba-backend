package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/shenjing023/llog"
)

// func UpdateTradeData() {
// 	// 先获取订阅数不为0的stock
// 	stocks, err := repository.GetSubscribedStocks()
// 	if err != nil {
// 		log.Errorf("get subscribed stocks error:%v", err)
// 		return
// 	}
// 	// 循环获取每个股票的最近拉取的时间
// 	for _, stock := range stocks {
// 		// 获取最近拉取的时间
// 		lastPulledTime, err := repository.GetStockLastPullTime(stock.ID)
// 		if err != nil {
// 			log.Errorf("get last pulled time error:%v", err)
// 			continue
// 		}
// 		// 获取当前时间
// 		now := time.Now()
// 		// 如果当前时间与最近拉取的时间相差大于一天，则拉取数据
// 		if now.Sub(lastPulledTime) > 24*time.Hour {
// 			// 拉取数据
// 			err = repository.PullTradeData(stock.Code)
// 			if err != nil {
// 				log.Errorf("pull trade data error:%v", err)
// 			}
// 		}
// 	}
// }

func getTradeData(symbol string, lastPulledTime int64) (data TradeData) {
	// 获取当前时间
	now := time.Now()
	// 两个时间相差的天数
	days := int(now.Sub(time.Unix(lastPulledTime, 0)).Hours() / 24)
	if days <= 0 {
		return
	}

	url := fmt.Sprintf("http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=%s,day,,,%d,qfq&r=0.%d", symbol, days, now.UnixNano())
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("get trade data error:%v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
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
	err = json.Unmarshal([]byte(s[1]), &data)
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
		Qfqday [][]string `json:"qfqday"`
	} `json:"data"`
}
