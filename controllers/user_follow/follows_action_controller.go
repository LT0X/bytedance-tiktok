package user_follow

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/services/follows_service"
	"strconv"
)

type ActionResp struct {
	StatusCode int64
	StatusMsg  string
}

func FollowsAction(c *gin.Context) {
	//解析参数
	userId, err1 := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	toUserId, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	//解析出错情况
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, ActionResp{
			StatusCode: -1,
			StatusMsg:  "解析错误",
		})
	}
	//处理请求
	followsService := follows_service.NewFollowsService(userId, toUserId, actionType)
	err := followsService.Do()
	if err != nil {
		c.JSON(http.StatusOK, ActionResp{
			StatusCode: -1,
			StatusMsg:  "解析错误",
		})
	}

	c.JSON(http.StatusOK, ActionResp{
		StatusCode: 1,
		StatusMsg:  "success",
	})
}
