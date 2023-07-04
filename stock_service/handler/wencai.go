package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"stock_service/conf"
	"stock_service/models"
	"stock_service/repository"
	"strings"
	"time"

	"encoding/json"

	log "github.com/shenjing023/llog"
)

var httpCli = &http.Client{
	Timeout: time.Duration(30) * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		MaxConnsPerHost:     1,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	},
}

var httpCli2 = &http.Client{
	Timeout: time.Duration(120) * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		MaxConnsPerHost:     1,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	},
}

// type WencaiStockData struct {
// 	Bull  int    `json:"bull"`
// 	Short string `json:"short"`
// }

type HTTPResponse[T any] struct {
	Data T      `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func getWencaiStock(ctx context.Context, stockSymbol string) (*models.WencaiStockData, error) {
	url := fmt.Sprintf("http://%s:%d/query/%s",
		conf.Cfg.Wencai.Host, conf.Cfg.Wencai.Port, stockSymbol)
	resp, err := httpCli.Get(url)
	if err != nil {
		log.Errorf("get wencai stock error: %+v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read wencai stock error: %+v", err)
		return nil, err
	}

	r := &HTTPResponse[models.WencaiStockData]{}
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Errorf("unmarshal wencai stock error: %+v", err)
		return nil, err
	}

	return &models.WencaiStockData{
		Bull:  r.Data.Bull,
		Short: r.Data.Short,
	}, nil
}

func getStockAnalyse(ctx context.Context, symbol, name string) (*models.WencaiStackAnalyse, error) {
	url := fmt.Sprintf("http://%s:%d/stock/analyse",
		conf.Cfg.Wencai.Host, conf.Cfg.Wencai.Port)
	data := map[string]any{
		"name":   name,
		"symbol": symbol,
	}

	payload, _ := json.Marshal(data)
	reqBody := bytes.NewBuffer(payload)

	resp, err := httpCli2.Post(url, "application/json", reqBody)
	if err != nil {
		log.Infof("get stock analyse http error: %+v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("get stock analyse read wencai stock error: %+v", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Errorf("get stock analyse http error, code: %+v, msg: %+s", resp.StatusCode, string(body))
		return nil, err
	}

	r := &HTTPResponse[models.WencaiStackAnalyse]{}
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Errorf("get stock analyse unmarshal wencai stock error: %+v", err)
		return nil, err
	}

	return &models.WencaiStackAnalyse{
		Result: r.Data.Result,
	}, nil
}

// 获取同花顺热榜数据
func getWencaiHot(ctx context.Context) ([]*models.WencaiHotStack, error) {
	u, err := url.Parse("https://dq.10jqka.com.cn/fuyao/hot_list_data/out/hot_list/v1/stock")
	if err != nil {
		log.Infof("wencai_hot 解析URL时出错：%s", err)
		return nil, err
	}
	// 构建请求参数
	params := url.Values{}
	params.Add("stock_type", "a")
	params.Add("type", "day")
	params.Add("list_type", "normal")

	u.RawQuery = params.Encode()
	resp, err := httpCli.Get(u.String())
	if err != nil {
		log.Infof("wencai_hot 发送GET请求时出错：%s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Infof("读取响应内容时出错：%s", err)
		return nil, err
	}

	type ResponseData struct {
		StatusCode int `json:"status_code"`
		Data       struct {
			StockList []*models.WencaiHotStack `json:"stock_list"`
		} `json:"data"`
		StatusMsg string `json:"status_msg"`
	}

	r := &ResponseData{}
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Errorf("unmarshal wencai hot stock error: %+v", err)
		return nil, err
	}

	if r.StatusCode != 0 {
		log.Errorf("wencai hot stock error: %s", r.StatusMsg)
		return nil, errors.New(r.StatusMsg)
	}

	return r.Data.StockList[:10], nil
}

func WencaiHot() {
	ctx := context.Background()
	// 获取同花顺热榜数据
	hotList, err := getWencaiHot(ctx)
	if err != nil {
		log.Errorf("get wencai hot error: %+v", err)
		return
	}

	today := time.Now().Format("2006-01-02")
	for _, hot := range hotList {
		symbol := map[int]string{
			33: "SZ",
			17: "SH",
		}[hot.Market] + hot.Code
		analyse, err := getStockAnalyse(ctx, symbol, hot.Name)
		if err != nil {
			log.Errorf("get stock analyse error: %+v", err)
			continue
		}
		wencaiData := getWencaiData(symbol)
		if wencaiData == nil {
			log.Errorf("get wencai data error")
			continue
		}
		tag := "[" + strings.Join(hot.Tag.ConceptTag, ",") + "]"
		if err := repository.InsertHotStock(ctx, symbol, hot.Name, today,
			analyse.Result, tag, hot.Order, wencaiData.Bull, wencaiData.Short); err != nil {
			log.Errorf("insert hot stock error: %+v", err)
		}
	}

	go repository.SaveHotStocks(ctx, today)
}
