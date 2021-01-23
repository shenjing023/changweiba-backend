package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	ErrTokenExpired     = errors.New("Token is expired")
	ErrTokenNotValidYet = errors.New("Token not active yet")
	ErrTokenMalformed   = errors.New("That's not even a token")
	ErrTokenInvalid     = errors.New("Couldn't handle this token")
	signKey             = "secret key"
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
	jwt.StandardClaims
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

var defaultOptions = options{
	signingMethod: jwt.SigningMethodHS256,
	signingKey:    signKey,
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(signKey), nil
	},
	expired:   3600,
	tokenType: "Bearer",
	claims:    jwt.StandardClaims{},
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
				return nil, ErrTokenInvalid
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
	o := defaultOptions
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
	expiresAt := now.Add(time.Duration(j.opts.expired) * time.Second).Unix()
	claims := CustomClaims{
		Attachment: attachment,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(j.opts.signingMethod, claims)
	tokenString, err := token.SignedString([]byte(j.opts.signingKey))
	if err != nil {
		return nil, err
	}
	return &JWTToken{
		ExpiresAt: expiresAt,
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
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.Attachment, nil
	}
	return ErrTokenInvalid, nil
}

// RefreshToken refresh token
func (j *JWTAuth) RefreshToken(tokenString string) (*JWTToken, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, j.opts.keyfunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.GenerateToken(claims.Attachment)
	}
	return nil, ErrTokenInvalid
}
