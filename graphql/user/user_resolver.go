package user

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql/models"
	"changweiba-backend/pkg/middleware"
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"net"
	"strings"
	"time"
)

const (
	AccountServiceError="system error"
)

type MyUserResolver struct {

}

func (u *MyUserResolver) RegisterUser(ctx context.Context,input models.NewUser,conn *grpc.ClientConn) (string, error){
	//获取客户端ip
	gc,err:=common.GinContextFromContext(ctx)
	if err!=nil{
		logs.Error(err.Error())
		return "", errors.New("system error")
	}
	ip,_,err:=net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	if err!=nil{
		logs.Error("get remote ip error:",err.Error())
		return "", errors.New("system error")
	}
	
	client:=pb.NewAccountClient(conn)
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	
	user:=pb.NewUserRequest{
		Name:input.Name,
		Password:input.Password,
		Ip:ip,
	}
	r,err :=client.RegisterUser(ctx,&user)
	if err!=nil{
		logs.Error("Register user error:",err.Error())
		return "", common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:AccountServiceError,
			codes.AlreadyExists:"该昵称已注册",
			codes.InvalidArgument:"昵称或密码不能为空",
		})
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
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	pbRequest:=pb.LoginRequest{
		Name:input.Name,
		Password:input.Password,
	}
	r,err:=client.Login(ctx,&pbRequest)
	if err!=nil{
		logs.Error("call Login error:",err.Error())
		return "", common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.InvalidArgument:AccountServiceError,
		})
	}
	//生成jwt
	jwt:=middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token,err:=jwt.GenerateToken(r.Id)
	if err!=nil{
		logs.Error("generate jwt token error:",err.Error())
		return "", errors.New(AccountServiceError)
	}
	return token.AccessToken, nil
}

func (u *MyUserResolver) GetUser(ctx context.Context,userId int,conn *grpc.ClientConn) (*models.User,error){
	client:=pb.NewAccountClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	pbUser:=pb.User{
		Id:int64(userId),
	}
	r,err:=client.GetUser(ctx,&pbUser)
	if err!=nil{
		logs.Error("get user error:",err.Error())
		return nil, err
	}
	if r.Id==0{
		//user id不存在
		return nil,errors.New("用户不存在")
	}
	return &models.User{
		ID:int(r.Id),
		Name:r.Name,
		Avatar:r.Avatar,
		Status:models.UserStatus(r.Status),
		Score:int(r.Score),
		BannedReason:r.BannedReason,
		Role:models.UserRole(r.Role),
	}, nil
}