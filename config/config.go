package config

import (
	"encoding/json"
	"os"
)

//项目配置文件及读取配置文件的相关功能

var ServConfig AppConfig

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
}

/**
 * MySQL配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	DataBase string `json:"database"`
}

//初始化服务配置
func InitConfig() *AppConfig {

	//打开配置json文件
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}

	//使用json解析器将文件内容解析到conf中
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)

	return &conf
}
