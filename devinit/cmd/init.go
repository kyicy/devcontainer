package cmd

import (
	"fmt"

	"github.com/kyicy/devcontainer/devinit/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	projectName string
)

var initCmd = &cobra.Command{
	Use:   "init [项目路径]",
	Short: "初始化新项目的 devcontainer 配置",
	Long:  `在指定目录创建完整的 devcontainer 配置文件。项目名称必须通过 --name 参数指定。`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "项目名称 (必填)")

	initCmd.MarkFlagRequired("name")
}

func runInit(cmd *cobra.Command, args []string) error {
	projectPath := "."
	if len(args) > 0 {
		projectPath = args[0]
	}

	devConfig := &generator.DevContainerConfig{
		ProjectName: projectName,
		DockerImage: "ghcr.io/kyicy/devcontainer:latest",
		RemoteUser:  "admin",
	}

	if err := generator.GenerateNonInteractive(projectPath, devConfig); err != nil {
		return fmt.Errorf("生成配置失败: %w", err)
	}

	fmt.Printf("\n✅ Devcontainer 配置已成功创建在: %s/.devcontainer\n", projectPath)
	fmt.Println("\n📋 使用的配置:")
	fmt.Printf("  项目名称: %s\n", projectName)
	fmt.Println("\n⚠️  重要提示:")
	fmt.Println("  workspaceFolder 已固定为 /home/admin")
	fmt.Println("  你必须根据项目需求手动修改 .devcontainer/devcontainer.json 中的 workspaceFolder")
	fmt.Println("\n📝 Git 和 GitHub 认证:")
	fmt.Println("  容器首次启动并运行 devdep.sh 时，会提示你输入:")
	fmt.Println("  - Git 用户名和邮箱")
	fmt.Println("  - GitHub 代理 (可选)")
	fmt.Println("  - GitHub Token (可选)")
	fmt.Println("\n下一步:")
	fmt.Println("1. 修改 workspaceFolder: .devcontainer/devcontainer.json")
	fmt.Println("2. 根据需要调整: .devcontainer/docker-compose.yml")
	fmt.Println("3. 在 VS Code 中重新打开容器")

	return nil
}
