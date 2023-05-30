package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"stock_service/conf"
	"stock_service/models"
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

// type WencaiStockData struct {
// 	Bull  int    `json:"bull"`
// 	Short string `json:"short"`
// }

type HTTPResponse struct {
	Data models.WencaiStockData `json:"data"`
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
}

func getWencaiStock(ctx context.Context, stockSymbol string) (*models.WencaiStockData, error) {
	urL := fmt.Sprintf("http://%s:%d/query/%s",
		conf.Cfg.Wencai.Host, conf.Cfg.Wencai.Port, stockSymbol)
	resp, err := httpCli.Get(urL)
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

	r := &HTTPResponse{}
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
