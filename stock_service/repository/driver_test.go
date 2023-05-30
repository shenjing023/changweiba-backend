package repository

import (
	"context"
	"fmt"
	"reflect"
	"stock_service/conf"
	"stock_service/repository/ent"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	er "github.com/shenjing023/vivy-polaris/errors"
)

func init() {
	fmt.Println("init")
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
	err := InsertStockTradeDate(ctx, 1, "2018-01-02", 13.0,
		10000, 11, 1, 111, 0, 0, "sss")
	// err := InsertStockTradeDate(ctx, 1, "2018-01-02", 13.0, 10000, 11)

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

func TestSubscribeStock(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
		symbol string
		name   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "stock not exist test1",
			args: args{
				ctx:    context.Background(),
				userID: 8,
				symbol: "SZ600220",
				name:   "*ST新联",
			},
			wantErr: false,
		},
		{
			name: "stock not exist test2",
			args: args{
				ctx:    context.Background(),
				userID: 8,
				symbol: "SZ600221",
				name:   "恒立实业",
			},
			wantErr: false,
		},
		{
			name: "stock exist test1",
			args: args{
				ctx:    context.Background(),
				userID: 7,
				symbol: "SZ600221",
				name:   "恒立实业",
			},
			wantErr: false,
		},
		{
			name: "subscribe stock limit test2",
			args: args{
				ctx:    context.Background(),
				userID: 8,
				symbol: "SZ600223",
				name:   "吉林敖东",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SubscribeStock(tt.args.ctx, tt.args.userID, tt.args.symbol, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnSubscribeStock(t *testing.T) {
	type args struct {
		ctx    context.Context
		symbol string
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "stock not exist test1",
			args: args{
				ctx:    context.Background(),
				userID: 8,
				symbol: "SZ600210",
			},
			wantErr: false,
		},
		{
			name: "user not exist test1",
			args: args{
				ctx:    context.Background(),
				userID: 9,
				symbol: "SZ600210",
			},
			wantErr: true,
		},
		{
			name: "normal test1",
			args: args{
				ctx:    context.Background(),
				userID: 8,
				symbol: "SZ600220",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnSubscribeStock(tt.args.ctx, tt.args.symbol, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UnSubscribeStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSubscribedStocksByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*ent.Stock
		wantErr bool
	}{
		{
			name: "normal test1",
			args: args{
				ctx:    context.Background(),
				userID: 8,
			},
			want:    []*ent.Stock{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSubscribedStocksByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscribedStocksByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscribedStocksByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
