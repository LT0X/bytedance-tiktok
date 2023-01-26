package dbUtil

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type mysqlConfig struct {
	Host      string `yaml:"Host"`
	Port      string `yaml:"Port"`
	User      string `yaml:"User"`
	PassWord  string `yaml:"PassWord"`
	DataBase  string `yaml:"DataBase"`
	CharSet   string `yaml:"CharSet"`
	ParseTime string `yaml:"ParseTime"`
	Loc       string `yaml:"Loc"`
}

type redisConfig struct {
	IP       string `yaml:"IP"`
	Port     int    `yaml:"Port"`
	PassWord string `yaml:"PassWord"`
	DataBase int    `yaml:"DataBase"`
}

type config struct {
	Mysql mysqlConfig `yaml:"mysql"`
	Redis redisConfig `yaml:"redis"`
}

var db *gorm.DB

var rdb *redis.Client

func init() {
	dataBytes, err := os.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	config := config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}

	//配置数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
		config.Mysql.User, config.Mysql.PassWord, config.Mysql.Host, config.Mysql.Port,
		config.Mysql.DataBase, config.Mysql.CharSet, config.Mysql.ParseTime,
		config.Mysql.Loc)

	//连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		fmt.Print("数据库连接出错了", err)
		return
	}
	//配置和获取数据库连接池
	//通过获取数据库连接池指针进行配置
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接池获取失败 ", err)

	}
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxIdleConns(20)

	//初始化redis配置
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port),
			Password: config.Redis.PassWord,
			DB:       config.Redis.DataBase,
		})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败", err)
	}
}

// GetDB 单例模式获取DB对象
// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过调用getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
func GetDB() *gorm.DB {
	return db
}

// GetRDB 单例返回rdb对象
// 使用后无需调用close,连接池会自动回收
// close是关闭连接池对象
func GetRDB() *redis.Client {
	return rdb
}
