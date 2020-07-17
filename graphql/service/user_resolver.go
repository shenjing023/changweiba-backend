package service

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	"changweiba-backend/graphql/models"
	"changweiba-backend/pkg/logs"
	"changweiba-backend/pkg/middleware"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
)

const (
	//系统错误
	AccountServiceError = "user system error"
)

// SignUp 用户登录
func SignUp(ctx context.Context, input models.NewUser) (string, error) {
	//获取客户端ip
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		logs.Error(err)
		return "", errors.New(AccountServiceError)
	}
	ip := gc.ClientIP()

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
		common.LogDaoError("get random avatar error:", err)
		return "", errors.New(AccountServiceError)
	}
	id, err := dao.InsertUser(input.Name, password, ip, avatar)
	if err != nil {
		common.LogDaoError(fmt.Sprintf("insert user[%s] error: ", input.Name), err)
		return "", common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Unknown:         AccountServiceError,
			common.AlreadyExists:   "改昵称已注册",
			common.Internal:        AccountServiceError,
			common.InvalidArgument: "昵称或密码不能为空",
		})
	}
	//生成jwt
	jwt := middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token, err := jwt.GenerateToken(id)
	if err != nil {
		logs.Error("generate jwt token error: ", err)
		return "", errors.New(AccountServiceError)
	}
	return token.AccessToken, nil
}

// SignIn 用户登录
func SignIn(ctx context.Context, input models.NewUser) (string, error) {
	dbUser, exist := dao.CheckUserExist(input.Name)
	if exist {
		dbPassword := dbUser.Password
		tmp, _ := encryptPassword(input.Password)
		if dbPassword != tmp {
			return "", errors.New("密码错误")
		}
	} else {
		return "", errors.New("用户名错误")
	}
	//生成jwt
	jwt := middleware.NewJWT(
		middleware.SetSigningKey(conf.Cfg.SignKey),
	)
	token, err := jwt.GenerateToken(dbUser.Id)
	if err != nil {
		logs.Error("generate jwt token error: ", err)
		return "", errors.New(AccountServiceError)
	}
	return token.AccessToken, nil
}

// GetUser 获取单个用户信息
func GetUser(ctx context.Context, userID int) (*models.User, error) {
	dbUser, err := dao.GetUser(int64(userID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get user[%d] error: ", userID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.NotFound: "该用户不存在",
			common.Internal: AccountServiceError,
		})
	}

	status_, role_ := models.UserStatusNormal, models.UserRoleNormal
	if dbUser.Status == 1 {
		status_ = models.UserStatusBanned
	}
	if dbUser.Role == 1 {
		role_ = models.UserRoleAdmin
	}
	return &models.User{
		ID:           int(dbUser.Id),
		Name:         dbUser.Name,
		Avatar:       dbUser.Avatar,
		Status:       status_,
		Score:        int(dbUser.Score),
		BannedReason: dbUser.BannedReason,
		Role:         role_,
	}, nil
}

// GetUsers 获取多个用户信息
func GetUsers(ctx context.Context, ids []int64) ([]*models.User, error) {
	dbUsers, err := dao.GetUsers(ids)
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get users[%v] error: ", ids), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: AccountServiceError,
		})
	}
	var users []*models.User
	for _, v := range dbUsers {
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
	b, err := dao.CheckUserIP(ip)
	if err != nil {
		common.LogDaoError("check ip error:", err)
		return false, AccountServiceError
	}
	if !b {
		return false, "该ip已注册多个用户"
	}
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
