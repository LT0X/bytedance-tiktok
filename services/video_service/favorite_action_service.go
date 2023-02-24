package video_service

import (
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
)

type FavoriteActionService struct {
	Uid        int64
	Vid        int64
	ActionType string
}

func NewFavoriteActionService(uid int64, vid int64, action string) *FavoriteActionService {
	return &FavoriteActionService{
		Uid:        uid,
		Vid:        vid,
		ActionType: action,
	}
}

func (f FavoriteActionService) Do() (*models.ResponseStatus, error) {

	resp := new(models.ResponseStatus)
	//开始操作
	err := cache.HandleUserFavorite(f.Uid, f.Vid, f.ActionType)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
