package post

import (
	"changweiba-backend/dao"
	"changweiba-backend/graphql/models"
	accountpb "changweiba-backend/rpc/account/pb"
	postpb "changweiba-backend/rpc/post/pb"
	"context"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"time"
	"changweiba-backend/graphql/user"
)

const (
	PostServiceError="post service system error"
)

type MyPostResolver struct {
	
}

func GetPost(ctx context.Context,postId int,conn *grpc.ClientConn) (*models.Post,error){
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
	u,_:=user.GetUser(ctx,int(r.Post.UserId),conn)

	return &models.Post{
		ID:int(r.Post.Id),
		User:u,
		Topic:r.Post.Topic,
		CreateAt:int(r.Post.CreateTime),
		LastAt:int(r.Post.LastUpdate),
		ReplyNum:int(r.Post.ReplyNum),
		Status:models.Status(r.Post.Status),
	}, nil
}

func GetCommentsByPostId(ctx context.Context, obj *models.Post, page int,
	pageSize int) ([]*models.Comment, error){
		dbComments,err:=dao.GetCommentsByPostId(int64(obj.ID),page,pageSize)
		if err!=nil{
			logs.Error(fmt.Sprintf("get comments by post_id failed: %+v",err))
			return nil, err
		}
		var comments []*models.Comment
		for _,v:=range dbComments{
			u,_:=user.GetUser(ctx,int(v.UserId),conn)
			comments=append(comments,&models.Comment{
				ID:int(v.Id),
				User:v.UserId
			})
		}
		return comments,nil
}