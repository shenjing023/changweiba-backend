package dao

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	redisClient *redis.Client
	dbOrm       *gorm.DB
)

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Cfg.Redis.Host, conf.Cfg.Redis.Port),
		Password: conf.Cfg.Redis.Password,
		DB:       0,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		logs.Error("connect to redis error: ", err)
		os.Exit(1)
	}

	var err error
	dbOrm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Cfg.DB.User, conf.Cfg.DB.Password, conf.Cfg.DB.Host, conf.Cfg.DB.Port, conf.Cfg.DB.Dbname))
	if err != nil {
		logs.Error("mysql connection error: ", err)
	}
	if conf.Cfg.DB.MaxIdle > 0 {
		dbOrm.DB().SetMaxIdleConns(conf.Cfg.DB.MaxIdle)
	}
	if conf.Cfg.DB.MaxOpen > 0 {
		dbOrm.DB().SetMaxOpenConns(conf.Cfg.DB.MaxOpen)
	}
	//日志
	f, err := os.Create(conf.Cfg.DB.LogFile)
	if err != nil {
		logs.Error("create sql log file failed:", err.Error())
		os.Exit(1)
	} else {
		dbOrm.SetLogger(log.New(f, "\r\n", 0))
		if conf.Cfg.Debug {
			dbOrm.LogMode(true)
		}
	}
}

func InsertUser(userName, password, ip, avatar string) (int64, error) {
	//先检查name是否存在
	var u User
	err := dbOrm.Where("name=?", userName).First(&u).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			now := time.Now().Unix()
			user := User{
				Name:       userName,
				Password:   password,
				Ip:         InetAtoi(ip),
				CreateTime: now,
				LastUpdate: now,
				Avatar:     avatar,
			}
			if dbOrm.Create(&user).Error != nil {
				return 0, common.NewDaoErr(common.Internal, err)
			} else {
				return user.Id, nil
			}
		} else {
			return 0, common.NewDaoErr(common.Internal, err)
		}
	}
	return 0, common.NewDaoErr(common.AlreadyExists, err)
}

func GetUser(userId int64) (*User, error) {
	var user User
	if err := dbOrm.First(&user, userId).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &user, nil
}

func CheckUserExist(name string) (*User, bool) {
	var user User
	if exist := dbOrm.Where("name=?", name).First(&user).RecordNotFound(); exist {
		return nil, false
	} else {
		return &user, true
	}
}

//随机获取一个头像url
func GetRandomAvatar() (url string, err error) {
	var avatars []Avatar
	if err = dbOrm.Select("url").Find(&avatars).Error; err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	if len(avatars) == 0 {
		return "", status.Error(codes.Internal, "there are no avatar data in db")
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := seed.Intn(len(avatars))
	url = avatars[index].Url
	return
}

func InsertPost(userId int64, topic string, content string) (int64, error) {
	session := dbOrm.Begin()
	now := time.Now().Unix()
	post := Post{
		UserId:     userId,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	if err := session.Create(&post).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}
	comment := Comment{
		UserId:     userId,
		PostId:     post.Id,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:      1,
	}
	if err := session.Create(&comment).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}

	session.Commit()
	go increasePostReplyNum(post.Id)
	return post.Id, nil
}

func InsertComment(userId int64, postId int64, content string) (int64, error) {
	session := dbOrm.Begin()
	//先获取楼层数
	var floor int64
	sql := "SELECT count(*) AS total FROM comment WHERE post_id=? FOR UPDATE"
	if err := session.Raw(sql, postId).Scan(&floor).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}

	now := time.Now().Unix()
	comment := Comment{
		UserId:     userId,
		PostId:     postId,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:      floor + 1,
	}
	if err := session.Create(&comment).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}

	session.Commit()
	go increasePostReplyNum(postId)
	return comment.Id, nil
}

func InsertReply(userId, postId, commentId, parentId int64, content string) (int64, error) {
	session := dbOrm.Begin()
	//先获取楼层数
	var floor int64
	sql := "SELECT count(*) AS total FROM reply WHERE comment_id=? FOR UPDATE"
	if err := session.Raw(sql, commentId).Scan(&floor).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}

	now := time.Now().Unix()
	reply := Reply{
		UserId:     userId,
		PostId:     postId,
		Content:    content,
		ParentId:   parentId,
		CommentId:  commentId,
		CreateTime: now,
		Status:     0,
		Floor:      floor + 1,
	}
	if err := session.Create(&reply).Error; err != nil {
		session.Rollback()
		return 0, status.Error(codes.Internal, err.Error())
	}

	session.Commit()
	go increasePostReplyNum(postId)
	return reply.Id, nil
}

//帖子回复数+1
func increasePostReplyNum(postId int64) {
	sql := "UPDATE post SET reply_num=reply_num+1,last_update=UNIX_TIMESTAMP() WHERE id=?"
	dbOrm.Exec(sql, postId)
}

func decreasePostReplyNum(postId int64) {
	sql := "UPDATE post SET reply_num=reply_num-1 WHERE id=?"
	dbOrm.Exec(sql, postId)
}

func DeletePost(postId int64) error {
	sql := "UPDATE post SET status=0 WHERE id=?"
	err := dbOrm.Exec(sql, postId).Error
	return status.Error(codes.Internal, err.Error())
}

func DeleteComment(commentId int64) (err error) {
	sql := "UPDATE comment SET status=0 WHERE id=?"
	err = dbOrm.Exec(sql, commentId).Error
	if err == nil {
		go func() {
			var postId int64
			sql = "SELECT post_id FROM comment WHERE id=?"
			if err = dbOrm.Raw(sql, commentId).Scan(&postId).Error; err == nil {
				go decreasePostReplyNum(postId)
			}
		}()
	}
	return err
}

func DeleteReply(replyId int64) (err error) {
	sql := "UPDATE reply SET status=0 WHERE id=?"
	err = dbOrm.Exec(sql, replyId).Error
	if err == nil {
		go func() {
			var postId int64
			sql = "SELECT post_id FROM reply WHERE id=?"
			if err = dbOrm.Raw(sql, replyId).Scan(&postId).Error; err == nil {
				go decreasePostReplyNum(postId)
			}
		}()
	}
	return err
}

func GetPost(postId int64) (*Post, error) {
	var post Post
	if err := dbOrm.First(&post, postId).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &post, nil
}

func GetPosts(page int64, pageSize int64) ([]*Post, int64, error) {
	var (
		posts      []*Post
		totalCount int64
	)
	if err := dbOrm.Raw("select count(*) from post where status=0").Scan(&totalCount).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	if err := dbOrm.Where("status=?", 0).Offset(pageSize * (page - 1)).Limit(pageSize).Order("last_update desc").Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return posts, totalCount, nil
}

func GetComment(commentId int64) (*Comment, error) {
	var comment Comment
	if err := dbOrm.First(&comment, commentId).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &comment, nil
}

func GetCommentsByPostId(postId int64, page int64, pageSize int64) ([]*Comment, int64, error) {
	var (
		comments   []*Comment
		totalCount int64
		tmp        []*Comment
	)
	if err := dbOrm.Where("post_id=?", postId).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Find(&comments).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return comments, totalCount, nil
}

func GetRepliesByCommentId(commentId int64, page int64, pageSize int64) ([]*Reply, int64, error) {
	var (
		replies    []*Reply
		totalCount int64
		tmp        []*Reply
	)
	if err := dbOrm.Where("comment_id=?", commentId).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Find(&replies).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return replies, totalCount, nil
}

func GetReply(replyId int64) (*Reply, error) {
	var reply Reply
	if err := dbOrm.First(&reply, replyId).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &reply, nil
}

//通过id获取user,id为post_id,comment_id,reply_id
//func GetUsersByIds(ids []int64,idType int) ([]*User,error){
//	var sql string
//	var orderField []string
//	for _,v:=range ids{
//		orderField=append(orderField,strconv.FormatInt(v,10))
//	}
//	switch idType {
//	case 0:
//		//post
//		sql=`SELECT
//				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason
//			FROM user t1
//			LEFT JOIN post t2
//				ON t2.user_id=t1.id
//			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
//		`
//	case 1:
//		//comment
//		sql=`SELECT
//				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason
//			FROM user t1
//			LEFT JOIN comment t2
//				ON t2.user_id=t1.id
//			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
//		`
//	case 2:
//		//reply
//		sql=`SELECT
//				t1.id,t1.name,t1.avatar,t1.status,t1.score,t1.role,t1.banned_reason
//			FROM user t1
//			LEFT JOIN reply t2
//				ON t2.user_id=t1.id
//			WHERE t2.id IN (?) ORDER BY FIELD(t2.id,?)
//		`
//	}
//	results,err:=dbEngine.Query(sql,ids,strings.Join(orderField,","))
//	if err!=nil{
//		return nil, status.Error(codes.Internal,err.Error())
//	}
//	//排序
//	var users []*User
//	j,l:=0,len(results)
//	for _,v:=range ids{
//		var u *User
//		if j+1>l{
//			users=append(users,u)
//			continue
//		}
//		if BytesToInt64(results[j]["id"])==v{
//			err:=mapstructure.Decode(v,u)
//			if err!=nil{
//				return nil, status.Error(codes.Internal,err.Error())
//			}
//			j++
//		}
//		users=append(users,u)
//	}
//	return users, nil
//}

//通过comment_ids获取reply
/*
	ids: comment_id list
	limit: 返回每个comment下的前limit个reply,order by create_time asc
*/
//func GetRepliesByCommentIds(ids []int64,limit int) ([][]*Reply,error){
//	if limit<=0 || limit>10{
//		return nil, status.Error(codes.Internal,"query reply_by_comment limit can not be <0 or >10")
//	}
//	var sql string
//	var orderField []string
//	for _,v:=range ids{
//		orderField=append(orderField,strconv.FormatInt(v,10))
//	}
//	sql=`
//		SELECT
//			t1.id,
//			t1.user_id,
//			t1.content,
//			t1.parent_id,
//			t1.create_time,
//			t1.floor,
//			t1.status
//		FROM
//			reply t1
//			LEFT JOIN reply t2 ON t1.comment_id=t2.comment_id
//			AND t1.create_time > t2.create_time
//		WHERE
//			t1.comment_id IN (?) AND t1.status=0
//		ORDER BY
//			t1.id,
//			t1.comment_id
//		HAVING
//			COUNT(t2.id) < ?
//		ORDER BY FIELD(t1.comment_id,?)
//	`
//	results,err:=dbEngine.Query(sql,ids,limit,strings.Join(orderField,","))
//	if err!=nil{
//		return nil, status.Error(codes.Internal,err.Error())
//	}
//	//排序
//	var replies [][]*Reply
//	j,l:=0,len(results)
//	for _,v:=range ids{
//		var temp []*Reply
//
//		for i:=0;i<limit;i++{
//			var r *Reply
//			if j+1>l{
//				temp=append(temp,r)
//				continue
//			}
//			if BytesToInt64(results[j]["id"])==v{
//				err:=mapstructure.Decode(results[j],r)
//				if err!=nil{
//					return nil, status.Error(codes.Internal,err.Error())
//				}
//				j++
//			}
//			temp=append(temp,r)
//		}
//		replies=append(replies,temp)
//	}
//	return replies, nil
//}

//通过post_ids获取comment
/*
	ids: post_id list
	limit: 返回每个post下的前limit个comment,order by create_time asc
*/
//func GetCommentsByPostIds(ids []int64,limit int) ([][]*Comment,error){
//	if limit<=0 || limit>10{
//		return nil, status.Error(codes.Internal,"get comment_by_post limit can not be <0 or >10")
//	}
//	var sql string
//	var orderField []string
//	for _,v:=range ids{
//		orderField=append(orderField,strconv.FormatInt(v,10))
//	}
//	sql=`
//		SELECT
//			t1.id,
//			t1.user_id,
//			t1.content,
//			t1.create_time,
//			t1.floor,
//			t1.status
//		FROM
//			comment t1
//			LEFT JOIN comment t2 ON t1.post_id=t2.post_id
//			AND t1.create_time > t2.create_time
//		WHERE
//			t1.post_id IN (?) AND t1.status=0
//		ORDER BY
//			t1.id,
//			t1.post_id
//		HAVING
//			COUNT(t2.id) < ?
//		ORDER BY FIELD(t1.post_id,?)
//	`
//	results,err:=dbEngine.Query(sql,ids,limit,strings.Join(orderField,","))
//	if err!=nil{
//		return nil, status.Error(codes.Internal,err.Error())
//	}
//	//排序
//	var comments [][]*Comment
//	j,l:=0,len(results)
//	for _,v:=range ids{
//		var temp []*Comment
//
//		for i:=0;i<limit;i++{
//			var r *Comment
//			if j+1>l{
//				temp=append(temp,r)
//				continue
//			}
//			if BytesToInt64(results[j]["id"])==v{
//				err:=mapstructure.Decode(results[j],r)
//				if err!=nil{
//					return nil, status.Error(codes.Internal,err.Error())
//				}
//				j++
//			}
//			temp=append(temp,r)
//		}
//		comments=append(comments,temp)
//	}
//	return comments, nil
//}

func GetUsers(ids []int64) ([]*User, error) {
	sql := `
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
	var results []*User
	if err := dbOrm.Raw(sql, ids, ids).Scan(&results).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	//排序
	var users []*User
	j, l := 0, len(results)
	for _, v := range ids {
		u := &User{}
		if j+1 > l {
			users = append(users, u)
			continue
		}
		if results[j].Id == v {
			users = append(users, results[j])
			j++
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func GetPostsByUserId(userId int64, page int64, pageSize int64) ([]*Post, int64, error) {
	var (
		posts      []*Post
		totalCount int64
		tmp        []*Post
	)
	if err := dbOrm.Where("user_id=?", userId).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Order("create_time desc").Find(&posts).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return posts, totalCount, nil
}

func GetCommentsByUserId(userId int64, page int64, pageSize int64) ([]*Comment, int64, error) {
	var (
		comments   []*Comment
		totalCount int64
		tmp        []*Comment
	)
	if err := dbOrm.Where("user_id=?", userId).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Order("create_time desc").Find(&comments).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return comments, totalCount, nil
}

func GetRepliesByUserId(userId int64, page int64, pageSize int64) ([]*Reply, int64, error) {
	var (
		replies    []*Reply
		totalCount int64
		tmp        []*Reply
	)
	if err := dbOrm.Where("user_id=?", userId).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Order("create_time desc").Find(&replies).Error; err != nil {
		return nil, 0, status.Error(codes.Internal, err.Error())
	}
	return replies, totalCount, nil
}

//ip地址int->string相互转换
func InetAtoi(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func InetItoa(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

//[]byte转int64
func BytesToInt64(buf []byte) int64 {
	r, _ := strconv.ParseInt(string(buf), 10, 64)
	return r
}

func BytesToInt32(buf []byte) int32 {
	r, _ := strconv.ParseInt(string(buf), 10, 32)
	return int32(r)
}
