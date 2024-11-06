package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"runtime"
)

// 版本号命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本号",
	Long:  "输出该命令行程序的版本号",
	Run: func(cmd *cobra.Command, args []string) {
		color.HiGreen("湖北大学校园网自动认证 v%d.%d.%d", 2, 2, 1)
		color.HiBlue("使用该版本Golang构建：%s", runtime.Version())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}