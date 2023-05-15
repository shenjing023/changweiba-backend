package handler

import (
	"context"
	"fmt"
	"gateway/conf"
	"gateway/models"
	"io"
	"net/http"
	"time"

	"gateway/common"

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

type WencaiStockData struct {
	Bull  int    `json:"bull"`
	Short string `json:"short"`
}

type HTTPResponse struct {
	Data WencaiStockData `json:"data"`
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
}

func WencaiStock(ctx context.Context, stockID int) (*models.WencaiStock, error) {
	urL := fmt.Sprintf("http://%s:%d/query/%d", conf.Cfg.Wencai.Host, conf.Cfg.Wencai.Port, stockID)
	resp, err := httpCli.Get(urL)
	if err != nil {
		log.Errorf("get wencai stock error: %+v", err)
		return nil, common.HTTPErrorConvert(err, 500)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read wencai stock error: %+v", err)
		return nil, common.HTTPErrorConvert(err, 500)
	}

	r := &HTTPResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Errorf("unmarshal wencai stock error: %+v", err)
		return nil, common.HTTPErrorConvert(err, 500)
	}

	return &models.WencaiStock{
		Bull:  r.Data.Bull,
		Short: r.Data.Short,
	}, nil
}
