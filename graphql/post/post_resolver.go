package post

import (
	"changweiba-backend/graphql/models"
	accountpb "changweiba-backend/rpc/account/pb"
	postpb "changweiba-backend/rpc/post/pb"
	"context"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"time"
)

const (
	PostServiceError="post service system error"
)

type MyPostResolver struct {
	
}

func (p *MyPostResolver) GetPost(ctx context.Context,postId int,conn *grpc.ClientConn) (*models.Post,error){
	client:=postpb.NewPostServiceClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	postRequest:=postpb.PostRequest{
		Id:int64(postId),
	}
	r,err:=client.GetPost(ctx,&postRequest)
	if err!=nil{
		logs.Error("get post error:",err.Error())
		return nil, err
	}
	if r.Post.Id==0{
		//不存在
		return nil,errors.New("post不存在")
	}
	pbUser:=accountpb.User{
		Id:r.Post.UserId,
	}
	client2:=accountpb.NewAccountClient(conn)
	r2,err:=client2.GetUser(ctx,&pbUser)
	if err!=nil{
		logs.Error(fmt.Sprintf("get user[%d] error:%+v",r.Post.UserId,err))
		return nil, err
	}
	user:=&models.User{
		ID:int(r2.Id),
		Name:r2.Name,
		Avatar:r2.Avatar,
		Status:models.UserStatus(r2.Status),
		Score:int(r2.Score),
		BannedReason:r2.BannedReason,
		Role:models.UserRole(r2.Role),
	}
	return &models.Post{
		ID:int(r.Post.Id),
		User:user,
		Topic:r.Post.Topic,
		CreateAt:int(r.Post.CreateTime),
		LastAt:int(r.Post.LastUpdate),
		ReplyNum:int(r.Post.ReplyNum),
		Status:models.Status(r.Post.Status),
	}, nil
}

func (p *MyPostResolver) GetCommentsByPostId(ctx context.Context, obj *models.Post, page int, 
	pageSize int) ([]*models.Comment, error){
		
}