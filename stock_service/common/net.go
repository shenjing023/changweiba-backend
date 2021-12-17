package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stock_service/conf"
	"time"
)

var (
	xqClient = &http.Client{
		Timeout: time.Duration(1) * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}
)

func xueqiuRequest(url, method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {

	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	req.Header.Add("Cookie", "xq_a_token="+conf.Cfg.XueqiuToken)
	resp, err := xqClient.Do(req)
	return resp, err
}

type StockData struct {
	Symbol string `json:"code"`
	Name   string `json:"query"`
	State  int    `json:"state"`
	Type   int    `json:"type"`
}

type SearchResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []StockData `json:"data,omitempty"`
}

func SearchStock(symbolorname string) ([]StockData, error) {
	url := fmt.Sprintf("http://xueqiu.com/query/v1/suggest_stock.json?q=%s&count=5", symbolorname)
	resp, err := xueqiuRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var searchResp SearchResp
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, err
	}
	return searchResp.Data, err
}
