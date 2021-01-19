package handler

import (
	"context"
	"errors"
	"gateway/models"
	"net"
	"strings"
	"time"

	"gateway/common"

	pb "gateway/pb"

	log "github.com/shenjing023/llog"
)

const (
	ServiceError = "gateway service internal error"
)

// SignUp 用户注册
func SignUp(ctx context.Context, input models.NewUser) (string, error) {
	//获取客户端ip
	gc, err := common.GinContextFromContext(ctx)
	if err != nil {
		log.Error("%+v", err)
		return "", errors.New(ServiceError)
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(gc.Request.RemoteAddr))
	if err != nil {
		log.Error("get remote ip error: ", err)
		return "", errors.New(ServiceError)
	}

	client := pb.NewAccountClient(AccountConn)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
}
