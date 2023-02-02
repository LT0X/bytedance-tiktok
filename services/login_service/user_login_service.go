package login_service

import (
	"errors"
	"qingxunyin/bytedance-tiktok/middlewares"
	"qingxunyin/bytedance-tiktok/models"
)

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginService struct {
	username string
	password string
	*LoginResponse
}

func NewUserLoginService(username string, password string) *UserLoginService {
	return &UserLoginService{
		username: username,
		password: password,
	}

}

func (u *UserLoginService) Do() (*LoginResponse, error) {

	//检查姓名格式
	err := u.checkUserName()
	if err != nil {
		return nil, err
	}
	//检查用户是否存在
	userLogin, err := models.GetUserLoginDao().QueryUserLogin(u.username, u.password)
	if err != nil {
		return nil, err
	}

	//颁发token和设置信息
	u.LoginResponse = new(LoginResponse)
	u.Token, err = middlewares.GetToken(*userLogin)
	u.UserId = userLogin.Id
	if err != nil {
		return nil, err
	}
	return u.LoginResponse, nil

}

func (u *UserLoginService) checkUserName() error {
	if len(u.username) > 20 {
		return errors.New("用户名长度过长")
	} else {
		return nil
	}
}
