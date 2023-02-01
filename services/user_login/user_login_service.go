package user_login

import (
	"errors"
	"qingxunyin/bytedance-tiktok/controllers/user_login"
	"qingxunyin/bytedance-tiktok/middlewares"
	"qingxunyin/bytedance-tiktok/models"
)

type UserLoginService struct {
	username string
	password string
	*user_login.UserLoginResponse
}

func NewUserLoginService(username string, password string) *UserLoginService {
	return &UserLoginService{
		username: username,
		password: password,
	}

}

func (u *UserLoginService) Do() (*user_login.UserLoginResponse, error) {

	//检查姓名格式
	err := u.checkUserName()
	if err != nil {
		return nil, err
	}
	//检查用户是否存在
	userLogin, err := models.GetUserLoginDao().QueryUserLogin(u.username, u.password)
	if err == nil {
		return nil, err
	}
	//颁发token和设置信息
	u.Token, err = middlewares.GetToken(*userLogin)
	u.UserId = userLogin.Id
	if err != nil {
		return nil, err
	}
	return u.UserLoginResponse, nil

}

func (u *UserLoginService) checkUserName() error {
	if len(u.username) > 20 {
		return errors.New("用户名长度过长")
	} else {
		return nil
	}
}
