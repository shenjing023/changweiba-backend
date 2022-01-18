package repository

import (
	"context"
	"stock_service/conf"
	"testing"
	"time"

	er "github.com/shenjing023/vivy-polaris/errors"
)

func init() {
	conf.Init("../conf/config.yaml")
	Init()
}

func TestGetSubscribeStocks(t *testing.T) {
	defer Close()
	ctx := context.Background()
	stocks, err := GetSubscribedStocks(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Log(stocks)
}

func TestInsertTradeDate(t *testing.T) {
	defer Close()
	ctx := context.Background()
	err := InsertStockTradeDate(ctx, 1, "2018-01-02", 13.0, 10000, 11)
	if err != nil {
		a, ok := err.(*er.Error)
		if ok {
			t.Log(a.Code)
		}
		t.Error(err)
	}
}

func TestGetStockLastPullTime(t *testing.T) {
	defer Close()
	ctx := context.Background()
	lastPullTime, err := GetStockLastPullTime(ctx, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(lastPullTime)

	now := time.Now()
	t.Log(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix())
}
