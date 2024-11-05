package cmd

import (
	"github.com/spf13/cobra"
	"hubu-wlan-connect/cmd/autostart"
)

// RootCmd 根命令
var RootCmd = &cobra.Command{
	Use:   "hubu-wlan",
	Short: "湖北大学校园网自动认证",
	Long:  "用于湖北大学校园网一键登录认证的程序，支持开机自启动",
	Run: func(cmd *cobra.Command, args []string) {
		authCmd.Run(cmd, args)
	},
}

func init() {
	RootCmd.CompletionOptions.HiddenDefaultCmd = true
	RootCmd.AddCommand(autostart.RootAutoStartCmd)
}