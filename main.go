package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/router"
)

func main() {

	go cache.CronVideoLikeToDB()
	r := gin.Default()

	router.InitRouter(r)

	r.Run(fmt.Sprintf(":%d", config.GetConf().Server.Port))

}
