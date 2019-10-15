//go:generate protoc  --go_out=plugins=grpc:./pb post.proto

package post

import (
	"changweiba-backend/dao"
	pb "changweiba-backend/rpc/post/pb"
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"strings"
)

const sysError  = "post service system error" 

type Post struct{

}

func (p *Post) NewPost(ctx context.Context,pr *pb.NewPostRequest) (*pb.NewPostResponse,error){
	if len(strings.TrimSpace(pr.Topic))==0 || len(strings.TrimSpace(pr.Content))==0{
		return nil,status.Error(codes.InvalidArgument,"topic or content can not be empty")
	}
	postId,err:=dao.InsertPost(pr.UserId,pr.Topic,pr.Content)
	if err!=nil{
		logs.Error("insert post error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.NewPostResponse{
		PostId:postId,
	},nil
}

func (p *Post) NewComment(ctx context.Context,cr *pb.NewCommentRequest) (*pb.NewCommentResponse,error){
	if len(strings.TrimSpace(cr.Content))==0{
		return nil,status.Error(codes.InvalidArgument,"comment content can not be empty")
	}
	commentId,err:=dao.InsertComment(cr.UserId,cr.PostId,cr.Content)
	if err!=nil{
		logs.Error("insert comment error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.NewCommentResponse{
		CommentId:commentId,
	}, nil
}

func (p *Post) NewReply(ctx context.Context,rr *pb.NewReplyRequest) (*pb.NewReplyResponse,error){
	if len(strings.TrimSpace(rr.Content))==0{
		return nil,status.Error(codes.InvalidArgument,"comment content can not be empty")
	}
	replyId,err:=dao.InsertReply(rr.UserId,rr.PostId,rr.CommentId,rr.ParentId,rr.Content)
	if err!=nil{
		logs.Error("insert reply error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.NewReplyResponse{
		ReplyId:replyId,
	}, nil
}

func (p *Post) DeletePost(ctx context.Context,dr *pb.DeleteRequest) (*pb.DeleteResponse,error){
	err:=dao.DeletePost(dr.Id)
	if err!=nil{
		logs.Error("delete post error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.DeleteResponse{
		Success:true,
	}, nil
}

func (p *Post) DeleteComment(ctx context.Context,dr *pb.DeleteRequest) (*pb.DeleteResponse,error){
	err:=dao.DeleteComment(dr.Id)
	if err!=nil{
		logs.Error("delete comment error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.DeleteResponse{
		Success:true,
	}, nil
}

func (p *Post) DeleteReply(ctx context.Context,dr *pb.DeleteRequest) (*pb.DeleteResponse,error){
	err:=dao.DeleteReply(dr.Id)
	if err!=nil{
		logs.Error("delete reply error: ",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	return &pb.DeleteResponse{
		Success:true,
	}, nil
}

func (p *Post) GetPost(ctx context.Context,pr *pb.PostRequest) (*pb.PostResponse,error){
	dbPost:=&dao.Post{
		Id:pr.Id,
	}
	has,err:=dao.GetPost(dbPost)
	if err!=nil{
		logs.Error("get post error:",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	if has{
		pbPost:=&pb.PostResponse{
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
		return pbPost,nil
	} else{
		logs.Info("get post failed: post is not exist")
		return &pb.PostResponse{}, nil
	}
}

func (p *Post) GetComment(ctx context.Context,cr *pb.CommentRequest) (*pb.CommentResponse,error){
	dbComment:=&dao.Comment{
		Id:cr.Id,
	}
	has,err:=dao.GetComment(dbComment)
	if err!=nil{
		logs.Error("get comment error:",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	if has{
		pbComment:=&pb.CommentResponse{
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
		return pbComment,nil
	} else{
		logs.Info("get comment failed: comment is not exist")
		return &pb.CommentResponse{}, nil
	}
}

func (p *Post) GetReply(ctx context.Context,rr *pb.ReplyRequest) (*pb.ReplyResponse,error){
	dbReply:=&dao.Reply{
		Id:rr.Id,
	}
	has,err:=dao.GetReply(dbReply)
	if err!=nil{
		logs.Error("get reply error:",err.Error())
		return nil,status.Error(codes.Internal,sysError)
	}
	if has{
		pbReply:=&pb.ReplyResponse{
			Reply:&pb.Reply{
				Id:dbReply.Id,
				UserId:dbReply.UserId,
				PostId:dbReply.PostId,
				CommentId:dbReply.CommentId,
				Content:dbReply.Content,
				CreateTime:dbReply.CreateTime,
				ParentId:dbReply.ParentId,
				Floor:int32(dbReply.Floor),
				Status:pb.Status(dbReply.Status),
			},
		}
		return pbReply,nil
	} else{
		logs.Info("get reply failed: reply is not exist")
		return &pb.ReplyResponse{}, nil
	}
}

func (p *Post) GetCommentsByPostId(ctx context.Context,cr *pb.CommentsRequest) (*pb.CommentsResponse,error){
	totalCount,err:=dao.GetCommentsCountByPostId(cr.PostId)
	if err!=nil{
		logs.Error(fmt.Sprintf("get comments_count by post_id failed: %+v",err))
		return nil, status.Error(codes.Internal,sysError)
	}
	dbComments,err:=dao.GetCommentsByPostId(cr.PostId,int(cr.Offset),int(cr.Limit))
	if err!=nil{
		logs.Error(fmt.Sprintf("get comments by post_id failed: %+v",err))
		return nil, status.Error(codes.Internal,sysError)
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
	return &pb.CommentsResponse{
		Comments:             comments,
		TotalCount:int32(totalCount),
	},nil
}

func (p *Post) GetRepliesByCommentId(ctx context.Context,rr *pb.RepliesRequest) (*pb.RepliesResponse,error){
	totalCount,err:=dao.GetRepliesCountByCommentId(rr.CommentId)
	if err!=nil{
		logs.Error(fmt.Sprintf("get replies_count by comment_id failed: %+v",err))
		return nil,status.Error(codes.Internal,sysError)
	}
	dbReplies,err:=dao.GetRepliesByCommentId(rr.CommentId,int(rr.Offset),int(rr.Limit))
	if err!=nil{
		logs.Error(fmt.Sprintf("get replies by comment_id failed: %+v",err))
		return nil,status.Error(codes.Internal,sysError)
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
	return &pb.RepliesResponse{
		Replies:replies,
		TotalCount:int32(totalCount),
	}, nil
}

func (p *Post) Posts(ctx context.Context,pr *pb.PostsRequest) (*pb.PostsResponse,error){
	dbPosts,err:=dao.GetPosts(int(pr.Offset),int(pr.Limit))
	if err!=nil{
		logs.Error(fmt.Sprintf("get posts failed: %+v",err))
		return nil,status.Error(codes.Internal,sysError)
	}
	totalCount,err:=dao.GetPostsCount()
	if err!=nil{
		logs.Error(fmt.Sprintf("get posts_count failed: %+v",err))
		return nil,status.Error(codes.Internal,sysError)
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
	return &pb.PostsResponse{
		Posts:posts,
		TotalCount:int32(totalCount),
	}, nil
}

func (p *Post) GetRepliesByCommentIds(ctx context.Context,rr *pb.RepliesByCommentsRequest) (*pb.
	RepliesByCommentsResponse,error){
	dbReplies,err:=dao.GetRepliesByCommentIds(rr.CommentIds,int(rr.Limit))
	if err!=nil{
		logs.Error(fmt.Sprintf("get replies by comment_id failed: %+v",err))
		return nil,err
	}
	var replies []*pb.RepliesByCommentsResponse_Replies_
	for _,v:=range dbReplies{
		var temp []*pb.Reply
		for _,vv:=range v{
			temp=append(temp,&pb.Reply{
				Id:vv.Id,
				UserId:vv.UserId,
				PostId:vv.PostId,
				CommentId:vv.CommentId,
				Content:vv.Content,
				CreateTime:vv.CreateTime,
				ParentId:vv.ParentId,
				Floor:vv.Floor,
			})
		}
		replies=append(replies,&pb.RepliesByCommentsResponse_Replies_{
			Replies_:temp,
		})
	}
	return &pb.RepliesByCommentsResponse{
		Replies:replies,
	},nil
}

func NewPostService(addr string,port int){
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",addr,port))
	if err!=nil{
		logs.Error(fmt.Sprintf("post service failed to listen: %+v",err))
		os.Exit(1)
	}
	grpcServer:=grpc.NewServer()
	pb.RegisterPostServiceServer(grpcServer,&Post{})
	if err=grpcServer.Serve(lis);err!=nil{
		logs.Error("run grpcserver postservice failed: %+v",err)
		os.Exit(1)
	}
}
