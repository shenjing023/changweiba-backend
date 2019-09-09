package post

import (
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/user"
	postpb "changweiba-backend/rpc/post/pb"
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"time"
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
	pageSize int,conn *grpc.ClientConn) ([]*models.Comment, error){
	client:=postpb.NewPostServiceClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	commentsRequest:=postpb.CommentsRequest{
		PostId:int64(obj.ID),
		Offset:int32(page),
		Limit:int32(pageSize),
	}
	r,err:=client.GetCommentsByPostId(ctx,&commentsRequest)
	if err!=nil{
		logs.Error("get comments error:",err.Error())
		return nil, err
	}
	var comments []*models.Comment
	for _,v:=range r.Comments{
		u,_:=user.GetUser(ctx,int(v.UserId),conn)
		comments=append(comments,&models.Comment{
			ID:int(v.Id),
			User:u,
			PostID:int(v.PostId),
			Content:v.Content,
			CreateAt:int(v.CreateTime),
			Floor:int(v.Floor),
			Status:models.Status(v.Status),
		})
	}
	return comments,nil
}

func GetRepliesByCommentId(ctx context.Context,obj *models.Comment,page int,pageSize int,
	conn *grpc.ClientConn) ([]*models.Reply,error){
	client:=postpb.NewPostServiceClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	repliesRequest:=postpb.RepliesRequest{
		CommentId:int64(obj.ID),
		Offset:int32(page),
		Limit:int32(pageSize),
	}
	r,err:=client.GetRepliesByCommentId(ctx,&repliesRequest)
	if err!=nil{
		logs.Error("get replies error:",err.Error())
		return nil, err
	}
	var replies []*models.Reply
	for _,v:=range r.Replies{
		u,_:=user.GetUser(ctx,int(v.UserId),conn)
		replies=append(replies,&models.Reply{
			ID:int(v.Id),
			User:u,
			PostID:int(v.PostId),
			CommentID:int(v.CommentId),
			Content:v.Content,
			CreateAt:int(v.CreateTime),
			Floor:int(v.Floor),
			Status:models.Status(v.Status),
		})
	}
	return replies,nil
}