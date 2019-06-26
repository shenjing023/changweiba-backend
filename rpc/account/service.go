//go:generate protoc --go_out=plugins=grpc:./pb account.proto

package account

import (
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, ur *pb.UserRequest) (*pb.User, error) {
	return &pb.User{Id: ur.Id}, nil
}

func (u *User) RegisterUser(ctx context.Context, ur *pb.NewUserRequest) (*pb.User, error) {
	if valid, err := u.checkNewUser(ur); !valid {
		return nil, err
	}

	return &pb.User{Id: "1334"}, nil
}

func (u *User) EditUser(ctx context.Context, pbUser *pb.User) (*pb.User, error) {
	return &pb.User{Id: pbUser.Id}, nil
}

func (u *User) checkNewUser(ur *pb.NewUserRequest) (bool, error) {
	if len(strings.TrimSpace(ur.User.Name)) == 0 {
		return false, errors.New("user name can not empty")
	}
	//检查该ip下的账号
	fmt.Println(ur.Ip)
	return true, nil
}

func NewAccountService(addr string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountServer(grpcServer, &User{})
	grpcServer.Serve(lis)
}