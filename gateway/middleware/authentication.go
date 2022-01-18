package middleware

import (
	"context"
	"gateway/common"
	"net/http"

	"github.com/cockroachdb/errors"

	"gateway/conf"

	"gateway/pkg/jwt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
)

var (
	accessTokenAuth  = jwt.NewJWTAuth()
	refreshTokenAuth = jwt.NewJWTAuth()
)

// AuthMiddleware authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}
		token := c.Request.Header.Get("auth")
		if token == "" {
			c.Next()
		} else {
			claims, err := ParseAccessToken(token)
			if err != nil {
				log.Error("parse token failed: ", err.Error())
				if errors.Is(err, common.ErrTokenExpired) {
					c.JSON(http.StatusBadRequest, []gqlError{{
						Message: "授权已过期",
						Extensions: map[string]interface{}{
							"code": common.InvalidArgument,
						},
						Path: c.GetStringSlice(queryNameKey),
					}})
					c.Abort()
					return
				}
				c.JSON(http.StatusBadRequest, []gqlError{{
					Message: "授权无效",
					Extensions: map[string]interface{}{
						"code": common.InvalidArgument,
					},
					Path: c.GetStringSlice(queryNameKey),
				}})
				c.Abort()
				return
			}
			//继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set("claims", claims)
			c.Next()
		}
	}
}

// GenerateAccessToken generate jwt access_token
func GenerateAccessToken(userID int64) (string, error) {
	token, err := accessTokenAuth.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

// ParseAccessToken parse access_token
func ParseAccessToken(token string) (interface{}, error) {
	return accessTokenAuth.ParseToken(token)
}

// GenerateRefreshToken generate jwt refresh_token
func GenerateRefreshToken(userID int64) (string, error) {
	token, err := refreshTokenAuth.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func RefreshRefreshToken(tokenString string) (string, error) {
	token, err := refreshTokenAuth.RefreshToken(tokenString)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

// ParseRefreshToken parse refresh_token
func ParseRefreshToken(token string) (interface{}, error) {
	return refreshTokenAuth.ParseToken(token)
}

// InitAuth init jwt token auth
func InitAuth() {
	accessTokenAuth = jwt.NewJWTAuth(
		jwt.WithSigningKey(conf.Cfg.AuthToken.Access.SignKey),
		jwt.WithExpired(conf.Cfg.AuthToken.Access.Expire),
	)
	refreshTokenAuth = jwt.NewJWTAuth(
		jwt.WithSigningKey(conf.Cfg.AuthToken.Refresh.SignKey),
		jwt.WithExpired(conf.Cfg.AuthToken.Refresh.Expire),
	)
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if _, err := common.GetUserIDFromContext(ctx); err != nil {
		return nil, common.NewGQLError(common.InvalidArgument, "未授权")
	}

	return next(ctx)
}

type gqlError struct {
	Message    string                 `json:"message"`
	Path       []string               `json:"path"`
	Extensions map[string]interface{} `json:"extensions"`
}
