package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

type serverConfig struct {
	IP   string `yaml:"IP"`
	Port int    `yaml:"Port"`
}

type pathConfig struct {
	StaticPath string `yaml:"StaticPath"`
}

type config struct {
	Mysql  mysqlConfig  `yaml:"mysql"`
	Redis  redisConfig  `yaml:"redis"`
	Server serverConfig `yaml:"server"`
	Path   pathConfig   `yaml:"path"`
}

//项目配置文件对象
var conf *config

//在导包的时候自动加载init函数读取配置文件初始化conf对象
func init() {

	dataBytes, err := os.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}

}

//单例模式获取配置文件对象
func GetConf() *config {
	return conf
}
