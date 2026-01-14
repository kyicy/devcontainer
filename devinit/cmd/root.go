package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devinit",
	Short: "快速初始化新项目的 devcontainer 配置",
	Long: `devinit 是一个强大的 CLI 工具，帮助你快速创建和配置项目的 devcontainer 环境。

支持功能：
  - 项目初始化
  - 配置文件管理`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(configCmd)
}
