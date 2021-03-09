package handler

import (
	"context"
	"gateway/models"
	pb "gateway/pb"
	"time"

	"gateway/common"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
)

// NewPost 新建post
func NewPost(ctx context.Context, input models.NewPost) (int, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Error("new post get userID from context error: ", err)
		return 0, err
	}
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	request := pb.NewPostRequest{
		UserId:  userID,
		Topic:   input.Topic,
		Content: input.Content,
	}
	r, err := client.NewPost(ctx, &request)
	if err != nil {
		log.Error("new post error: ", err)
		return 0, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: ServiceError,
		})
	}
	return int(r.PostId), nil
}

// Posts 帖子列表
func Posts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.PostsRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
	}
	r, err := client.GetPosts(ctx, &request)
	if err != nil {
		log.Error("get posts error: ", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: ServiceError,
		})
	}
	var posts []*models.Post
	for _, v := range r.Posts {
		posts = append(posts, &models.Post{
			ID:        int(v.Id),
			Topic:     v.Topic,
			CreatedAt: int(v.CreateTime),
			UpdatedAt: int(v.LastUpdate),
			ReplyNum:  int(v.ReplyNum),
			Status:    models.PostStatus(v.Status),
			User: &models.User{
				ID: int(v.UserId),
			},
		})
	}
	log.Info("posts:", posts)
	return &models.PostConnection{
		Nodes:      posts,
		TotalCount: int(r.TotalCount),
	}, nil
}
