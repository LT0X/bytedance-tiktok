package test

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestAdds(t *testing.T) {
	xx := time.Now()
	fmt.Print(xx)
	_, err := GetSnapshot("D:\\GoProject\\src\\qingxunyin\\bytedance-tiktok\\static\\video\\bear.mp4", "D:\\GoProject\\src\\qingxunyin\\bytedance-tiktok\\static\\picture\\xx", 1)
	if err != nil {
		return
	}
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {

	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
