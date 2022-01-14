package main

import (
	"flag"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

	"os"

	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"cw_account_service/repository"

	"cw_account_service/handler"

	log "github.com/shenjing023/llog"
	"github.com/shenjing023/vivy-polaris/contrib/registry"
	"github.com/shenjing023/vivy-polaris/options"
	vp_server "github.com/shenjing023/vivy-polaris/server"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// runAccountService create and run new service
func runAccountService(configPath string) {
	conf.Init(configPath)
	repository.Init()

	etcdConf := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Cfg.Etcd.Host, conf.Cfg.Etcd.Port)},
		DialTimeout: time.Second * 5,
	}
	r, err := registry.NewEtcdRegister(etcdConf, pb.Account_ServiceDesc, "127.0.0.1", fmt.Sprintf("%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed register server: %+v", err)
	}
	defer r.Deregister()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	s := vp_server.NewServer(options.WithDebug(conf.Cfg.Debug))
	pb.RegisterAccountServer(s, &handler.User{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %+v", err)
		}
	}()
	log.Info("service start success")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
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
	runAccountService(*execDir + "/conf/config.yaml")

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
