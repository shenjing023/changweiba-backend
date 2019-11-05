//go:generate protoc  --plugin=protoc-gen-micro=C:\GoProjects/bin/protoc-gen-micro.exe --micro_out=./pb --go_out=plugins=grpc:./pb post.proto

package post

import (
	"changweiba-backend/dao"
	pb "changweiba-backend/rpc/post/pb"
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/micro/go-micro"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

const sysError  = "post service system error" 

type Post struct{

}

func (p *Post) NewPost(ctx context.Context,pr *pb.NewPostRequest,resp *pb.NewPostResponse) error{
	if len(strings.TrimSpace(pr.Topic))==0 || len(strings.TrimSpace(pr.Content))==0{
		return status.Error(codes.InvalidArgument,"topic or content can not be empty")
	}
	postId,err:=dao.InsertPost(pr.UserId,pr.Topic,pr.Content)
	if err!=nil{
		logs.Error("insert post error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.NewPostResponse{
		PostId:int64(postId),
	}
	return nil
}

func (p *Post) NewComment(ctx context.Context,cr *pb.NewCommentRequest,resp *pb.NewCommentResponse) error{
	if len(strings.TrimSpace(cr.Content))==0{
		return status.Error(codes.InvalidArgument,"comment content can not be empty")
	}
	commentId,err:=dao.InsertComment(cr.UserId,cr.PostId,cr.Content)
	if err!=nil{
		logs.Error("insert comment error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.NewCommentResponse{
		CommentId:int64(commentId),
	}
	return nil
}

func (p *Post) NewReply(ctx context.Context,rr *pb.NewReplyRequest,resp *pb.NewReplyResponse) error{
	if len(strings.TrimSpace(rr.Content))==0{
		return status.Error(codes.InvalidArgument,"comment content can not be empty")
	}
	replyId,err:=dao.InsertReply(rr.UserId,rr.PostId,rr.CommentId,rr.ParentId,rr.Content)
	if err!=nil{
		logs.Error("insert reply error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.NewReplyResponse{
		ReplyId:replyId,
	}
	return nil
}

func (p *Post) DeletePost(ctx context.Context,dr *pb.DeleteRequest,resp *pb.DeleteResponse) error{
	err:=dao.DeletePost(dr.Id)
	if err!=nil{
		logs.Error("delete post error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.DeleteResponse{
		Success:true,
	}
	return nil
}

func (p *Post) DeleteComment(ctx context.Context,dr *pb.DeleteRequest,resp *pb.DeleteResponse) error{
	err:=dao.DeleteComment(dr.Id)
	if err!=nil{
		logs.Error("delete comment error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.DeleteResponse{
		Success:true,
	}
	return nil
}

func (p *Post) DeleteReply(ctx context.Context,dr *pb.DeleteRequest,resp *pb.DeleteResponse) error{
	err:=dao.DeleteReply(dr.Id)
	if err!=nil{
		logs.Error("delete reply error: ",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.DeleteResponse{
		Success:true,
	}
	return nil
}

func (p *Post) GetPost(ctx context.Context,pr *pb.PostRequest,resp *pb.PostResponse) error{
	dbPost,err:=dao.GetPost(pr.Id)
	if err!=nil{
		logs.Error("get post error:",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.PostResponse{
		Post:&pb.Post{
			Id:dbPost.Id,
			UserId:dbPost.UserId,
			Topic:dbPost.Topic,
			CreateTime:dbPost.CreateTime,
			LastUpdate:dbPost.LastUpdate,
			ReplyNum:dbPost.ReplyNum,
			Status:pb.Status(dbPost.Status),
		},
	}
	return nil
}

func (p *Post) GetComment(ctx context.Context,cr *pb.CommentRequest,resp *pb.CommentResponse) error{
	dbComment,err:=dao.GetComment(cr.Id)
	if err!=nil{
		logs.Error("get comment error:",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.CommentResponse{
		Comment:&pb.Comment{
			Id:dbComment.Id,
			UserId:dbComment.UserId,
			PostId:dbComment.PostId,
			Content:dbComment.Content,
			CreateTime:dbComment.CreateTime,
			Floor:dbComment.Floor,
			Status:pb.Status(dbComment.Status),
		},
	}
	return nil
}

func (p *Post) GetReply(ctx context.Context,rr *pb.ReplyRequest,resp *pb.ReplyResponse) error{
	dbReply,err:=dao.GetReply(rr.Id)
	if err!=nil{
		logs.Error("get reply error:",err.Error())
		return status.Error(codes.Internal,sysError)
	}
	resp=&pb.ReplyResponse{
		Reply:&pb.Reply{
			Id:dbReply.Id,
			UserId:dbReply.UserId,
			PostId:dbReply.PostId,
			CommentId:dbReply.CommentId,
			Content:dbReply.Content,
			CreateTime:dbReply.CreateTime,
			ParentId:dbReply.ParentId,
			Floor:dbReply.Floor,
			Status:pb.Status(dbReply.Status),
		},
	}
	return nil
}

func (p *Post) GetCommentsByPostId(ctx context.Context,cr *pb.CommentsRequest,resp *pb.CommentsResponse) error{
	dbComments,totalCount,err:=dao.GetCommentsByPostId(cr.PostId,cr.Page,cr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get comments by post_id failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	var comments []*pb.Comment
	for _,v:=range dbComments{
		comments=append(comments,&pb.Comment{
			Id:                   v.Id,
			UserId:               v.UserId,
			PostId:               v.PostId,
			Content:              v.Content,
			CreateTime:           v.CreateTime,
			Floor:                v.Floor,
			Status:               pb.Status(v.Status),
		})
	}
	resp=&pb.CommentsResponse{
		Comments: comments,
		TotalCount: totalCount,
	}
	return nil
}

func (p *Post) GetRepliesByCommentId(ctx context.Context,rr *pb.RepliesRequest,resp *pb.RepliesResponse) error{
	dbReplies,totalCount,err:=dao.GetRepliesByCommentId(rr.CommentId,rr.Page,rr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get replies by comment_id failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	var replies []*pb.Reply
	for _,v:=range dbReplies{
		replies=append(replies,&pb.Reply{
			Id:                   v.Id,
			UserId:               v.UserId,
			PostId:               v.PostId,
			CommentId:            v.CommentId,
			Content:              v.Content,
			CreateTime:           v.CreateTime,
			ParentId:             v.ParentId,
			Floor:                v.Floor,
			Status:               pb.Status(v.Status),
		})
	}
	resp=&pb.RepliesResponse{
		Replies:replies,
		TotalCount:totalCount,
	}
	return nil
}

func (p *Post) Posts(ctx context.Context,pr *pb.PostsRequest,resp *pb.PostsResponse) error{
	dbPosts,totalCount,err:=dao.GetPosts(pr.Page,pr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get posts failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	
	var posts []*pb.Post
	for _,v:=range dbPosts{
		posts=append(posts,&pb.Post{
			Id:v.Id,
			UserId:v.UserId,
			Topic:v.Topic,
			CreateTime:v.CreateTime,
			LastUpdate:v.LastUpdate,
			ReplyNum:v.ReplyNum,
			Status:pb.Status(v.Status),
		})
	}
	resp=&pb.PostsResponse{
		Posts:posts,
		TotalCount:totalCount,
	}
	return nil
}

//func (p *Post) GetRepliesByCommentIds(ctx context.Context,rr *pb.RepliesByCommentsRequest) (*pb.
//	RepliesByCommentsResponse,error){
//	dbReplies,err:=dao.GetRepliesByCommentIds(rr.CommentIds,int(rr.Limit))
//	if err!=nil{
//		logs.Error(fmt.Sprintf("get replies by comment_id failed: %+v",err))
//		return nil,err
//	}
//	var replies []*pb.RepliesByCommentsResponse_Replies_
//	for _,v:=range dbReplies{
//		var temp []*pb.Reply
//		for _,vv:=range v{
//			temp=append(temp,&pb.Reply{
//				Id:vv.Id,
//				UserId:vv.UserId,
//				PostId:vv.PostId,
//				CommentId:vv.CommentId,
//				Content:vv.Content,
//				CreateTime:vv.CreateTime,
//				ParentId:vv.ParentId,
//				Floor:vv.Floor,
//			})
//		}
//		replies=append(replies,&pb.RepliesByCommentsResponse_Replies_{
//			Replies_:temp,
//		})
//	}
//	return &pb.RepliesByCommentsResponse{
//		Replies:replies,
//	},nil
//}

func (p *Post) GetPostsByUserId(ctx context.Context, pr *pb.PostsByUserIdRequest, resp *pb.PostsByUserIdResponse) error{
	dbPosts,totalCount,err:=dao.GetPostsByUserId(pr.UserId,pr.Page,pr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get posts by user_id failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	var posts []*pb.Post
	for _,v:=range dbPosts{
		posts=append(posts,&pb.Post{
			Id:v.Id,
			UserId:v.UserId,
			Topic:v.Topic,
			CreateTime:v.CreateTime,
			LastUpdate:v.LastUpdate,
			ReplyNum:v.ReplyNum,
			Status:pb.Status(v.Status),
		})
	}
	resp=&pb.PostsByUserIdResponse{
		Posts:posts,
		TotalCount:totalCount,
	}
	return nil
}

func (p *Post) GetCommentsByUserId(ctx context.Context, cr *pb.CommentsByUserIdRequest,resp *pb.CommentsByUserIdResponse) error{
	dbComments,totalCount,err:=dao.GetCommentsByUserId(cr.UserId,cr.Page,cr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get comments by user_id failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	var comments []*pb.Comment
	for _,v:=range dbComments{
		comments=append(comments,&pb.Comment{
			Id:v.Id,
			UserId:v.UserId,
			Content:v.Content,
			CreateTime:v.CreateTime,
			PostId:v.PostId,
			Floor:v.Floor,
			Status:pb.Status(v.Status),
		})
	}
	resp=&pb.CommentsByUserIdResponse{
		Comments:comments,
		TotalCount:totalCount,
	}
	return nil
}

func (p *Post) GetRepliesByUserId(ctx context.Context, rr *pb.RepliesByUserIdRequest, resp *pb.RepliesByUserIdResponse) error{
	dbReplies,totalCount,err:=dao.GetRepliesByUserId(rr.UserId,rr.Page,rr.PageSize)
	if err!=nil{
		logs.Error(fmt.Sprintf("get replies by user_id failed: %+v",err))
		return status.Error(codes.Internal,sysError)
	}
	var replies []*pb.Reply
	for _,v:=range dbReplies{
		replies=append(replies,&pb.Reply{
			Id:v.Id,
			UserId:v.UserId,
			CommentId:int64(v.CommentId),
			PostId:int64(v.PostId),
			Content:v.Content,
			CreateTime:v.CreateTime,
			ParentId:int64(v.ParentId),
			Floor:int64(v.Floor),
			Status:pb.Status(v.Status),
		})
	}
	resp=&pb.RepliesByUserIdResponse{
		Replies:replies,
		TotalCount:int64(totalCount),
	}
	return nil
}

func NewPostService(addr string,port int){
	service:=micro.NewService(
		micro.Name("post"),
	)
	service.Init()
	pb.RegisterPostServiceHandler(service.Server(),new(Post))
	if err:=service.Run();err!=nil{
		
	}
	
	//lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",addr,port))
	//if err!=nil{
	//	logs.Error(fmt.Sprintf("post service failed to listen: %+v",err))
	//	os.Exit(1)
	//}
	//grpcServer:=grpc.NewServer()
	//pb.RegisterPostServiceServer(grpcServer,&Post{})
	//if err=grpcServer.Serve(lis);err!=nil{
	//	logs.Error("run grpcserver postservice failed: %+v",err)
	//	os.Exit(1)
	//}
}
