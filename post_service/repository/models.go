package repository

// Posts 帖子表
type Post struct {
	ID         int64 `grom:"column:id"`
	UserID     int64 `grom:"column:user_id"`
	Topic      string
	CreateTime int64
	LastUpdate int64
	ReplyNum   int64
	Status     int64
}

func (Post) TableName() string {
	return "cw_post"
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

func (Comment) TableName() string {
	return "cw_comment"
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

func (Reply) TableName() string {
	return "cw_reply"
}
