package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	config2 "qingxunyin/bytedance-tiktok/config"
)

// redis操作连接
// 使用后无需调用close,会自动回收
var rdb *redis.Client

func init() {

	config := config2.GetConf()
	//初始化redis配置
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port),
			Password: config.Redis.PassWord,
			DB:       config.Redis.DataBase,
		})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败", err)
	}
}