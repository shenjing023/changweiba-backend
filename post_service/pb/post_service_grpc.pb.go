// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: post_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	NewPost(ctx context.Context, in *NewPostRequest, opts ...grpc.CallOption) (*NewPostResponse, error)
	NewComment(ctx context.Context, in *NewCommentRequest, opts ...grpc.CallOption) (*NewCommentResponse, error)
	NewReply(ctx context.Context, in *NewReplyRequest, opts ...grpc.CallOption) (*NewReplyResponse, error)
	DeletePosts(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteComments(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteReplies(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	GetReply(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*ReplyResponse, error)
	GetCommentsByPostId(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
	GetRepliesByCommentId(ctx context.Context, in *RepliesRequest, opts ...grpc.CallOption) (*RepliesResponse, error)
	// rpc GetPosts(PostsRequest) returns(PostsResponse){}
	// rpc GetRepliesByCommentIds(RepliesByCommentsRequest) returns(RepliesByCommentsResponse){}
	GetPostsByUserId(ctx context.Context, in *PostsByUserIdRequest, opts ...grpc.CallOption) (*PostsByUserIdResponse, error)
	GetCommentsByUserId(ctx context.Context, in *CommentsByUserIdRequest, opts ...grpc.CallOption) (*CommentsByUserIdResponse, error)
	GetRepliesByUserId(ctx context.Context, in *RepliesByUserIdRequest, opts ...grpc.CallOption) (*RepliesByUserIdResponse, error)
	GetPostFirstComment(ctx context.Context, in *FirstCommentRequest, opts ...grpc.CallOption) (*FirstCommentResponse, error)
	GetAllPosts(ctx context.Context, in *AllPostsRequest, opts ...grpc.CallOption) (*PostsResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) NewPost(ctx context.Context, in *NewPostRequest, opts ...grpc.CallOption) (*NewPostResponse, error) {
	out := new(NewPostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/NewPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) NewComment(ctx context.Context, in *NewCommentRequest, opts ...grpc.CallOption) (*NewCommentResponse, error) {
	out := new(NewCommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/NewComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) NewReply(ctx context.Context, in *NewReplyRequest, opts ...grpc.CallOption) (*NewReplyResponse, error) {
	out := new(NewReplyResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/NewReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePosts(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/DeletePosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteComments(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/DeleteComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteReplies(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/DeleteReplies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetReply(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*ReplyResponse, error) {
	out := new(ReplyResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetCommentsByPostId(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetCommentsByPostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetRepliesByCommentId(ctx context.Context, in *RepliesRequest, opts ...grpc.CallOption) (*RepliesResponse, error) {
	out := new(RepliesResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetRepliesByCommentId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostsByUserId(ctx context.Context, in *PostsByUserIdRequest, opts ...grpc.CallOption) (*PostsByUserIdResponse, error) {
	out := new(PostsByUserIdResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetPostsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetCommentsByUserId(ctx context.Context, in *CommentsByUserIdRequest, opts ...grpc.CallOption) (*CommentsByUserIdResponse, error) {
	out := new(CommentsByUserIdResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetCommentsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetRepliesByUserId(ctx context.Context, in *RepliesByUserIdRequest, opts ...grpc.CallOption) (*RepliesByUserIdResponse, error) {
	out := new(RepliesByUserIdResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetRepliesByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostFirstComment(ctx context.Context, in *FirstCommentRequest, opts ...grpc.CallOption) (*FirstCommentResponse, error) {
	out := new(FirstCommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetPostFirstComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllPosts(ctx context.Context, in *AllPostsRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	NewPost(context.Context, *NewPostRequest) (*NewPostResponse, error)
	NewComment(context.Context, *NewCommentRequest) (*NewCommentResponse, error)
	NewReply(context.Context, *NewReplyRequest) (*NewReplyResponse, error)
	DeletePosts(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteComments(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteReplies(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetPost(context.Context, *PostRequest) (*PostResponse, error)
	GetComment(context.Context, *CommentRequest) (*CommentResponse, error)
	GetReply(context.Context, *ReplyRequest) (*ReplyResponse, error)
	GetCommentsByPostId(context.Context, *CommentsRequest) (*CommentsResponse, error)
	GetRepliesByCommentId(context.Context, *RepliesRequest) (*RepliesResponse, error)
	// rpc GetPosts(PostsRequest) returns(PostsResponse){}
	// rpc GetRepliesByCommentIds(RepliesByCommentsRequest) returns(RepliesByCommentsResponse){}
	GetPostsByUserId(context.Context, *PostsByUserIdRequest) (*PostsByUserIdResponse, error)
	GetCommentsByUserId(context.Context, *CommentsByUserIdRequest) (*CommentsByUserIdResponse, error)
	GetRepliesByUserId(context.Context, *RepliesByUserIdRequest) (*RepliesByUserIdResponse, error)
	GetPostFirstComment(context.Context, *FirstCommentRequest) (*FirstCommentResponse, error)
	GetAllPosts(context.Context, *AllPostsRequest) (*PostsResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) NewPost(context.Context, *NewPostRequest) (*NewPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPost not implemented")
}
func (UnimplementedPostServiceServer) NewComment(context.Context, *NewCommentRequest) (*NewCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewComment not implemented")
}
func (UnimplementedPostServiceServer) NewReply(context.Context, *NewReplyRequest) (*NewReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewReply not implemented")
}
func (UnimplementedPostServiceServer) DeletePosts(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosts not implemented")
}
func (UnimplementedPostServiceServer) DeleteComments(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComments not implemented")
}
func (UnimplementedPostServiceServer) DeleteReplies(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReplies not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) GetComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedPostServiceServer) GetReply(context.Context, *ReplyRequest) (*ReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReply not implemented")
}
func (UnimplementedPostServiceServer) GetCommentsByPostId(context.Context, *CommentsRequest) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByPostId not implemented")
}
func (UnimplementedPostServiceServer) GetRepliesByCommentId(context.Context, *RepliesRequest) (*RepliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepliesByCommentId not implemented")
}
func (UnimplementedPostServiceServer) GetPostsByUserId(context.Context, *PostsByUserIdRequest) (*PostsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByUserId not implemented")
}
func (UnimplementedPostServiceServer) GetCommentsByUserId(context.Context, *CommentsByUserIdRequest) (*CommentsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByUserId not implemented")
}
func (UnimplementedPostServiceServer) GetRepliesByUserId(context.Context, *RepliesByUserIdRequest) (*RepliesByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepliesByUserId not implemented")
}
func (UnimplementedPostServiceServer) GetPostFirstComment(context.Context, *FirstCommentRequest) (*FirstCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostFirstComment not implemented")
}
func (UnimplementedPostServiceServer) GetAllPosts(context.Context, *AllPostsRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_NewPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).NewPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/NewPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).NewPost(ctx, req.(*NewPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_NewComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).NewComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/NewComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).NewComment(ctx, req.(*NewCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_NewReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).NewReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/NewReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).NewReply(ctx, req.(*NewReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeletePosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePosts(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeleteComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteComments(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteReplies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteReplies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeleteReplies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteReplies(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetReply(ctx, req.(*ReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetCommentsByPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetCommentsByPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetCommentsByPostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetCommentsByPostId(ctx, req.(*CommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetRepliesByCommentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetRepliesByCommentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetRepliesByCommentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetRepliesByCommentId(ctx, req.(*RepliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetPostsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostsByUserId(ctx, req.(*PostsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetCommentsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetCommentsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetCommentsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetCommentsByUserId(ctx, req.(*CommentsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetRepliesByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepliesByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetRepliesByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetRepliesByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetRepliesByUserId(ctx, req.(*RepliesByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostFirstComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirstCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostFirstComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetPostFirstComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostFirstComment(ctx, req.(*FirstCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllPosts(ctx, req.(*AllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPost",
			Handler:    _PostService_NewPost_Handler,
		},
		{
			MethodName: "NewComment",
			Handler:    _PostService_NewComment_Handler,
		},
		{
			MethodName: "NewReply",
			Handler:    _PostService_NewReply_Handler,
		},
		{
			MethodName: "DeletePosts",
			Handler:    _PostService_DeletePosts_Handler,
		},
		{
			MethodName: "DeleteComments",
			Handler:    _PostService_DeleteComments_Handler,
		},
		{
			MethodName: "DeleteReplies",
			Handler:    _PostService_DeleteReplies_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _PostService_GetComment_Handler,
		},
		{
			MethodName: "GetReply",
			Handler:    _PostService_GetReply_Handler,
		},
		{
			MethodName: "GetCommentsByPostId",
			Handler:    _PostService_GetCommentsByPostId_Handler,
		},
		{
			MethodName: "GetRepliesByCommentId",
			Handler:    _PostService_GetRepliesByCommentId_Handler,
		},
		{
			MethodName: "GetPostsByUserId",
			Handler:    _PostService_GetPostsByUserId_Handler,
		},
		{
			MethodName: "GetCommentsByUserId",
			Handler:    _PostService_GetCommentsByUserId_Handler,
		},
		{
			MethodName: "GetRepliesByUserId",
			Handler:    _PostService_GetRepliesByUserId_Handler,
		},
		{
			MethodName: "GetPostFirstComment",
			Handler:    _PostService_GetPostFirstComment_Handler,
		},
		{
			MethodName: "GetAllPosts",
			Handler:    _PostService_GetAllPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_service.proto",
}
