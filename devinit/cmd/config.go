package cmd

import (
	"fmt"
	"os"

	"github.com/kyicy/devcontainer/devinit/pkg/config"
	"github.com/kyicy/devcontainer/devinit/pkg/util"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [命令]",
	Short: "管理 devcontainer 配置",
	Long:  `查看和修改 devcontainer 配置文件`,
}

var setupConfigCmd = &cobra.Command{
	Use:   "setup",
	Short: "交互式设置用户默认配置",
	Long:  `通过交互式向导设置 Git 用户信息、GitHub Token 等默认配置，保存到 ~/.devinit.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚀 开始设置 devinit 默认配置")
		fmt.Println("========================")
		fmt.Println()

		// 加载现有配置(如果有)
		existingConfig, _ := config.LoadUserConfig()

		// 交互式输入
		gitUser := util.PromptString("Git 用户名", existingConfig.GitUser)
		gitEmail := util.PromptString("Git 邮箱", existingConfig.GitEmail)
		githubToken := util.PromptString("GitHub Token (可选)", existingConfig.GithubToken)
		gitBranch := util.PromptString("Git 默认分支", existingConfig.GitBranch)
		githubProxy := util.PromptString("GitHub 代理地址", existingConfig.GithubProxy)

		fmt.Println()
		fmt.Println("📋 配置确认")
		fmt.Println("========================")
		fmt.Printf("Git 用户名: %s\n", gitUser)
		fmt.Printf("Git 邮箱: %s\n", gitEmail)
		if githubToken != "" {
			fmt.Printf("GitHub Token: %s\n", maskToken(githubToken))
		} else {
			fmt.Println("GitHub Token: (未设置)")
		}
		fmt.Printf("Git 默认分支: %s\n", gitBranch)
		fmt.Printf("GitHub 代理: %s\n", githubProxy)
		fmt.Println()

		if !util.PromptBool("确认保存以上配置?", true) {
			fmt.Println("❌ 已取消")
			return nil
		}

		// 保存配置
		newConfig := &config.UserConfig{
			GitUser:     gitUser,
			GitEmail:    gitEmail,
			GithubToken: githubToken,
			GitBranch:   gitBranch,
			GithubProxy: githubProxy,
			RemoteUser:  "admin",
			Workspace:   "/home/admin/gopath/src",
		}

		if err := config.SaveUserConfig(newConfig); err != nil {
			return fmt.Errorf("保存配置失败: %w", err)
		}

		configPath, _ := config.GetUserConfigPath()
		fmt.Printf("✅ 配置已保存到: %s\n", configPath)
		fmt.Println()
		fmt.Println("💡 提示:")
		fmt.Println("  - 运行 'devinit init --name <项目名>' 时将自动使用这些默认值")
		fmt.Println("  - 可以通过命令行参数覆盖默认值")
		fmt.Println("  - 随时可以运行 'devinit config setup' 重新配置")

		return nil
	},
}

// maskToken 隐藏 Token 的大部分内容
func maskToken(token string) string {
	if len(token) <= 8 {
		return "***"
	}
	return token[:4] + "..." + token[len(token)-4:]
}

var viewUserConfigCmd = &cobra.Command{
	Use:   "view-user",
	Short: "查看用户默认配置",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := config.GetUserConfigPath()
		if err != nil {
			return err
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Println("❌ 用户配置文件不存在")
			fmt.Println("💡 运行 'devinit config setup' 来创建配置")
			return nil
		}

		cfg, err := config.LoadUserConfig()
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		fmt.Println("📋 当前用户配置")
		fmt.Println("========================")
		fmt.Printf("配置文件: %s\n", configPath)
		fmt.Printf("Git 用户名: %s\n", cfg.GitUser)
		fmt.Printf("Git 邮箱: %s\n", cfg.GitEmail)
		if cfg.GithubToken != "" {
			fmt.Printf("GitHub Token: %s\n", maskToken(cfg.GithubToken))
		} else {
			fmt.Println("GitHub Token: (未设置)")
		}
		fmt.Printf("Git 默认分支: %s\n", cfg.GitBranch)
		fmt.Printf("GitHub 代理: %s\n", cfg.GithubProxy)
		fmt.Printf("容器用户: %s\n", cfg.RemoteUser)
		fmt.Printf("工作目录: %s\n", cfg.Workspace)

		return nil
	},
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
	configCmd.AddCommand(setupConfigCmd)
	configCmd.AddCommand(viewUserConfigCmd)
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setEnvCmd)
	configCmd.AddCommand(addExtensionCmd)
}
