package autostart

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootAutoStartCmd 配置开机自动启动项的子命令
var RootAutoStartCmd = &cobra.Command{
	Use:   "auto-start",
	Short: "配置/移除开机自启动",
	Long:  "用于将程序添加到开机自启，或者移除开机自启的子命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请执行hubu-wlan auto-start -h查看具体的开机自启动管理命令！")
	},
}