//go:generate go run github.com/vektah/dataloaden UserLoader int64 *changweiba-backend/graphql/models.User
//go:generate go run github.com/vektah/dataloaden CommentLoader int []*changweiba-backend/graphql/models.Comment
//go:generate go run github.com/vektah/dataloaden ReplyLoader int []*changweiba-backend/graphql/models.Reply

package dataloader

import (
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/user"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

var ctxKey ="dataloaderCtx" 

type loaders struct {
	UsersByIds  *UserLoader
}

func LoaderMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		ldrs:=loaders{}
		ldrs.UsersByIds=&UserLoader{
			wait:1*time.Millisecond,
			maxBatch:100,
			fetch:userLoaderFunc,
		}
		c.Set(ctxKey,ldrs)
		c.Next()
	}
}

func userLoaderFunc (keys []int64,params interface{}) ([]*models.User,[]error){
	//rpc调用
	ctx:=context.Background()
	users,err:=user.GetUsers(ctx,keys)
	var errs []error
	if err!=nil{
		for i:=0;i<len(keys);i++{
			errs=append(errs,err)
		}
	}
	return users,errs
}

func CtxLoaders(ctx context.Context) loaders{
	return ctx.Value(ctxKey).(loaders)
}
