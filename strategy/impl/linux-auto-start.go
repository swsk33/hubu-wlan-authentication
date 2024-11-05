package impl

import (
	"bufio"
	"fmt"
	"gitee.com/swsk33/sclog"
	"hubu-wlan-connect/util"
	"os"
	"path/filepath"
)

// 自启动desktop文件所在路径
var autoStartFilePath string

func init() {
	userFolder, _ := os.UserHomeDir()
	autoStartFilePath = filepath.Join(userFolder, ".config", "autostart", "hubu-wlan.desktop")
}

// LinuxAutoStart Linux操作系统实现自启动的实现类
type LinuxAutoStart struct {
}

// AddAutoStart 将自己本身添加至开机自启动程序
func (autoStart *LinuxAutoStart) AddAutoStart(exePath string) error {
	// 创建桌面配置文件
	// 创建之前先删除
	_ = autoStart.RemoveAutoStart()
	// 释放内嵌模板
	e := util.ExtractEmbedFile("config-template/linux-autostart.desktop", autoStartFilePath)
	if e != nil {
		return e
	}
	// 进行读取
	file, e := os.OpenFile(autoStartFilePath, os.O_APPEND|os.O_WRONLY, 0755)
	if e != nil {
		return e
	}
	// 准备写入
	writer := bufio.NewWriter(file)
	_, e = writer.WriteString(fmt.Sprintf("\nExec=\"%s\" auth", exePath))
	_, e = writer.WriteString(fmt.Sprintf("\nPath=%s", filepath.Dir(exePath)))
	defer func() {
		_ = writer.Flush()
		_ = file.Close()
	}()
	if e != nil {
		sclog.ErrorLine("写入文件时出现错误！")
		return e
	}
	return nil
}

// RemoveAutoStart 移除开机启动
func (autoStart *LinuxAutoStart) RemoveAutoStart() error {
	e := os.Remove(autoStartFilePath)
	return e
}