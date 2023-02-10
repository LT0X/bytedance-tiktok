package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/message_service"
	"strconv"
)

//message douyin_relation_action_request {
//required string token = 1; // 用户鉴权token
//required int64 to_user_id = 2; // 对方用户id
//required int32 action_type = 3; // 1-发送消息
//required string content = 4; // 消息内容
//}

type ActionMessageResponse struct {
	models.ResponseStatus
}

type ActionMessageHandler struct {
	FromUid  int64
	ToUserId int64
	Content  string
	response ActionMessageResponse
	*gin.Context
}

func ActionMessageController(c *gin.Context) {

	handler := new(ActionMessageHandler)
	handler.Context = c
	//解析参数
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	id, _ := c.Get("user_id")
	uid, ok := id.(int64)
	if !ok {
		handler.SendResponse(-1, "解析uid错误")
		return
	}
	handler.FromUid = uid
	handler.Content = content
	i, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		handler.SendResponse(-1, "解析uid错误")
		return
	}
	handler.ToUserId = i

	//开始业务逻辑
	err = message_service.NewActionMessageService(handler.ToUserId, handler.FromUid, handler.Content).Do()
	if err != nil {
		handler.SendResponse(-1, err.Error())
	}

	//开始返回响应参数
	handler.SendResponse(0, "success")

}

func (a *ActionMessageHandler) SendResponse(code int32, msg string) {
	a.JSON(http.StatusOK, models.ResponseStatus{
		StatusCode: code,
		StatusMsg:  msg,
	})

}
