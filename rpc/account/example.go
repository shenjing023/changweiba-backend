package account

import (
	pb "changweiba-backend/rpc/account/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Register() {
	conn, err := grpc.Dial("localhost:9112", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAccountClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u := pb.NewUserRequest{
		Name:     "asda",
		Password: "asdad",
		Ip: "122",
	}
	c, err := client.RegisterUser(ctx, &u)
	if err != nil {
		log.Fatalf("%v.GetUser(_) = _, %v: ", client, err)
	}
	fmt.Println(c.Id)
}
