package handler

import (
	"fmt"

	"gateway/conf"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc"
)

//rpc连接
var (
	AccountConn *grpc.ClientConn
	PostConn    *grpc.ClientConn
)

// InitGRPCConn init grpc conn
func InitGRPCConn() {
	var err error
	AccountConn, err = grpc.Dial(fmt.Sprintf("%s:%d", conf.Cfg.Account.Host, conf.Cfg.Account.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to accountRPC dial: %+v", err))
	}
	PostConn, err = grpc.Dial(fmt.Sprintf("%s:%d", conf.Cfg.Post.Host, conf.Cfg.Post.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to postsRPC dial: %+v", err))
	}
}

// StopGRPCConn 关闭rpc连接
func StopGRPCConn() {
	if AccountConn != nil {
		AccountConn.Close()
	}
	if PostConn != nil {
		PostConn.Close()
	}
}
