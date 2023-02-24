package video_service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/util/videoutil"
	"time"
)

type VideoResponse struct {
	*models.ResponseStatus
}

var (
	videoFormatMap = map[string]struct{}{
		".mp4": {},
		".avi": {},
		".mov": {},
	}
)

type PublishVideoService struct {
	File  *multipart.FileHeader
	Title string
	Uid   int64
	c     *gin.Context
}

func NewPublishVideoService(file *multipart.FileHeader, title string, uid int64, c *gin.Context) *PublishVideoService {
	return &PublishVideoService{
		File:  file,
		Title: title,
		Uid:   uid,
		c:     c,
	}

}

func (p *PublishVideoService) Do() (*models.ResponseStatus, error) {
	response := VideoResponse{
		ResponseStatus: &models.ResponseStatus{},
	}
	//检查文件格式
	suffix := filepath.Ext(p.File.Filename)   //得到后缀
	if _, ok := videoFormatMap[suffix]; !ok { //判断是否为视频格式
		return nil, errors.New("不支持视频格式")
	}
	//为视频文件重新命名文件名
	fileName, err := p.GetVideoName()
	videoName := fileName + suffix
	if err != nil {
		return nil, err
	}

	//保存视频
	video := new(models.Video)
	staticPath := config.GetConf().Path.StaticPath
	videoPath := filepath.Join(staticPath+"/video/", videoName)
	err = p.c.SaveUploadedFile(p.File, videoPath)
	if err != nil {
		return nil, err
	}
	//获取视频第一帧为视频封面并保存

	_, err = videoutil.GetVideoPicture(videoPath, staticPath+"/picture/"+fileName, 1)
	if err != nil {
		return nil, err
	}
	//获取视频url和封面url
	video.CoverUrl = videoutil.GetPictureUrl(fileName + ".png")
	video.PlayUrl = videoutil.GetVideoUrl(videoName)

	//更新数据库数据
	video.Title = p.Title
	video.UserInfoId = p.Uid
	video.UploadTime = time.Now()
	err = models.GetVideoDao().AddVideo(video)

	if err != nil {
		return nil, err
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	return response.ResponseStatus, nil
}

func (p *PublishVideoService) GetVideoName() (string, error) {

	//根据用户id 和视频总数返回特定的文件名
	var filename string
	count, err := models.GetVideoDao().QueryUserVideoCount(p.Uid)
	if err != nil {
		return "", err
	}
	filename = fmt.Sprintf("%d-%d", p.Uid, count)
	return filename, nil

}
