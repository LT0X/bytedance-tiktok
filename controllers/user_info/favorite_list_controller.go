package user_info

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/info_service"
	"strconv"
)

type UserFavoriteResponse struct {
	models.ResponseStatus
	*info_service.FavoriteResponse
}

type FavoriteListHandler struct {
	uid int64
	UserFavoriteResponse
	*gin.Context
}

func FavoriteListController(c *gin.Context) {
	handler := &FavoriteListHandler{
		Context: c,
	}
	//解析参数
	id := c.Query("user_id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "解析uid错误")
		return
	}

	//开始业务逻辑,查询用户喜欢的视频,
	handler.FavoriteResponse, err = info_service.NewFavoriteListService(uid).Do()
	if err != nil {
		handler.SendResponse(-1, err.Error())
		return
	}
	//开始发送响应数据
	handler.SendResponse(0, "success")

}

func (f *FavoriteListHandler) SendResponse(code int32, msg string) {
	if f.FavoriteResponse == nil {
		f.StatusCode = code
		f.StatusMsg = msg
		f.JSON(http.StatusOK, f.ResponseStatus)
	} else {
		f.StatusCode = code
		f.StatusMsg = msg
		f.JSON(http.StatusOK, f.UserFavoriteResponse)
	}

}
