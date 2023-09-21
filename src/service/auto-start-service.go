package service

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

// 开机自启动逻辑

const (
	// 注册表名
	appName = "HubuWLANAuth"
	// 注册表键所在路径
	regKey = "HKLM\\Software\\Microsoft\\Windows\\CurrentVersion\\Run"
)

// 程序的启动命令
var startCommand string

// SetupAppPath 初始化程序自身所在路径以完成启动命令初始化
func SetupAppPath() {
	path, _ := os.Executable()
	dirPath := filepath.Dir(path)
	// 启动命令将延迟3s执行
	startCommand = fmt.Sprintf("cmd /c timeout /t 3 && cd /d \"%s\" && \"%s\"", dirPath, path)
}

// SetAutoStart 将自己本身添加至开机自启动程序
func SetAutoStart() {
	cmd := exec.Command("reg", "add", regKey, "/v", appName, "/t", "REG_SZ", "/d", startCommand, "/f")
	err := cmd.Run()
	if err != nil {
		color.Red("添加开机启动项失败！请以管理员身份重新运行该程序！")
		return
	}
	color.Green("添加开机启动项成功！")
}

// RemoveAutoStart 移除开机启动
func RemoveAutoStart() {
	cmd := exec.Command("reg", "delete", regKey, "/v", appName, "/f")
	err := cmd.Run()
	if err != nil {
		color.Red("移除开机启动项失败！请以管理员身份重新运行该程序！")
		return
	}
	color.Green("移除开机启动项成功！")
}