//go:generate protoc --go_out=plugins=grpc:./pb account.proto

package account

import (
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const (
	ServiceError ="account service system error"
	UserNotFund="user can not fund"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, ur *pb.User) (*pb.User, error) {
	dbUser:=&dao.User{
		Id:ur.Id,
		Name:ur.Name,
	}
	has,err:=dao.GetUser(dbUser)
	if err!=nil{
		logs.Error("get user error:",err.Error())
		return nil,errors.New(ServiceError)
	}
	if has{
		status,role:="NORMAL","NORMAL"
		if dbUser.Status==1{
			status="BANNED"
		}
		if dbUser.Role==1{
			role="ADMIN"
		}
		pbUser:=&pb.User{
			Id:dbUser.Id,
			Name:dbUser.Name,
			Password:dbUser.Password,
			Avatar:dbUser.Avatar,
			Status:status,
			Score:dbUser.Score,
			BannedReason:dbUser.BannedReason,
			CreateTime:dbUser.CreateTime,
			LastUpdate:dbUser.LastUpdate,
			Role:role,
		}
		return pbUser,nil
	} else{
		logs.Info("get user failed:",UserNotFund)
		return &pb.User{},nil
	}
}

func (u *User) RegisterUser(ctx context.Context, ur *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	if valid, err := u.checkNewUser(ur); !valid {
		return nil, err
	}
	
	password,err:=u.encryptPassword(ur.Password)
	if err!=nil{
		logs.Error("generate crypto password error:",err.Error())
		return nil, errors.New(ServiceError)
	}
	//头像url
	avatar,err:=dao.GetRandomAvatar()
	if err!=nil{
		logs.Error("get random avatar error:",err.Error())
		return nil, errors.New(ServiceError)
	}
	id,err:=dao.InsertUser(ur.Name,password,ur.Ip,avatar)
	if err!=nil{
		logs.Error("insert user error:",err.Error())
		return nil, errors.New(ServiceError)
	}
	return &pb.NewUserResponse{
		Id:id,
	}, nil
}

func (u *User) EditUser(ctx context.Context, pbUser *pb.User) (*pb.User, error) {
	return &pb.User{Id: pbUser.Id}, nil
}

func (u *User) Login(ctx context.Context,param *pb.LoginRequest) (*pb.LoginResponse, error){
	dbUser:=&dao.User{
		Name:param.Name,
	}
	has,err:=dao.GetUser(dbUser)
	if err!=nil{
		logs.Error("get user by name error:",err.Error())
		return nil,errors.New(ServiceError)
	}
	if has{
		dbPassword:=dbUser.Password
		tmp,_:=u.encryptPassword(param.Password)
		if dbPassword==tmp{
			return &pb.LoginResponse{
				Id:dbUser.Id,
			}, nil
		} else{
			return nil,errors.New("密码错误")
		}
	} else{
		return nil,errors.New("用户名错误")
	}
}

func (u *User) checkNewUser(ur *pb.NewUserRequest) (bool, error) {
	if len(strings.TrimSpace(ur.Name)) == 0 || len(strings.TrimSpace(ur.Password))==0{
		return false, errors.New("user name or password can not be empty")
	}
	//检查该ip下的账号
	fmt.Println(ur.Ip)
	return true, nil
}

//密码加盐加密
func (u *User) encryptPassword(password string) (string,error){
	dk,err:=scrypt.Key([]byte(password),[]byte(conf.Cfg.Salt),1<<15,8,1,32)
	if err!=nil{
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk)[:32],nil
}

func NewAccountService(addr string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountServer(grpcServer, &User{})
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("new grpcserver failed: %v", err)
	}
}
