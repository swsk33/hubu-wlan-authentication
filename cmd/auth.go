package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"hubu-wlan-connect/service"
	"strconv"
	"time"
)

// 认证命令
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "一键认证登录",
	Long:  "通过该命令即可读取配置文件的账户信息，并一键登录校园网",
	Run: func(cmd *cobra.Command, args []string) {
		service.DoLoginRetry()
		delay := 3
		color.HiMagenta("程序将在%s秒后自动退出...", strconv.Itoa(delay))
		time.Sleep(time.Duration(delay) * time.Second)
	},
}

func init() {
	RootCmd.AddCommand(authCmd)
}