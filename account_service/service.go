package account_service

import (
	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"fmt"
	"net"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc"
)

type User struct {
	pb.UnimplementedAccountServer
}

func NewAccountService(configPath string) {
	conf.Init(configPath)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServer(s, &User{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
