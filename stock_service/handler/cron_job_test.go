package handler

import (
	"context"
	"reflect"
	"stock_service/conf"
	"stock_service/repository"
	"testing"
)

func init() {
	conf.Init("../conf/config.yaml")
	repository.Init()
}

func Test_getTradeData(t *testing.T) {
	type args struct {
		symbol         string
		lastPulledTime int64
	}
	tests := []struct {
		name     string
		args     args
		wantData *TradeData
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				symbol:         "sh600000",
				lastPulledTime: 1684490209,
			},
			wantData: &TradeData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData, _ := getTradeData(tt.args.symbol); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("getTradeData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestUpdateTradeData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateTradeData()
		})
	}
}

func TestGetWencaiHot(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getWencaiHot(tt.args.ctx)
		})
	}
}
