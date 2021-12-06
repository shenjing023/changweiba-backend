package handler

import (
	"context"
	"gateway/middleware"
	"gateway/models"
	"time"

	"gateway/common"

	pb "gateway/pb"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
)

const (
	// ServiceError service error
	ServiceError = "gateway service internal error"
)

// SignUp 用户注册
func SignUp(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	//获取客户端ip
	// gc, err := common.GinContextFromContext(ctx)
	// if err != nil {
	// 	log.Error("%+v", err)
	// 	return nil, errors.New(ServiceError)
	// }
	// ip, _, err := net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	// if err != nil {
	// 	log.Error("get remote ip error: ", err)
	// 	return nil, errors.New(ServiceError)
	// }

	client := pb.NewAccountClient(AccountConn)
	log.Infof("account target:%s", AccountConn.Target())
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	user := pb.SignUpRequest{
		Name:     input.Name,
		Password: input.Password,
	}
	resp, err := client.SignUp(ctx, &user)
	if err != nil {
		log.Error("SignUp user error: ", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.AlreadyExists:   "该昵称已注册",
			codes.InvalidArgument: "昵称或密码不能为空",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Error("generate access_token error: ", err)
		return nil, common.NewGQLError(common.Internal, ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Error("generate refresh_token error: ", err)
		return nil, common.NewGQLError(common.Internal, ServiceError)
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
		log.Error("SignIn user error: ", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        ServiceError,
			codes.NotFound:        "昵称不正确",
			codes.InvalidArgument: "密码不正确",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Error("generate access_token error: ", err)
		return nil, common.NewGQLError(common.Internal, ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Error("generate refresh_token error: ", err)
		return nil, common.NewGQLError(common.Internal, ServiceError)
	}

	return &models.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// UsersByIDsLoaderFunc 批量获取用户的dataloader func
func UsersByIDsLoaderFunc(ctx context.Context, keys []int64) (users []*models.User, errs []error) {
	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ur := pb.UsersByUserIdsRequest{
		Ids: keys,
	}
	resp, err := client.GetUsersByUserIds(ctx, &ur)
	if err != nil {
		log.Error("get user by ids error: ", err)
		for i := 0; i < len(keys); i++ {
			errs = append(errs, err)
		}
		return
	}

	for _, v := range resp.Users {
		users = append(users, &models.User{
			ID:           int(v.Id),
			Name:         v.Name,
			Avatar:       v.Avatar,
			Status:       models.UserStatus(v.Status.String()),
			Role:         models.UserRole(v.Role.String()),
			Score:        int(v.Score),
			BannedReason: v.BannedReason,
		})
	}
	return
}
