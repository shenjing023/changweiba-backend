//go:generate go run github.com/vektah/dataloaden UserLoader int *changweiba-backend/graphql/models.User
//go:generate go run github.com/vektah/dataloaden CommentLoader int *changweiba-backend/graphql/models.Comment
//go:generate go run github.com/vektah/dataloaden ReplyLoader int *changweiba-backend/graphql/models.Reply

package dataloader

import (
	"changweiba-backend/graphql/models"
	"net/http"
)

func LoaderMiddleware(next http.Handler) http.Handler{
	
}

func commentLoaderFunc(keys []int) ([]*models.Comment,[]error){
	
}
