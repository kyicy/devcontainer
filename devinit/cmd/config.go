package cmd

import (
	"fmt"

	"github.com/kyicy/devcontainer/devinit/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [命令]",
	Short: "管理 devcontainer 配置",
	Long:  `查看和修改 devcontainer 配置文件`,
}

var viewConfigCmd = &cobra.Command{
	Use:   "view [项目路径]",
	Short: "查看当前配置",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectPath := "."
		if len(args) > 0 {
			projectPath = args[0]
		}

		cfg, err := config.LoadConfig(projectPath)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		fmt.Println("当前 devcontainer 配置:")
		fmt.Println("========================")
		fmt.Printf("项目名称: %s\n", cfg.Name)
		fmt.Printf("Docker 镜像: %s\n", cfg.Image)
		fmt.Printf("工作目录: %s\n", cfg.WorkspaceFolder)
		fmt.Printf("容器用户: %s\n", cfg.RemoteUser)
		fmt.Println("\n环境变量:")
		for key, value := range cfg.ContainerEnv {
			fmt.Printf("  %s: %s\n", key, value)
		}

		return nil
	},
}

var setEnvCmd = &cobra.Command{
	Use:   "set-env [key] [value]",
	Short: "设置环境变量",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		if err := config.SetEnvVariable(".", key, value); err != nil {
			return fmt.Errorf("设置环境变量失败: %w", err)
		}

		fmt.Printf("✅ 环境变量 %s 已设置\n", key)
		return nil
	},
}

var addExtensionCmd = &cobra.Command{
	Use:   "add-extension [扩展ID]",
	Short: "添加 VS Code 扩展",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		extensionID := args[0]

		if err := config.AddVSCodeExtension(".", extensionID); err != nil {
			return fmt.Errorf("添加扩展失败: %w", err)
		}

		fmt.Printf("✅ VS Code 扩展 %s 已添加\n", extensionID)
		return nil
	},
}

func init() {
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setEnvCmd)
	configCmd.AddCommand(addExtensionCmd)
}
