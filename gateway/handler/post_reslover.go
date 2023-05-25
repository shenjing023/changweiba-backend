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
		log.Errorf("new post get userID from context error: %+v", err)
		return 0, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request := pb.NewPostRequest{
		UserId:  userID,
		Title:   input.Title,
		Content: input.Content,
	}
	r, err := client.NewPost(ctx, &request)
	if err != nil {
		log.Errorf("new post error: %+v", err)
		return 0, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	return int(r.PostId), nil
}

// Posts 帖子列表
func AllPosts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.AllPostsRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
	}
	r, err := client.GetAllPosts(ctx, &request)
	if err != nil {
		log.Errorf("get posts error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var posts []*models.Post
	for _, v := range r.Posts {
		posts = append(posts, &models.Post{
			ID:        int(v.Id),
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: int(v.CreateTime),
			UpdatedAt: int(v.UpdateTime),
			ReplyNum:  int(v.ReplyNum),
			Status:    models.PostStatus(v.Status),
			User: &models.User{
				ID: int(v.UserId),
			},
			FirstComment: &models.Comment{
				PostID: int(v.Id),
			},
		})
	}

	return &models.PostConnection{
		Nodes:      posts,
		TotalCount: int(r.TotalCount),
	}, nil
}

func Posts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("posts get userID from context error: %+v", err)
		return nil, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	request := pb.PostsByUserIdRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
		UserId:   userID,
	}
	r, err := client.GetPostsByUserId(ctx, &request)
	if err != nil {
		log.Errorf("get posts error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var posts []*models.Post
	for _, v := range r.Posts {
		posts = append(posts, &models.Post{
			ID:        int(v.Id),
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: int(v.CreateTime),
			UpdatedAt: int(v.UpdateTime),
			ReplyNum:  int(v.ReplyNum),
			Status:    models.PostStatus(v.Status),
			User: &models.User{
				ID: int(v.UserId),
			},
			FirstComment: &models.Comment{
				PostID: int(v.Id),
			},
		})
	}

	return &models.PostConnection{
		Nodes:      posts,
		TotalCount: int(r.TotalCount),
	}, nil
}

// FirstCommentLoaderFunc 批量获取帖子第一个评论的dataloader func
func FirstCommentLoaderFunc(ctx context.Context, keys []int64) (comments []*models.Comment, errs []error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	request := pb.FirstCommentRequest{
		PostIds: keys,
	}
	r, err := client.GetPostFirstComment(ctx, &request)
	if err != nil {
		log.Errorf("get first comment error: %+v", err)
		for i := 0; i < len(keys); i++ {
			errs = append(errs, err)
		}
		return
	}
	for _, v := range r.Comments {
		if v.Id == 0 && v.Content == "" {
			comments = append(comments, nil)
		} else {
			comments = append(comments, &models.Comment{
				ID:      int(v.Id),
				Content: v.Content,
			})
		}
	}
	return
}

func NewComment(ctx context.Context, input models.NewComment) (int, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("new comment get userID from context error: %+v", err)
		return 0, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request := pb.NewCommentRequest{
		UserId:  userID,
		PostId:  int64(input.PostID),
		Content: input.Content,
	}
	r, err := client.NewComment(ctx, &request)
	if err != nil {
		log.Errorf("new comment error: %+v", err)
		return 0, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	return int(r.CommentId), nil
}

func NewReply(ctx context.Context, input models.NewReply) (int, error) {
	userID, err := common.GetUserIDFromContext(ctx)
	if err != nil {
		log.Errorf("new reply get userID from context error: %+v", err)
		return 0, common.NewGQLError(common.Internal, common.ServiceError)
	}
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request := pb.NewReplyRequest{
		UserId:    userID,
		PostId:    int64(input.PostID),
		Content:   input.Content,
		CommentId: int64(input.CommentID),
	}
	if input.ParentID == nil {
		request.ParentId = 0
	} else {
		request.ParentId = int64(*input.ParentID)
	}
	r, err := client.NewReply(ctx, &request)
	if err != nil {
		log.Errorf("new reply error: %+v", err)
		return 0, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	return int(r.ReplyId), nil
}

func GetCommentsByPostID(ctx context.Context, postID int, page int, pageSize int) (*models.CommentConnection, error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	request := pb.CommentsRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
		PostId:   int64(postID),
	}
	r, err := client.GetCommentsByPostId(ctx, &request)
	if err != nil {
		log.Errorf("get post comments error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var comments []*models.Comment
	for _, v := range r.Comments {
		comments = append(comments, &models.Comment{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			PostID:    int(v.PostId),
			Content:   v.Content,
			CreatedAt: int(v.CreateTime),
			Floor:     int(v.Floor),
		})
	}
	return &models.CommentConnection{
		Nodes:      comments,
		TotalCount: int(r.TotalCount),
	}, nil
}

func GetRepliesByCommentID(ctx context.Context, commentID int, page int, pageSize int) (*models.ReplyConnection, error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	request := pb.RepliesRequest{
		Page:      int64(page),
		PageSize:  int64(pageSize),
		CommentId: int64(commentID),
	}
	r, err := client.GetRepliesByCommentId(ctx, &request)
	if err != nil {
		log.Errorf("get comment replies error: %+v", err)
		return nil, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	var replies []*models.Reply
	for _, v := range r.Replies {
		replies = append(replies, &models.Reply{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			Content:   v.Content,
			CreatedAt: int(v.CreateTime),
			Floor:     int(v.Floor),
			CommentID: int(v.CommentId),
			Parent: &models.Reply{
				ID: int(v.ParentId),
			},
		})
	}
	return &models.ReplyConnection{
		Nodes:      replies,
		TotalCount: int(r.TotalCount),
	}, nil
}

func DeletePost(ctx context.Context, postID int) (bool, error) {
	client := pb.NewPostServiceClient(PostConn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	request := pb.DeleteRequest{
		Ids: []int64{int64(postID)},
	}
	_, err := client.DeletePosts(ctx, &request)
	if err != nil {
		log.Errorf("delete post error: %+v", err)
		return false, common.GRPCErrorConvert(err, map[codes.Code]string{
			codes.Internal: common.ServiceError,
		})
	}
	return true, nil
}
