package dbutil

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qingxunyin/bytedance-tiktok/config"
)

var db *gorm.DB

func init() {

	config := config.GetConf()
	//配置数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
		config.Mysql.User, config.Mysql.PassWord, config.Mysql.Host, config.Mysql.Port,
		config.Mysql.DataBase, config.Mysql.CharSet, config.Mysql.ParseTime,
		config.Mysql.Loc)

	//连接数据库
	var err error
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

}

// GetDB 单例模式获取DB对象
// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过调用getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
func GetDB() *gorm.DB {
	return db
}
