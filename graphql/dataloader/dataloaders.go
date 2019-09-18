//go:generate go run github.com/vektah/dataloaden UserLoader int64 *changweiba-backend/graphql/models.User
//go:generate go run github.com/vektah/dataloaden CommentLoader int []*changweiba-backend/graphql/models.Comment
//go:generate go run github.com/vektah/dataloaden ReplyLoader int []*changweiba-backend/graphql/models.Reply

package dataloader

import (
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/user"
	"context"
	"net/http"
	"time"
)

var ctxKey ="dataloaderCtx" 

type loaders struct {
	UsersByIds  *UserLoader
}

func LoaderMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		ldrs:=loaders{}
		ldrs.UsersByIds=&UserLoader{
			wait:1*time.Millisecond,
			maxBatch:100,
			fetch:userLoaderFunc,
		}
		dlCtx:=context.WithValue(r.Context(),ctxKey,ldrs)
		next.ServeHTTP(w,r.WithContext(dlCtx))
	})
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

func commentLoaderFunc(keys []int) ([][]*models.Comment,[]error){
	
}

func CtxLoaders(ctx context.Context) loaders{
	return ctx.Value(ctxKey).(loaders)
}
