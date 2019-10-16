package dao

import (
	"changweiba-backend/conf"
	"encoding/binary"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/goinggo/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	dbEngine *xorm.Engine
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

func InsertPost(userId int64,topic string,content string) (int64,error){
	session := dbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0,err
	}
	now:=time.Now().Unix()
	post:=Post{
		UserId:     userId,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	_,err:=session.InsertOne(&post)
	if err!=nil{
		session.Rollback()
		return 0,err
	}
	comment:=Comment{
		UserId:     userId,
		PostId:     post.Id,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:		1,
	}
	_,err=session.InsertOne(&comment)
	if err!=nil{
		session.Rollback()
		return 0,err
	}
	go increasePostReplyNum(post.Id)
	session.Commit()
	return post.Id,nil
}


func InsertComment(userId int64,postId int64,content string) (int64,error){
	session := dbEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return 0,err
	}
	//先获取楼层数
	sql:="SELECT count(*) AS total FROM comment WHERE post_id=? FOR UPDATE"
	results,err:=session.Query(sql,postId)
	floor,err:=strconv.Atoi(string(results[0]["total"]))
	if err!=nil{
		session.Rollback()
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
	_,err=session.InsertOne(&comment)
	if err!=nil{
		session.Rollback()
		return 0,err
	}
	go increasePostReplyNum(postId)
	session.Commit()
	return comment.Id,nil
}

func InsertReply(userId,postId,commentId,parentId int64,content string) (int64,error){
	session := dbEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return 0,err
	}
	//先获取楼层数
	sql:="SELECT count(*) AS total FROM reply WHERE comment_id=? FOR UPDATE"
	results,err:=session.Query(sql,commentId)
	floor,err:=strconv.Atoi(string(results[0]["total"]))
	if err!=nil{
		session.Rollback()
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
		Floor:int32(floor+1),
	}
	_,err=session.InsertOne(&reply)
	if err!=nil{
		session.Rollback()
		return 0, err
	}
	go increasePostReplyNum(postId)
	session.Commit()
	return reply.Id,nil
}

//帖子回复数+1
func increasePostReplyNum(postId int64){
	sql:="UPDATE post SET reply_num=reply_num+1,last_update=UNIX_TIMESTAMP() WHERE id=?"
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
		return false, err
	}
	return has,nil
}

func GetPosts(page int,pageSize int) ([]*Post,error){
	var posts []*Post
	err:=dbEngine.Where("status = ?",0).Desc("last_update").Limit(pageSize,page).Find(&posts)
	return posts,err
}

func GetPostsCount() (int64,error){
	post:=new(Post)
	return dbEngine.Where("status = ?",0).Count(post)
}

func GetComment(comment *Comment) (bool,error){
	has,err:=dbEngine.Get(comment)
	if err!=nil{
		return false, err
	}
	return has,nil
}

func GetCommentsByPostId(postId int64,page int,pageSize int) ([]*Comment,error){
	var comments []*Comment
	err:=dbEngine.Where("post_id=?",postId).Limit(pageSize,page).Find(&comments)
	return comments,err
}

func GetCommentsCountByPostId(postId int64) (int64,error){
	comment:=&Comment{}
	return dbEngine.Where("post_id=?",postId).Count(comment)
}

func GetRepliesCountByPostId(postId int64) (int64,error){
	reply:=&Reply{}
	return dbEngine.Where("post_id",postId).Count(reply)
}

func GetRepliesCountByCommentId(commentId int64) (int64,error){
	reply:=&Reply{}
	return dbEngine.Where("comment_id",commentId).Count(reply)
}

func GetRepliesByCommentId(commentId int64,page int,pageSize int) ([]*Reply,error){
	replies:=make([]*Reply,pageSize)
	err:=dbEngine.Where("comment_id=?",commentId).Limit(page,pageSize).Find(&replies)
	return replies,err
}

func GetReply(reply *Reply) (bool,error){
	has,err:=dbEngine.Get(reply)
	if err!=nil{
		return false, status.Error(codes.Internal,err.Error())
	}
	return has,nil
}

//通过id获取user,id为post_id,comment_id,reply_id
func GetUsersByIds(ids []int64,idType int) ([]*User,error){
	var sql string
	var orderField []string
	for _,v:=range ids{
		orderField=append(orderField,strconv.FormatInt(v,10))
	}
	switch idType {
	case 0:
		//post
		sql=`SELECT 
				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason 
			FROM user t1 
			LEFT JOIN post t2 
				ON t2.user_id=t1.id 
			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
		`
	case 1:
		//comment
		sql=`SELECT 
				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason 
			FROM user t1 
			LEFT JOIN comment t2 
				ON t2.user_id=t1.id 
			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
		`
	case 2:
		//reply
		sql=`SELECT 
				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason 
			FROM user t1 
			LEFT JOIN reply t2 
				ON t2.user_id=t1.id 
			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
		`
	}
	results,err:=dbEngine.Query(sql,ids,strings.Join(orderField,","))
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	//排序
	var users []*User
	j,l:=0,len(results)
	for _,v:=range ids{
		var u *User
		if j+1>l{
			users=append(users,u)
			continue
		}
		if BytesToInt64(results[j]["id"])==v{
			err:=mapstructure.Decode(v,u)
			if err!=nil{
				return nil, status.Error(codes.Internal,err.Error())
			}
			j++
		}
		users=append(users,u)
	}
	return users, nil
}

//通过comment_ids获取reply
/*
	ids: comment_id list
	limit: 返回每个comment下的前limit个reply,order by create_time asc
 */
func GetRepliesByCommentIds(ids []int64,limit int) ([][]*Reply,error){
	if limit<=0 || limit>10{
		return nil, status.Error(codes.Internal,"query reply_by_comment limit can not be <0 or >10")
	}
	var sql string
	var orderField []string
	for _,v:=range ids{
		orderField=append(orderField,strconv.FormatInt(v,10))
	}
	sql=`
		SELECT 
			t1.id,
			t1.user_id,
			t1.content,
			t1.parent_id,
			t1.create_time,
			t1.floor,
			t1.status 
		FROM 
			reply t1 
			LEFT JOIN reply t2 ON t1.comment_id=t2.comment_id 
			AND t1.create_time > t2.create_time 
		WHERE 
			t1.comment_id IN (?) AND t1.status=0 
		ORDER BY 
			t1.id,
			t1.comment_id 
		HAVING 
			COUNT(t2.id) < ? 
		ORDER BY FIELD(t1.comment_id,?)
	`
	results,err:=dbEngine.Query(sql,ids,limit,strings.Join(orderField,","))
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	//排序
	var replies [][]*Reply
	j,l:=0,len(results)
	for _,v:=range ids{
		var temp []*Reply
		
		for i:=0;i<limit;i++{
			var r *Reply
			if j+1>l{
				temp=append(temp,r)
				continue
			}
			if BytesToInt64(results[j]["id"])==v{
				err:=mapstructure.Decode(results[j],r)
				if err!=nil{
					return nil, status.Error(codes.Internal,err.Error())
				}
				j++
			}
			temp=append(temp,r)
		}
		replies=append(replies,temp)
	}
	return replies, nil
}

//通过post_ids获取comment
/*
	ids: post_id list
	limit: 返回每个post下的前limit个comment,order by create_time asc
*/
func GetCommentsByPostIds(ids []int64,limit int) ([][]*Comment,error){
	if limit<=0 || limit>10{
		return nil, status.Error(codes.Internal,"get comment_by_post limit can not be <0 or >10")
	}
	var sql string
	var orderField []string
	for _,v:=range ids{
		orderField=append(orderField,strconv.FormatInt(v,10))
	}
	sql=`
		SELECT 
			t1.id,
			t1.user_id,
			t1.content,
			t1.create_time,
			t1.floor,
			t1.status 
		FROM 
			comment t1 
			LEFT JOIN comment t2 ON t1.post_id=t2.post_id 
			AND t1.create_time > t2.create_time 
		WHERE 
			t1.post_id IN (?) AND t1.status=0 
		ORDER BY 
			t1.id,
			t1.post_id 
		HAVING 
			COUNT(t2.id) < ? 
		ORDER BY FIELD(t1.post_id,?)
	`
	results,err:=dbEngine.Query(sql,ids,limit,strings.Join(orderField,","))
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	//排序
	var comments [][]*Comment
	j,l:=0,len(results)
	for _,v:=range ids{
		var temp []*Comment

		for i:=0;i<limit;i++{
			var r *Comment
			if j+1>l{
				temp=append(temp,r)
				continue
			}
			if BytesToInt64(results[j]["id"])==v{
				err:=mapstructure.Decode(results[j],r)
				if err!=nil{
					return nil, status.Error(codes.Internal,err.Error())
				}
				j++
			}
			temp=append(temp,r)
		}
		comments=append(comments,temp)
	}
	return comments, nil
}

func GetUsers(ids []int64) ([]*User,error){
	var sql string
	var orderField []string
	for _,v:=range ids{
		orderField=append(orderField,strconv.FormatInt(v,10))
	}
	sql=`
		SELECT 
			id,
			name,
			avatar,
			score,
			role,
			banned_reason,
			create_time,
			last_update 
		FROM 
			user 
		WHERE 
			id IN (?) 
		ORDER BY FIELD(id,?)
	`
	results,err:=dbEngine.Query(sql,ids,strings.Join(orderField,","))
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	//排序
	var users []*User
	j,l:=0,len(results)
	for _,v:=range ids{
		var u *User
		if j+1>l{
			users=append(users,u)
			continue
		}
		if BytesToInt64(results[j]["id"])==v{
			err:=mapstructure.Decode(v,u)
			if err!=nil{
				return nil, status.Error(codes.Internal,err.Error())
			}
			j++
		}
		users=append(users,u)
	}
	return users, nil
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

//[]byte转int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
