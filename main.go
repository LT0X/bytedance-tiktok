package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/router"
)

func main() {

	r := gin.Default()

	router.InitRouter(r)

	r.Run(fmt.Sprintf(":%d", config.GetConf().Server.Port))

}
