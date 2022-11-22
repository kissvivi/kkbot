package kkbot

//打卡

// User       FmUser        `json:"user"`
type FmUserInfo struct {
	FmUser
	UserName string `gorm:"column:username;type:varchar(200);not null" json:"username"`
}

type dd struct {
}
