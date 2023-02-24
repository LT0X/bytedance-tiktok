package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/video_service"
	"strconv"
)

type CommentActionResponse struct {
	models.ResponseStatus
	*video_service.ActionResponse
}

type CommentActionHandler struct {
	CommentActionResponse
	*gin.Context
}

func CommentActionController(c *gin.Context) {

	handler := &CommentActionHandler{
		Context: c,
	}
	//解析参数
	temp, _ := c.Get("user_id")
	uid, ok := temp.(int64)
	if !ok {
		handler.sendResponse(-1, "解析uid错误")
		return
	}
	id := c.Query("video_id")
	vid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.sendResponse(-1, "uid解析错误")
		return
	}
	id = c.Query("comment_id")
	var cid int64
	if id != "" {
		cid, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			handler.sendResponse(-1, "uid解析错误")
			return
		}
	} else {
		cid = 0
	}

	actionType := c.Query("action_type")
	comment := c.Query("comment_text")
	//开始进行业务逻辑
	handler.ActionResponse, err = video_service.
		NewCommentActionService(uid, vid, actionType, comment, cid).Do()
	if err != nil {
		handler.sendResponse(-1, err.Error())
	}

	//发送响应数据
	handler.sendResponse(0, "")

}

func (c *CommentActionHandler) sendResponse(code int32, msg string) {

	if c.ActionResponse == nil {
		c.StatusCode = code
		c.StatusMsg = msg
		c.JSON(http.StatusOK, c.ResponseStatus)
	} else {
		c.StatusCode = code
		c.StatusMsg = msg
		c.JSON(http.StatusOK, c.CommentActionResponse)
	}

}
