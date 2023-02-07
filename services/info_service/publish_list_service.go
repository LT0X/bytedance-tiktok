package info_service

import "qingxunyin/bytedance-tiktok/models"

type ListResponse struct {
	VideoList *[]models.Video `json:"video_list"`
}

type PublishListService struct {
	Uid int64
}

func NewPublishListService(uid int64) *PublishListService {
	return &PublishListService{
		Uid: uid,
	}
}

func (p *PublishListService) Do() (*ListResponse, error) {

	response := new(ListResponse)
	//查询用户视频列表
	list, err := models.GetVideoDao().QueryUserVideoById(p.Uid)
	if err != nil {
		return nil, err
	}
	response.VideoList = list
	//为每个视频赋值作者信息
	userInfo, err := models.GetUserInfoDao().QueryUserInfoById(p.Uid)
	for index, _ := range *list {
		(*list)[index].Author = userInfo
	}

	return response, nil
}
