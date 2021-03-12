package dataloader

//go:generate go run github.com/shenjing023/dataloaden generate -s UserLoader -k int64 -v *gateway/models.User
//go:generate go run github.com/vektah/dataloaden CommentLoader int []*changweiba-backend/graphql/models.Comment
//go:generate go run github.com/vektah/dataloaden ReplyLoader int []*changweiba-backend/graphql/models.Reply

import (
	"gateway/handler"
	"time"
)

type loaders struct {
	UsersByIDs *UserLoader
}

var Loader loaders

// Init init dataloader
func Init() {
	Loader = loaders{}
	Loader.UsersByIDs = &UserLoader{
		wait:       1 * time.Millisecond,
		maxBatch:   100,
		fetch:      handler.UsersByIDsLoaderFunc,
		expiration: time.Minute,
	}
}
