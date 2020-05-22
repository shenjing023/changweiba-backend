package service

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/rpc_conn"
	"changweiba-backend/pkg/middleware"
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/codes"
	"net"
	"strings"
	"time"
)

const (
	AccountServiceError = "system error"
)

type MyUserResolver struct {
}

func SignUp(ctx context.Context, input models.NewUser) (string, error) {
	//获取客户端ip
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		logs.Error("%+v", err)
		return "", errors.New(AccountServiceError)
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	if err != nil {
		logs.Error("get remote ip error:", err.Error())
		return "", errors.New(AccountServiceError)
	}
	//client := pb.NewAccountClient(rpc_conn.AccountConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	//user := pb.NewUserRequest{
	//	Name:     input.Name,
	//	Password: input.Password,
	//	Ip:       ip,
	//}
	//r, err := client.RegisterUser(ctx, &user)
	//if err != nil {
	//	logs.Error("register user error:", err.Error())
	//	return "", common.GRPCErrorConvert(err, map[codes.Code]string{
	//		codes.Internal:        AccountServiceError,
	//		codes.AlreadyExists:   "该昵称已注册",
	//		codes.InvalidArgument: "昵称或密码不能为空",
	//	})
	//}

	if b, msg := checkNewUser(input.Name, input.Password, ip); !b {
		return "", errors.New(msg)
	}
	password, err := encryptPassword(input.Password)
	if err != nil {
		logs.Error("generate crypto password error: ", err)
		return "", errors.New(AccountServiceError)
	}
	//头像url
	avatar, err := dao.GetRandomAvatar()
	if err != nil {
		logs.Error("get random avatar error:", err)
		return "", errors.New(AccountServiceError)
	}
	id, err := dao.InsertUser(input.Name, password, ip, avatar)
	if err != nil {
		logs.Error("insert user error:", err)
		return "", errors.New("error")
	}
	//生成jwt
	jwt := middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token, err := jwt.GenerateToken(id)
	if err != nil {
		logs.Error("generate jwt token error: ", err)
		return "", err
	}
	return token.AccessToken, nil
}

func LoginUser(ctx context.Context, input models.NewUser) (string, error) {
	client := pb.NewAccountClient(rpc_conn.AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	pbRequest := pb.LoginRequest{
		Name:     input.Name,
		Password: input.Password,
	}
	r, err := client.Login(ctx, &pbRequest)
	if err != nil {
		logs.Error("call Login error:", err.Error())
		return "", common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.InvalidArgument: AccountServiceError,
		})
	}
	//生成jwt
	jwt := middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token, err := jwt.GenerateToken(r.Id)
	if err != nil {
		logs.Error("generate jwt token error:", err.Error())
		return "", errors.New(AccountServiceError)
	}
	return token.AccessToken, nil
}

func GetUser(ctx context.Context, userId int) (*models.User, error) {
	client := pb.NewAccountClient(rpc_conn.AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	pbUser := pb.User{
		Id: int64(userId),
	}
	r, err := client.GetUser(ctx, &pbUser)
	if err != nil {
		logs.Error("get user error:", err.Error())
		return nil, err
	}
	if r.Id == 0 {
		//user id不存在
		return nil, errors.New("用户不存在")
	}
	status_, role_ := models.UserStatusNormal, models.UserRoleNormal
	if r.Status == 1 {
		status_ = models.UserStatusBanned
	}
	if r.Role == 1 {
		role_ = models.UserRoleAdmin
	}
	return &models.User{
		ID:           int(r.Id),
		Name:         r.Name,
		Avatar:       r.Avatar,
		Status:       status_,
		Score:        int(r.Score),
		BannedReason: r.BannedReason,
		Role:         role_,
	}, nil
}

func GetUsers(ctx context.Context, ids []int64) ([]*models.User, error) {
	client := pb.NewAccountClient(rpc_conn.AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	r, err := client.GetUsersByUserIds(ctx, &pb.UsersByUserIdsRequest{Ids: ids})
	if err != nil {
		logs.Error("get users by user_ids error:", err.Error())
		return nil, err
	}
	var users []*models.User
	for _, v := range r.Users {
		status_, role_ := models.UserStatusNormal, models.UserRoleNormal
		if v.Status == 1 {
			status_ = models.UserStatusBanned
		}
		if v.Role == 1 {
			role_ = models.UserRoleAdmin
		}
		users = append(users, &models.User{
			ID:           int(v.Id),
			Name:         v.Name,
			Avatar:       v.Avatar,
			Status:       status_,
			Score:        int(v.Score),
			BannedReason: v.BannedReason,
			Role:         role_,
		})
	}
	return users, nil
}

/*
	检查请求的ip
*/
func checkNewUser(name string, password string, ip string) (bool, string) {
	if len(strings.TrimSpace(name)) == 0 || len(strings.TrimSpace(password)) == 0 {
		return false, "user name or password can not be empty"
	}
	//检查该ip下的账号
	fmt.Println(ip)
	return true, ""
}

//密码加盐加密
func encryptPassword(password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(conf.Cfg.Salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk)[:32], nil
}
