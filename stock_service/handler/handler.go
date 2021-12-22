package handler

import (
	"context"
	"stock_service/common"
	"stock_service/pb"
	"stock_service/repository"
	"strings"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type StockService struct {
	pb.UnimplementedStockServiceServer
}

// ServiceErr2GRPCErr serviceErr covert to GRPCErr
func ServiceErr2GRPCErr(err error) error {
	if e, ok := err.(*common.ServiceErr); ok {
		if e.Code == common.Internal {
			log.Errorf("Service Internal Error: %v", e.Err)
		}
		if _, ok := common.ErrMap[e.Code]; ok {
			return status.Error(common.ErrMap[e.Code], e.Err.Error())
		}
		return status.Error(codes.Unknown, e.Err.Error())
	}
	return status.Error(codes.Unknown, err.Error())
}

func (StockService) SubscribeStock(ctx context.Context, req *pb.SubscribeStockRequest) (*emptypb.Empty, error) {
	if err := repository.SubscribeStock(req.StockId, req.UserId); err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	return new(emptypb.Empty), nil
}

func (StockService) UnSubscribeStock(ctx context.Context, req *pb.UnSubscribeStockRequest) (*emptypb.Empty, error) {
	if err := repository.UnSubscribeStock(req.StockId, req.UserId); err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	return new(emptypb.Empty), nil
}

func (StockService) SearchStock(ctx context.Context, req *pb.SearchStockRequest) (*pb.SearchStockReply, error) {
	symbolorname := strings.TrimSpace(req.Symbolorname)
	data, err := SearchStock(symbolorname)
	if err != nil {
		log.Errorf("SearchStock Error1: %v", err)
		return nil, ServiceErr2GRPCErr(err)
	}
	var symbols []string
	for _, d := range data {
		symbols = append(symbols, d.Symbol)
	}
	stocks, err := repository.GetStockBySymbols(symbols...)
	if err != nil {
		log.Errorf("SearchStock Error2: %v", err)
		return nil, ServiceErr2GRPCErr(err)
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
		_stocks, err := repository.InsertStocks(_symbols, names)
		if err != nil {
			log.Errorf("SearchStock Error3: %v", err)
			return nil, ServiceErr2GRPCErr(err)
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
	return &pb.SearchStockReply{
		Stocks: replyStocks,
	}, nil
}

func (StockService) SubscribedStocks(ctx context.Context, req *pb.SubscribeStocksRequest) (*pb.SubscribeStocksReply, error) {
	stocks, err := repository.GetSubscribedStocksByUserID(req.UserId)
	if err != nil {
		log.Errorf("SubscribedStocks Error: %v", err)
		return nil, ServiceErr2GRPCErr(err)
	}
	var replyStocks []*pb.StockInfo
	for _, s := range stocks {
		replyStocks = append(replyStocks, &pb.StockInfo{
			Id:     int64(s.ID),
			Symbol: s.Symbol,
			Name:   s.Name,
		})
	}
	return &pb.SubscribeStocksReply{
		Data: replyStocks,
	}, nil
}
