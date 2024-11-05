package autostart

import (
	"gitee.com/swsk33/sclog"
	"github.com/spf13/cobra"
	"hubu-wlan-connect/config"
	"hubu-wlan-connect/strategy/context"
	"runtime"
)

// 添加开机自启的子命令
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加开机自启动",
	Long:  "将自动认证程序添加到开机自启动项中，使得程序可以开机时自动运行",
	Run: func(cmd *cobra.Command, args []string) {
		strategy, e := context.GetStrategy(runtime.GOOS)
		if e != nil {
			sclog.ErrorLine("获取开机自启动策略出现错误！")
			sclog.ErrorLine(e.Error())
			return
		}
		e = (*strategy).AddAutoStart(config.SelfPath)
		if e != nil {
			sclog.ErrorLine("添加开机自启动出错！")
			sclog.ErrorLine(e.Error())
			return
		}
		sclog.InfoLine("添加开机自启动完成！")
	},
}

func init() {
	RootAutoStartCmd.AddCommand(addCmd)
}