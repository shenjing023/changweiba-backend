package handler

import (
	"context"
	"errors"
	"gateway/middleware"
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
func SignUp(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	//获取客户端ip
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		log.Error("%+v", err)
		return nil, errors.New(ServiceError)
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	if err != nil {
		log.Error("get remote ip error: ", err)
		return nil, errors.New(ServiceError)
	}

	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user := pb.SignUpRequest{
		Name:     input.Name,
		Password: input.Password,
		Ip:       ip,
	}
	resp, err := client.SignUp(ctx, &user)
	if err != nil {
		log.Error("SignUp user error: %v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.AlreadyExists:   "该昵称已注册",
			codes.InvalidArgument: "昵称或密码不能为空",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Error("generate access_token error: %v", err)
		return nil, errors.New(ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Error("generate refresh_token error: %v", err)
		return nil, errors.New(ServiceError)
	}

	return &models.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// SignIn 登录
func SignIn(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user := pb.SignInRequest{
		Name:     input.Name,
		Password: input.Password,
	}
	resp, err := client.SignIn(ctx, &user)
	if err != nil {
		log.Error("SignUp user error: %v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.NotFound:        "昵称不正确",
			codes.InvalidArgument: "密码不正确",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Error("generate access_token error: %v", err)
		return nil, errors.New(ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Error("generate refresh_token error: %v", err)
		return nil, errors.New(ServiceError)
	}

	return &models.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
