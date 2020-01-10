package rpc_conn

import (
	accountpb "changweiba-backend/rpc/account/pb"
	postpb "changweiba-backend/rpc/post/pb"
	"github.com/micro/go-micro"
)

//rpc连接
var (
	//AccountConn *grpc.ClientConn
	//PostConn *grpc.ClientConn
	PostClient postpb.PostService
	AccountClient accountpb.AccountService
)

func InitRPCConnection(){
	//var err error
	//AccountConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Account.Port),grpc.WithInsecure())
	//if err!=nil{
	//	fmt.Println(2222)
	//	log.Fatal(fmt.Sprintf("fail to accountRPC dial: %+v",err))
	//}
	//PostConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Post.Port),grpc.WithInsecure())
	//if err!=nil{
	//	log.Fatal(fmt.Sprintf("fail to postRPC dial: %+v",err))
	//}
	
	postService:=micro.NewService(micro.Name("srv.post.client"))
	postService.Init()
	PostClient=postpb.NewPostService("srv.post",postService.Client())
	accountService:=micro.NewService(micro.Name("srv.account.client"))
	accountService.Init()
	AccountClient=accountpb.NewAccountService("srv.account",accountService.Client())
}

//关闭rpc连接
func StopRPCConnection(){
	//if AccountConn!=nil{
	//	AccountConn.Close()
	//}
	//if PostConn!=nil{
	//	PostConn.Close()
	//}
}
