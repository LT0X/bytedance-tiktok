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
	Uid      int64
	ToUserId int64
	UserInfoResponse
	*gin.Context
}

func UserInfoController(c *gin.Context) {

	handler := &UserInfoHandler{
		Context: c,
	}
	//解析参数
	id := c.Query("user_id")
	toUserId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "uid解析错误")
		return
	}

	temp, _ := c.Get("user_id")
	uid, ok := temp.(int64)
	if !ok {
		handler.SendResponse(-1, "解析错误")
		return
	}
	handler.ToUserId = toUserId
	handler.Uid = uid
	//开始业务逻辑，查找对应的UserInfo
	res, err := info_service.NewUserInfoService(handler.Uid, handler.ToUserId).Do()
	if err != nil {
		handler.SendResponse(-1, err.Error())
		return
	}
	handler.InfoResponse = res
	// 开始返回响应数据
	handler.SendResponse(0, "success")
	return
}

func (u UserInfoHandler) SendResponse(code int32, msg string) {
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
