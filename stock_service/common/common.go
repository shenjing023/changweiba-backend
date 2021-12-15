package common

import (
	"google.golang.org/grpc/codes"
)

type ErrorCode uint8

const (
	Unknown ErrorCode = iota
	AlreadyExists
	NotFound
	Internal
	InvalidArgument
)

var (
	// ErrMap 对应的grpc error code
	ErrMap = map[ErrorCode]codes.Code{
		Unknown:         codes.Unknown,
		AlreadyExists:   codes.AlreadyExists,
		NotFound:        codes.NotFound,
		Internal:        codes.Internal,
		InvalidArgument: codes.InvalidArgument,
	}
)

type ServiceErr struct {
	Code ErrorCode
	Err  error
}

func (d *ServiceErr) Error() string {
	return d.Err.Error()
}

func NewServiceErr(code ErrorCode, err error) *ServiceErr {
	return &ServiceErr{code, err}
}

// Err2ServiceErr normal error to ServiceErr
func Err2ServiceErr(err error) (*ServiceErr, bool) {
	if de, ok := err.(*ServiceErr); ok {
		return de, true
	}
	return NewServiceErr(Unknown, err), false
}
