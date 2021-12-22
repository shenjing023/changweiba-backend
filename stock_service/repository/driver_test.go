package repository

import (
	"stock_service/conf"
	"testing"
	"time"
)

func init() {
	conf.Init("../conf/config.yaml")
	Init()
}

func TestGetSubscribeStocks(t *testing.T) {
	defer Close()
	stocks, err := GetSubscribedStocks()
	if err != nil {
		t.Error(err)
	}
	t.Log(stocks)
}

func TestInsertTradeDate(t *testing.T) {
	defer Close()
	err := InsertStockTradeDate(1, "2018-01-02", 13.0, 10000, 11)
	if err != nil {
		t.Error(err)
	}
}

func TestGetStockLastPullTime(t *testing.T) {
	defer Close()
	lastPullTime, err := GetStockLastPullTime(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(lastPullTime)

	now := time.Now()
	t.Log(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix())
}
