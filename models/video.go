package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
	"time"
)

type Video struct {
	Id            int64 `json:"id"`
	UserInfoId    int64
	Author        *UserInfo `json:"author" gorm:"-"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64
	IsFavorite    bool   `json:"is_follow" gorm:"-"`
	Title         string `json:"title"`
	UploadTime    time.Time
}

type VideoDao struct {
}

var videoDao *VideoDao

const MaxVideoNum = 30

func GetVideoDao() *VideoDao {
	return videoDao
}

// GetVideoList 按投稿时间倒序返回限制数量的视频
func (*VideoDao) GetVideoList(lastTime time.Time) (*[]Video, error) {
	DB := dbutil.GetDB()
	var videoList []Video
	return &videoList, DB.Table("videos").Order("upload_time Desc").
		Limit(MaxVideoNum).Where("upload_time < ?", lastTime).
		Find(&videoList).Error
}

// QueryUserVideoCount 查询用户的视频数量
func (*VideoDao) QueryUserVideoCount(uid int64) (int64, error) {
	DB := dbutil.GetDB()
	count := new(int64)
	return *count, DB.Table("videos").
		Where("user_info_id = ? ", uid).
		Count(count).Error
}

// AddVideo 更新视频数据库
func (*VideoDao) AddVideo(video *Video) error {
	DB := dbutil.GetDB()
	return DB.Table("videos").Create(video).Error
}

// QueryUserVideoById 查询用户视频
func (*VideoDao) QueryUserVideoById(uid int64) (*[]Video, error) {
	DB := dbutil.GetDB()
	var videoList []Video
	return &videoList, DB.Table("videos").Order("upload_time Desc").
		Limit(MaxVideoNum).Where("user_info_id = ?", uid).
		Find(&videoList).Error
}
