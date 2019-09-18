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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		pbUser:=&pb.User{
			Id:dbUser.Id,
			Name:dbUser.Name,
			Password:dbUser.Password,
			Avatar:dbUser.Avatar,
			Status:pb.Status(dbUser.Status),
			Score:dbUser.Score,
			BannedReason:dbUser.BannedReason,
			CreateTime:dbUser.CreateTime,
			LastUpdate:dbUser.LastUpdate,
			Role:pb.Role(dbUser.Role),
		}
		return pbUser,nil
	} else{
		logs.Info("get user failed:",UserNotFund)
		return &pb.User{},nil
	}
}

func (u *User) RegisterUser(ctx context.Context, ur *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	msg, _ := u.checkNewUser(ur)
	if len(msg)!=0{
		return nil,status.Error(codes.InvalidArgument,msg)
	}
	
	password,err:=u.encryptPassword(ur.Password)
	if err!=nil{
		logs.Error("generate crypto password error:",err.Error())
		return nil, status.Error(codes.Internal,ServiceError)
	}
	//头像url
	avatar,err:=dao.GetRandomAvatar()
	if err!=nil{
		logs.Error("get random avatar error:",err.Error())
		return nil, status.Error(codes.Internal,ServiceError)
	}
	id,err:=dao.InsertUser(ur.Name,password,ur.Ip,avatar)
	if err!=nil{
		logs.Error("insert user error:",err.Error())
		return nil, err
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
		return nil,status.Error(codes.Internal,ServiceError)
	}
	if has{
		dbPassword:=dbUser.Password
		tmp,_:=u.encryptPassword(param.Password)
		if dbPassword==tmp{
			return &pb.LoginResponse{
				Id:dbUser.Id,
			}, nil
		} else{
			return nil,status.Error(codes.InvalidArgument,"密码错误")
		}
	} else{
		return nil,status.Error(codes.InvalidArgument,"用户名错误")
	}
}

func (u *User) checkNewUser(ur *pb.NewUserRequest) (string, error) {
	if len(strings.TrimSpace(ur.Name)) == 0 || len(strings.TrimSpace(ur.Password))==0{
		return "user name or password can not be empty",nil
	}
	//检查该ip下的账号
	fmt.Println(ur.Ip)
	return "", nil
}

func (u *User) GetUsersByIds(ctx context.Context,ur *pb.UsersRequest) (*pb.UsersResponse,error){
	dbUsers,err:=dao.GetUsersByIds(ur.Ids,int(ur.IdType))
	if err!=nil{
		logs.Error("get users by ids error:",err.Error())
		return nil,status.Error(codes.Internal,ServiceError)
	}
	var users []*pb.User
	for _,v:=range dbUsers{
		users=append(users,&pb.User{
			Id:v.Id,
			Name:v.Name,
			Avatar:v.Avatar,
			Status:pb.Status(v.Status),
			Score:v.Score,
			BannedReason:v.BannedReason,
			CreateTime:v.CreateTime,
			LastUpdate:v.LastUpdate,
			Role:pb.Role(v.Role),
		})
	}
	return &pb.UsersResponse{
		Users:users,
	}, nil
}

func (u *User) GetUsersByUserIds(ctx context.Context, ur *pb.UsersByUserIdsRequest) (*pb.UsersByUserIdsResponse,error){
	dbUsers,err:=dao.GetUsers(ur.Ids)
	if err!=nil{
		logs.Error("get users by ids error:",err.Error())
		return nil,status.Error(codes.Internal,ServiceError)
	}
	var users []*pb.User
	for _,v:=range dbUsers{
		users=append(users,&pb.User{
			Id:v.Id,
			Name:v.Name,
			Avatar:v.Avatar,
			Status:pb.Status(v.Status),
			Score:v.Score,
			BannedReason:v.BannedReason,
			CreateTime:v.CreateTime,
			LastUpdate:v.LastUpdate,
			Role:pb.Role(v.Role),
		})
	}
	return &pb.UsersByUserIdsResponse{
		Users:users,
	}, nil
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
