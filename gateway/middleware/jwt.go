package middleware

import (
	"github.com/pkg/errors"
	log "github.com/shenjing023/llog"

	//"github.com/davecgh/go-spew/spew"

	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

// 一些常量
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
	SignKey          = "NORMAL" //服务器保存
)

type postParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func systemError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": -1,
		"msg":  "system error",
	})
	ctx.Abort()
	return
}

type options struct {
	signingMethod jwtgo.SigningMethod
	signingKey    string
	keyfunc       jwtgo.Keyfunc
	expired       int
	tokenType     string
	claims        jwtgo.Claims
}

type CustomClaims struct {
	Attachment interface{} `json:"attachment"`
	jwtgo.StandardClaims
}

type Option func(*options)

type JWTAuth struct {
	opts *options
}

type JWTToken struct {
	ExpiresAt   int64
	TokenType   string
	AccessToken string
}

var defaultOptions = options{
	signingMethod: jwtgo.SigningMethodHS256,
	signingKey:    SignKey,
	keyfunc: func(t *jwtgo.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, TokenInvalid
		}
		return []byte(SignKey), nil
	},
	expired:   3600,
	tokenType: "Bearer",
	claims:    jwtgo.StandardClaims{},
}

func NewJWT(opts ...Option) *JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{
		opts: &o,
	}
}

// SetSigningMethod 设定签名方式
func SetSigningMethod(method jwtgo.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// SetSigningKey 设定签名key
func SetSigningKey(key string) Option {
	return func(o *options) {
		o.signingKey = key
		o.keyfunc = func(t *jwtgo.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwtgo.SigningMethodHMAC); !ok {
				return nil, TokenInvalid
			}
			return []byte(key), nil
		}
	}
}

// SetKeyfunc 设定验证key的回调函数
//func SetKeyFunc(keyFunc jwtgo.Keyfunc) Option {
//	return func(o *options) {
//		o.keyfunc = keyFunc
//	}
//}

// SetExpired 设定令牌过期时长(单位秒，默认3600)
func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func (j *JWTAuth) GenerateToken(attachment interface{}) (*JWTToken, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(j.opts.expired) * time.Second).Unix()
	claims := CustomClaims{
		Attachment: attachment,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	}

	token := jwtgo.NewWithClaims(j.opts.signingMethod, claims)
	tokenString, err := token.SignedString([]byte(j.opts.signingKey))
	if err != nil {
		return nil, err
	}
	return &JWTToken{
		ExpiresAt:   expiresAt,
		TokenType:   j.opts.tokenType,
		AccessToken: tokenString,
	}, nil
}

func (j *JWTAuth) ParseToken(tokenString string) (interface{}, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, j.opts.keyfunc)
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.Attachment, nil
	}
	return TokenInvalid, nil
}

// 更新token
func (j *JWTAuth) RefreshToken(tokenString string) (*JWTToken, error) {
	jwtgo.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, j.opts.keyfunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwtgo.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.GenerateToken(claims.Attachment)
	}
	return nil, TokenInvalid
}

/*
	JWTAuth 中间件，检查token
*/
func JWTMiddleware(signKey string) gin.HandlerFunc {
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
			j := NewJWT(SetSigningKey(signKey))
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			if err != nil {
				log.Error("parse token failed:", err.Error())
				if err == TokenExpired {
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

/*
获取查询的深度
*/
func getQueryFieldDeep(set ast.SelectionSet, deep int) int {
	if set == nil {
		return deep
	}
	deep++
	max := 0
	for _, v := range set {
		if tmp, ok := v.(*ast.Field); ok {
			d := getQueryFieldDeep(tmp.SelectionSet, deep)
			if d > max {
				max = d
			}
		}
	}
	return max
}

/*
	指定的请求路径是否在queryName里
*/
func checkQuery(c *gin.Context) bool {
	queryName := c.GetStringSlice("queryName")
	for _, v := range queryName {
		if v == "posts" || v == "signIn" || v == "signUp" {
			return true
		}
	}
	return false
}
