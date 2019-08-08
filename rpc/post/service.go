//go:generate protoc --go_out=plugins=grpc:./pb post.proto

package post

import (
	"changweiba-backend/dao"
	pb "changweiba-backend/rpc/post/pb"
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"strings"
)

type Post struct{

}

func (p *Post) NewPost(ctx context.Context,pr *pb.NewPostRequest) (*pb.NewPostResponse,error){
	if len(strings.TrimSpace(pr.Topic))==0{
		return nil,status.Error(codes.InvalidArgument,"topic can not be empty")
	}
	postId,err:=dao.InsertPost(pr.UserId,pr.Topic)
	if err!=nil{
		logs.Error("insert post error: ",err.Error())
		return nil,status.Error(codes.Internal,"post service system error")
	}
	_,err=dao.InsertComment(pr.UserId,postId,pr.Content)
	if err!=nil{
		logs.Error("insert comment error: ",err.Error())
		return nil,status.Error(codes.Internal,"post service system error")
	}
	return &pb.NewPostResponse{
		PostId:postId,
	},nil
}

func (p *Post) NewComment(ctx context.Context,cr *pb.NewCommentRequest) (*pb.NewCommentResponse,error){
	if len(strings.TrimSpace(cr.Content))==0{
		return nil,status.Error(codes.InvalidArgument,"comment content can not be empty")
	}
	commentId,err:=dao.InsertComment(cr.UserId,cr.PostId,cr.Content)
	if err!=nil{
		logs.Error("insert comment error: ",err.Error())
		return nil,status.Error(codes.Internal,"post service system error")
	}
	return &pb.NewCommentResponse{
		CommentId:commentId,
	}, nil
}

func (p *Post) NewReply(ctx context.Context,request *pb.NewReplyRequest) (*pb.NewReplyResponse,error){

}

func NewPostService(addr string,port int){
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",addr,port))
	if err!=nil{
		logs.Error(fmt.Sprintf("post service failed to listen: %+v",err))
		os.Exit(1)
	}
	grpcServer:=grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer,&Post{})
	if err=grpcServer.Serve(lis);err!=nil{
		logs.Error("run grpcserver postservice failed: %+v",err)
		os.Exit(1)
	}
}
