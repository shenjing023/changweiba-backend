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
	clientv3 "go.etcd.io/etcd/client/v3"
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

	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}
	r, err := NewRegister(etcdConf, "svc", conf.Cfg.SvcName, "127.0.0.1", fmt.Sprintf("%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed register serve: %v", err)
	}
	log.Info("service start success")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("signal %d received and shutdown service", quit)
	r.Close()
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
