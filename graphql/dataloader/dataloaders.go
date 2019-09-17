//go:generate go run github.com/vektah/dataloaden UserLoader int *changweiba-backend/graphql/models.User
//go:generate go run github.com/vektah/dataloaden CommentLoader int []*changweiba-backend/graphql/models.Comment
//go:generate go run github.com/vektah/dataloaden ReplyLoader int []*changweiba-backend/graphql/models.Reply

package dataloader

import (
	"changweiba-backend/graphql/models"
	"context"
	"net/http"
)

var ctxKey ="normalCtx" 

type loaders struct {
	
}

func LoaderMiddleware(next http.Handler) http.Handler{
	
}

func userLoaderFunc (keys []int) ([]*models.User,[]error){
	
}

func commentLoaderFunc(keys []int) ([][]*models.Comment,[]error){
	
}

func CtxLoaders(ctx context.Context) loaders{
	
}
