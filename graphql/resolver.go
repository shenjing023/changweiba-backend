//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql/generated"
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/post"
	"changweiba-backend/graphql/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//rpc连接
var (
	accountConn *grpc.ClientConn
	postConn *grpc.ClientConn
)

func InitRPCConnection(){
	var err error
	accountConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Account.Port),grpc.WithInsecure())
	if err!=nil{
		log.Fatal("fail to accountRPC dial: %+v",err)
	}
	postConn,err=grpc.Dial(fmt.Sprintf("localhost:%d",conf.Cfg.Post.Port),grpc.WithInsecure())
	if err!=nil{
		log.Fatal("fail to postRPC dial: %+v",err)
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
	return &userResolver{
		Resolver:r,
		myPostResolver:&post.MyPostResolver{},
	}
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
func (r *mutationResolver) ReportUser(ctx context.Context, input models.ReportUser) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewPost(ctx context.Context, input models.NewPost) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewComment(ctx context.Context, input models.NewComment) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewReply(ctx context.Context, input models.NewReply) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePost(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{
	*Resolver
	myUserResolver *user.MyUserResolver
	myPostResolver *post.MyPostResolver
}

func (r *queryResolver) User(ctx context.Context, userID int) (*models.User, error) {
	return r.myUserResolver.GetUser(ctx,userID,accountConn)
}

func (r *queryResolver) Post(ctx context.Context, postID int) (*models.Post, error){
	return r.myPostResolver.GetPost(ctx,postID,postConn)
}

func (r *queryResolver) Posts(ctx context.Context, page int, pageSize int) ([]*models.Post, error){
	panic("not implemented")
}

func (r *queryResolver) Comment(ctx context.Context, commentID int) (*models.Comment, error){
	panic("not implemented")
}

func (r *queryResolver) Reply(ctx context.Context, replyID int) (*models.Reply, error){
	panic("not implemented")
}

type userResolver struct {
	*Resolver
	myPostResolver *post.MyPostResolver
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User, page int, pageSize int) ([]*models.Post, error) {
	panic("not implemented")
}

func (r *userResolver) Comments(ctx context.Context, obj *models.User, page int, pageSize int) ([]*models.Comment, 
	error){
	panic("not implemented")
}

func (r *userResolver) Replies(ctx context.Context, obj *models.User, page int, pageSize int) ([]*models.Reply, error){
	panic("not implemented")
}

type postResolver struct {
	*Resolver
	myPostResolver *post.MyPostResolver
}

func (r *postResolver) Comments(ctx context.Context, obj *models.Post, page int, pageSize int) ([]*models.Comment, error) {
	return r.myPostResolver.GetCommentsByPostId(ctx,obj,page,pageSize)
}
