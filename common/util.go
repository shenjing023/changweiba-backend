package common

import (
	"changweiba-backend/pkg/logs"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const GinContextKey = "GinContextKey"

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
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		return nil, errors.New("retrieve gin.Context error")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}

//记录dao错误到日志
func LogDaoError(prefix string, err error) {
	de, ok := FromError(err)
	if !ok {
		logs.Error(prefix, err)
	}
	if de.Code == Unknown || de.Code == Internal {
		logs.Error(prefix, de.Err)
	} else {
		logs.Error(prefix, de)
	}
}
