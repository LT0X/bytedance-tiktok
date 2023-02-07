package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/video_service"
)

type PublishVideoResponse struct {
	*models.ResponseStatus
}

func PublishVideoController(c *gin.Context) {

	//参数解析
	temp, _ := c.Get("user_id")
	uid, ok := temp.(int64)
	if !ok {
		sendResponse(c, -1, "uid解析错误")
		return
	}
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		sendResponse(c, -1, err.Error())
		return
	}
	//执行业务逻辑
	response, err := video_service.NewPublishVideoService(data, title, uid, c).Do()
	if err != nil {
		sendResponse(c, -1, err.Error())
		return
	}
	//发送响应报告
	sendResponse(c, response.StatusCode, response.StatusMsg)
}

func sendResponse(c *gin.Context, code int32, msg string) {
	response := PublishVideoResponse{
		ResponseStatus: &models.ResponseStatus{
			StatusCode: code,
			StatusMsg:  msg,
		},
	}
	c.JSON(http.StatusOK, response)
}
