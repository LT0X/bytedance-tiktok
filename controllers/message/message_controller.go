package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/message_service"
	"strconv"
)

type UserMessageResponse struct {
	models.ResponseStatus
	MessageList *[]models.Message `json:"message_list"`
}

type UserMessageHandler struct {
	Uid      int64
	ToUserId int64
	UserMessageResponse
	*gin.Context
}

func UserMessageController(c *gin.Context) {

	handler := &UserMessageHandler{
		Context: c,
	}
	//解析参数
	id, _ := c.Get("user_id")
	uid, ok := id.(int64)
	if !ok {
		handler.SendResponse(-1, "解析uid错误")
		return
	}
	temp := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "解析uid错误")
		return
	}

	handler.Uid = uid
	handler.ToUserId = toUserId

	times := c.Query("pre_msg_time")
	//初次访问记录
	preTime, err := strconv.ParseInt(times, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "解析preTime错误")
	}

	//开始业务逻辑,开始查询双方的聊天记录
	response, err := message_service.NewMessageService(handler.Uid, handler.ToUserId, preTime).Do()
	if err != nil {
		handler.SendResponse(-1, err.Error())
		return
	}
	handler.MessageList = response.MessageList
	handler.SendResponse(0, "success")

}

func (u *UserMessageHandler) SendResponse(code int32, msg string) {
	if u.MessageList == nil {
		u.StatusCode = code
		u.StatusMsg = msg
		u.JSON(http.StatusOK, u.ResponseStatus)
	} else {
		u.StatusCode = code
		u.StatusMsg = msg
		u.JSON(http.StatusOK, u.UserMessageResponse)
	}
}
