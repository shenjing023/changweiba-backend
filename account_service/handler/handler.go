package handler

import (
	"context"
	"cw_account_service/conf"
	pb "cw_account_service/pb"
	"cw_account_service/repository"
	"encoding/base64"
	"strings"

	"github.com/cockroachdb/errors"

	log "github.com/shenjing023/llog"
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
		log.Errorf("check new_user error: %+v", err)
		return nil, err
	}

	password, err := encryptPassword(sr.Password)
	if err != nil {
		log.Errorf("generate crypto password error: %+v", err)
		return nil, er.NewServiceErr(er.Internal, err)
	}
	//头像url
	avatar, err := repository.GetRandomAvatar(ctx)
	if err != nil {
		log.Errorf("get random avatar error: %+v", err)
		return nil, err
	}
	id, err := repository.InsertUser(ctx, sr.Name, password, avatar)
	if err != nil {
		log.Errorf("insert user error: %+v", err)
		return nil, err
	}
	resp := &pb.SignUpResponse{
		Id: id,
	}
	return resp, nil
}

// SignIn 登录
func (u *User) SignIn(ctx context.Context, sr *pb.SignInRequest) (*pb.SignInResponse, error) {
	dbUser, err := repository.GetUserByName(ctx, sr.Name)
	if err != nil {
		log.Errorf("get user by name error: %+v", err)
		return nil, err
	}
	dbPassword := dbUser.Password
	tmp, _ := encryptPassword(sr.Password)
	if dbPassword != tmp {
		return nil, er.NewServiceErr(er.InvalidArgument,
			errors.New("password incorrect"))
	}
	return &pb.SignInResponse{
		Id: int64(dbUser.ID),
	}, nil
}

// GetUser 获取user信息
func (u *User) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	dbUser, err := repository.GetUserByID(ctx, user.Id)
	if err != nil {
		log.Errorf("get user by id error: %+v", err)
		return nil, err
	}
	ban := ""
	if dbUser.Status != 0 {
		ban, _ = repository.GetBannedReason(ctx, int64(dbUser.Status))
	}
	return &pb.User{
		Id:           int64(dbUser.ID),
		Name:         dbUser.NickName,
		Avatar:       dbUser.Avatar,
		Status:       pb.UserStatusEnum_Status(dbUser.Status),
		Score:        dbUser.Score,
		BannedReason: ban,
		Role:         pb.UserRoleEnum_Role(dbUser.Role),
	}, nil
}

// GetUsersByUserIds 通过用户id批量获取用户信息
func (u *User) GetUsersByUserIds(ctx context.Context, ur *pb.UsersByUserIdsRequest) (*pb.UsersByUserIdsResponse, error) {
	dbUsers, err := repository.GetUsers(ctx, ur.Ids)
	if err != nil {
		log.Errorf("get users by ids error: %+v", err)
		return nil, err
	}
	var users []*pb.User
	for _, v := range dbUsers {
		role := pb.UserRoleEnum_Role(v.Role)
		if v.Role == 1 {
			role = pb.UserRoleEnum_ADMIN
		}
		ban := ""
		if v.Status != 0 {
			ban, _ = repository.GetBannedReason(ctx, int64(v.Status))
		}
		users = append(users, &pb.User{
			Id:           int64(v.ID),
			Name:         v.NickName,
			Avatar:       v.Avatar,
			Status:       pb.UserStatusEnum_Status(v.Status),
			Score:        v.Score,
			BannedReason: ban,
			Role:         role,
		})
	}
	return &pb.UsersByUserIdsResponse{
		Users: users,
	}, nil
}

func checkNewUser(ctx context.Context, sr *pb.SignUpRequest) error {
	if len(strings.TrimSpace(sr.Name)) == 0 || len(strings.TrimSpace(sr.Password)) == 0 {
		return er.NewServiceErr(er.InvalidArgument,
			errors.New("user name or password can not be empty"))
	}
	if exist, err := repository.CheckUserExistByName(ctx, sr.Name); err != nil {
		return err
	} else if exist {
		return er.NewServiceErr(er.AlreadyExists,
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
