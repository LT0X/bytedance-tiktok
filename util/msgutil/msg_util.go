package msgutil

import "qingxunyin/bytedance-tiktok/models"

var Judge map[int64]bool
var FirstJudge map[int64]bool

// ListResponse 为了应对消息叠加的妥协
type ListResponse struct {
	Res  *models.ResponseStatus
	List *[]models.Msg `json:"message_list"`
}

func init() {
	Judge = make(map[int64]bool)
	FirstJudge = make(map[int64]bool)
}
