package handler

import (
	"context"
	pb "cw_post_service/pb"
	"cw_post_service/repository"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/reply"
	"strings"

	"github.com/cockroachdb/errors"
	log "github.com/shenjing023/llog"
	er "github.com/shenjing023/vivy-polaris/errors"
	"google.golang.org/grpc/codes"
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
	if len(strings.TrimSpace(pr.Title)) == 0 || len(strings.TrimSpace(pr.Content)) == 0 {
		return nil, er.NewServiceErr(codes.InvalidArgument,
			errors.New("user name or password can not be empty"))
	}
	postID, err := repository.InsertPost(ctx, pr.UserId, pr.Title, pr.Content)
	if err != nil {
		log.Errorf("insert post error: %+v", err)
		return nil, err
	}
	// // 插入一楼
	// if _, err := repository.InsertComment(ctx, pr.UserId, postID, pr.Content); err != nil {
	// 	log.Errorf("insert post[%d] first comment error: %+v", postID, err)
	// 	go repository.DeletePost(ctx, postID)
	// 	return nil, err
	// }
	return &pb.NewPostResponse{
		PostId: postID,
	}, nil
}

// GetPost get post info
func (PostService) GetPost(ctx context.Context, pr *pb.PostRequest) (*pb.PostResponse, error) {
	dbPost, err := repository.GetPostByID(ctx, pr.Id)
	if err != nil {
		log.Errorf("get post error: %+v", err)
		return nil, err
	}
	if dbPost == nil {
		return nil, er.NewServiceErr(codes.NotFound,
			errors.New("post not found"))
	}
	return &pb.PostResponse{
		Post: &pb.Post{
			Id:         int64(dbPost.ID),
			UserId:     int64(dbPost.AuthorID),
			Title:      dbPost.Title,
			CreateTime: dbPost.CreatedAt.UnixMilli(),
			UpdateTime: dbPost.UpdatedAt.UnixMilli(),
			ReplyNum:   int64(dbPost.ReplyNum),
			Status:     convertPostStatus(dbPost.Status),
			Content:    dbPost.Content,
		},
	}, nil
}

func convertPostStatus(status post.Status) pb.PostStatusEnum_Status {
	switch status {
	case post.StatusNORMAL:
		return pb.PostStatusEnum_NORMAL
	case post.StatusBANNED:
		return pb.PostStatusEnum_BANNED
	case post.StatusDELETED:
		return pb.PostStatusEnum_DELETE
	default:
		return pb.PostStatusEnum_NORMAL
	}
}

func convertCommentStatus(status comment.Status) pb.CommentStatusEnum_Status {
	switch status {
	case comment.StatusNORMAL:
		return pb.CommentStatusEnum_NORMAL
	case comment.StatusBANNED:
		return pb.CommentStatusEnum_BANNED
	case comment.StatusDELETED:
		return pb.CommentStatusEnum_DELETE
	default:
		return pb.CommentStatusEnum_NORMAL
	}
}

func convertReplyStatus(status reply.Status) pb.ReplyStatusEnum_Status {
	switch status {
	case reply.StatusNORMAL:
		return pb.ReplyStatusEnum_NORMAL
	case reply.StatusBANNED:
		return pb.ReplyStatusEnum_BANNED
	case reply.StatusDELETED:
		return pb.ReplyStatusEnum_DELETE
	default:
		return pb.ReplyStatusEnum_NORMAL
	}
}

// GetPosts get posts info by page and page_size
func (PostService) GetAllPosts(ctx context.Context, pr *pb.AllPostsRequest) (*pb.PostsResponse, error) {
	dbPosts, err := repository.GetPosts(ctx, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get all posts error: %+v", err)
		return nil, er.NewInternalError()
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         int64(v.ID),
			UserId:     int64(v.AuthorID),
			Title:      v.Title,
			CreateTime: v.CreatedAt.UnixMilli(),
			UpdateTime: v.UpdatedAt.UnixMilli(),
			ReplyNum:   int64(v.ReplyNum),
			Status:     convertPostStatus(v.Status),
			Content:    v.Content,
		})
	}
	totalCount, err := repository.GetPostsTotalCount(ctx)
	if err != nil {
		log.Errorf("get posts total count error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.PostsResponse{
		Posts:      posts,
		TotalCount: totalCount,
	}, nil
}

func (PostService) NewComment(ctx context.Context, pr *pb.NewCommentRequest) (*pb.NewCommentResponse, error) {
	commentID, err := repository.InsertComment(ctx, pr.UserId,
		pr.PostId, pr.Content)
	if err != nil {
		log.Errorf("insert comment error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.NewCommentResponse{
		CommentId: commentID,
	}, nil
}

func (PostService) NewReply(ctx context.Context, pr *pb.NewReplyRequest) (*pb.NewReplyResponse, error) {
	replyID, err := repository.InsertReply(ctx, pr.UserId, pr.PostId,
		pr.CommentId, pr.ParentId, pr.Content)
	if err != nil {
		log.Errorf("insert reply error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.NewReplyResponse{
		ReplyId: replyID,
	}, nil
}

func (PostService) GetPostFirstComment(ctx context.Context, pr *pb.FirstCommentRequest) (*pb.FirstCommentResponse, error) {
	dbComments, err := repository.GetPostFirstComment(ctx, pr.PostIds)
	if err != nil {
		log.Errorf("get first comment error: %+v", err)
		return nil, er.NewInternalError()
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		if v.Status != comment.StatusDELETED {
			// 被删了
			comments = append(comments, &pb.Comment{})
		} else {
			comments = append(comments, &pb.Comment{
				Id:      int64(v.ID),
				Content: v.Content,
				Status:  convertCommentStatus(v.Status),
			})
		}
	}
	return &pb.FirstCommentResponse{
		Comments: comments,
	}, nil
}

func (PostService) GetCommentsByPostId(ctx context.Context, pr *pb.CommentsRequest) (*pb.CommentsResponse, error) {
	dbComments, err := repository.GetCommentsByPostID(ctx, pr.PostId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get post comments error: %+v", err)
		return nil, er.NewInternalError()
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		comments = append(comments, &pb.Comment{
			Id:         int64(v.ID),
			Content:    v.Content,
			Status:     convertCommentStatus(v.Status),
			UserId:     int64(v.AuthorID),
			CreateTime: v.CreatedAt.UnixMilli(),
			Floor:      int64(v.Floor),
			PostId:     pr.PostId,
		})
	}
	totalCount, err := repository.GetPostCommentTotalCount(ctx, pr.PostId)
	if err != nil {
		log.Errorf("get post comments total count error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.CommentsResponse{
		TotalCount: totalCount,
		Comments:   comments,
	}, nil
}

func (PostService) GetRepliesByCommentId(ctx context.Context, pr *pb.RepliesRequest) (*pb.RepliesResponse, error) {
	dbReplies, err := repository.GetRepliesByCommentID(ctx, pr.CommentId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get comment replies error: %+v", err)
		return nil, er.NewInternalError()
	}
	var replies []*pb.Reply
	for _, v := range dbReplies {
		replies = append(replies, &pb.Reply{
			Id:         int64(v.ID),
			Content:    v.Content,
			Status:     convertReplyStatus(v.Status),
			CreateTime: v.CreatedAt.UnixMilli(),
			ParentId:   int64(v.ParentID),
			UserId:     int64(v.AuthorID),
			CommentId:  pr.CommentId,
		})
	}
	totalCount, err := repository.GetCommentReplyTotalCount(ctx, pr.CommentId)
	if err != nil {
		log.Errorf("get comment replies total count error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.RepliesResponse{
		TotalCount: totalCount,
		Replies:    replies,
	}, nil
}

func (PostService) GetPostsByUserId(ctx context.Context, pr *pb.PostsByUserIdRequest) (*pb.PostsByUserIdResponse, error) {
	dbPosts, err := repository.GetPostsByUserId(ctx, pr.UserId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get user posts error: %+v", err)
		return nil, er.NewInternalError()
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         int64(v.ID),
			UserId:     int64(v.AuthorID),
			Title:      v.Title,
			CreateTime: v.CreatedAt.UnixMilli(),
			UpdateTime: v.UpdatedAt.UnixMilli(),
			ReplyNum:   int64(v.ReplyNum),
			Status:     convertPostStatus(v.Status),
			Content:    v.Content,
			Pin:        int64(v.Pin),
		})
	}
	totalCount, err := repository.GetUserPostCount(ctx, pr.UserId)
	if err != nil {
		log.Errorf("get posts total count error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.PostsByUserIdResponse{
		Posts:      posts,
		TotalCount: totalCount,
	}, nil
}

func (PostService) DeletePosts(ctx context.Context, pr *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	for _, v := range pr.Ids {
		err := repository.DeletePost(ctx, v)
		if err != nil {
			log.Errorf("delete post error: %+v", err)
			return nil, er.NewInternalError()
		}
	}
	return &pb.DeleteResponse{
		Success: true,
	}, nil
}

func (PostService) PinPost(ctx context.Context, pr *pb.PinPostRequest) (*pb.PinPostResponse, error) {
	err := repository.PinPost(ctx, pr.PostId, int(pr.PinStatus))
	if err != nil {
		log.Errorf("pin post error: %+v", err)
		return nil, er.NewInternalError()
	}
	return &pb.PinPostResponse{
		Success: true,
	}, nil
}

func (PostService) GetPinPosts(ctx context.Context, pr *pb.PinPostsRequest) (*pb.PinPostsResponse, error) {
	dbPosts, err := repository.GetPinPostsByUserId(ctx, pr.UserId)
	if err != nil {
		log.Errorf("get user pin posts error: %+v", err)
		return nil, er.NewInternalError()
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         int64(v.ID),
			UserId:     int64(v.AuthorID),
			Title:      v.Title,
			CreateTime: v.CreatedAt.UnixMilli(),
			UpdateTime: v.UpdatedAt.UnixMilli(),
			ReplyNum:   int64(v.ReplyNum),
			Status:     convertPostStatus(v.Status),
			Content:    v.Content,
			Pin:        int64(v.Pin),
		})
	}

	return &pb.PinPostsResponse{
		Posts:      posts,
		TotalCount: int64(len(posts)),
	}, nil
}
