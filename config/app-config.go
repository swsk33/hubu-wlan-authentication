package config

import (
	"gitee.com/swsk33/sclog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// 配置文件读取(Viper)

// AppConfig 全局配置对象
type AppConfig struct {
	// 用户名
	Username string `mapstructure:"username"`
	// 密码
	Password string `mapstructure:"password"`
	// 重试次数
	Retry int `mapstructure:"retry"`
	// 初次启动延迟秒数
	Delay int `mapstructure:"delay"`
}

// GlobalConfig 全局配置对象
var GlobalConfig AppConfig

// SelfPath 可执行文件自身路径
var SelfPath string

func init() {
	var e error
	// 初始化自己路径
	SelfPath, e = os.Executable()
	if e != nil {
		sclog.ErrorLine("获取自身路径失败！")
		sclog.ErrorLine(e.Error())
		os.Exit(1)
	}
	// 设定配置文件名
	viper.SetConfigName("config")
	// 设定配置文件类型
	viper.SetConfigType("yaml")
	// 设定配置文件可能在的文件夹
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(SelfPath))
	viper.AddConfigPath("/etc/hubu-wlan")
	e = viper.ReadInConfig()
	if e != nil {
		sclog.ErrorLine("找不到配置文件！请将配置文件放在运行目录或者和可执行文件同级目录下，名为config.yaml")
		sclog.ErrorLine(e.Error())
		os.Exit(1)
	}
	// 反序列化配置
	e = viper.Unmarshal(&GlobalConfig)
	if e != nil {
		sclog.ErrorLine("反序列化配置出错！")
		sclog.ErrorLine(e.Error())
		os.Exit(1)
	}
	sclog.InfoLine("加载配置完成！")
}