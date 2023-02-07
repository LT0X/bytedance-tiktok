package user_info

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/info_service"
	"strconv"
)

type UserInfoResponse struct {
	models.ResponseStatus
	*info_service.InfoResponse
}

type UserInfoHandler struct {
	Uid int64
	UserInfoResponse
	*gin.Context
}

func UserInfoController(c *gin.Context) {

	handler := &UserInfoHandler{
		Context: c,
	}
	//解析参数
	uid := c.Query("user_id")
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		handler.sendResponse(-1, "uid解析错误")
		return
	}
	handler.Uid = id
	//开始业务逻辑，查找对应的UserInfo
	res, err := info_service.NewUserInfoService(handler.Uid).Do()
	if err != nil {
		handler.sendResponse(-1, err.Error())
		return
	}
	handler.InfoResponse = res
	// 开始返回响应数据
	handler.sendResponse(0, "")
	return
}

func (u UserInfoHandler) sendResponse(code int32, msg string) {
	if u.InfoResponse == nil {
		u.JSON(http.StatusOK, models.ResponseStatus{
			StatusCode: code,
			StatusMsg:  msg,
		})
	} else {
		u.StatusCode = code
		u.StatusMsg = msg
		u.JSON(http.StatusOK, u.UserInfoResponse)
	}
}
