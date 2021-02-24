package main

import (
	"flag"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

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
	registerSignalHandler()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &handler.PostService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runPostService(*execDir + "/conf/config.yaml")
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		log.Infof("Signal %d received", sig)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			repository.Close()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
}
