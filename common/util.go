package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func ServiceErrorConvert(err error, conf map[ErrorCode]string) error {
	de, ok := FromError(err)
	if !ok {
		return errors.New("system error")
	}
	var errMsg = de.Error()
	for k, v := range conf {
		if k == de.Code {
			errMsg = v
			break
		}
	}
	return errors.New(errMsg)
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

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
