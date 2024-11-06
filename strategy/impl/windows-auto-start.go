package impl

import (
	"fmt"
	"os/exec"
)

const (
	// 注册表名
	appName = "HubuWLANAuth"
	// 注册表键所在路径
	regKey = "HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run"
)

// WindowsAutoStart Windows操作系统实现自启动的实现类
type WindowsAutoStart struct {
}

// AddAutoStart 将自己本身添加至开机自启动程序
func (autoStart *WindowsAutoStart) AddAutoStart(exePath string) error {
	cmd := exec.Command("reg", "add", regKey, "/v", appName, "/t", "REG_SZ", "/d", fmt.Sprintf("\"%s\"", exePath), "/f")
	e := cmd.Run()
	if e != nil {
		return e
	}
	return nil
}

// RemoveAutoStart 移除开机启动
func (autoStart *WindowsAutoStart) RemoveAutoStart() error {
	cmd := exec.Command("reg", "delete", regKey, "/v", appName, "/f")
	e := cmd.Run()
	if e != nil {
		return e
	}
	return nil
}