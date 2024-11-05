package context

import (
	"errors"
	"hubu-wlan-connect/strategy"
	"hubu-wlan-connect/strategy/impl"
)

// 管理不同开机自启动策略的上下文

// 存放策略对象的容器
var strategyMap map[string]strategy.AutoStart

// 初始化策略容器
func init() {
	strategyMap = make(map[string]strategy.AutoStart)
	strategyMap["windows"] = &impl.WindowsAutoStart{}
	strategyMap["linux"] = &impl.LinuxAutoStart{}
}

// GetStrategy 获得策略对象
//
// os 传入操作系统名称
//
// 返回策略对象，若出现错误返回错误对象
func GetStrategy(os string) (*strategy.AutoStart, error) {
	value, exists := strategyMap[os]
	if !exists {
		return nil, errors.New("指定的操作系统类型不存在！")
	}
	return &value, nil
}