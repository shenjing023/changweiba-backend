//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql/generated"
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//rpc连接
var accountConn *grpc.ClientConn

func InitRPCConnection(){
	var err error
	accountConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Account.Port),grpc.WithInsecure())
	if err!=nil{
		log.Fatal("fail to accountRPC dial: %+v",err)
	}
}

//关闭rpc连接
func StopRPCConnection(){
	if accountConn!=nil{
		accountConn.Close()
	}
}

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{
		Resolver:r,
		myUserResolver:&user.MyUserResolver{},
	}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		Resolver:r,
		myUserResolver:&user.MyUserResolver{},
	}
}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

func (r *Resolver) Post() generated.PostResolver {
	return &postResolver{r}
}

type mutationResolver struct{
	*Resolver
	myUserResolver *user.MyUserResolver
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input models.NewUser) (string, error) {
	if _,err:=common.GinContextFromContext(ctx);err!=nil{
		return "", err
	}
	return r.myUserResolver.RegisterUser(ctx,input,accountConn)
}
func (r *mutationResolver) LoginUser(ctx context.Context, input models.NewUser) (string, error) {
	return r.myUserResolver.LoginUser(ctx,input,accountConn)
}
func (r *mutationResolver) EditUser(ctx context.Context, input models.EditUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) ReportUser(ctx context.Context, input models.ReportUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewPost(ctx context.Context, input models.NewPost) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewComment(ctx context.Context, input models.NewComment) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewReply(ctx context.Context, input models.NewReply) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditPost(ctx context.Context, input models.EditPost) (string, error) {
	panic("not implemented")
}

type queryResolver struct{
	*Resolver
	myUserResolver *user.MyUserResolver
}

func (r *queryResolver) User(ctx context.Context, userID int) (*models.User, error) {
	return r.myUserResolver.GetUser(ctx,userID,accountConn)
}

type userResolver struct {
	*Resolver
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	panic("not implemented")
}

type postResolver struct {
	*Resolver
}

func (r *postResolver) Comments(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	panic("not implemented")
}
