//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) RegisterUser(ctx context.Context, input NewUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) LoginUser(ctx context.Context, input NewUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditUser(ctx context.Context, input EditUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) ReportUser(ctx context.Context, input ReportUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewPost(ctx context.Context, input NewPost) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewComment(ctx context.Context, input NewComment) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewReply(ctx context.Context, input NewReply) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditPost(ctx context.Context, input EditPost) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, userID string) (*User, error) {
	panic("not implemented")
}
