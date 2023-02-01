package user_login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/user_login"
)

type UserLoginResponse struct {
	models.ResponseStatus
	*LoginResponse
}

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserLoginHandler(c *gin.Context) {

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
	}
	//返回对应的错误
	response, error := user_login.NewUserLoginService(username, password).Do()
	if error != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			ResponseStatus: models.ResponseStatus{
				StatusCode: -1,
				StatusMsg:  error.Error(),
			},
		})
	}
	response.ResponseStatus.StatusCode = 0
	c.JSON(http.StatusOK, response)
}
