// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Comment struct {
	ID       int    `json:"id"`
	User     *User  `json:"user"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	CreateAt int    `json:"create_at"`
	// 第几楼
	Floor   int              `json:"floor"`
	Status  Status           `json:"status"`
	Replies *ReplyConnection `json:"replies"`
}

type CommentConnection struct {
	Nodes      []*Comment `json:"nodes"`
	TotalCount int        `json:"total_count"`
}

type DeletePost struct {
	ID int `json:"id"`
}

type EditUser struct {
	Name     *string     `json:"name"`
	Password *string     `json:"password"`
	Avatar   *string     `json:"avatar"`
	Status   *UserStatus `json:"status"`
	Role     *UserRole   `json:"role"`
}

type NewComment struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

type NewPost struct {
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

type NewReply struct {
	PostID    int    `json:"post_id"`
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	ParentID  int    `json:"parent_id"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Post struct {
	ID       int    `json:"id"`
	User     *User  `json:"user"`
	Topic    string `json:"topic"`
	CreateAt int    `json:"create_at"`
	// 最后回复时间
	LastAt int `json:"last_at"`
	// 帖子评论+回复的总数
	ReplyNum int                `json:"reply_num"`
	Status   Status             `json:"status"`
	Comments *CommentConnection `json:"comments"`
	// 最后评论或回复的用户
	LastReplyUser *User `json:"last_reply_user"`
}

type PostConnection struct {
	Nodes      []*Post `json:"nodes"`
	TotalCount int     `json:"total_count"`
}

type Reply struct {
	ID        int    `json:"id"`
	User      *User  `json:"user"`
	PostID    int    `json:"post_id"`
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	CreateAt  int    `json:"create_at"`
	// 父回复
	Parent *Reply `json:"parent"`
	// 楼中楼的第几楼
	Floor  int    `json:"floor"`
	Status Status `json:"status"`
}

type ReplyConnection struct {
	Nodes      []*Reply `json:"nodes"`
	TotalCount int      `json:"total_count"`
}

type ReportUser struct {
	UserID         int    `json:"user_id"`
	ReportedUserID string `json:"reported_user_id"`
	Reason         string `json:"reason"`
}

type User struct {
	ID       int    `json:"id"`
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
	BannedReason string             `json:"banned_reason"`
	Posts        *PostConnection    `json:"posts"`
	Comments     *CommentConnection `json:"comments"`
	Replies      *ReplyConnection   `json:"replies"`
}

type Status string

const (
	StatusNormal Status = "NORMAL"
	StatusBanned Status = "BANNED"
)

var AllStatus = []Status{
	StatusNormal,
	StatusBanned,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusNormal, StatusBanned:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
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