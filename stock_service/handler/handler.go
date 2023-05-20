package handler

import (
	"context"
	"stock_service/pb"
	"stock_service/repository"
	"strings"

	log "github.com/shenjing023/llog"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type StockService struct {
	pb.UnimplementedStockServiceServer
}

func (StockService) SubscribeStock(ctx context.Context, req *pb.SubscribeStockRequest) (*emptypb.Empty, error) {
	if err := repository.SubscribeStock(ctx, req.UserId, req.Symbol, req.Name); err != nil {
		log.Errorf("SubscribeStock Error: %+v", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (StockService) UnSubscribeStock(ctx context.Context, req *pb.UnSubscribeStockRequest) (*emptypb.Empty, error) {
	if err := repository.UnSubscribeStock(ctx, req.Symbol, req.UserId); err != nil {
		log.Errorf("UnSubscribeStock Error: %+v", err)
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (StockService) SearchStock(ctx context.Context, req *pb.SearchStockRequest) (*pb.SearchStockResponse, error) {
	symbolorname := strings.TrimSpace(req.Symbolorname)
	data, err := SearchStock(symbolorname)
	if err != nil {
		log.Errorf("SearchStock Error1: %v", err)
		return nil, err
	}
	var symbols []string
	for _, d := range data {
		symbols = append(symbols, d.Symbol)
	}
	stocks, err := repository.GetStockBySymbols(ctx, symbols...)
	if err != nil {
		log.Errorf("SearchStock Error2: %+v", err)
		return nil, err
	}
	var (
		_symbols []string
		names    []string
		index    []int
	)
	for i, s := range stocks {
		// 不存在的股票，插入db
		if s.Symbol == "" {
			_symbols = append(_symbols, data[i].Symbol)
			names = append(names, data[i].Name)
			index = append(index, i)
		}
	}
	if len(_symbols) > 0 {
		_stocks, err := repository.InsertStocks(ctx, _symbols, names)
		if err != nil {
			log.Errorf("SearchStock Error3: %+v", err)
			return nil, err
		}
		for i, s := range _stocks {
			stocks[index[i]] = s
		}
	}
	var replyStocks []*pb.StockInfo
	for _, s := range stocks {
		replyStocks = append(replyStocks, &pb.StockInfo{
			Id:     int64(s.ID),
			Symbol: s.Symbol,
			Name:   s.Name,
		})
	}
	return &pb.SearchStockResponse{
		Stocks: replyStocks,
	}, nil
}

func (StockService) SubscribedStocks(ctx context.Context, req *pb.SubscribeStocksRequest) (*pb.SubscribeStocksResponse, error) {
	stocks, err := repository.GetSubscribedStocksByUserID(ctx, req.UserId)
	if err != nil {
		log.Errorf("SubscribedStocks Error: %+v", err)
		return nil, err
	}
	var replyStocks []*pb.StockInfo
	for _, s := range stocks {
		replyStocks = append(replyStocks, &pb.StockInfo{
			Id:     int64(s.ID),
			Symbol: s.Symbol,
			Name:   s.Name,
		})
	}
	return &pb.SubscribeStocksResponse{
		Data: replyStocks,
	}, nil
}

func (StockService) StockTradeData(ctx context.Context, req *pb.StockTradeDataRequest) (*pb.StockTradeDataResponse, error) {
	data, err := repository.GetStockTradeDate(ctx, uint64(req.Id))
	if err != nil {
		log.Errorf("StockTradeData Error: %+v", err)
		return nil, err
	}
	var replyData []*pb.TradeData
	for _, d := range data {
		replyData = append(replyData, &pb.TradeData{
			Date:        d.TDate,
			Close:       d.Close,
			Volume:      int64(d.Volume),
			XueqiuCount: d.XueqiuCommentCount,
		})
	}
	return &pb.StockTradeDataResponse{
		TradeData: replyData,
		Info:      &pb.StockInfo{Id: req.Id},
	}, nil
}
