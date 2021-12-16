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

type Stock struct {
	symbol string
	name   string
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
	return nil, nil
}

func (StockService) UnSubscribeStock(ctx context.Context, req *pb.UnSubscribeStockRequest) (*emptypb.Empty, error) {
	if err := repository.UnSubscribeStock(req.StockId, req.UserId); err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	return nil, nil
}

func (StockService) SearchStock(ctx context.Context, req *pb.SearchStockRequest) (*pb.SearchStockReply, error) {
	symbolorname := strings.TrimSpace(req.Symbolorname)
	data, err := common.SearchStock(symbolorname)
	if err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}

}
