package main

//go:generate rm -rf generated
//go:generate go run github.com/99designs/gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"gateway/dataloader"
	"gateway/generated"
	"gateway/handler"
	"gateway/models"
)

type Resolver struct{}

func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	return dataloader.Loader.UsersByIDs.Load(ctx, int64(obj.User.ID))
}

func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment, page int, pageSize int) (*models.ReplyConnection, error) {
	return handler.GetRepliesByCommentID(ctx, obj.ID, page, pageSize)
}

func (r *mutationResolver) SignUp(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	return handler.SignUp(ctx, input)
}

func (r *mutationResolver) SignIn(ctx context.Context, input models.NewUser) (*models.AuthToken, error) {
	return handler.SignIn(ctx, input)
}

func (r *mutationResolver) EditUser(ctx context.Context, input models.EditUser) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) ReportUser(ctx context.Context, input models.ReportUser) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) NewPost(ctx context.Context, input models.NewPost) (int, error) {
	return handler.NewPost(ctx, input)
}

func (r *mutationResolver) NewComment(ctx context.Context, input models.NewComment) (int, error) {
	return handler.NewComment(ctx, input)
}

func (r *mutationResolver) NewReply(ctx context.Context, input models.NewReply) (int, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeletePost(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) GetAccessToken(ctx context.Context, input string) (string, error) {
	panic("not implemented")
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	return dataloader.Loader.UsersByIDs.Load(ctx, int64(obj.User.ID))
}

func (r *postResolver) Comments(ctx context.Context, obj *models.Post, page int, pageSize int) (*models.CommentConnection, error) {
	return handler.GetCommentsByPostID(ctx, obj.ID, page, pageSize)
}

func (r *postResolver) LastReplyUser(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic("not implemented")
}

func (r *postResolver) FirstComment(ctx context.Context, obj *models.Post) (*models.Comment, error) {
	return dataloader.Loader.FirstComment.Load(ctx, int64(obj.FirstComment.PostID))
}

func (r *queryResolver) User(ctx context.Context, userID int) (*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Post(ctx context.Context, postID int) (*models.Post, error) {
	panic("not implemented")
}

func (r *queryResolver) Posts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error) {
	return handler.Posts(ctx, page, pageSize)
}

func (r *queryResolver) Comment(ctx context.Context, commentID int) (*models.Comment, error) {
	panic("not implemented")
}

func (r *queryResolver) Comments(ctx context.Context, postID int, page int, pageSize int) (*models.CommentConnection, error) {
	panic("not implemented")
}

func (r *queryResolver) Reply(ctx context.Context, replyID int) (*models.Reply, error) {
	panic("not implemented")
}

func (r *queryResolver) Replies(ctx context.Context, commentID int, page int, pageSize int) (*models.ReplyConnection, error) {
	panic("not implemented")
}

func (r *replyResolver) User(ctx context.Context, obj *models.Reply) (*models.User, error) {
	return dataloader.Loader.UsersByIDs.Load(ctx, int64(obj.User.ID))
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User, page int, pageSize int) (*models.PostConnection, error) {
	panic("not implemented")
}

func (r *userResolver) Comments(ctx context.Context, obj *models.User, page int, pageSize int) (*models.CommentConnection, error) {
	panic("not implemented")
}

func (r *userResolver) Replies(ctx context.Context, obj *models.User, page int, pageSize int) (*models.ReplyConnection, error) {
	panic("not implemented")
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Reply returns generated.ReplyResolver implementation.
func (r *Resolver) Reply() generated.ReplyResolver { return &replyResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type replyResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
