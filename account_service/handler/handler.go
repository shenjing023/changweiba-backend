package handler

import (
	"context"
	"cw_account_service/common"
	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"cw_account_service/repository"
	"encoding/base64"
	"errors"
	"strings"

	log "github.com/shenjing023/llog"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ServiceError = "account service internal error"
)

// User user struct
type User struct {
	pb.UnimplementedAccountServer
}

// ServiceErr2GRPCErr serviceErr covert to GRPCErr
func ServiceErr2GRPCErr(err error) error {
	if e, ok := err.(*common.ServiceErr); ok {
		if e.Code == common.Internal {
			log.Errorf("Service Internal Error: %v", e.Err)
		}
		if _, ok := common.ErrMap[e.Code]; ok {
			return status.Error(common.ErrMap[e.Code], e.Err.Error())
		}
		return status.Error(codes.Unknown, e.Err.Error())
	}
	return status.Error(codes.Unknown, err.Error())
}

// SignUp 注册
func (u *User) SignUp(ctx context.Context, sr *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if err := checkNewUser(sr); err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}

	password, err := encryptPassword(sr.Password)
	if err != nil {
		log.Error("generate crypto password error:", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	//头像url
	avatar, err := repository.GetRandomAvatar()
	if err != nil {
		log.Error("get random avatar error:", err.Error())
		return nil, ServiceErr2GRPCErr(err)
	}
	id, err := repository.InsertUser(sr.Name, password, avatar)
	if err != nil {
		log.Error("insert user error:", err.Error())
		return nil, ServiceErr2GRPCErr(err)
	}
	resp := &pb.SignUpResponse{
		Id: id,
	}
	return resp, nil
}

// SignIn 登录
func (u *User) SignIn(ctx context.Context, sr *pb.SignInRequest) (*pb.SignInResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Info("md: ", md)
	}
	dbUser, err := repository.GetUserByName(sr.Name)
	if err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	dbPassword := dbUser.Password
	tmp, _ := encryptPassword(sr.Password)
	if dbPassword != tmp {
		return nil, ServiceErr2GRPCErr(common.NewServiceErr(common.InvalidArgument,
			errors.New("password incorrect")))
	}
	return &pb.SignInResponse{
		Id: dbUser.ID,
	}, nil
}

// GetUser 获取user信息
func (u *User) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	dbUser, err := repository.GetUserByID(user.Id)
	if err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	return &pb.User{
		Id:           dbUser.ID,
		Name:         dbUser.Name,
		Avatar:       dbUser.Avatar,
		Status:       pb.UserStatusEnum_Status(dbUser.Status),
		Score:        dbUser.Score,
		BannedReason: dbUser.BannedReason,
		Role:         pb.UserRoleEnum_Role(dbUser.Role),
	}, nil
}

// GetUsersByUserIds 通过用户id批量获取用户信息
func (u *User) GetUsersByUserIds(ctx context.Context, ur *pb.UsersByUserIdsRequest) (*pb.UsersByUserIdsResponse, error) {
	dbUsers, err := repository.GetUsers(ur.Ids)
	if err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	var users []*pb.User
	for _, v := range dbUsers {
		status, role := pb.UserStatusEnum_Status(v.Status), pb.UserRoleEnum_Role(v.Role)
		if v.Status == 1 {
			status = pb.UserStatusEnum_BANNED
		}
		if v.Role == 1 {
			role = pb.UserRoleEnum_ADMIN
		}
		users = append(users, &pb.User{
			Id:           v.ID,
			Name:         v.Name,
			Avatar:       v.Avatar,
			Status:       status,
			Score:        v.Score,
			BannedReason: v.BannedReason,
			Role:         role,
		})
	}
	return &pb.UsersByUserIdsResponse{
		Users: users,
	}, nil
}

func checkNewUser(sr *pb.SignUpRequest) error {
	if len(strings.TrimSpace(sr.Name)) == 0 || len(strings.TrimSpace(sr.Password)) == 0 {
		return common.NewServiceErr(common.InvalidArgument,
			errors.New("user name or password can not be empty"))
	}
	if exist, err := repository.CheckUserExistByName(sr.Name); err != nil {
		return err
	} else if exist {
		return common.NewServiceErr(common.AlreadyExists,
			errors.New("user name already exist"))
	}
	return nil
}

// encryptPassword 密码加盐加密
func encryptPassword(password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(conf.Cfg.Salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk)[:32], nil
}
