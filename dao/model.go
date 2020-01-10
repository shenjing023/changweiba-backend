package dao

type User struct {
	Id int64
	Name string
	Password string
	Avatar string
	Status int64
	Score int64
	BannedReason string
	CreateTime int64
	LastUpdate int64
	Role int64
	Ip int64
}

type Avatar struct {
	Id int64
	Url string
	Status int64
}

type Post struct {
	Id int64
	UserId int64
	Topic string
	CreateTime int64
	LastUpdate int64
	ReplyNum int64
	Status int64
}

type Comment struct {
	Id int64
	UserId int64
	PostId int64
	Content string
	CreateTime int64
	Floor int64
	Status int64
}

type Reply struct {
	Id int64
	UserId int64
	PostId int64
	CommentId int64
	Content string
	CreateTime int64
	ParentId int64
	Floor int64
	Status int64
}
