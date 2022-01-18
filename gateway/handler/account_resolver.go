package handler

import (
	"context"
	"errors"
	"gateway/middleware"
	"gateway/models"
	"time"

	"gateway/common"

	pb "gateway/pb"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
)

// SignUp 用户注册
func SignUp(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	user := pb.SignUpRequest{
		Name:     input.Name,
		Password: input.Password,
	}
	resp, err := client.SignUp(ctx, &user)
	if err != nil {
		log.Errorf("signUp user error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        common.ServiceError,
			codes.AlreadyExists:   "该昵称已注册",
			codes.InvalidArgument: "昵称或密码不能为空",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Errorf("generate access_token error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Errorf("generate refresh_token error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
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
		log.Errorf("SignIn user error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal:        common.ServiceError,
			codes.NotFound:        "昵称不正确",
			codes.InvalidArgument: "密码不正确",
		})
	}

	// 生成jwt token
	accessToken, err := middleware.GenerateAccessToken(resp.Id)
	if err != nil {
		log.Errorf("generate access_token error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
	}
	refreshToken, err := middleware.GenerateRefreshToken(resp.Id)
	if err != nil {
		log.Errorf("generate refresh_token error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
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
		log.Errorf("get user by ids error: %+v", err)
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

func GetAccessToken(ctx context.Context, input string) (string, error) {
	// 先验证refresh_token
	token, err := middleware.RefreshRefreshToken(input)
	if err != nil {
		if errors.Is(err, common.ErrTokenExpired) {
			return "", common.NewGQLError(common.InvalidArgument, "授权已过期")
		} else if errors.Is(err, common.ErrTokenNotValidYet) {
			return "", common.NewGQLError(common.InvalidArgument, "授权未生效")
		} else if errors.Is(err, common.ErrTokenMalformed) {
			return "", common.NewGQLError(common.InvalidArgument, "token无效")
		} else if errors.Is(err, common.ErrTokenInvalid) {
			return "", common.NewGQLError(common.InvalidArgument, "token无效")
		}
		log.Errorf("refresh token error: %+v", err)
		return "", common.NewGQLError(common.Internal, common.ServiceError)
	}
	return token, nil
}
