package autostart

import (
	"gitee.com/swsk33/sclog"
	"github.com/spf13/cobra"
	"hubu-wlan-connect/strategy/context"
	"runtime"
)

// 移除开机自启的子命令
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "移除开机自启动",
	Long:  "将自动认证程序从开机自启动项中移除",
	Run: func(cmd *cobra.Command, args []string) {
		strategy, e := context.GetStrategy(runtime.GOOS)
		if e != nil {
			sclog.ErrorLine("获取开机自启动策略出现错误！")
			sclog.ErrorLine(e.Error())
			return
		}
		e = (*strategy).RemoveAutoStart()
		if e != nil {
			sclog.ErrorLine("移除开机自启动出错！")
			sclog.ErrorLine(e.Error())
			return
		}
		sclog.InfoLine("移除开机自启动完成！")
	},
}

func init() {
	RootAutoStartCmd.AddCommand(removeCmd)
}