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
		log.Errorf("search stock error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
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

func SubscribeStock(ctx context.Context, symbol, name string) (bool, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("subscribeStock get userID from context error: %+v", err)
		return false, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.SubscribeStockRequest{
		Symbol: symbol,
		UserId: int64(userID),
		Name:   name,
	}
	_, err = client.SubscribeStock(ctx, &request)
	if err != nil {
		log.Errorf("subscribe stock error: %+v", err)
		return false, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        common.ServiceError,
			codes.InvalidArgument: "stock id is not exist",
		})
	}
	return true, nil
}

func UnSubscribeStock(ctx context.Context, symbol string) (bool, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("unscribeStock get userID from context error: %+v", err)
		return false, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.UnSubscribeStockRequest{
		Symbol: symbol,
		UserId: int64(userID),
	}
	_, err = client.UnSubscribeStock(ctx, &request)
	if err != nil {
		log.Errorf("unscribe stock error: %+v", err)
		return false, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        common.ServiceError,
			codes.InvalidArgument: "stock id is not exist",
		})
	}
	return true, nil
}

func SubscribedStocks(ctx context.Context) (*models.StockConnection, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("subscribedStock get userID from context error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.SubscribeStocksRequest{
		UserId: int64(userID),
	}
	resp, err := client.SubscribedStocks(ctx, &request)
	if err != nil {
		log.Errorf("subscribed stock error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var stocks []*models.Stock
	for _, stock := range resp.Data {
		stocks = append(stocks, &models.Stock{
			Symbol: stock.Symbol,
			Name:   stock.Name,
			ID:     int(stock.Id),
			Bull:   int(stock.Bull),
			Short:  stock.Short,
		})
	}
	return &models.StockConnection{
		Nodes:      stocks,
		TotalCount: len(stocks),
	}, nil
}

func StockTrades(ctx context.Context, stockID int) (*models.TradeDateConnection, error) {
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.StockTradeDataRequest{
		Id: int64(stockID),
	}
	resp, err := client.StockTradeData(ctx, &request)
	if err != nil {
		log.Errorf("get stock trades error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var trades []*models.TradeDate
	for _, trade := range resp.TradeData {
		trades = append(trades, &models.TradeDate{
			Date:   trade.Date,
			Close:  float64(trade.Close),
			Volume: int(trade.Volume),
			Xq:     int(trade.XueqiuCount),
			Open:   float64(trade.Open),
			Max:    float64(trade.Max),
			Min:    float64(trade.Min),
			Bull:   int(trade.Bull),
			Short:  trade.Short,
		})
	}
	return &models.TradeDateConnection{
		Nodes:      trades,
		TotalCount: len(trades),
		ID:         stockID,
	}, nil
}

func HotStocks(ctx context.Context, date string) (*models.HotStockConnection, error) {
	client := pb.NewStockServiceClient(StockConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.HotStocksRequest{
		Date: date,
	}
	resp, err := client.HotStocks(ctx, &request)
	if err != nil {
		log.Errorf("get hot stocks error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var stocks []*models.HotStock
	for _, stock := range resp.HotStocks {
		stocks = append(stocks, &models.HotStock{
			Symbol:  stock.Symbol,
			Name:    stock.Name,
			Bull:    int(stock.Bull),
			Short:   stock.Short,
			Order:   int(stock.Order),
			Analyse: stock.Analyse,
			Tag:     stock.Tag,
			Date:    stock.Date,
		})
	}
	return &models.HotStockConnection{
		Nodes:      stocks,
		TotalCount: len(stocks),
	}, nil
}
