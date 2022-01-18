package handler

import (
	"context"
	pb "cw_post_service/pb"
	"cw_post_service/repository"
	"strings"

	"github.com/cockroachdb/errors"
	log "github.com/shenjing023/llog"
	er "github.com/shenjing023/vivy-polaris/errors"
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
		return nil, er.NewServiceErr(er.InvalidArgument, errors.New("topic or content can not be empty"))
	}
	postID, err := repository.InsertPost(ctx, pr.UserId, pr.Topic)
	if err != nil {
		log.Errorf("insert post error: %+v", err)
		return nil, err
	}
	// 插入一楼
	if _, err := repository.InsertComment(ctx, pr.UserId, postID, pr.Content); err != nil {
		log.Errorf("insert post[%d] first comment error: %+v", postID, err)
		go repository.DeletePost(ctx, postID)
		return nil, err
	}
	return &pb.NewPostReply{
		PostId: postID,
	}, nil
}

// GetPost get post info
func (PostService) GetPost(ctx context.Context, pr *pb.PostRequest) (*pb.PostReply, error) {
	return nil, er.NewServiceErr(er.Unimplemented, errors.New("method GetPost not implemented"))
}

// GetPosts get posts info by page and page_size
func (PostService) GetPosts(ctx context.Context, pr *pb.PostsRequest) (*pb.PostsReply, error) {
	dbPosts, err := repository.GetPosts(ctx, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get posts error: %+v", err)
		return nil, err
	}

	var posts []*pb.Post
	for _, v := range dbPosts {
		posts = append(posts, &pb.Post{
			Id:         int64(v.ID),
			UserId:     int64(v.UserID),
			Topic:      v.Topic,
			CreateTime: v.CreateAt,
			LastUpdate: v.UpdateAt,
			ReplyNum:   v.ReplyNum,
			Status:     pb.PostStatusEnum_Status(v.Status),
		})
	}
	totalCount, err := repository.GetPostsTotalCount(ctx)
	if err != nil {
		log.Errorf("get posts total count error: %+v", err)
		return nil, err
	}
	return &pb.PostsReply{
		Posts:      posts,
		TotalCount: totalCount,
	}, nil
}

func (PostService) NewComment(ctx context.Context, pr *pb.NewCommentRequest) (*pb.NewCommentReply, error) {
	// TODO 解析content 文本 图片 视频
	commentID, err := repository.InsertComment(ctx, pr.UserId, pr.PostId, pr.Content)
	if err != nil {
		log.Errorf("insert comment error: %+v", err)
		return nil, err
	}
	return &pb.NewCommentReply{
		CommentId: commentID,
	}, nil
}

func (PostService) NewReply(ctx context.Context, pr *pb.NewReplyRequest) (*pb.NewReplyReply, error) {
	replyID, err := repository.InsertReply(ctx, pr.UserId, pr.PostId,
		pr.CommentId, pr.ParentId, pr.Content)
	if err != nil {
		log.Errorf("insert reply error: %+v", err)
		return nil, err
	}
	return &pb.NewReplyReply{
		ReplyId: replyID,
	}, nil
}

func (PostService) GetPostFirstComment(ctx context.Context, pr *pb.FirstCommentRequest) (*pb.FirstCommentReply, error) {
	dbComments, err := repository.GetPostFirstComment(ctx, pr.PostIds)
	if err != nil {
		log.Errorf("get first comment error: %+v", err)
		return nil, err
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		if v.Status != 0 {
			// 被删了
			comments = append(comments, &pb.Comment{})
		} else {
			comments = append(comments, &pb.Comment{
				Id:      int64(v.ID),
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
	dbComments, err := repository.GetCommentsByPostID(ctx, pr.PostId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get post comments error: %+v", err)
		return nil, err
	}
	var comments []*pb.Comment
	for _, v := range dbComments {
		comments = append(comments, &pb.Comment{
			Id:      int64(v.ID),
			Content: v.Content,
			Status:  pb.PostStatusEnum_Status(v.Status),
		})
	}
	totalCount, err := repository.GetPostCommentTotalCount(ctx, pr.PostId)
	if err != nil {
		log.Errorf("get post comments total count error: %+v", err)
		return nil, err
	}
	return &pb.CommentsReply{
		TotalCount: totalCount,
		Comments:   comments,
	}, nil
}

func (PostService) GetRepliesByCommentId(ctx context.Context, pr *pb.RepliesRequest) (*pb.RepliesReply, error) {
	dbReplies, err := repository.GetRepliesByCommentID(ctx, pr.CommentId, int(pr.Page), int(pr.PageSize))
	if err != nil {
		log.Errorf("get comment replies error: %+v", err)
		return nil, err
	}
	var replies []*pb.Reply
	for _, v := range dbReplies {
		replies = append(replies, &pb.Reply{
			Id:         int64(v.ID),
			Content:    v.Content,
			Status:     pb.PostStatusEnum_Status(v.Status),
			CreateTime: v.CreateAt,
			ParentId:   int64(v.ParentID),
			Floor:      int64(v.Floor),
			UserId:     int64(v.UserID),
			CommentId:  pr.CommentId,
		})
	}
	totalCount, err := repository.GetCommentReplyTotalCount(ctx, pr.CommentId)
	if err != nil {
		log.Errorf("get comment replies total count error: %+v", err)
		return nil, err
	}
	return &pb.RepliesReply{
		TotalCount: totalCount,
		Replies:    replies,
	}, nil
}
