package main

import (
	"flag"
	"fmt"
	"net"
	"os/signal"
	"syscall"

	"os"

	"cw_post_service/conf"
	pb "cw_post_service/pb"
	"cw_post_service/repository"

	"cw_post_service/handler"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc"
)

// runPostService create and run new service
func runPostService(configPath string) {
	conf.Init(configPath)
	repository.Init()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &handler.PostService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("signal %d received and shutdown service", quit)
	s.GracefulStop()
	stopService()
}

func stopService() {
	repository.Close()
}

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runPostService(*execDir + "/conf/config.yaml")
}
