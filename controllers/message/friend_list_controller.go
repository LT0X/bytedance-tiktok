package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
)

type FriendListResponse struct {
	models.ResponseStatus
	UserList []FriendUser `json:"user_list"`
}

type FriendUser struct {
	models.UserInfo
	Message string `json:"message"`
	MsgType int64  `json:"msgType"`
}

// FriendList 测试用随便写的
func FriendList(c *gin.Context) {

	Friend := new(FriendUser)
	user, _ := models.GetUserInfoDao().QueryUserInfoById(25)
	Friend.UserInfo = *user
	Friend.Message = "sdfsdf"
	Friend.MsgType = 0
	c.JSON(http.StatusOK, FriendListResponse{
		ResponseStatus: models.ResponseStatus{
			StatusCode: 0,
		},
		UserList: []FriendUser{*Friend},
	})
}
