//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"changweiba-backend/common"
	"changweiba-backend/graphql/dataloader"
	"changweiba-backend/graphql/generated"
	"changweiba-backend/graphql/models"
	"changweiba-backend/graphql/post"
	"changweiba-backend/graphql/rpc_conn"
	"changweiba-backend/graphql/user"
	"context"
	
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func InitRPCConnection(){
	rpc_conn.InitRPCConnection()
}

func StopRPCConnection(){
	rpc_conn.StopRPCConnection()
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

func (r *Resolver) Comment() generated.CommentResolver {
	return &commentResolver{r}
}

func (r *Resolver) Reply() generated.ReplyResolver {
	return &replyResolver{r}
}

type mutationResolver struct{
	*Resolver
	myUserResolver *user.MyUserResolver
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input models.NewUser) (string, error) {
	if _,err:=common.GinContextFromContext(ctx);err!=nil{
		return "", err
	}
	return user.RegisterUser(ctx,input)
}
func (r *mutationResolver) LoginUser(ctx context.Context, input models.NewUser) (string, error) {
	return user.LoginUser(ctx,input)
}
func (r *mutationResolver) EditUser(ctx context.Context, input models.EditUser) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) ReportUser(ctx context.Context, input models.ReportUser) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) NewPost(ctx context.Context, input models.NewPost) (int, error) {
	return post.NewPost(ctx,input)
}
func (r *mutationResolver) NewComment(ctx context.Context, input models.NewComment) (int, error) {
	return post.NewComment(ctx,input)
}
func (r *mutationResolver) NewReply(ctx context.Context, input models.NewReply) (int, error) {
	return post.NewReply(ctx,input)
}
func (r *mutationResolver) DeletePost(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{
	*Resolver
}

func (r *queryResolver) User(ctx context.Context, userID int) (*models.User, error) {
	return user.GetUser(ctx,userID)
}

func (r *queryResolver) Post(ctx context.Context, postID int) (*models.Post, error){
	return post.GetPost(ctx,postID)
}

func (r *queryResolver) Posts(ctx context.Context, page int, pageSize int) (*models.PostConnection, error){
	return post.GetPosts(ctx,page,pageSize)
}

func (r *queryResolver) Comment(ctx context.Context, commentID int) (*models.Comment, error){
	panic("not implemented")
}

func (r *queryResolver) Comments(ctx context.Context, postId int,page int,pageSize int) (*models.CommentConnection, error){
	panic("not implemented")
}

func (r *queryResolver) Reply(ctx context.Context, replyID int) (*models.Reply, error){
	panic("not implemented")
}

func (r *queryResolver) Replies(ctx context.Context, commentID int,page int,pageSize int) (*models.ReplyConnection, 
	error){
	panic("not implemented")
}

type userResolver struct {
	*Resolver
	myPostResolver *post.MyPostResolver
}

func (r *userResolver) Posts(ctx context.Context, obj *models.User, page int, pageSize int) (*models.PostConnection, error) {
	panic("not implemented")
}

func (r *userResolver) Comments(ctx context.Context, obj *models.User, page int, pageSize int) (*models.CommentConnection, 
	error){
	panic("not implemented")
}

func (r *userResolver) Replies(ctx context.Context, obj *models.User, page int, pageSize int) (*models.ReplyConnection, error){
	panic("not implemented")
}

type postResolver struct {
	*Resolver
}

func (r *postResolver) Comments(ctx context.Context, obj *models.Post, page int, pageSize int) (*models.CommentConnection, error) {
	return post.GetCommentsByPostId(ctx,obj,page,pageSize)
}

func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error){
	return dataloader.CtxLoaders(ctx).UsersByIds.Load(int64(obj.User.ID),nil)
}

func (r *postResolver) LastReplyUser(ctx context.Context, obj *models.Post) (*models.User, error){
	return dataloader.CtxLoaders(ctx).UsersByIds.Load(int64(obj.User.ID),nil)
}

type commentResolver struct {
	*Resolver
}

func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment, page int,pageSize int) (*models.ReplyConnection, error){
	return post.GetRepliesByCommentId(ctx,obj,page,pageSize)
}

func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error){
	return dataloader.CtxLoaders(ctx).UsersByIds.Load(int64(obj.User.ID),nil)
}

type replyResolver struct{
	*Resolver
}

func (r *replyResolver) User(ctx context.Context, obj *models.Reply) (*models.User, error){
	return dataloader.CtxLoaders(ctx).UsersByIds.Load(int64(obj.User.ID),nil)
}
