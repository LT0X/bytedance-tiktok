package user_login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/login_service"
)

type UserRegisterResponse struct {
	models.ResponseStatus
	*login_service.RegisterResponse
}

func UserRegisterHandler(c *gin.Context) {

	response := new(UserRegisterResponse)
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
	//进行业务逻辑，进行用户注册
	response.RegisterResponse, err = login_service.
		NewRegisterService(username, password).Do()
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
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(http.StatusOK, response)
}
