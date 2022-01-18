package jwt

import (
	"gateway/common"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v4"
)

const (
	signKey = "secret key"
)

// options jwt
type options struct {
	signingMethod jwt.SigningMethod
	signingKey    string
	keyfunc       jwt.Keyfunc
	expired       int //second unit
	tokenType     string
	claims        jwt.Claims
}

// CustomClaims custom claim
type CustomClaims struct {
	Attachment interface{} `json:"attachment"`
	jwt.RegisteredClaims
}

// Option set option function
type Option func(*options)

// JWTAuth auth struct
type JWTAuth struct {
	opts *options
}

// JWTToken jwt token struct
type JWTToken struct {
	ExpiresAt int64
	TokenType string
	Token     string
}

// WithSigningMethod 设定签名方式
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithSigningKey 设定签名key
func WithSigningKey(key string) Option {
	return func(o *options) {
		o.signingKey = key
		o.keyfunc = func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, common.ErrTokenInvalid
			}
			return []byte(key), nil
		}
	}
}

// WithExpired 设定令牌过期时长(单位秒，默认3600)
func WithExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// NewJWTAuth create auth
func NewJWTAuth(opts ...Option) *JWTAuth {
	o := options{
		signingMethod: jwt.SigningMethodHS256,
		signingKey:    signKey,
		keyfunc: func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, common.ErrTokenInvalid
			}
			return []byte(signKey), nil
		},
		expired:   3600,
		tokenType: "Bearer",
		claims:    jwt.RegisteredClaims{},
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{
		opts: &o,
	}
}

// GenerateToken generate new token
func (j *JWTAuth) GenerateToken(attachment interface{}) (*JWTToken, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(j.opts.expired) * time.Second)
	claims := CustomClaims{
		Attachment: attachment,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(j.opts.signingMethod, claims)
	tokenString, err := token.SignedString([]byte(j.opts.signingKey))
	if err != nil {
		return nil, errors.Wrap(err, "jwt error")
	}
	return &JWTToken{
		ExpiresAt: expiresAt.Unix(),
		TokenType: j.opts.tokenType,
		Token:     tokenString,
	}, nil
}

// ParseToken parse token
func (j *JWTAuth) ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, j.opts.keyfunc)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, common.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, common.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, common.ErrTokenNotValidYet
			} else {
				return nil, common.ErrTokenInvalid
			}
		}
		return nil, common.ErrTokenInternal
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.Attachment, nil
	}
	return common.ErrTokenInvalid, nil
}

// RefreshToken refresh token
func (j *JWTAuth) RefreshToken(tokenString string) (*JWTToken, error) {
	attachment, err := j.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return j.GenerateToken(attachment)
}
