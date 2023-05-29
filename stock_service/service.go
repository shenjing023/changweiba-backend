package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"stock_service/conf"
	"stock_service/handler"
	"stock_service/repository"
	"stock_service/repository/ent"
	"syscall"
	"time"

	log "github.com/shenjing023/llog"

	pb "stock_service/pb"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shenjing023/vivy-polaris/contrib/registry"
	"github.com/shenjing023/vivy-polaris/contrib/tracing"
	"github.com/shenjing023/vivy-polaris/options"
	vp_server "github.com/shenjing023/vivy-polaris/server"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// pwd, _ := os.Getwd()
	// execDir := flag.String("d", pwd, "execute directory")
	// flag.Parse()
	// runStockService(*execDir + "/conf/config.yaml")

	client, err := ent.Open("mysql", "root:123456@(127.0.0.1:3306)/changweiba?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Debug().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// runStockService create and run new service
func runStockService(configPath string) {
	conf.Init(configPath)
	repository.Init()

	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}
	r, err := registry.NewEtcdRegister(etcdConf, pb.StockService_ServiceDesc, "127.0.0.1", fmt.Sprintf("%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed register server: %+v", err)
	}
	defer r.Deregister()

	tp, err := tracing.NewJaegerTracerProvider(conf.Cfg.JaegerCollectURL, "account-server")
	if err != nil {
		log.Fatalf("new JaegerTracerProvider error: %+v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("Error shutting down tracer provider: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	s := vp_server.NewServer(options.WithDebug(conf.Cfg.Debug), options.WithServerTracing(tp))
	pb.RegisterStockServiceServer(s, &handler.StockService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %+v", err)
		}
	}()
	handler.RunCronJob()
	log.Info("service start success")

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
