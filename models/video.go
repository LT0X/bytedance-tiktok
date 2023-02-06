package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
	"time"
)

type Video struct {
	Id            int64
	UserInfoId    int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool `json:"is_follow" gorm:"-"`
	Title         string
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
