package user

import (
	"changweiba-backend/graphql"
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"time"
)

type MyUserResolver struct {

}

func (u *MyUserResolver) RegisterUser(ctx context.Context,input graphql.NewUser,conn *grpc.ClientConn) (*graphql.User,error){
	client:=pb.NewAccountClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	user:=pb.NewUserRequest{
		User:&pb.User{
			Name:input.Name,
			Password:input.Password,
		},
		Ip:"122",
	}
	r,err :=client.RegisterUser(ctx,&user)
	if err!=nil{
		logs.Error("Register user error:",err.Error())
		return nil, err
	}
	return &graphql.User{
		ID:r.Id,
		Name:r.Name,
		Avatar:r.Avatar,
		BannedReason:r.BannedReason,
		Score:int(r.Score),
		Status:graphql.UserStatus(r.Status),
		Role:graphql.UserRole(r.Role),
	}, nil
}

func (u *MyUserResolver) LoginUser(ctx context.Context,input graphql.NewUser) (string,error){
	panic("not implemented")
}