package handler

import (
	"context"
	"cw_account_service/common"
	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"cw_account_service/repository"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	log "github.com/shenjing023/llog"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/codes"
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
	id, err := repository.InsertUser(sr.Name, password, sr.Ip, avatar)
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
	dbUser, err := repository.GetUserByName(sr.Name)
	if err != nil {
		return nil, ServiceErr2GRPCErr(err)
	}
	dbPassword := dbUser.Password
	tmp, _ := encryptPassword(sr.Password)
	if dbPassword != tmp {
		return nil, common.NewServiceErr(common.InvalidArgument,
			errors.New("password incorrect"))
	}
	return &pb.SignInResponse{
		Id: dbUser.ID,
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
	//检查该ip下的账号
	fmt.Println(sr.Ip)
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
