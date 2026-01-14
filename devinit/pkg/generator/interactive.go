package generator

import (
	"fmt"
	"os"
)

// GenerateNonInteractive 非交互式生成配置
func GenerateNonInteractive(projectPath string, config *DevContainerConfig) error {
	if config.ProjectName == "" {
		return fmt.Errorf("项目名称不能为空")
	}

	if len(config.Extensions) == 0 {
		config.Extensions = []string{"anthropic.claude-code"}
	}

	return generateFiles(projectPath, config)
}

// generateFiles 生成配置文件
func generateFiles(projectPath string, config *DevContainerConfig) error {
	serviceName := config.ProjectName + "_dev"

	// 创建 .devcontainer 目录
	devcontainerPath := projectPath + "/.devcontainer"
	if err := os.MkdirAll(devcontainerPath, 0o755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 生成 devcontainer.json
	if err := generateDevcontainerJSON(devcontainerPath, config, serviceName); err != nil {
		return err
	}

	// 生成 docker-compose.yml
	if err := generateDockerCompose(devcontainerPath, config, serviceName); err != nil {
		return err
	}

	return nil
}
