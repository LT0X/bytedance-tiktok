package user_follow

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/follows_service"
	"strconv"
)

type FollowsListResp struct {
	StatusCode int64
	StatusMsg  string
	List       *[]models.UserInfo
}

func GetFollowsList(c *gin.Context) {
	//解析参数
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	//解析出错情况
	if err != nil {
		c.JSON(http.StatusOK, FollowsListResp{
			StatusCode: -1,
			StatusMsg:  "解析错误",
			List:       nil,
		})
	}
	//处理请求
	service := follows_service.NewFollowsListService(userId)
	list, err1 := service.Do()
	if err1 != nil {
		c.JSON(http.StatusOK, FollowsListResp{
			StatusCode: -1,
			StatusMsg:  "获取关注列表失败",
			List:       nil,
		})
	}

	c.JSON(http.StatusOK, FollowsListResp{
		StatusCode: 1,
		StatusMsg:  "success",
		List:       list,
	})
}
