package main

import (
	"cw_account_service/conf"
	"flag"
	"fmt"
	"net"

	"os"

	pb "cw_account_service/pb"

	"cw_account_service/handler"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc"
)

// runAccountService create and run new service
func runAccountService(configPath string) {
	conf.Init(configPath)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServer(s, &handler.User{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runAccountService(*execDir + "/conf/config.yaml")
}
