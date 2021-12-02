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
func (PostService) NewPost(ctx context.Context, pr *pb.NewPostRequest) (*pb.NewPostReply, error) {
	if len(strings.TrimSpace(pr.Topic)) == 0 || len(strings.TrimSpace(pr.Content)) == 0 {
		return nil, status.Error(codes.InvalidArgument, "topic or content can not be empty")
	}
	postID, err := repository.InsertPost(pr.UserId, pr.Topic)
	if err != nil {
		log.Error("insert post error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	// 插入一楼
	if _, err := repository.InsertComment(pr.UserId, postID, pr.Content); err != nil {
		log.Errorf("insert post[%d] first comment error: %v", postID, err)
		go repository.DeletePost(postID)
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.NewPostReply{
		PostId: postID,
	}, nil
}

// GetPost get post info
func (PostService) GetPost(ctx context.Context, pr *pb.PostRequest) (*pb.PostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}

// GetPosts get posts info by page and page_size
func (PostService) GetPosts(ctx context.Context, pr *pb.PostsRequest) (*pb.PostsReply, error) {
	dbPosts, err := repository.GetPosts(int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Error("get posts error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         v.ID,
			UserId:     v.UserID,
			Topic:      v.Topic,
			CreateTime: v.CreateAt,
			LastUpdate: v.UpdateAt,
			ReplyNum:   v.ReplyNum,
			Status:     pb.PostStatusEnum_Status(v.Status),
		})
	}
	totalCount, err := repository.GetPostsTotalCount()
	if err != nil {
		log.Error("get posts total count error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.PostsReply{
		Posts:      posts,
		TotalCount: totalCount,
	}, nil
}

func (PostService) NewComment(ctx context.Context, pr *pb.NewCommentRequest) (*pb.NewCommentReply, error) {
	// TODO 解析content 文本 图片 视频
	commentID, err := repository.InsertComment(pr.UserId, pr.PostId, pr.Content)
	if err != nil {
		log.Error("insert comment error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.NewCommentReply{
		CommentId: commentID,
	}, nil
}

func (PostService) NewReply(ctx context.Context, pr *pb.NewReplyRequest) (*pb.NewReplyReply, error) {
	replyID, err := repository.InsertReply(pr.UserId, pr.PostId,
		pr.CommentId, pr.ParentId, pr.Content)
	if err != nil {
		log.Error("insert reply error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.NewReplyReply{
		ReplyId: replyID,
	}, nil
}

func (PostService) GetPostFirstComment(ctx context.Context, pr *pb.FirstCommentRequest) (*pb.FirstCommentReply, error) {
	dbComments, err := repository.GetPostFirstComment(pr.PostIds)
	if err != nil {
		log.Error("get first comment error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		if v.Status != 0 {
			// 被删了
			comments = append(comments, &pb.Comment{})
		} else {
			comments = append(comments, &pb.Comment{
				Id:      v.ID,
				Content: v.Content,
				Status:  pb.PostStatusEnum_Status(v.Status),
			})
		}
	}
	return &pb.FirstCommentReply{
		Comments: comments,
	}, nil
}

func (PostService) GetCommentsByPostId(ctx context.Context, pr *pb.CommentsRequest) (*pb.CommentsReply, error) {
	dbComments, err := repository.GetCommentsByPostID(pr.PostId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Error("get post comments error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		comments = append(comments, &pb.Comment{
			Id:      v.ID,
			Content: v.Content,
			Status:  pb.PostStatusEnum_Status(v.Status),
		})
	}
	totalCount, err := repository.GetPostCommentTotalCount(pr.PostId)
	if err != nil {
		log.Error("get post comments total count error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.CommentsReply{
		TotalCount: totalCount,
		Comments:   comments,
	}, nil
}

func (PostService) GetRepliesByCommentId(ctx context.Context, pr *pb.RepliesRequest) (*pb.RepliesReply, error) {
	dbReplies, err := repository.GetRepliesByCommentID(pr.CommentId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Error("get comment replies error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	var replies []*pb.Reply
	for _, v := range dbReplies {
		replies = append(replies, &pb.Reply{
			Id:         v.ID,
			Content:    v.Content,
			Status:     pb.PostStatusEnum_Status(v.Status),
			CreateTime: v.CreateAt,
			ParentId:   v.ParentID,
			Floor:      v.Floor,
			UserId:     v.UserID,
			CommentId:  pr.CommentId,
		})
	}
	totalCount, err := repository.GetCommentReplyTotalCount(pr.CommentId)
	if err != nil {
		log.Error("get comment replies total count error: ", err.Error())
		return nil, status.Error(codes.Internal, ServiceError)
	}
	return &pb.RepliesReply{
		TotalCount: totalCount,
		Replies:    replies,
	}, nil
}
