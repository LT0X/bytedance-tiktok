package models

import (
	"errors"
	"qingxunyin/bytedance-tiktok/util/dbutil"
)

// UserLogin 用户登录信息，
type UserLogin struct {
	Id         int64
	UserInfoId int64
	Username   string
	Password   string
}

type UserLoginDao struct {
}

var (
	userLoginDao *UserLoginDao
)

func GetUserLoginDao() *UserLoginDao {
	return userLoginDao
}

func (*UserLoginDao) AddUserLogin(userLogin *UserLogin) error {
	if userLogin == nil {
		return ErrNullPointer
	}
	DB := dbutil.GetDB()
	return DB.Create(userLogin).Error
}

func (*UserLoginDao) IsUsernameExist(username string) bool {
	DB := dbutil.GetDB()
	userLogin := UserLogin{}
	DB.Table("user_logins").Where("username = ?", username).
		First(&userLogin)
	if userLogin.Id == 0 {
		return false
	} else {
		return true
	}
}

func (*UserLoginDao) QueryUserLogin(username string, password string) (*UserLogin, error) {
	DB := dbutil.GetDB()
	userLogin := UserLogin{}
	DB.Table("user_logins").
		Where("username = ? and password = ?", username, password).
		First(&userLogin)
	if userLogin.Id == 0 {
		return nil, errors.New("用户和密码不存在")
	}
	return &userLogin, nil

}
