package user_info

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/services/info_service"
	"strconv"
)

type PublishListResponse struct {
	models.ResponseStatus
	*info_service.ListResponse
}

func PublishListController(c *gin.Context) {

	response := new(PublishListResponse)
	//解析参数
	id := c.Query("user_id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "解析uid错误"
		sendResponse(c, response)
		return
	}
	//进行业务逻辑的处理，进行查询用户视频
	response.ListResponse, err = info_service.NewPublishListService(uid).Do()
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		sendResponse(c, response)
		return
	}
	response.StatusCode = 0
	sendResponse(c, response)

}

func sendResponse(c *gin.Context, response *PublishListResponse) {
	if response.ListResponse == nil {
		c.JSON(http.StatusOK, response.ResponseStatus)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
