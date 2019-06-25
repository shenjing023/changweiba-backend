// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type Comment struct {
	ID       string `json:"id"`
	User     *User  `json:"user"`
	PostID   string `json:"post_id"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
	// 第几楼
	Floor   int        `json:"floor"`
	Status  PostStatus `json:"status"`
	Replies []*Reply   `json:"replies"`
}

type EditPost struct {
	ID     string     `json:"id"`
	Status PostStatus `json:"status"`
}

type EditUser struct {
	ID       string      `json:"id"`
	Name     *string     `json:"name"`
	Password *string     `json:"password"`
	Avatar   *string     `json:"avatar"`
	Status   *UserStatus `json:"status"`
	Role     *UserRole   `json:"role"`
	Score    *int        `json:"score"`
}

type NewComment struct {
	UserID  string `json:"user_id"`
	PostID  string `json:"post_id"`
	Content string `json:"content"`
}

type NewPost struct {
	UserID  string `json:"user_id"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

type NewReply struct {
	UserID      string    `json:"user_id"`
	PostID      string    `json:"post_id"`
	CommentID   string    `json:"comment_id"`
	Content     string    `json:"content"`
	ReplyUserID string    `json:"reply_user_id"`
	Type        ReplyType `json:"type"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Post struct {
	ID       string `json:"id"`
	User     *User  `json:"user"`
	Topic    string `json:"topic"`
	CreateAt string `json:"create_at"`
	// 最后回复时间
	LastAt string `json:"last_at"`
	// 帖子回复的数量
	Reply    int        `json:"reply"`
	Status   PostStatus `json:"status"`
	Comments []*Comment `json:"comments"`
}

type Reply struct {
	ID        string `json:"id"`
	User      *User  `json:"user"`
	PostID    string `json:"post_id"`
	CommentID string `json:"comment_id"`
	Content   string `json:"content"`
	CreateAt  string `json:"create_at"`
	// 回复谁
	ReplyUserID string `json:"reply_user_id"`
	// 楼中楼的第几楼
	Floor int `json:"floor"`
	// 回复类型
	Type ReplyType `json:"type"`
}

type ReportUser struct {
	UserID         string `json:"user_id"`
	ReportedUserID string `json:"reported_user_id"`
	Reason         string `json:"reason"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	// 头像
	Avatar string `json:"avatar"`
	// 状态
	Status UserStatus `json:"status"`
	// 用户角色
	Role UserRole `json:"role"`
	// 当前分数
	Score int `json:"score"`
	// 被封原因
	BannedReason string  `json:"banned_reason"`
	Posts        []*Post `json:"posts"`
}

type PostStatus string

const (
	PostStatusNormal PostStatus = "NORMAL"
	PostStatusBanned PostStatus = "BANNED"
)

var AllPostStatus = []PostStatus{
	PostStatusNormal,
	PostStatusBanned,
}

func (e PostStatus) IsValid() bool {
	switch e {
	case PostStatusNormal, PostStatusBanned:
		return true
	}
	return false
}

func (e PostStatus) String() string {
	return string(e)
}

func (e *PostStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PostStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PostStatus", str)
	}
	return nil
}

func (e PostStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReplyType string

const (
	ReplyTypeReplycomment ReplyType = "REPLYCOMMENT"
	ReplyTypeReplyfloor   ReplyType = "REPLYFLOOR"
)

var AllReplyType = []ReplyType{
	ReplyTypeReplycomment,
	ReplyTypeReplyfloor,
}

func (e ReplyType) IsValid() bool {
	switch e {
	case ReplyTypeReplycomment, ReplyTypeReplyfloor:
		return true
	}
	return false
}

func (e ReplyType) String() string {
	return string(e)
}

func (e *ReplyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReplyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReplyType", str)
	}
	return nil
}

func (e ReplyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleNormal UserRole = "NORMAL"
	UserRoleAdmin  UserRole = "ADMIN"
)

var AllUserRole = []UserRole{
	UserRoleNormal,
	UserRoleAdmin,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleNormal, UserRoleAdmin:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusNormal UserStatus = "NORMAL"
	UserStatusBanned UserStatus = "BANNED"
)

var AllUserStatus = []UserStatus{
	UserStatusNormal,
	UserStatusBanned,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusNormal, UserStatusBanned:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
