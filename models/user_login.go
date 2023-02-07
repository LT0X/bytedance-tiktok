package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"qingxunyin/bytedance-tiktok/config"
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

// GetUserLoginDao 单例返回对象
func GetUserLoginDao() *UserLoginDao {
	return userLoginDao
}

// AddUserLogin 更新user_login表
func (*UserLoginDao) AddUserLogin(userLogin *UserLogin) error {
	if userLogin == nil {
		return ErrNullPointer
	}
	DB := dbutil.GetDB()
	return DB.Create(userLogin).Error
}

// AddRegisterUser 添加注册用户
func (*UserLoginDao) AddRegisterUser(userLogin *UserLogin) error {
	//开启事务进行多表插入
	tx := dbutil.GetDB().Begin()
	tx.Begin()
	userInfo := UserInfo{
		Name: userLogin.Username,
		Avatar: fmt.Sprintf("http://%s:%d/static/avatar/%s",
			config.GetConf().Server.IP, config.GetConf().Server.Port, "user.png"),
	}

	err := GetUserInfoDao().AddUserInfoDao(&userInfo)
	//发生错误，回滚事务
	if err != nil {
		tx.Rollback()
		return err
	}
	userLogin.UserInfoId = userInfo.Id
	err = GetUserLoginDao().AddUserLogin(userLogin)
	//发生错误，回滚事务
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
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
		Where("username = ?", username).
		First(&userLogin)
	judge := equalsPassword(password, userLogin.Password)
	if userLogin.Id == 0 || judge {
		return nil, errors.New("用户或密码不正确")
	}
	return &userLogin, nil

}

func equalsPassword(password, encryptPassword string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	// 对比密码是否正确会返回一个异常，只要异常是 nil 就证明密码正确
	return err == nil
}
