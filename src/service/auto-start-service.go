package service

import (
	"github.com/fatih/color"
	"hubu-wlan-connect/strategy/context"
	"os"
	"runtime"
)

// 开机自启动逻辑

// 程序自身路径
var selfPath string

// SetupAppPath 初始化程序自身所在路径以完成启动命令初始化
func SetupAppPath() {
	selfPath, _ = os.Executable()
}

// SetAutoStart 将自己本身添加至开机自启动程序
func SetAutoStart() {
	strategy, e1 := context.GetStrategy(runtime.GOOS)
	if e1 != nil {
		color.Red(e1.Error())
		return
	}
	e2 := (*strategy).AddAutoStart(selfPath)
	if e2 != nil {
		color.Red("添加开机启动项失败！请以管理员身份重新运行该程序！")
		return
	}
	color.Green("添加开机启动项成功！")
}

// RemoveAutoStart 移除开机启动
func RemoveAutoStart() {
	strategy, e1 := context.GetStrategy(runtime.GOOS)
	if e1 != nil {
		color.Red(e1.Error())
		return
	}
	e2 := (*strategy).RemoveAutoStart()
	if e2 != nil {
		color.Red("移除开机启动项失败！请以管理员身份重新运行该程序！")
		return
	}
	color.Green("移除开机启动项成功！")
}