package handler

import (
	"context"
	"gateway/common"
	"gateway/models"
	pb "gateway/pb"
	"time"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
)

func SearchStock(ctx context.Context, symbolorname string) (*models.StockConnection, error) {
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.SearchStockRequest{
		Symbolorname: symbolorname,
	}
	resp, err := client.SearchStock(ctx, &request)
	if err != nil {
		log.Errorf("search stock error: %v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: ServiceError,
		})
	}
	var stocks []*models.Stock
	for _, stock := range resp.Stocks {
		stocks = append(stocks, &models.Stock{
			Symbol: stock.Symbol,
			Name:   stock.Name,
			ID:     int(stock.Id),
		})
	}
	return &models.StockConnection{
		Nodes:      stocks,
		TotalCount: len(stocks),
	}, nil
}

func SubscribeStock(ctx context.Context, intput int) (bool, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Error("subscribeStock get userID from context error: ", err)
		return false, err
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.SubscribeStockRequest{
		StockId: int64(intput),
		UserId:  int64(userID),
	}
	_, err = client.SubscribeStock(ctx, &request)
	if err != nil {
		log.Errorf("subscribe stock error: %v", err)
		return false, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.InvalidArgument: "stock id is not exist",
		})
	}
	return true, nil
}

func UnSubscribeStock(ctx context.Context, intput int) (bool, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Error("unscribeStock get userID from context error: ", err)
		return false, err
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.UnSubscribeStockRequest{
		StockId: int64(intput),
		UserId:  int64(userID),
	}
	_, err = client.UnSubscribeStock(ctx, &request)
	if err != nil {
		log.Errorf("unscribe stock error: %v", err)
		return false, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.InvalidArgument: "stock id is not exist",
		})
	}
	return true, nil
}

func SubscribedStocks(ctx context.Context) (*models.StockConnection, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Error("subscribedStock get userID from context error: ", err)
		return nil, err
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.SubscribeStocksRequest{
		UserId: int64(userID),
	}
	resp, err := client.SubscribedStocks(ctx, &request)
	if err != nil {
		log.Errorf("subscribed stock error: %v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: ServiceError,
		})
	}
	var stocks []*models.Stock
	for _, stock := range resp.Data {
		stocks = append(stocks, &models.Stock{
			Symbol: stock.Symbol,
			Name:   stock.Name,
			ID:     int(stock.Id),
		})
	}
	return &models.StockConnection{
		Nodes:      stocks,
		TotalCount: len(stocks),
	}, nil
}
