package dao

import (
	"changweiba-backend/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	dbEngine *xorm.Engine
	commentMutex sync.Mutex
	replyMutex sync.Mutex
)

func Init(){
	var err error
	dataSourceName:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",conf.Cfg.DB.User,conf.Cfg.DB.Passwd,
		conf.Cfg.DB.Host, 
		conf.Cfg.DB.Port,conf.Cfg.DB.Dbname)
	dbEngine, err = xorm.NewEngine("mysql", dataSourceName)
	if err!=nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	if err=dbEngine.Ping();err!=nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	if conf.Cfg.DB.MaxIdleConns>0{
		dbEngine.SetMaxIdleConns(conf.Cfg.DB.MaxIdleConns)
	}
	if conf.Cfg.DB.MaxOpenConns>0{
		dbEngine.SetMaxOpenConns(conf.Cfg.DB.MaxOpenConns)
	}
	//日志
	f,err:=os.Create(conf.Cfg.DB.LogFile)
	if err!=nil{
		logs.Error("create sql file failed:",err.Error())
	} else{
		dbEngine.SetLogger(xorm.NewSimpleLogger(f))
		dbEngine.ShowSQL(true)	//不能忽略
	}
	
}

func InsertUser(userName,password,ip,avatar string) (int64,error){
	//先检查name是否存在
	u:=User{
		Name:userName,
	}
	has,err:=GetUser(&u)
	if err!=nil{
		return 0, err
	}
	if has{
		return 0,status.Error(codes.AlreadyExists,"this user has already exists")
	}
	
	now:=time.Now().Unix()
	user:=User{
		Name:userName,
		Password:password,
		Ip:InetAtoi(ip),
		CreateTime:now,
		LastUpdate:now,
		Avatar:avatar,
	}
	_,err=dbEngine.InsertOne(&user)
	if err!=nil{
		return 0,status.Error(codes.Internal,err.Error())
	}
	//spew.Dump(user)
	return user.Id,nil
}

func GetUser(user *User)(bool,error){
	has,err:=dbEngine.Get(user)
	if err!=nil{
		return false, status.Error(codes.Internal,err.Error())
	}
	return has,nil
}

//随机获取一个头像url
func GetRandomAvatar() (url string,err error){
	var avatars []Avatar
	if err=dbEngine.Cols("url").Find(&avatars);err!=nil{
		return "",status.Error(codes.Internal,err.Error())
	}
	if len(avatars)==0{
		return "",status.Error(codes.Internal,"there are no avatar data in db")
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	index:=seed.Intn(len(avatars))
	url=avatars[index].Url
	return 
}

func InsertPost(userId int64,topic string) (int64,error){
	now:=time.Now().Unix()
	post:=Post{
		UserId:     userId,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	_,err:=dbEngine.InsertOne(&post)
	if err!=nil{
		return 0,err
	}
	return post.Id,nil
}

func InsertComment(userId int64,postId int64,content string) (int64,error){
	commentMutex.Lock()
	defer commentMutex.Unlock()
	//先获取楼层数
	sql:="SELECT count(*) AS total FROM comment WHERE post_id=?"
	results,err:=dbEngine.Query(sql,postId)
	floor,err:=strconv.Atoi(string(results[0]["total"]))
	if err!=nil{
		return 0, err
	}
	now:=time.Now().Unix()
	comment:=Comment{
		UserId:     userId,
		PostId:     postId,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:int32(floor+1),
	}
	_,err=dbEngine.InsertOne(&comment)
	if err!=nil{
		return 0,err
	}
	go increasePostReplyNum(postId)
	return comment.Id,nil
}

func InsertReply(userId,postId,commentId,parentId int64,content string) (int64,error){
	replyMutex.Lock()
	defer replyMutex.Unlock()
	//先获取楼层数
	sql:="SELECT count(*) AS total FROM reply WHERE comment_id=?"
	results,err:=dbEngine.Query(sql,commentId)
	floor,err:=strconv.Atoi(string(results[0]["total"]))
	if err!=nil{
		return 0, err
	}
	now:=time.Now().Unix()
	reply:=Reply{
		UserId:userId,
		PostId:postId,
		Content:content,
		ParentId:parentId,
		CommentId:commentId,
		CreateTime:now,
		Status:0,
		Floor:int32(floor),
	}
	_,err=dbEngine.InsertOne(&reply)
	if err!=nil{
		return 0, err
	}
	go increasePostReplyNum(postId)
	return reply.Id,nil
}

//帖子回复数+1
func increasePostReplyNum(postId int64){
	sql:="UPDATE post SET reply_num=reply_num+1 WHERE id=?"
	dbEngine.Exec(sql,postId)
}

func decreasePostReplyNum(postId int64){
	sql:="UPDATE post SET reply_num=reply_num-1 WHERE id=?"
	dbEngine.Exec(sql,postId)
}

func DeletePost(postId int64) error{
	sql:="UPDATE post SET status=0 WHERE id=?"
	_,err:=dbEngine.Exec(sql,postId)
	if err==nil{
		go decreasePostReplyNum(postId)
	}
	return err
}

func DeleteComment(commentId int64) error{
	sql:="UPDATE comment SET status=0 WHERE id=?"
	_,err:=dbEngine.Exec(sql,commentId)
	if err==nil{
		go func() {
			sql="SELECT post_id FROM comment WHERE id=?"
			results,_:=dbEngine.Query(sql,commentId)
			postId,_:=strconv.ParseInt(string(results[0]["post_id"]),10,64)
			go decreasePostReplyNum(postId)
		}()
	}
	return err
}

func DeleteReply(replyId int64) error{
	sql:="UPDATE reply SET status=0 WHERE id=?"
	_,err:=dbEngine.Exec(sql,replyId)
	if err==nil{
		go func() {
			sql="SELECT post_id FROM reply WHERE id=?"
			results,_:=dbEngine.Query(sql,replyId)
			postId,_:=strconv.ParseInt(string(results[0]["post_id"]),10,64)
			go decreasePostReplyNum(postId)
		}()
	}
	return err
}

func GetPost(post *Post) (bool,error){
	has,err:=dbEngine.Get(post)
	if err!=nil{
		return false, status.Error(codes.Internal,err.Error())
	}
	return has,nil
}

func GetComment(comment *Comment) (bool,error){
	has,err:=dbEngine.Get(comment)
	if err!=nil{
		return false, status.Error(codes.Internal,err.Error())
	}
	return has,nil
}

func GetReply(reply *Reply) (bool,error){
	has,err:=dbEngine.Get(reply)
	if err!=nil{
		return false, status.Error(codes.Internal,err.Error())
	}
	return has,nil
}

//ip地址int->string相互转换
func InetAtoi(ip string) int64{
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func InetItoa(ip int64) string{
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
