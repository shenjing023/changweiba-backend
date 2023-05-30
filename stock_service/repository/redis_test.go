package repository

import (
	"context"
	"reflect"
	"stock_service/models"
	"testing"
)

func TestSetWencaiData(t *testing.T) {
	type args struct {
		ctx     context.Context
		stockID int
		date    string
		data    *models.WencaiStockData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:     context.Background(),
				stockID: 1,
				date:    "2022-01-01",
				data: &models.WencaiStockData{
					Bull:  1,
					Short: "1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetWencaiData(tt.args.ctx, tt.args.stockID, tt.args.date, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SetWencaiData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetWencaiData(t *testing.T) {
	type args struct {
		ctx     context.Context
		stockID int
		date    string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.WencaiStockData
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:     context.Background(),
				stockID: 1,
				date:    "2022-01-01",
			},
			want: &models.WencaiStockData{
				Bull:  1,
				Short: "1",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ctx:     context.Background(),
				stockID: 2,
				date:    "2022-01-01",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWencaiData(tt.args.ctx, tt.args.stockID, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWencaiData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWencaiData() = %v, want %v", got, tt.want)
			}
		})
	}
}
