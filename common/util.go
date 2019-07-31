package common

import (
	"errors"
	"google.golang.org/grpc/status"
)

func GRPCErrorConvert(err error) error{
	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		return errors.New("system error")
	}
	return errors.New(st.Message())
}
