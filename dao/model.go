package dao

type User struct {
	Id int64	//`xorm:"autoincr"`
	Name string
	Password string
	Avatar string
	Status int
	Score int32
	BannedReason string
	CreateTime int64
	LastUpdate int64
	Role int
	Ip int64
}

type Avatar struct {
	Id int64
	Url string
	Status int
}
