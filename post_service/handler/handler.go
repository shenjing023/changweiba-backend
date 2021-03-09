package handler

import (
	"context"
	pb "cw_post_service/pb"
	"cw_post_service/repository"
	"strings"

	log "github.com/shenjing023/llog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// ServiceError Internal Error
	ServiceError = "post service internal error"
)

// PostService post_service struct
type PostService struct {
	pb.UnimplementedPostServiceServer
}

// NewPost new post
func (PostService) NewPost(ctx context.Context, pr *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	if len(strings.TrimSpace(pr.Topic)) == 0 || len(strings.TrimSpace(pr.Content)) == 0 {
		return nil, status.Error(codes.InvalidArgument, "topic or content can not be empty")
	}
	postID, err := repository.InsertPost(pr.UserId, pr.Topic, pr.Content)
	if err != nil {
		log.Error("insert post error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.NewPostResponse{
		PostId: postID,
	}, nil
}

// GetPost get post info
func (PostService) GetPost(ctx context.Context, pr *pb.PostRequest) (*pb.PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}

// GetPosts get posts info by page and page_size
func (PostService) GetPosts(ctx context.Context, pr *pb.PostsRequest) (*pb.PostsResponse, error) {
	dbPosts, err := repository.GetPosts(int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Error("get posts failed: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         v.ID,
			UserId:     v.UserID,
			Topic:      v.Topic,
			CreateTime: v.CreateTime,
			LastUpdate: v.LastUpdate,
			ReplyNum:   v.ReplyNum,
			Status:     pb.PostStatusEnum_Status(v.Status),
		})
	}
	totalCount, err := repository.GetPostsTotalCount()
	if err != nil {
		log.Error("get posts total count failed: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.PostsResponse{
		Posts:      posts,
		TotalCount: totalCount,
	}, nil
}
