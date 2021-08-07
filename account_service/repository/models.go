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
}

func (User) TableName() string {
	return "cw_user"
}

// Avatar user avatar table
type Avatar struct {
	ID     int64  `grom:"column:id"`
	URL    string `grom:"column:url"`
	Status int64
}

func (Avatar) TableName() string {
	return "cw_avatar"
}
