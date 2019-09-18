package post

import (
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/rpc_conn"
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

func GetPost(ctx context.Context,postId int) (*models.Post,error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
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

	return &models.Post{
		ID:int(r.Post.Id),
		User:&models.User{
			ID:int(r.Post.UserId),
		},
		Topic:r.Post.Topic,
		CreateAt:int(r.Post.CreateTime),
		LastAt:int(r.Post.LastUpdate),
		ReplyNum:int(r.Post.ReplyNum),
		Status:models.Status(r.Post.Status),
	}, nil
}

func GetCommentsByPostId(ctx context.Context, obj *models.Post, page int,
	pageSize int) ([]*models.Comment, error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
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
		comments=append(comments,&models.Comment{
			ID:int(v.Id),
			User:&models.User{
				ID:int(v.UserId),
			},
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
		replies=append(replies,&models.Reply{
			ID:int(v.Id),
			User:&models.User{
				ID:int(v.UserId),
			},
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

func GetPosts(ctx context.Context, page int, pageSize int) ([]*models.Post, error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	postsRequest:=postpb.PostsRequest{
		Offset:int32(page),
		Limit:int32(pageSize),
	}
	r,err:=client.Posts(ctx,&postsRequest)
	if err!=nil{
		logs.Error("get posts error:",err.Error())
		return nil, err
	}
	var posts []*models.Post
	for _,v:=range r.Posts{
		posts=append(posts,&models.Post{
			ID:int(v.Id),
			Topic:v.Topic,
			CreateAt:int(v.CreateTime),
			LastAt:int(v.LastUpdate),
			ReplyNum:int(v.ReplyNum),
			Status:models.Status(v.Status),
			User:&models.User{
				ID:int(v.UserId),
			},
		})
	}
	return posts,nil
}