package post

import (
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/rpc_conn"
	postpb "changweiba-backend/rpc/post/pb"
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
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
	pageSize int) (*models.CommentConnection, error){
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
	return &models.CommentConnection{
		Nodes:comments,
		TotalCount:0,
	},nil
}

func GetRepliesByCommentId(ctx context.Context,obj *models.Comment,page int,pageSize int) (*models.ReplyConnection,error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
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
	return &models.ReplyConnection{
		Nodes:replies,
		TotalCount:0,
	},nil
}

func GetPosts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
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
	return &models.PostConnection{
		Nodes:posts,
		TotalCount:0,
	},nil
}

//
func NewPost(ctx context.Context,post models.NewPost) (int,error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	request:=postpb.NewPostRequest{
		UserId:int64(post.UserID),
		Topic:post.Topic,
		Content:post.Content,
	}
	r,err:=client.NewPost(ctx,&request)
	if err!=nil{
		logs.Error("create post error:",err.Error())
		return 0, err
	}
	return int(r.PostId),nil
}

func NewComment(ctx context.Context,comment models.NewComment) (int,error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	request:=postpb.NewCommentRequest{
		UserId:int64(comment.UserID),
		PostId:int64(comment.PostID),
		Content:comment.Content,
	}
	r,err:=client.NewComment(ctx,&request)
	if err!=nil{
		logs.Error("create comment error:",err.Error())
		return 0, err
	}
	return int(r.CommentId),nil
}

func NewReply(ctx context.Context,reply models.NewReply) (int,error){
	client:=postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx,cancel:=context.WithTimeout(ctx,10*time.Second)
	defer cancel()
	request:=postpb.NewReplyRequest{
		UserId:int64(reply.UserID),
		PostId:int64(reply.PostID),
		Content:reply.Content,
		CommentId:int64(reply.CommentID),
		ParentId:int64(reply.ParentID),
	}
	r,err:=client.NewReply(ctx,&request)
	if err!=nil{
		logs.Error("create reply error:",err.Error())
		return 0, err
	}
	return int(r.ReplyId),nil
}