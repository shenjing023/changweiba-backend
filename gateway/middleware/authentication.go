package middleware

import (
	"gateway/common"
	"net/http"

	"gateway/conf"

	"github.com/gin-gonic/gin"
	log "github.com/shenjing023/llog"
)

var (
	accessTokenAuth  = common.NewJWTAuth()
	refreshTokenAuth = common.NewJWTAuth()
)

// AuthMiddleware authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}
		flag := checkQuery(c)
		token := c.Request.Header.Get("token")
		if token == "" && !flag {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		if !flag {
			claims, err := ParseAccessToken(token)
			if err != nil {
				log.Error("parse token failed:", err.Error())
				if err == common.ErrTokenExpired {
					c.JSON(http.StatusOK, gin.H{
						"status": -1,
						"msg":    "授权已过期",
					})
					c.Abort()
					return
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "token无效",
				})
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

// ParseRefreshToken parse refresh_token
func ParseRefreshToken(token string) (interface{}, error) {
	return refreshTokenAuth.ParseToken(token)
}

// InitAuth init jwt token auth
func InitAuth() {
	accessTokenAuth = common.NewJWTAuth(
		common.WithSigningKey(conf.Cfg.AuthToken.Access.SignKey),
		common.WithExpired(conf.Cfg.AuthToken.Access.Expire),
	)
	refreshTokenAuth = common.NewJWTAuth(
		common.WithSigningKey(conf.Cfg.AuthToken.Refresh.SignKey),
		common.WithExpired(conf.Cfg.AuthToken.Refresh.Expire),
	)
}
