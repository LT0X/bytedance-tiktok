package models

// UserLogin 用户登录信息，
type UserLogin struct {
	Id         int64
	UserInfoId int64
	Username   string
	Password   string
}
