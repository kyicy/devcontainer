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
