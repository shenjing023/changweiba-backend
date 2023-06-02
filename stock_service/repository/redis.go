package repository

import (
	"context"
	"fmt"
	"time"

	"stock_service/models"

	"github.com/cockroachdb/errors"
	"github.com/go-redis/redis/v8"
	er "github.com/shenjing023/vivy-polaris/errors"
)

const (
	WCKEY = "wencai"
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
