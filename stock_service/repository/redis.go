package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"stock_service/models"
	"stock_service/repository/ent"

	"github.com/cockroachdb/errors"
	"github.com/go-redis/redis/v8"
	er "github.com/shenjing023/vivy-polaris/errors"
)

const (
	WCKEY = "wencai"
	HOT   = "hot"
)

func SetWencaiData(ctx context.Context, stockID int, date string, data *models.WencaiStockData) error {
	key := fmt.Sprintf("%s:%d:%s", WCKEY, stockID, date)
	_, err := redisClient.SetNX(ctx, key, data, time.Hour*24*7).Result()
	if err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	return nil
}

func GetWencaiData(ctx context.Context, stockID int, date string) (*models.WencaiStockData, error) {
	key := fmt.Sprintf("%s:%d:%s", WCKEY, stockID, date)
	data := new(models.WencaiStockData)
	err := redisClient.Get(ctx, key).Scan(data)
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}

	return data, nil
}

func SaveHotStocks(ctx context.Context, date string) error {
	data, err := GetHotStocks1(ctx, date)
	if err != nil {
		return err
	}
	_byte, _ := json.Marshal(data)
	key := fmt.Sprintf("%s:%s:%s", WCKEY, HOT, date)
	_, err = redisClient.SetNX(ctx, key, string(_byte), time.Hour*24*7).Result()
	if err != nil {
		return er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	return nil
}

func GetHotStocks2(ctx context.Context, date string) ([]*ent.Hot, error) {
	key := fmt.Sprintf("%s:%s:%s", WCKEY, HOT, date)
	data := new([]*ent.Hot)
	value, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, er.NewServiceErr(er.Internal, errors.Wrap(err, "redis error"))
	}
	json.Unmarshal([]byte(value), &data)
	return *data, nil
}
