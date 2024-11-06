package main

import (
	"embed"
	"gitee.com/swsk33/sclog"
	"github.com/fatih/color"
	"hubu-wlan-connect/cmd"
	"hubu-wlan-connect/service"
	"hubu-wlan-connect/util"
	"os"
	"strconv"
	"time"
)

// 嵌入的配置文件模板
//
//go:embed config-template/linux-autostart.desktop
var embedConfig embed.FS

func init() {
	util.SetupEmbedFile(&embedConfig)
}

func main() {
	// 若未传参，则直接进行认证
	if len(os.Args) == 1 {
		service.DoLoginRetry()
		delay := 3
		color.HiMagenta("程序将在%s秒后自动退出...", strconv.Itoa(delay))
		time.Sleep(time.Duration(delay) * time.Second)
		return
	}
	// 否则，调用Cobra命令行逻辑
	e := cmd.RootCmd.Execute()
	if e != nil {
		sclog.ErrorLine("执行命令出错！")
		sclog.ErrorLine(e.Error())
	}
}