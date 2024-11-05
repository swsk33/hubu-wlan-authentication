package main

import (
	"embed"
	"gitee.com/swsk33/sclog"
	"hubu-wlan-connect/cmd"
	"hubu-wlan-connect/util"
)

// 嵌入的配置文件模板
//
//go:embed config-template/linux-autostart.desktop
var embedConfig embed.FS

func init() {
	util.SetupEmbedFile(&embedConfig)
}

func main() {
	e := cmd.RootCmd.Execute()
	if e != nil {
		sclog.ErrorLine("执行命令出错！")
		sclog.ErrorLine(e.Error())
	}
}