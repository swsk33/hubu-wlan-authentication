package service

import (
	"gitee.com/swsk33/sclog"
	"hubu-wlan-connect/config"
	"hubu-wlan-connect/strategy/context"
	"runtime"
)

// 开机自启动逻辑

// SetAutoStart 将自己本身添加至开机自启动程序
func SetAutoStart() {
	strategy, e := context.GetStrategy(runtime.GOOS)
	if e != nil {
		sclog.ErrorLine(e.Error())
		return
	}
	e = (*strategy).AddAutoStart(config.SelfPath)
	if e != nil {
		sclog.ErrorLine("添加开机启动项失败！请以管理员身份重新运行该程序！")
		sclog.ErrorLine(e.Error())
		return
	}
	sclog.InfoLine("添加开机启动项成功！")
}

// RemoveAutoStart 移除开机启动
func RemoveAutoStart() {
	strategy, e := context.GetStrategy(runtime.GOOS)
	if e != nil {
		sclog.ErrorLine(e.Error())
		return
	}
	e = (*strategy).RemoveAutoStart()
	if e != nil {
		sclog.ErrorLine("移除开机启动项失败！请以管理员身份重新运行该程序！")
		sclog.ErrorLine(e.Error())
		return
	}
	sclog.InfoLine("移除开机启动项成功！")
}