package common

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserInfo struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

// GinContextFromContext normal ctx covert to gin ctx
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContext)
	if ginContext == nil {
		return nil, errors.New("retrieve gin.Context error")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}

// GRPCErrorConvert grpc error convert to service error
func GRPCErrorConvert(err error, conf map[codes.Code]string) error {
	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		return &gqlerror.Error{
			Message: "system error",
			Extensions: map[string]interface{}{
				"code": Internal,
			},
		}
	}
	var (
		errMsg = st.Message()
		code   = Unknown
	)
	for k, v := range conf {
		if k == st.Code() {
			errMsg = v
			code = ErrMap[k]
			break
		}
	}
	return &gqlerror.Error{
		Message: errMsg,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

func HTTPErrorConvert(err error, code int) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

// GetUserIDFromContext get user_id from context
func GetUserIDFromContext(ctx context.Context) (*UserInfo, error) {
	gctx, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tmp, ok := gctx.Value("claims").(map[string]any)
	if !ok {
		return nil, errors.New("get user_info from request ctx error")
	}
	user := UserInfo{
		ID:   int64(tmp["id"].(float64)),
		Role: tmp["role"].(string),
	}
	return &user, nil
}

// NewGQLError new graphql error return to front end
func NewGQLError(code ErrorCode, msg string) error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}
