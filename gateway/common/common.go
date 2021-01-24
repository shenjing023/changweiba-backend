package common

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GinContextFromContext normal ctx covert to gin ctx
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		return nil, errors.New("retrieve gin.Context error")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}

// GRPCErrorConvert grpc error convert to service error
func GRPCErrorConvert(err error, conf map[codes.Code]string) error {
	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		return errors.New("system error")
	}
	var errMsg = st.Message()
	for k, v := range conf {
		if k == st.Code() {
			errMsg = v
			break
		}
	}
	return errors.New(errMsg)
}
