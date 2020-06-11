package service

import (
	"changweiba-backend/common"
	"changweiba-backend/dao"
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/rpc_conn"
	postpb "changweiba-backend/rpc/post/pb"
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/micro/go-micro"
	"github.com/pkg/errors"
	"time"
)

const (
	ServiceError = "post service system error"
)

type MyPostResolver struct {
}

func GetPost(ctx context.Context, postId int) (*models.Post, error) {
	//service := micro.NewService(micro.Name("post.client"))
	//service.Init()
	//client := postpb.NewPostService("post", service.Client())
	//client.GetPost()

	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//postRequest := postpb.PostRequest{
	//	Id: int64(postId),
	//}
	//r, err := client.GetPost(ctx, &postRequest)

	dbPost, err := dao.GetPost(int64(postId))
	if err != nil {
		logs.Error("get post error:", err)
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

func GetCommentsByPostId(ctx context.Context, postId int, page int,
	pageSize int) (*models.CommentConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//commentsRequest := postpb.CommentsRequest{
	//	PostId: int64(postId),
	//	Offset: int32(page),
	//	Limit:  int32(pageSize),
	//}
	//r, err := client.GetCommentsByPostId(ctx, &commentsRequest)

	dbComments, totalCount, err := dao.GetCommentsByPostId(int64(postId), int64(page), int64(pageSize))
	if err != nil {
		logs.Error("get comments error:", err)
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

func GetRepliesByCommentId(ctx context.Context, commentId int, page int, pageSize int) (*models.ReplyConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//repliesRequest := postpb.RepliesRequest{
	//	CommentId: int64(commentId),
	//	Offset:    int32(page),
	//	Limit:     int32(pageSize),
	//}
	//r, err := client.GetRepliesByCommentId(ctx, &repliesRequest)

	dbReplies, totalCount, err := dao.GetRepliesByCommentId(int64(commentId), int64(page), int64(pageSize))
	if err != nil {
		logs.Error("get replies error:", err.Error())
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

func GetPosts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//postsRequest := postpb.PostsRequest{
	//	Offset: int32(page),
	//	Limit:  int32(pageSize),
	//}
	//r, err := client.Posts(ctx, &postsRequest)

	dbPosts, totalCount, err := dao.GetPosts(int64(page), int64(pageSize))
	if err != nil {
		logs.Error("get posts error:", err)
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

func GetPostsByUserId(ctx context.Context, userId int, page int, pageSize int) (*models.PostConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//request := postpb.PostsByUserIdRequest{
	//	Offset: int32(page),
	//	Limit:  int32(pageSize),
	//	UserId: int64(userId),
	//}
	//r, err := client.GetPostsByUserId(ctx, &request)

	dbPosts, totalCount, err := dao.GetPostsByUserId(int64(userId), int64(page), int64(pageSize))
	if err != nil {
		logs.Error(fmt.Sprintf("get posts by user_id[%d] error: ", userId), err)
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

func GetCommentsByUserId(ctx context.Context, userId int, page int,
	pageSize int) (*models.CommentConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//request := postpb.CommentsByUserIdRequest{
	//	UserId: int64(userId),
	//	Offset: int32(page),
	//	Limit:  int32(pageSize),
	//}
	//r, err := client.GetCommentsByUserId(ctx, &request)

	dbComments, totalCount, err := dao.GetCommentsByUserId(int64(userId), int64(page), int64(pageSize))
	if err != nil {
		logs.Error(fmt.Sprintf("get comments by user_id[%d] error: ", userId), err)
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

func GetRepliesByUserId(ctx context.Context, userId int, page int, pageSize int) (*models.ReplyConnection, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//request := postpb.RepliesByUserIdRequest{
	//	UserId: int64(userId),
	//	Offset: int32(page),
	//	Limit:  int32(pageSize),
	//}
	//r, err := client.GetRepliesByUserId(ctx, &request)

	dbReplies, totalCount, err := dao.GetRepliesByUserId(int64(userId), int64(page), int64(pageSize))
	if err != nil {
		logs.Error(fmt.Sprintf("get replies by user_id[%d] error: ", userId), err)
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

func GetCommentById(ctx context.Context, commentId int) (*models.Comment, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//request := postpb.CommentRequest{
	//	Id: int64(commentId),
	//}
	//r, err := client.GetComment(ctx, &request)

	dbComment, err := dao.GetComment(int64(commentId))
	if err != nil {
		logs.Error(fmt.Sprintf("get comment[%d] error: ", commentId), err)
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

func GetReplyById(ctx context.Context, replyId int) (*models.Reply, error) {
	//client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()
	//request := postpb.ReplyRequest{
	//	Id: int64(replyId),
	//}
	//r, err := client.GetReply(ctx, &request)

	dbReply, err := dao.GetReply(int64(replyId))
	if err != nil {
		logs.Error(fmt.Sprintf("get reply[%d] error: ", replyId), err)
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

func DeletePost(ctx context.Context, postId int) (bool, error) {
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.DeleteRequest{
		Id: int64(postId),
	}
	r, err := client.DeletePost(ctx, &request)
	if err != nil {
		logs.Error("delete post error:", err.Error())
		return false, err
	}
	return r.Success, nil
}

func DeleteComment(ctx context.Context, commentId int) (bool, error) {
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.DeleteRequest{
		Id: int64(commentId),
	}
	r, err := client.DeleteComment(ctx, &request)
	if err != nil {
		logs.Error("delete comment error:", err.Error())
		return false, err
	}
	return r.Success, nil
}

func DeleteReply(ctx context.Context, replyId int) (bool, error) {
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.DeleteRequest{
		Id: int64(replyId),
	}
	r, err := client.DeleteReply(ctx, &request)
	if err != nil {
		logs.Error("delete reply error:", err.Error())
		return false, err
	}
	return r.Success, nil
}

//
func NewPost(ctx context.Context, post models.NewPost) (int, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return 0, err
	}
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.NewPostRequest{
		UserId:  userId,
		Topic:   post.Topic,
		Content: post.Content,
	}
	r, err := client.NewPost(ctx, &request)
	if err != nil {
		logs.Error("create post error:", err.Error())
		return 0, err
	}
	return int(r.PostId), nil
}

func NewComment(ctx context.Context, comment models.NewComment) (int, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return 0, err
	}
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.NewCommentRequest{
		UserId:  userId,
		PostId:  int64(comment.PostID),
		Content: comment.Content,
	}
	r, err := client.NewComment(ctx, &request)
	if err != nil {
		logs.Error("create comment error:", err.Error())
		return 0, err
	}
	return int(r.CommentId), nil
}

func NewReply(ctx context.Context, reply models.NewReply) (int, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return 0, err
	}
	client := postpb.NewPostServiceClient(rpc_conn.PostConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	request := postpb.NewReplyRequest{
		UserId:    userId,
		PostId:    int64(reply.PostID),
		Content:   reply.Content,
		CommentId: int64(reply.CommentID),
		ParentId:  int64(reply.ParentID),
	}
	r, err := client.NewReply(ctx, &request)
	if err != nil {
		logs.Error("create reply error:", err.Error())
		return 0, err
	}
	return int(r.ReplyId), nil
}

func getUserIdFromContext(ctx context.Context) (int64, error) {
	gctx, err := common.GinContextFromContext(ctx)
	if err != nil {
		logs.Error(err.Error())
		return 0, errors.New(ServiceError)
	}
	userId, ok := gctx.Value("claims").(float64)
	if !ok {
		logs.Error("get user_id from request ctx error")
		logs.Info(fmt.Sprintf("ctx claims: %+v", gctx.Value("claims")))
		return 0, errors.New(ServiceError)
	}
	return int64(userId), nil
}
