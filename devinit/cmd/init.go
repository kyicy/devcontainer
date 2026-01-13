package cmd

import (
	"fmt"

	"github.com/kyicy/devcontainer/devinit/pkg/config"
	"github.com/kyicy/devcontainer/devinit/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	projectName string
	remoteUser  string
	gitEmail    string
	gitUser     string
	githubToken string
	gitBranch   string
	githubProxy string
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
	initCmd.Flags().StringVarP(&remoteUser, "user", "u", "", "容器用户 (默认从配置文件读取)")
	initCmd.Flags().StringVar(&gitEmail, "git-email", "", "Git 邮箱 (默认从配置文件读取)")
	initCmd.Flags().StringVar(&gitUser, "git-user", "", "Git 用户名 (默认从配置文件读取)")
	initCmd.Flags().StringVar(&githubToken, "github-token", "", "GitHub Token (默认从配置文件读取)")
	initCmd.Flags().StringVar(&gitBranch, "git-branch", "", "Git 分支 (默认从配置文件读取)")
	initCmd.Flags().StringVar(&githubProxy, "github-proxy", "", "GitHub 代理 (默认从配置文件读取)")

	initCmd.MarkFlagRequired("name")
}

func runInit(cmd *cobra.Command, args []string) error {
	projectPath := "."
	if len(args) > 0 {
		projectPath = args[0]
	}

	// 加载用户配置
	userConfig, err := config.LoadUserConfig()
	if err != nil {
		return fmt.Errorf("加载用户配置失败: %w", err)
	}

	// 检查是否需要提示用户先设置配置
	if !config.IsConfigExists() {
		fmt.Println("⚠️  检测到您还未设置用户默认配置")
		fmt.Println("💡 建议先运行 'devinit config setup' 来设置默认配置")
		fmt.Println("   这样可以避免每次都输入相同的参数")
		fmt.Println()
	}

	// 使用命令行参数覆盖配置文件中的值
	if remoteUser == "" {
		remoteUser = userConfig.RemoteUser
	}
	if gitEmail == "" {
		gitEmail = userConfig.GitEmail
	}
	if gitUser == "" {
		gitUser = userConfig.GitUser
	}
	if githubToken == "" {
		githubToken = userConfig.GithubToken
	}
	if gitBranch == "" {
		gitBranch = userConfig.GitBranch
	}
	if githubProxy == "" {
		githubProxy = userConfig.GithubProxy
	}

	// 验证必填字段
	if gitUser == "" || gitEmail == "" {
		return fmt.Errorf("Git 用户名和邮箱不能为空，请通过参数指定或运行 'devinit config setup' 配置默认值")
	}

	devConfig := &generator.DevContainerConfig{
		ProjectName: projectName,
		DockerImage: "ghcr.io/kyicy/devcontainer:latest",
		RemoteUser:  remoteUser,
		GitEmail:    gitEmail,
		GitUser:     gitUser,
		GithubToken: githubToken,
		GitBranch:   gitBranch,
		GithubProxy: githubProxy,
	}

	if err := generator.GenerateNonInteractive(projectPath, devConfig); err != nil {
		return fmt.Errorf("生成配置失败: %w", err)
	}

	fmt.Printf("\n✅ Devcontainer 配置已成功创建在: %s/.devcontainer\n", projectPath)
	fmt.Println("\n📋 使用的配置:")
	fmt.Printf("  项目名称: %s\n", projectName)
	fmt.Printf("  Git 用户: %s <%s>\n", gitUser, gitEmail)
	if githubToken != "" {
		fmt.Println("  GitHub Token: *** (已设置)")
	}
	fmt.Println("\n⚠️  重要提示:")
	fmt.Println("  workspaceFolder 已固定为 /home/admin")
	fmt.Println("  你必须根据项目需求手动修改 .devcontainer/devcontainer.json 中的 workspaceFolder")
	fmt.Println("\n下一步:")
	fmt.Println("1. 修改 workspaceFolder: .devcontainer/devcontainer.json")
	fmt.Println("2. 根据需要调整: .devcontainer/docker-compose.yml")
	fmt.Println("3. 在 VS Code 中重新打开容器")

	return nil
}
