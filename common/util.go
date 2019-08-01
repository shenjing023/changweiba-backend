package common

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCErrorConvert(err error,conf map[codes.Code]string) error{
	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		return errors.New("system error")
	}
	var errMsg =st.Message()
	for k,v:=range conf{
		if k==st.Code(){
			errMsg=v
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
		return nil, errors.New("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}
