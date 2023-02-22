package user_follow

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/follows_service"
	"strconv"
)

type FollowersListResp struct {
	StatusCode int64              `json:"status_code"`
	StatusMsg  string             `json:"status_msg"`
	List       *[]models.UserInfo `json:"user_list"`
}

func GetFollowersList(c *gin.Context) {
	//解析参数
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	//解析出错情况
	if err != nil {
		c.JSON(http.StatusOK, FollowersListResp{
			StatusCode: -1,
			StatusMsg:  "解析错误",
			List:       nil,
		})
	}
	//处理请求
	service := follows_service.NewFollowersListService(userId)
	list, err1 := service.Do()
	if err1 != nil {
		c.JSON(http.StatusOK, FollowersListResp{
			StatusCode: -1,
			StatusMsg:  "获取粉丝列表失败",
			List:       nil,
		})
	}

	c.JSON(http.StatusOK, FollowersListResp{
		StatusCode: 1,
		StatusMsg:  "success",
		List:       list,
	})
}
