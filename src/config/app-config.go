package config

import (
	"bufio"
	"errors"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

// 配置文件读取
// 配置文件格式：
//
// 第1行：学号
// 第2行：密码
// 第3行：失败重试次数

// 配置文件位置
const configFilePath = "config.cfg"

// 配置定义结构体（单例）
type configObject struct {
	// 学号
	Number string
	// 密码
	Password string
	// 失败重试次数
	Retry int
}

// 初始化全局配置结构体对象
var totalConfig configObject

// GetConfig 获取全局配置对象
func GetConfig() *configObject {
	return &totalConfig
}

// SetupConfig 读取并初始化配置
func SetupConfig() error {
	file, e1 := os.Open(configFilePath)
	if e1 != nil {
		color.HiRed("无法读取配置！配置文件%s可能不存在！", configFilePath)
		return e1
	}
	reader := bufio.NewReader(file)
	// 第一行是学号
	l1, _, _ := reader.ReadLine()
	totalConfig.Number = strings.TrimSpace(string(l1))
	// 第二行是密码
	l2, _, _ := reader.ReadLine()
	totalConfig.Password = strings.TrimSpace(string(l2))
	// 第三行是重试次数
	l3, _, _ := reader.ReadLine()
	var e2 error
	totalConfig.Retry, e2 = strconv.Atoi(strings.TrimSpace(string(l3)))
	if e2 != nil || totalConfig.Retry < 0 {
		color.HiRed("配置文件中重试次数配置（第3行）必须是不小于0的整数！")
		return errors.New("")
	}
	_ = file.Close()
	color.Green("已完成配置读取！")
	return nil
}