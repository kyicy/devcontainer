package cmd

import (
	"fmt"

	"github.com/kyicy/devcontainer/devinit/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	projectName     string
	workspaceFolder string
	remoteUser      string
	gitEmail        string
	gitUser         string
	githubToken     string
	gitBranch       string
	githubProxy     string
	nonInteractive  bool
)

var initCmd = &cobra.Command{
	Use:   "init [项目路径]",
	Short: "初始化新项目的 devcontainer 配置",
	Long:  `在指定目录创建完整的 devcontainer 配置文件`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "项目名称")
	initCmd.Flags().StringVarP(&workspaceFolder, "workspace", "w", "/home/admin/gopath/src", "工作目录")
	initCmd.Flags().StringVarP(&remoteUser, "user", "u", "admin", "容器用户")
	initCmd.Flags().StringVar(&gitEmail, "git-email", "", "Git 邮箱")
	initCmd.Flags().StringVar(&gitUser, "git-user", "", "Git 用户名")
	initCmd.Flags().StringVar(&githubToken, "github-token", "", "GitHub Token")
	initCmd.Flags().StringVar(&gitBranch, "git-branch", "master", "Git 分支")
	initCmd.Flags().StringVar(&githubProxy, "github-proxy", "http://host.docker.internal:7890", "GitHub 代理")
	initCmd.Flags().BoolVarP(&nonInteractive, "non-interactive", "y", false, "非交互模式")
}

func runInit(cmd *cobra.Command, args []string) error {
	projectPath := "."
	if len(args) > 0 {
		projectPath = args[0]
	}

	config := &generator.DevContainerConfig{
		ProjectName:     projectName,
		DockerImage:     "ghcr.io/kyicy/devcontainer:latest",
		WorkspaceFolder: workspaceFolder,
		RemoteUser:      remoteUser,
		GitEmail:        gitEmail,
		GitUser:         gitUser,
		GithubToken:     githubToken,
		GitBranch:       gitBranch,
		GithubProxy:     githubProxy,
	}

	if nonInteractive {
		if err := generator.GenerateNonInteractive(projectPath, config); err != nil {
			return fmt.Errorf("生成配置失败: %w", err)
		}
	} else {
		if err := generator.GenerateInteractive(projectPath, config); err != nil {
			return fmt.Errorf("生成配置失败: %w", err)
		}
	}

	fmt.Printf("\n✅ Devcontainer 配置已成功创建在: %s/.devcontainer\n", projectPath)
	fmt.Println("\n下一步:")
	fmt.Println("1. 检查配置文件: .devcontainer/devcontainer.json")
	fmt.Println("2. 根据需要调整: .devcontainer/docker-compose.yml")
	fmt.Println("3. 在 VS Code 中重新打开容器")

	return nil
}
