package common

import (
	"errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
	ErrTokenInternal    = errors.New("token internal error")
)

const (
	// ServiceError service error
	ServiceError = "gateway service internal error"
	GinContext   = "GinContextKey"
)

// ErrorCode service error code
type ErrorCode uint32

const (
	Unknown          ErrorCode = 10000
	InvalidArgument  ErrorCode = 400
	NotFound         ErrorCode = 444
	AlreadyExists    ErrorCode = 455
	PermissionDenied ErrorCode = 555
	Internal         ErrorCode = 500
)

// ErrMap service error code map to grpc service code
var ErrMap = map[codes.Code]ErrorCode{
	codes.Unknown:         Unknown,
	codes.AlreadyExists:   AlreadyExists,
	codes.NotFound:        NotFound,
	codes.Internal:        Internal,
	codes.InvalidArgument: InvalidArgument,
}
