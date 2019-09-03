package dao

type User struct {
	Id int64	//`xorm:"autoincr"`
	Name string
	Password string
	Avatar string
	Status int32
	Score int32
	BannedReason string
	CreateTime int64
	LastUpdate int64
	Role int32
	Ip int64
}

type Avatar struct {
	Id int64
	Url string
	Status int32
}

type Post struct {
	Id int64
	UserId int64
	Topic string
	CreateTime int64
	LastUpdate int64
	ReplyNum int32
	Status int32
}

type Comment struct {
	Id int64
	UserId int64
	PostId int64
	Content string
	CreateTime int64
	Floor int32
	Status int32
}

type Reply struct {
	Id int64
	UserId int64
	PostId int64
	CommentId int64
	Content string
	CreateTime int64
	ParentId int64
	Floor int32
	Status int32
}
