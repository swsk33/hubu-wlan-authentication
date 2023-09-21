package config

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// 配置文件读取(Viper)

// 配置项名常量
const (
	// USERNAME 用户名
	USERNAME = "username"
	// PASSWORD 密码
	PASSWORD = "password"
	// RETRY 重试次数
	RETRY = "retry"
	// DELAY 初次启动延迟秒数
	DELAY = "delay"
)

// SelfPath 可执行文件自身路径
var SelfPath string

// SetupConfig 读取并初始化配置
func SetupConfig() error {
	// 初始化自己路径
	SelfPath, _ = os.Executable()
	// 设定配置文件名
	viper.SetConfigName("config")
	// 设定配置文件类型
	viper.SetConfigType("yaml")
	// 设定配置文件可能在的文件夹
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(SelfPath))
	err := viper.ReadInConfig()
	if err != nil {
		color.Red("找不到配置文件！请将配置文件放在运行目录或者和可执行文件同级目录下，名为config.yaml\n")
		return err
	}
	color.Green("加载配置完成！")
	return nil
}