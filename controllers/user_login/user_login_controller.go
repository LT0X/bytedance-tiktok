package user_login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/login_service"
)

type UserLoginResponse struct {
	models.ResponseStatus
	*login_service.LoginResponse
}

func UserLoginHandler(c *gin.Context) {

	response := new(UserLoginResponse)
	var err error
	//解析参数
	username := c.Query("username")
	temp, _ := c.Get("password")
	password, ok := temp.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			ResponseStatus: models.ResponseStatus{
				StatusCode: -1,
				StatusMsg:  "密码解析出错",
			},
		})
		return
	}
	//进行业务逻辑,进行登录验证
	response.LoginResponse, err = login_service.
		NewUserLoginService(username, password).Do()
	//返回对应的错误
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			ResponseStatus: models.ResponseStatus{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//返回响应参数
	response.ResponseStatus.StatusCode = 0
	c.JSON(http.StatusOK, response)
	return
}
