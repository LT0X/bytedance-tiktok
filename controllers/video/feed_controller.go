package video

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/middlewares"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/video_service"
	"strconv"
	"time"
)

type UserFeedResponse struct {
	models.ResponseStatus
	*video_service.FeedResponse
}

type UserFeedHandler struct {
	Token      string
	Uid        int64
	LatestTime time.Time
	*UserFeedResponse
	*gin.Context
}

func UserFeedController(c *gin.Context) {

	handler := new(UserFeedHandler)
	handler.UserFeedResponse = new(UserFeedResponse)
	handler.Context = c
	//解析参数
	handler.Token = c.Query("token")
	timeStamp := c.Query("latest_time")
	intTime, err := strconv.ParseInt(timeStamp, 10, 64)
	if err == nil {
		handler.LatestTime = time.Unix(0, intTime*1e6) //转换前端传来的数字为时间戳
	}

	if handler.Token != "" {
		//有登录状态的时候进入
		handler.Uid, err = handler.ParseToken()
		if err != nil {
			handler.SendResponse()
			return
		}
		handler.FeedResponse, err = video_service.
			NewFeedService(handler.Uid, handler.LatestTime).Do()
		if err != nil {
			handler.SendResponse()
			return
		}
		handler.SendResponse()
	} else {
		//游客状态
		handler.FeedResponse, err = video_service.
			NewFeedService(handler.Uid, handler.LatestTime).Do()
		if err != nil {
			handler.SendResponse()
			return
		}
		//发送响应信息
		handler.StatusCode = 0
		handler.StatusMsg = "success"
		handler.SendResponse()
	}
}

func (u UserFeedHandler) ParseToken() (int64, error) {

	tokenStruck, ok := middlewares.ParseToken(u.Token)
	if !ok {
		u.StatusCode = 403
		u.StatusMsg = "token错误"
		return -1, errors.New("token错误")
	}
	//token超时
	if time.Now().Unix() > tokenStruck.ExpiresAt {
		u.StatusCode = 403
		u.StatusMsg = "token过期"
		return -1, errors.New("token过期")
	}
	return tokenStruck.Uid, nil
}

func (u UserFeedHandler) SendResponse() {
	if u.FeedResponse == nil {
		u.JSON(http.StatusOK, models.ResponseStatus{
			StatusCode: u.StatusCode,
			StatusMsg:  u.StatusMsg,
		})
	} else {
		u.JSON(http.StatusOK, u.UserFeedResponse)
	}
}
