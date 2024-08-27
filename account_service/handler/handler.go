package handler

import (
	"context"
	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"cw_account_service/repository"
	user "cw_account_service/repository/ent/user"
	"encoding/base64"
	"log/slog"
	"strings"

	"github.com/cockroachdb/errors"
	"google.golang.org/grpc/codes"

	er "github.com/shenjing023/vivy-polaris/errors"
	"golang.org/x/crypto/scrypt"
)

const (
	ServiceError = "account service internal error"
)

// User user struct
type User struct {
	pb.UnimplementedAccountServiceServer
}

// SignUp 注册
func (u *User) SignUp(ctx context.Context, sr *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if err := checkNewUser(ctx, sr); err != nil {
		slog.Error("check new_user", "error", err)
		return nil, err
	}

	password, err := encryptPassword(sr.Password)
	if err != nil {
		slog.Error("generate crypto password", "error", err)
		return nil, er.NewInternalError()
	}
	//头像url
	avatar, err := repository.GetRandomAvatar(ctx)
	if err != nil {
		slog.Error("get random avatar", "error", err)
		return nil, er.NewInternalError()
	}
	id, err := repository.InsertUser(ctx, sr.Name, password, avatar)
	if err != nil {
		slog.Error("insert user", "error", err)
		return nil, er.NewInternalError()
	}
	resp := &pb.SignUpResponse{
		Id:   id,
		Role: pb.UserRoleEnum_NORMAL,
	}
	return resp, nil
}

// SignIn 登录
func (u *User) SignIn(ctx context.Context, sr *pb.SignInRequest) (*pb.SignInResponse, error) {
	dbUser, err := repository.GetUserByName(ctx, sr.Name)
	if err != nil {
		slog.Error("get user by name", "error", err)
		return nil, er.NewInternalError()
	}
	if dbUser == nil {
		return nil, er.NewServiceErr(codes.InvalidArgument,
			errors.New("user not exist"))
	}
	dbPassword := dbUser.Password
	tmp, _ := encryptPassword(sr.Password)
	if dbPassword != tmp {
		return nil, er.NewServiceErr(codes.InvalidArgument,
			errors.New("password incorrect"))
	}
	return &pb.SignInResponse{
		Id:   int64(dbUser.ID),
		Role: convertUserRole(dbUser.Role),
	}, nil
}

// GetUser 获取user信息
func (u *User) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	dbUser, err := repository.GetUserByID(ctx, user.Id)
	if err != nil {
		slog.Error("get user by id", "error", err)
		return nil, er.NewInternalError()
	}
	if dbUser == nil {
		return nil, er.NewServiceErr(codes.InvalidArgument,
			errors.New("user not exist"))
	}
	return &pb.User{
		Id:     int64(dbUser.ID),
		Name:   dbUser.Name,
		Avatar: dbUser.Avatar,
		Status: convertUserStatus(dbUser.Status),
		Score:  int64(dbUser.Score),
		Role:   convertUserRole(dbUser.Role),
	}, nil
}

func convertUserStatus(status user.Status) pb.UserStatusEnum_Status {
	switch status {
	case user.StatusNORMAL:
		return pb.UserStatusEnum_NORMAL
	case user.StatusBANNED:
		return pb.UserStatusEnum_BANNED
	default:
		return pb.UserStatusEnum_NORMAL
	}
}

func convertUserRole(role user.Role) pb.UserRoleEnum_Role {
	switch role {
	case user.RoleUSER:
		return pb.UserRoleEnum_NORMAL
	case user.RoleADMIN:
		return pb.UserRoleEnum_ADMIN
	default:
		return pb.UserRoleEnum_NORMAL
	}
}

// GetUsersByUserIds 通过用户id批量获取用户信息
func (u *User) GetUsersByUserIds(ctx context.Context, ur *pb.UsersByUserIdsRequest) (*pb.UsersByUserIdsResponse, error) {
	dbUsers, err := repository.GetUsers(ctx, ur.Ids)
	if err != nil {
		slog.Error("get users by ids", "error", err)
		return nil, er.NewInternalError()
	}
	var users []*pb.User
	for _, v := range dbUsers {
		users = append(users, &pb.User{
			Id:     int64(v.ID),
			Name:   v.Name,
			Avatar: v.Avatar,
			Status: convertUserStatus(v.Status),
			Score:  int64(v.Score),
			Role:   convertUserRole(v.Role),
		})
	}
	return &pb.UsersByUserIdsResponse{
		Users: users,
	}, nil
}

func checkNewUser(ctx context.Context, sr *pb.SignUpRequest) error {
	if len(strings.TrimSpace(sr.Name)) == 0 || len(strings.TrimSpace(sr.Password)) == 0 {
		return er.NewServiceErr(codes.InvalidArgument,
			errors.New("user name or password can not be empty"))
	}
	if exist, err := repository.CheckUserExistByName(ctx, sr.Name); err != nil {
		return err
	} else if exist {
		return er.NewServiceErr(codes.AlreadyExists,
			errors.Newf("user_name[%s] already exist", sr.Name))
	}
	return nil
}

// encryptPassword 密码加盐加密
func encryptPassword(password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(conf.Cfg.Salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", errors.Wrap(err, "scrypt error")
	}
	return base64.StdEncoding.EncodeToString(dk)[:32], nil
}
