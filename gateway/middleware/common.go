package middleware

import (
	"context"

	"gateway/common"

	"github.com/gin-gonic/gin"
)

// GinContextToContextMiddleware gin ctx middleware
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), common.GinContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
