package rpc_conn

import (
	"changweiba-backend/conf"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

//rpc连接
var (
	AccountConn *grpc.ClientConn
	PostConn *grpc.ClientConn
)

func InitRPCConnection(){
	var err error
	AccountConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Account.Port),grpc.WithInsecure())
	if err!=nil{
		log.Fatal(fmt.Sprintf("fail to accountRPC dial: %+v",err))
	}
	PostConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Post.Port),grpc.WithInsecure())
	if err!=nil{
		log.Fatal(fmt.Sprintf("fail to postRPC dial: %+v",err))
	}
}

//关闭rpc连接
func StopRPCConnection(){
	if AccountConn!=nil{
		AccountConn.Close()
	}
	if PostConn!=nil{
		PostConn.Close()
	}
}
