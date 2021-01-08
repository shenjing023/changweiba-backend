package repository

// User user table
type User struct {
	ID           int64 `grom:"column:id"`
	Name         string
	Password     string
	Avatar       string
	Status       int64
	Score        int64
	BannedReason string
	CreateTime   int64
	LastUpdate   int64
	Role         uint8
	IP           int64 `grom:"column:ip"`
}

// Avatar user avatar table
type Avatar struct {
	ID     int64  `grom:"column:id"`
	URL    string `grom:"column:url"`
	Status int64
}

// Post 帖子表
type Post struct {
	ID         int64 `grom:"column:id"`
	UserID     int64 `grom:"column:user_id"`
	Topic      string
	CreateTime int64
	LastUpdate int64
	ReplyNum   int64
	Status     int64
}

// Comment 评论表
type Comment struct {
	ID         int64 `grom:"column:id"`
	UserID     int64 `grom:"column:user_id"`
	PostID     int64 `grom:"column:post_id"`
	Content    string
	CreateTime int64
	Floor      int64
	Status     int64
}

// Reply 评论回复表
type Reply struct {
	ID         int64 `grom:"column:id"`
	UserID     int64 `grom:"column:user_id"`
	PostID     int64 `grom:"column:post_id"`
	CommentID  int64 `grom:"column:comment_id"`
	Content    string
	CreateTime int64
	ParentID   int64 `grom:"column:parent_id"`
	Floor      int64
	Status     int64
}
