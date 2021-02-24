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
