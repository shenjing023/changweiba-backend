package handler

import (
	"context"
	"errors"
	"gateway/models"
	"net"
	"strings"
	"time"

	"gateway/common"

	pb "gateway/pb"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
)

const (
	ServiceError = "gateway service internal error"
)

// SignUp 用户注册
func SignUp(ctx context.Context, input models.NewUser) (string, error) {
	//获取客户端ip
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		log.Error("%+v", err)
		return "", errors.New(ServiceError)
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	if err != nil {
		log.Error("get remote ip error: ", err)
		return "", errors.New(ServiceError)
	}

	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user:=pb.SignUpRequest{
		Name: input.Name,
		Password: input.Password,
		Ip: ip,
	}
	resp,err:=client.SignUp(ctx,&user)
	if err!=nil{
		log.Error("SignUp user error: %v",err)
		return "",common.GRPCErrorConvert(err,map[codes.Code]string{
			codes.Internal:ServiceError,
			codes.AlreadyExists:"该昵称已注册",
			codes.InvalidArgument:"昵称或密码不能为空",
		})
	}

	// 生成jwt token
	
}
