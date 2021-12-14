package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"stock_service/conf"
	"stock_service/handler"
	"stock_service/repository"
	"syscall"
	"time"

	log "github.com/shenjing023/llog"

	pb "stock_service/pb"

	_ "github.com/go-sql-driver/mysql"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	runStockService(*execDir + "/conf/config.yaml")

	// client, err := ent.Open("mysql", "root:123456@(10.0.0.214:6033)/liuwei1?parseTime=true")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to mysql: %v", err)
	// }
	// defer client.Close()
	// // Run the auto migration tool.
	// if err := client.Debug().Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
}

// runStockService create and run new service
func runStockService(configPath string) {
	conf.Init(configPath)
	repository.Init()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &handler.StockService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}
	r, err := NewRegister(etcdConf, "svc-"+conf.Cfg.SvcName, conf.Cfg.SvcName, "127.0.0.1", fmt.Sprintf("%d", conf.Cfg.Port))
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
