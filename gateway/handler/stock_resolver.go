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
		log.Error("search stock error: ", err)
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
