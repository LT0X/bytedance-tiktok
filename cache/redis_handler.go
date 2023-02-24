package cache

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"log"
	config2 "qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/models"
	"strconv"
	"strings"
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

func GetUserFavorite(uid int64) (map[string]string, error) {
	result, err := rdb.HGetAll(fmt.Sprintf("user%v", uid)).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetIsFavorite(vid int64, uid int64) (bool, error) {
	key := fmt.Sprintf("%v:%v", vid, uid)
	result, err := rdb.HGet("video_user_like", key).Result()
	if err != nil {
		return false, err
	}
	if result == "1" {
		return true, nil
	} else {
		return false, nil
	}
}

func GetIsFollow() {

}

func HandleUserFavorite(uid int64, vid int64, action string) error {

	key := fmt.Sprintf("%v:%v", vid, uid)
	videokey := fmt.Sprintf("%v", vid)
	if action == "1" {
		//开启事务
		pipe := rdb.TxPipeline()
		//进行点赞操作
		pipe.HSet("video_user_like", key, 1)
		//增加视频点赞统计
		pipe.HIncrBy("video_like_count", videokey, 1)
		//更新用户的redis列表
		pipe.HSet(fmt.Sprintf("user%v", uid), key, vid)
		// 执行事务
		_, err := pipe.Exec()
		if err != nil {
			return err
		}
	} else if action == "2" {
		//取消点赞操作
		pipe := rdb.TxPipeline()
		//进行点赞操作
		pipe.HSet("video_user_like", key, 0)
		//增加视频点赞统计
		pipe.HIncrBy("video_like_count", videokey, -1)

		pipe.HDel(fmt.Sprintf("user%v", uid), key)
		// 执行事务
		_, err := pipe.Exec()
		if err != nil {
			return err
		}
	}
	return errors.New("错误操作")
}

func GetUserPreMsgTime(key string) (string, error) {
	result, err := rdb.HGet("user_msg_time", key).Result()
	if err != nil {
		return "", err
	}
	return result, err
}
func SetUserPreMsgTime(key string, value string) error {
	_, err := rdb.HSet("user_msg_time", key, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func CronVideoLikeToDB() {
	//开启定时查询redis进行持久化
	c := cron.New()
	_, err := c.AddFunc("@every 1m", func() {

		//从redis获取用户点赞信息
		userLike, err := rdb.HGetAll("video_user_like").Result()
		if err != nil {
			log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
			log.Println("redis持久化失败：", err.Error())
			return
		}
		//从redis获取视频总点赞数
		likeCount, err := rdb.HGetAll("video_like_count").Result()
		if err != nil {
			log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
			log.Println("redis持久化失败：", err.Error())
			return
		}
		if len(userLike) == 0 || len(likeCount) == 0 {
			return
		} else {
			//进行数据库持久化,
			for index, value := range likeCount {

				vid, _ := strconv.ParseInt(index, 10, 64)
				count, _ := strconv.ParseInt(value, 10, 32)
				//更新用户信息
				models.GetUserInfoDao().UpdateUserTotalCount(vid, count)
				err = models.GetVideoDao().UpdateVideoLikeCount(count, vid)
				if err != nil {
					log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
					log.Println("redis持久化失败：", err.Error())
					return
				}
			}
			for index, value := range userLike {
				ids := strings.Split(":", index)
				if value == "1" {
					//更新新的用户视频，
					vid, _ := strconv.ParseInt(ids[0], 10, 64)
					uid, _ := strconv.ParseInt(ids[1], 10, 32)

					err := models.GetUserInfoDao().UpdateUserFavor(uid, vid)
					if err != nil {
						log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
						log.Println("redis持久化失败：", err.Error())
						return
					}
				} else {
					vid, _ := strconv.ParseInt(ids[0], 10, 64)
					uid, _ := strconv.ParseInt(ids[1], 10, 32)
					err := models.GetUserInfoDao().DeleteUserFavor(uid, vid)
					if err != nil {
						log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
						log.Println("redis持久化失败：", err.Error())
						return
					}
				}
			}

		}
	})

	if err != nil {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println(err.Error())
		return
	}
	go c.Start()
	defer c.Stop()
}
