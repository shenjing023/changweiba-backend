package service

import (
	"changweiba-backend/common"
	"changweiba-backend/dao"
	"changweiba-backend/graphql/models"
	"changweiba-backend/pkg/logs"
	"context"
	"fmt"

	"github.com/pkg/errors"
)

const (
	//系统错误
	ServiceError = "post system error"
)

// GetPost 获取帖子信息
func GetPost(ctx context.Context, postID int) (*models.Post, error) {
	dbPost, err := dao.GetPost(int64(postID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get post[%d] error: ", postID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.NotFound: "该帖子不存在",
			common.Internal: ServiceError,
		})
	}

	return &models.Post{
		ID: int(dbPost.Id),
		User: &models.User{
			ID: int(dbPost.UserId),
		},
		Topic:    dbPost.Topic,
		CreateAt: int(dbPost.CreateTime),
		LastAt:   int(dbPost.LastUpdate),
		ReplyNum: int(dbPost.ReplyNum),
		Status:   models.Status(dbPost.Status),
	}, nil
}

// GetCommentsByPostID 获取该帖子下的评论
func GetCommentsByPostID(ctx context.Context, postID int, page int,
	pageSize int) (*models.CommentConnection, error) {
	dbComments, totalCount, err := dao.GetCommentsByPostId(int64(postID), int64(page), int64(pageSize))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get comments under post[%d] error: ", postID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该帖子不存在",
		})
	}
	var comments []*models.Comment
	for _, v := range dbComments {
		comments = append(comments, &models.Comment{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			PostID:   int(v.PostId),
			Content:  v.Content,
			CreateAt: int(v.CreateTime),
			Floor:    int(v.Floor),
			Status:   models.Status(v.Status),
		})
	}
	return &models.CommentConnection{
		Nodes:      comments,
		TotalCount: int(totalCount),
	}, nil
}

// GetRepliesByCommentID 获取该评论下的回复
func GetRepliesByCommentID(ctx context.Context, commentID int, page int, pageSize int) (*models.ReplyConnection, error) {
	dbReplies, totalCount, err := dao.GetRepliesByCommentId(int64(commentID), int64(page), int64(pageSize))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get replies under comment[%d] error: ", commentID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该评论不存在",
		})
	}
	var replies []*models.Reply
	for _, v := range dbReplies {
		replies = append(replies, &models.Reply{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			PostID:    int(v.PostId),
			CommentID: int(v.CommentId),
			Content:   v.Content,
			CreateAt:  int(v.CreateTime),
			Floor:     int(v.Floor),
			Status:    models.Status(v.Status),
			Parent: &models.Reply{
				ID: int(v.ParentId),
			},
		})
	}
	return &models.ReplyConnection{
		Nodes:      replies,
		TotalCount: int(totalCount),
	}, nil
}

// GetPosts 获取帖子信息
func GetPosts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	dbPosts, totalCount, err := dao.GetPosts(int64(page), int64(pageSize))
	if err != nil {
		logs.Error("get posts error: ", err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
		})
	}
	var posts []*models.Post
	for _, v := range dbPosts {
		posts = append(posts, &models.Post{
			ID:       int(v.Id),
			Topic:    v.Topic,
			CreateAt: int(v.CreateTime),
			LastAt:   int(v.LastUpdate),
			ReplyNum: int(v.ReplyNum),
			Status:   models.Status(v.Status),
			User: &models.User{
				ID: int(v.UserId),
			},
		})
	}
	return &models.PostConnection{
		Nodes:      posts,
		TotalCount: int(totalCount),
	}, nil
}

// GetPostsByUserId 获取该用户的帖子
func GetPostsByUserId(ctx context.Context, userID int, page int, pageSize int) (*models.PostConnection, error) {
	dbPosts, totalCount, err := dao.GetPostsByUserId(int64(userID), int64(page), int64(pageSize))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get posts by user_id[%d] error: ", userID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该用户不存在",
		})
	}
	var posts []*models.Post
	for _, v := range dbPosts {
		posts = append(posts, &models.Post{
			ID:       int(v.Id),
			Topic:    v.Topic,
			CreateAt: int(v.CreateTime),
			LastAt:   int(v.LastUpdate),
			ReplyNum: int(v.ReplyNum),
			Status:   models.Status(v.Status),
			User: &models.User{
				ID: int(v.UserId),
			},
		})
	}
	return &models.PostConnection{
		Nodes:      posts,
		TotalCount: int(totalCount),
	}, nil
}

// GetCommentsByUserId 获取该用户的评论
func GetCommentsByUserId(ctx context.Context, userID int, page int,
	pageSize int) (*models.CommentConnection, error) {
	dbComments, totalCount, err := dao.GetCommentsByUserId(int64(userID), int64(page), int64(pageSize))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get comments by user_id[%d] error: ", userID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该用户不存在",
		})
	}
	var comments []*models.Comment
	for _, v := range dbComments {
		comments = append(comments, &models.Comment{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			PostID:   int(v.PostId),
			Content:  v.Content,
			CreateAt: int(v.CreateTime),
			Floor:    int(v.Floor),
			Status:   models.Status(v.Status),
		})
	}
	return &models.CommentConnection{
		Nodes:      comments,
		TotalCount: int(totalCount),
	}, nil
}

// GetRepliesByUserID 获取该用户的回复
func GetRepliesByUserID(ctx context.Context, userID int, page int, pageSize int) (*models.ReplyConnection, error) {
	dbReplies, totalCount, err := dao.GetRepliesByUserId(int64(userID), int64(page), int64(pageSize))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get replies by user_id[%d] error: ", userID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.NotFound: "该用户不存在",
			common.Internal: ServiceError,
		})
	}
	var replies []*models.Reply
	for _, v := range dbReplies {
		replies = append(replies, &models.Reply{
			ID: int(v.Id),
			User: &models.User{
				ID: int(v.UserId),
			},
			PostID:    int(v.PostId),
			CommentID: int(v.CommentId),
			Content:   v.Content,
			CreateAt:  int(v.CreateTime),
			Floor:     int(v.Floor),
			Status:    models.Status(v.Status),
			Parent: &models.Reply{
				ID: int(v.ParentId),
			},
		})
	}
	return &models.ReplyConnection{
		Nodes:      replies,
		TotalCount: int(totalCount),
	}, nil
}

// GetCommentByID 根据id获取评论信息
func GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	dbComment, err := dao.GetComment(int64(commentID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get comment[%d] error: ", commentID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该评论不存在",
		})
	}

	status_ := models.StatusNormal
	if dbComment.Status == 1 {
		status_ = models.StatusBanned
	}
	return &models.Comment{
		ID: int(dbComment.Id),
		User: &models.User{
			ID: int(dbComment.UserId),
		},
		PostID:   int(dbComment.PostId),
		Content:  dbComment.Content,
		CreateAt: int(dbComment.CreateTime),
		Floor:    int(dbComment.Floor),
		Status:   status_,
	}, nil
}

// GetReplyByID 根据id获取回复
func GetReplyByID(ctx context.Context, replyID int) (*models.Reply, error) {
	dbReply, err := dao.GetReply(int64(replyID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("get reply[%d] error: ", replyID), err)
		return nil, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该回复不存在",
		})
	}

	status_ := models.StatusNormal
	if dbReply.Status == 1 {
		status_ = models.StatusBanned
	}
	return &models.Reply{
		ID: int(dbReply.Id),
		User: &models.User{
			ID: int(dbReply.UserId),
		},
		PostID:    int(dbReply.PostId),
		CommentID: int(dbReply.CommentId),
		Content:   dbReply.Content,
		CreateAt:  int(dbReply.CreateTime),
		Floor:     int(dbReply.Floor),
		Status:    status_,
	}, nil
}

// DeletePost 删除帖子
func DeletePost(ctx context.Context, postID int) (bool, error) {
	err := dao.DeletePost(int64(postID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("delete post[%d] error: ", postID), err)
		return false, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该帖子不存在",
		})
	}
	return true, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, commentID int) (bool, error) {
	err := dao.DeleteComment(int64(commentID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("delete comment[%d] error: ", commentID), err)
		return false, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该评论不存在",
		})
	}
	return true, nil
}

// DeleteReply 删除回复
func DeleteReply(ctx context.Context, replyID int) (bool, error) {
	err := dao.DeleteReply(int64(replyID))
	if err != nil {
		common.LogDaoError(fmt.Sprintf("delete reply[%d] error: ", replyID), err)
		return false, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
			common.NotFound: "该回复不存在",
		})
	}
	return true, nil
}

// NewPost 新建帖子
func NewPost(ctx context.Context, post models.NewPost) (int, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return 0, errors.New(ServiceError)
	}

	postID, err := dao.InsertPost(userID, post.Topic, post.Content)
	if err != nil {
		logs.Error("create post error: ", err)
		return 0, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
		})
	}
	return int(postID), nil
}

// NewComment 新建评论
func NewComment(ctx context.Context, comment models.NewComment) (int, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return 0, errors.New(ServiceError)
	}

	commentID, err := dao.InsertComment(userID, int64(comment.PostID), comment.Content)
	if err != nil {
		common.LogDaoError("create comment error: ", err)
		return 0, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal:        ServiceError,
			common.InvalidArgument: err.Error(),
		})
	}
	return int(commentID), nil
}

// NewReply 新建回复
func NewReply(ctx context.Context, reply models.NewReply) (int, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return 0, errors.New(ServiceError)
	}

	replyID, err := dao.InsertReply(userID, int64(reply.PostID), int64(reply.CommentID), int64(reply.ParentID), reply.Content)
	if err != nil {
		logs.Error("create reply error:", err)
		return 0, common.ServiceErrorConvert(err, map[common.ErrorCode]string{
			common.Internal: ServiceError,
		})
	}
	return int(replyID), nil
}

func getUserIDFromContext(ctx context.Context) (int64, error) {
	gctx, err := common.GinContextFromContext(ctx)
	if err != nil {
		logs.Error("get gin_context from context error: ", err)
		return 0, err
	}
	userID, ok := gctx.Value("claims").(float64)
	if !ok {
		logs.Error("get user_id from request ctx error")
		return 0, errors.New(ServiceError)
	}
	return int64(userID), nil
}
