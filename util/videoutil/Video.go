package videoutil

import (
	"fmt"
	"qingxunyin/bytedance-tiktok/config"
)

func GetVideoUrl(videoName string) string {
	url := fmt.Sprintf("http://%s:%d/static/video/%s",
		config.GetConf().Server.IP, config.GetConf().Server.Port, videoName)
	return url
}

func GetPictureUrl(pictureName string) string {
	url := fmt.Sprintf("http://%s:%d/static/picture/%s",
		config.GetConf().Server.IP, config.GetConf().Server.Port, pictureName)
	return url
}
