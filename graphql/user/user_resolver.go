package user

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql/models"
	"changweiba-backend/pkg/middleware"
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"time"
)

type MyUserResolver struct {

}

func (u *MyUserResolver) RegisterUser(ctx context.Context,input models.NewUser,conn *grpc.ClientConn) (string, error){
	client:=pb.NewAccountClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	user:=pb.NewUserRequest{
		Name:input.Name,
		Password:input.Password,
		Ip:"127.0.0.1",
	}
	r,err :=client.RegisterUser(ctx,&user)
	if err!=nil{
		logs.Error("Register user error:",err.Error())
		return "", err
	}
	//生成jwt
	jwt:=middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
		)
	token,err:=jwt.GenerateToken(r.Id)
	if err!=nil{
		logs.Error("generate jwt token error:",err.Error())
		return "", err
	}
	return token.AccessToken, nil
}

func (u *MyUserResolver) LoginUser(ctx context.Context,input models.NewUser,conn *grpc.ClientConn) (string,error){
	client:=pb.NewAccountClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	pbRequest:=pb.LoginRequest{
		Name:input.Name,
		Password:input.Password,
	}
	r,err:=client.Login(ctx,&pbRequest)
	if err!=nil{
		logs.Error("call Login error:",err.Error())
		return "", common.GRPCErrorConvert(err)
	}
	//生成jwt
	jwt:=middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token,err:=jwt.GenerateToken(r.Id)
	if err!=nil{
		logs.Error("generate jwt token error:",err.Error())
		return "", err
	}
	return token.AccessToken, nil
}