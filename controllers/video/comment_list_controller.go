package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/video_service"
	"strconv"
)

type CommentListResponse struct {
	models.ResponseStatus
	*video_service.CommentResponse
}

type CommentListHandler struct {
	VideoId int64
	CommentListResponse
	*gin.Context
}

func CommentListController(c *gin.Context) {

	handler := new(CommentListHandler)
	handler.Context = c
	//解析参数
	id := c.Query("video_id")
	videoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.SendResponse(-1, err.Error())
	}
	handler.VideoId = videoId

	//开始业务逻辑，查询视频评论列表
	handler.CommentResponse, err = video_service.NewCommentListService(handler.VideoId).Do()
	if err != nil {
		handler.SendResponse(-1, err.Error())
	}
	//发送请求
	handler.SendResponse(0, "")

}

func (c CommentListHandler) SendResponse(code int32, msg string) {
	if c.CommentResponse == nil {
		c.JSON(http.StatusOK, models.ResponseStatus{
			StatusCode: code,
			StatusMsg:  msg,
		})
	} else {
		c.StatusCode = code
		c.StatusMsg = msg
		c.JSON(http.StatusOK, c.CommentListResponse)
	}
}
