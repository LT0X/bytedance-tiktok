package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/video_service"
	"strconv"
)

type FavoriteActionResponse struct {
	*models.ResponseStatus
}

type FavoriteActionHandler struct {
	FavoriteActionResponse
	*gin.Context
}

func FavoriteActionController(c *gin.Context) {

	handler := &FavoriteActionHandler{
		Context: c,
	}
	//解析参数
	id, _ := c.Get("user_id")
	uid, ok := id.(int64)
	if !ok {
		handler.SendResponse(-1, "解析uid错误")
		return
	}
	i := c.Query("video_id")
	vid, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "解析id错误")
	}
	acton := c.Query("action_type")

	//开始业务逻辑
	handler.ResponseStatus, err = video_service.NewFavoriteActionService(uid, vid, acton).Do()

	//发送响应参数
	handler.SendResponse(0, "success")

}

func (f *FavoriteActionHandler) SendResponse(code int32, msg string) {
	if f.ResponseStatus == nil {
		resp := &models.ResponseStatus{
			StatusCode: code,
			StatusMsg:  msg,
		}
		f.JSON(http.StatusOK, resp)
	} else {
		f.StatusCode = code
		f.StatusMsg = msg
		f.JSON(http.StatusOK, f.ResponseStatus)
	}

}
