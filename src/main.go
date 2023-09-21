package main

import (
	"github.com/fatih/color"
	"hubu-wlan-connect/config"
	"hubu-wlan-connect/service"
	"os"
	"time"
)

// 完成操作后延迟5秒退出
func waitAndExit() {
	color.HiMagenta("程序将在5s后自动退出...")
	time.Sleep(5 * time.Second)
}

// 不传入参数启动时，则直接进行网络认证操作
//
// 传入参数时，则是添加/移除开机启动项操作
// 用法：
// hubu-wlan enable-auto-start 启用开机自启
// hubu-wlan disable-auto-start 移除开机自启
func main() {
	args := os.Args
	if len(args) > 1 {
		// 进行启用或者禁用开机自启操作
		service.SetupAppPath()
		option := args[1]
		switch option {
		case "enable-auto-start":
			service.SetAutoStart()
		case "disable-auto-start":
			service.RemoveAutoStart()
		default:
			color.Red("命令行参数错误！")
		}
	} else {
		// 读取配置
		e1 := config.SetupConfig()
		if e1 != nil {
			waitAndExit()
			os.Exit(1)
		}
		// 执行自动登录尝试
		service.DoLoginRetry()
	}
	waitAndExit()
}