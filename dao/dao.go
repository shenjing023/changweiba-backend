package dao

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/pkg/logs"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql驱动
	"github.com/mitchellh/mapstructure"
)

var (
	redisClient *redis.Client
	dbOrm       *gorm.DB
)

/*
* 数据库连接初始化
 */
func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Cfg.Redis.Host, conf.Cfg.Redis.Port),
		Password: conf.Cfg.Redis.Password,
		DB:       0,
	})
	if _, err := redisClient.Ping(redisClient.Context()).Result(); err != nil {
		fmt.Printf("connect to redis error: %+v", err)
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
	if f, err := os.Create(conf.Cfg.DB.LogFile); err != nil {
		logs.Error("create sql log file failed:", err.Error())
		os.Exit(1)
	} else {
		// 待优化
		dbOrm.SetLogger(log.New(f, "\r\n", 0))
		if conf.Cfg.Debug {
			dbOrm.LogMode(true)
		}
	}
	// 全局禁用表名复数形式
	dbOrm.SingularTable(true)
}

//InsertUser 插入用户
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
				return 0, common.NewDaoErr(common.Internal, errors.WithStack(err))
			}
			return user.Id, nil
		}
		return 0, common.NewDaoErr(common.Internal, errors.WithStack(err))
	}
	return 0, common.NewDaoErr(common.AlreadyExists, errors.New("user already exist"))
}

// GetUser 根据id获取用户
func GetUser(userID int64) (*User, error) {
	var user User
	if err := dbOrm.First(&user, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, common.NewDaoErr(common.NotFound, err)
		}
		return nil, common.NewDaoErr(common.Internal, err)
	}
	return &user, nil
}

// CheckUserExist 检查用户是否存在,有则返回相关信息
func CheckUserExist(name string) (*User, bool, error) {
	var user User
	// 先检查redis,hgetall的返回没有redis.Nil,先用exists
	if val, err := redisClient.Exists(redisClient.Context(), "user_"+name).Result(); err != nil {
		return nil, false, common.NewDaoErr(common.Internal, err)
	} else {
		if val == 1 {
			//存在
			if val, err := redisClient.HGetAll(redisClient.Context(), "user_"+name).Result(); err != nil {
				return nil, false, common.NewDaoErr(common.Internal, err)
			} else {
				if err = mapstructure.WeakDecode(val, &user); err != nil {
					logs.Error(errors.Wrap(err, "ss"))
					return nil, false, common.NewDaoErr(common.Internal, err)
				}
				return &user, true, nil
			}
		}
	}
	fmt.Println(1111)

	//redis不存在
	if err := dbOrm.Where("name=?", name).First(&user).Error; err != nil {
		return nil, false, common.NewDaoErr(common.Internal, err)
	} else if gorm.IsRecordNotFoundError(err) {
		return nil, false, common.NewDaoErr(common.NotFound, err)
	}
	if err := redisClient.HSet(redisClient.Context(), "user_"+name, map[string]interface{}{
		"id":       user.Id,
		"name":     user.Name,
		"password": user.Password,
		"role":     user.Role,
	}).Err(); err != nil {

	}
	return &user, true, nil
}

//GetRandomAvatar 随机获取一个头像url
func GetRandomAvatar() (url string, err error) {
	var avatars []Avatar
	if err = dbOrm.Select("url").Find(&avatars).Error; err != nil {
		return "", common.NewDaoErr(common.Internal, errors.WithStack(err))
	}
	if len(avatars) == 0 {
		return "", common.NewDaoErr(common.Internal, errors.New("there are no avatar data in db"))
	}
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := seed.Intn(len(avatars))
	url = avatars[index].Url
	return
}

// InsertPost 插入新帖子
func InsertPost(userID int64, topic string, content string) (int64, error) {
	session := dbOrm.Begin()
	now := time.Now().Unix()
	post := Post{
		UserId:     userID,
		Topic:      topic,
		CreateTime: now,
		LastUpdate: now,
		ReplyNum:   0,
		Status:     0,
	}
	if err := session.Create(&post).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}
	comment := Comment{
		UserId:     userID,
		PostId:     post.Id,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:      1,
	}
	if err := session.Create(&comment).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}

	session.Commit()
	go increasePostReplyNum(post.Id)
	return post.Id, nil
}

// InsertComment 对某个帖子插入新评论
func InsertComment(userID int64, postID int64, content string) (int64, error) {
	// 先检查帖子是否存在或删除,待优化，可redis
	var count int64
	if err := dbOrm.Model(&Post{}).Where("id = ? AND status = 0", postID).Count(&count).Error; err != nil {
		return 0, common.NewDaoErr(common.Internal, err)
	}
	if count == 0 {
		return 0, common.NewDaoErr(common.InvalidArgument, errors.New("该帖子不存在或已删除"))
	}

	session := dbOrm.Begin()
	//先获取楼层数
	type Floor struct {
		Total int64
	}
	var floor Floor
	sql := "SELECT count(*) AS total FROM comment WHERE post_id=? FOR UPDATE"
	if err := session.Raw(sql, postID).Scan(&floor).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}

	now := time.Now().Unix()
	comment := Comment{
		UserId:     userID,
		PostId:     postID,
		Content:    content,
		CreateTime: now,
		Status:     0,
		Floor:      floor.Total + 1,
	}
	if err := session.Create(&comment).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}

	session.Commit()
	go increasePostReplyNum(postID)
	return comment.Id, nil
}

// InsertReply 给评论插入回复
func InsertReply(userID, postID, commentID, parentID int64, content string) (int64, error) {
	// 先检查参数
	var count int64
	if err := dbOrm.Model(&Post{}).Where("id = ? AND status = 0", postID).Count(&count).Error; err != nil {
		return 0, common.NewDaoErr(common.Internal, err)
	}
	if count == 0 {
		return 0, common.NewDaoErr(common.InvalidArgument, errors.New("该帖子不存在或已删除"))
	}
	count = 0
	if err := dbOrm.Model(&Comment{}).Where("id = ? AND status = 0", commentID).Count(&count).Error; err != nil {
		return 0, common.NewDaoErr(common.Internal, err)
	}
	if count == 0 {
		return 0, common.NewDaoErr(common.InvalidArgument, errors.New("该评论不存在或已删除"))
	}
	if parentID == 0 {
		count = 0
		if err := dbOrm.Model(&Reply{}).Where("id = ? AND status = 0", parentID).Count(&count).Error; err != nil {
			return 0, common.NewDaoErr(common.Internal, err)
		}
		if count == 0 {
			return 0, common.NewDaoErr(common.InvalidArgument, errors.New("该回复不存在或已删除"))
		}
	}

	session := dbOrm.Begin()
	//先获取楼层数
	type Floor struct {
		Total int64
	}
	var floor Floor
	sql := "SELECT count(*) AS total FROM reply WHERE comment_id=? FOR UPDATE"
	if err := session.Raw(sql, commentID).Scan(&floor).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}

	now := time.Now().Unix()
	reply := Reply{
		UserId:     userID,
		PostId:     postID,
		Content:    content,
		ParentId:   parentID,
		CommentId:  commentID,
		CreateTime: now,
		Status:     0,
		Floor:      floor.Total + 1,
	}
	if err := session.Create(&reply).Error; err != nil {
		session.Rollback()
		return 0, common.NewDaoErr(common.Internal, err)
	}

	session.Commit()
	go increasePostReplyNum(postID)
	return reply.Id, nil
}

//帖子回复数+1
func increasePostReplyNum(postID int64) {
	sql := "UPDATE post SET reply_num=reply_num+1,last_update=UNIX_TIMESTAMP() WHERE id=?"
	dbOrm.Exec(sql, postID)
}

func decreasePostReplyNum(postID int64) {
	sql := "UPDATE post SET reply_num=reply_num-1,last_update=UNIX_TIMESTAMP() WHERE id=?"
	dbOrm.Exec(sql, postID)
}

// DeletePost 删除帖子
func DeletePost(postID int64) error {
	sql := "UPDATE post SET status=0 WHERE id=?"
	err := dbOrm.Exec(sql, postID).Error
	if err != nil {
		return common.NewDaoErr(common.Internal, err)
	}
	return nil
}

func DeleteComment(commentID int64) (err error) {
	sql := "UPDATE comment SET status=0 WHERE id=?"
	err = dbOrm.Exec(sql, commentID).Error
	if err == nil {
		go func() {
			var postID int64
			sql = "SELECT post_id FROM comment WHERE id=?"
			if err = dbOrm.Raw(sql, commentID).Scan(&postID).Error; err == nil {
				go decreasePostReplyNum(postID)
			}
		}()
		return nil
	}
	return common.NewDaoErr(common.Internal, err)
}

func DeleteReply(replyID int64) (err error) {
	sql := "UPDATE reply SET status=0 WHERE id=?"
	err = dbOrm.Exec(sql, replyID).Error
	if err == nil {
		go func() {
			var postId int64
			sql = "SELECT post_id FROM reply WHERE id=?"
			if err = dbOrm.Raw(sql, replyID).Scan(&postId).Error; err == nil {
				go decreasePostReplyNum(postId)
			}
		}()
		return nil
	}
	return common.NewDaoErr(common.Internal, err)
}

func GetPost(postID int64) (*Post, error) {
	var post Post
	if err := dbOrm.First(&post, postID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, common.NewDaoErr(common.NotFound, err)
		}
		return nil, common.NewDaoErr(common.Internal, err)
	}
	return &post, nil
}

func GetPosts(page int64, pageSize int64) ([]*Post, int64, error) {
	var (
		posts      []*Post
		totalCount int64
	)
	if err := dbOrm.Raw("select count(*) from post where status=0").Scan(&totalCount).Error; err != nil {
		return nil, 0, common.NewDaoErr(common.Internal, err)
	}
	if err := dbOrm.Where("status=?", 0).Offset(pageSize * (page - 1)).Limit(pageSize).Order("last_update desc").Error; err != nil {
		return nil, 0, common.NewDaoErr(common.Internal, err)
	}
	return posts, totalCount, nil
}

func GetComment(commentID int64) (*Comment, error) {
	var comment Comment
	if err := dbOrm.First(&comment, commentID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, common.NewDaoErr(common.NotFound, err)
		}
		return nil, common.NewDaoErr(common.Internal, err)
	}
	return &comment, nil
}

func GetCommentsByPostId(postID int64, page int64, pageSize int64) ([]*Comment, int64, error) {
	var (
		comments   []*Comment
		totalCount int64
		tmp        []*Comment
	)
	if err := dbOrm.Where("post_id=?", postID).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Find(&comments).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, common.NewDaoErr(common.NotFound, err)
		}
		return nil, 0, common.NewDaoErr(common.Internal, err)
	}
	return comments, totalCount, nil
}

func GetRepliesByCommentId(commentID int64, page int64, pageSize int64) ([]*Reply, int64, error) {
	var (
		replies    []*Reply
		totalCount int64
		tmp        []*Reply
	)
	if err := dbOrm.Where("comment_id=?", commentID).Find(&tmp).Count(&totalCount).Limit(pageSize).Offset(pageSize * (page - 1)).Find(&replies).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, common.NewDaoErr(common.NotFound, err)
		}
		return nil, 0, common.NewDaoErr(common.Internal, err)
	}
	return replies, totalCount, nil
}

func GetReply(replyId int64) (*Reply, error) {
	var reply Reply
	if err := dbOrm.First(&reply, replyId).Error; err != nil {
		return nil, common.NewDaoErr(common.Internal, err)
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

// GetUsers 获取多个user
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
		return nil, common.NewDaoErr(common.Internal, err)
	}
	//可能有的id不存在,需要再排序
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
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, common.NewDaoErr(common.NotFound, err)
		}
		return nil, 0, common.NewDaoErr(common.Internal, err)
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
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, common.NewDaoErr(common.NotFound, err)
		}
		return nil, 0, common.NewDaoErr(common.Internal, err)
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
		return nil, 0, common.NewDaoErr(common.Internal, err)
	}
	return replies, totalCount, nil
}

//CheckUserIP 检测user表中的ip是否超过次数，防止恶意注册用户
func CheckUserIP(ip string) (bool, error) {
	var (
		tmp        []*User
		totalCount int64
	)
	if err := dbOrm.Where("ip=?", InetAtoi(ip)).Find(&tmp).Count(&totalCount).Error; err != nil {
		return false, common.NewDaoErr(common.Internal, err)
	}
	if totalCount > 3 {
		return false, nil
	}
	return true, nil
}

//InetAtoi ip地址int->string相互转换
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
