package video_service

import (
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
	"time"
)

type FeedResponse struct {
	VideoList *[]models.Video `json:"video_list"`
	NextTime  int64           `json:"next_time"`
}

type FeedService struct {
	Uid        int64
	LatestTime time.Time
	*FeedResponse
}

func NewFeedService(uid int64, latestTime time.Time) *FeedService {
	return &FeedService{
		Uid:        uid,
		LatestTime: latestTime,
	}
}

func (f *FeedService) Do() (*FeedResponse, error) {
	response := new(FeedResponse)
	//游客登录，无数据时，我们也应该填充数据
	f.FillData()
	//返回视频的操作
	list, err := models.GetVideoDao().GetVideoList(f.LatestTime)
	if err != nil {
		return nil, err
	}
	response.VideoList = list

	for index, _ := range *list {
		//为每个视频赋值作者信息
		user, err := models.GetUserInfoDao().QueryUserInfoById((*list)[index].UserInfoId)
		if err != nil {
			return nil, err
		}
		(*list)[index].Author = user
		//如果是用户登录,需要检查视频和用户是否关注已经点赞
		if f.Uid > 0 {
			cache.GetIsFavorite((*list)[index])
		}
	}

	//填入下一次的时间戳
	if len(*list) > 0 {
		time := (*list)[len(*list)-1].UploadTime
		response.NextTime = (time).UnixNano() / 1e6
	} else {
		response.NextTime = time.Now().UnixNano() / 1e6
	}
	return response, nil
}

func (f *FeedService) FillData() {

	if f.LatestTime.IsZero() {
		//当前端无数据时填入
		f.LatestTime = time.Now()

	}
}
