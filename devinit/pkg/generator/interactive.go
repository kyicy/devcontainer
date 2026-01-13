package generator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GenerateInteractive 交互式生成配置
func GenerateInteractive(projectPath string, config *DevContainerConfig) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🚀 欢迎使用 devinit 交互式配置向导")
	fmt.Println("=====================================")
	fmt.Println()

	// 项目名称
	if config.ProjectName == "" {
		fmt.Print("请输入项目名称: ")
		input, _ := reader.ReadString('\n')
		config.ProjectName = strings.TrimSpace(input)
		if config.ProjectName == "" {
			return fmt.Errorf("项目名称不能为空")
		}
	}

	// 工作目录
	var input string
	fmt.Printf("请输入容器工作目录 [%s]: ", config.WorkspaceFolder)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		config.WorkspaceFolder = input
	}

	// 容器用户
	fmt.Printf("请输入容器用户 [%s]: ", config.RemoteUser)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		config.RemoteUser = input
	}

	fmt.Println()
	fmt.Println("📧 Git 配置")

	// Git 用户
	if config.GitUser == "" {
		fmt.Print("请输入 Git 用户名: ")
		input, _ = reader.ReadString('\n')
		config.GitUser = strings.TrimSpace(input)
	}

	// Git 邮箱
	if config.GitEmail == "" {
		fmt.Print("请输入 Git 邮箱: ")
		input, _ = reader.ReadString('\n')
		config.GitEmail = strings.TrimSpace(input)
	}

	// GitHub Token
	if config.GithubToken == "" {
		fmt.Print("请输入 GitHub Token (可选，直接回车跳过): ")
		input, _ = reader.ReadString('\n')
		config.GithubToken = strings.TrimSpace(input)
	}

	// Git 分支
	fmt.Printf("请输入默认 Git 分支 [%s]: ", config.GitBranch)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		config.GitBranch = input
	}

	fmt.Println()
	fmt.Println("🔌 VS Code 扩展")
	fmt.Println("常用扩展:")
	fmt.Println("  - golang.go (Go 语言支持)")
	fmt.Println("  - ms-python.python (Python 支持)")
	fmt.Println("  - ms-vscode.vscode-typescript-next (TypeScript)")
	fmt.Println("  - rust-lang.rust-analyzer (Rust)")
	fmt.Println()

	extensions := []string{"golang.go", "eamodio.gitlens", "anthropic.claude-code"}
	fmt.Print("请输入需要的 VS Code 扩展 (用逗号分隔，直接回车使用默认): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		config.Extensions = strings.Split(input, ",")
		for i := range config.Extensions {
			config.Extensions[i] = strings.TrimSpace(config.Extensions[i])
		}
	} else {
		config.Extensions = extensions
	}

	fmt.Println()
	fmt.Println("📝 生成配置文件...")
	return generateFiles(projectPath, config)
}

// GenerateNonInteractive 非交互式生成配置
func GenerateNonInteractive(projectPath string, config *DevContainerConfig) error {
	if config.ProjectName == "" {
		return fmt.Errorf("项目名称不能为空")
	}

	if len(config.Extensions) == 0 {
		config.Extensions = []string{"golang.go", "eamodio.gitlens", "anthropic.claude-code"}
	}

	return generateFiles(projectPath, config)
}

// generateFiles 生成配置文件
func generateFiles(projectPath string, config *DevContainerConfig) error {
	serviceName := config.ProjectName + "_dev"

	// 创建 .devcontainer 目录
	devcontainerPath := projectPath + "/.devcontainer"
	if err := os.MkdirAll(devcontainerPath, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 创建 mapping 目录
	mappingPath := devcontainerPath + "/mapping"
	if err := os.MkdirAll(mappingPath, 0755); err != nil {
		return fmt.Errorf("创建 mapping 目录失败: %w", err)
	}

	// 生成 devcontainer.json
	if err := generateDevcontainerJSON(devcontainerPath, config, serviceName); err != nil {
		return err
	}

	// 生成 docker-compose.yml
	if err := generateDockerCompose(devcontainerPath, config, serviceName); err != nil {
		return err
	}

	// 生成 post-create.sh
	if err := generatePostCreateScript(mappingPath, config); err != nil {
		return err
	}

	// 生成 devcontainer-dependencies
	if err := generateDependenciesFile(mappingPath); err != nil {
		return err
	}

	return nil
}
