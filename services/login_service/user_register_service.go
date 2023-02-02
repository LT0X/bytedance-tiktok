package login_service

import (
	"errors"
	"qingxunyin/bytedance-tiktok/middlewares"
	"qingxunyin/bytedance-tiktok/models"
)

type RegisterResponse struct {
	UserId int64
	Token  string
}

type UserRegisterService struct {
	username string
	password string
	*RegisterResponse
}

func NewRegisterService(username string, password string) *UserRegisterService {
	return &UserRegisterService{
		username: username,
		password: password,
	}
}

func (u *UserRegisterService) Do() (*RegisterResponse, error) {

	//检查姓名格式
	err := u.checkUserName()
	if err != nil {
		return nil, err
	}

	//检查用户名是否存在
	judge := models.GetUserLoginDao().IsUsernameExist(u.username)
	if judge {
		return nil, errors.New("用户名已经存在")
	}
	//更新注册信息
	userLogin := models.UserLogin{
		Username: u.username,
		Password: u.password,
	}

	err = models.GetUserLoginDao().AddRegisterUser(&userLogin)
	if err != nil {
		return nil, err
	}
	//颁发token和设置信息
	u.RegisterResponse = new(RegisterResponse)
	u.Token, err = middlewares.GetToken(userLogin)
	u.UserId = userLogin.Id
	if err != nil {
		return nil, err
	}
	return u.RegisterResponse, nil
}

func (u *UserRegisterService) checkUserName() error {
	if len(u.username) > 20 {
		return errors.New("用户名长度过长")
	} else {
		return nil
	}
}
