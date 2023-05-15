// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Comment struct {
	ID        int    `json:"id"`
	User      *User  `json:"user"`
	PostID    int    `json:"postId"`
	Content   string `json:"content"`
	CreatedAt int    `json:"createdAt"`
	// 第几楼
	Floor   int              `json:"floor"`
	Status  PostStatus       `json:"status"`
	Replies *ReplyConnection `json:"replies,omitempty"`
}

type CommentConnection struct {
	Nodes      []*Comment `json:"nodes,omitempty"`
	TotalCount int        `json:"totalCount"`
}

type DeletePost struct {
	ID int `json:"id"`
}

type EditUser struct {
	Name     *string     `json:"name,omitempty"`
	Password *string     `json:"password,omitempty"`
	Avatar   *string     `json:"avatar,omitempty"`
	Status   *UserStatus `json:"status,omitempty"`
	Role     *UserRole   `json:"role,omitempty"`
}

type NewComment struct {
	PostID  int    `json:"postId"`
	Content string `json:"content"`
}

type NewPost struct {
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

type NewReply struct {
	PostID    int    `json:"postId"`
	CommentID int    `json:"commentId"`
	Content   string `json:"content"`
	ParentID  *int   `json:"parentId,omitempty"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Post struct {
	ID        int    `json:"id"`
	User      *User  `json:"user"`
	Topic     string `json:"topic"`
	CreatedAt int    `json:"createdAt"`
	// 最后回复时间
	UpdatedAt int `json:"updatedAt"`
	// 帖子评论+回复的总数
	ReplyNum int                `json:"replyNum"`
	Status   PostStatus         `json:"status"`
	Comments *CommentConnection `json:"comments"`
	// 最后评论或回复的用户
	LastReplyUser *User `json:"lastReplyUser"`
	// 一楼的评论，首页会用到
	FirstComment *Comment `json:"firstComment"`
}

type PostConnection struct {
	Nodes      []*Post `json:"nodes,omitempty"`
	TotalCount int     `json:"totalCount"`
}

type Reply struct {
	ID        int    `json:"id"`
	User      *User  `json:"user"`
	CommentID int    `json:"commentId"`
	Content   string `json:"content"`
	CreatedAt int    `json:"createdAt"`
	// 父回复
	Parent *Reply `json:"parent"`
	// 楼中楼的第几楼
	Floor  int        `json:"floor"`
	Status PostStatus `json:"status"`
}

type ReplyConnection struct {
	Nodes      []*Reply `json:"nodes,omitempty"`
	TotalCount int      `json:"totalCount"`
}

type ReportUser struct {
	UserID         int    `json:"userId"`
	ReportedUserID string `json:"reportedUserId"`
	Reason         string `json:"reason"`
}

type Stock struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type StockConnection struct {
	Nodes      []*Stock `json:"nodes,omitempty"`
	TotalCount int      `json:"totalCount"`
}

type TradeDate struct {
	Date   string  `json:"date"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
	Xq     int     `json:"xq"`
}

type TradeDateConnection struct {
	Nodes      []*TradeDate `json:"nodes,omitempty"`
	TotalCount int          `json:"totalCount"`
	ID         int          `json:"id"`
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
	Posts        *PostConnection    `json:"posts,omitempty"`
	Comments     *CommentConnection `json:"comments,omitempty"`
	Replies      *ReplyConnection   `json:"replies,omitempty"`
}

type WencaiStock struct {
	Bull  int    `json:"bull"`
	Short string `json:"short"`
}

type PostStatus string

const (
	PostStatusNormal PostStatus = "NORMAL"
	PostStatusDelete PostStatus = "DELETE"
)

var AllPostStatus = []PostStatus{
	PostStatusNormal,
	PostStatusDelete,
}

func (e PostStatus) IsValid() bool {
	switch e {
	case PostStatusNormal, PostStatusDelete:
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
