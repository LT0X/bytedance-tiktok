package info_service

import (
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
)

type FavoriteResponse struct {
	VideoList *[]models.Video `json:"video_list"`
}

type FavoriteListService struct {
	UserId int64
}

func NewFavoriteListService(uid int64) *FavoriteListService {
	return &FavoriteListService{
		UserId: uid,
	}
}

func (f FavoriteListService) Do() (*FavoriteResponse, error) {

	response := new(FavoriteResponse)

	//从redis查询用户喜爱列表
	favorite, err := cache.GetUserFavorite(f.UserId)

	if err != nil {
		return nil, err
	}
	favor := make([]string, 30)
	for _, value := range favorite {
		favor = append(favor, value)
	}

	//开始查询用户喜欢数据
	list, err := models.GetVideoDao().QueryUserFavorite(favor)
	if err != nil {
		return nil, err
	}
	response.VideoList = list

	//为每一个视频赋值作者信息
	for index, _ := range *response.VideoList {
		uid := (*response.VideoList)[index].UserInfoId
		UserInfo, err := models.GetUserInfoDao().QueryUserInfoById(uid)
		if err != nil {
			return nil, err
		}
		(*response.VideoList)[index].Author = UserInfo
	}

	return response, nil
}
