package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// DevContainerConfig 表示从 devcontainer.json 加载的配置
type DevContainerConfig struct {
	Name              string                 `json:"name"`
	DockerComposeFile string                 `json:"dockerComposeFile"`
	Service           string                 `json:"service"`
	WorkspaceFolder   string                 `json:"workspaceFolder"`
	PostCreateCommand string                 `json:"postCreateCommand"`
	ContainerEnv      map[string]string      `json:"containerEnv"`
	RemoteUser        string                 `json:"remoteUser"`
	Image             string                 `json:"image"`
	Customizations    map[string]interface{} `json:"customizations"`
}

// DockerComposeConfig 表示 docker-compose.yml 的配置
type DockerComposeConfig struct {
	Services map[string]ServiceConfig `yaml:"services"`
	Volumes  map[string]interface{}   `yaml:"volumes"`
}

// ServiceConfig 表示 Docker 服务配置
type ServiceConfig struct {
	Image       string            `yaml:"image"`
	Volumes     []string          `yaml:"volumes"`
	Command     string            `yaml:"command"`
	Environment map[string]string `yaml:"environment"`
}

// LoadConfig 从项目目录加载配置
func LoadConfig(projectPath string) (*DevContainerConfig, error) {
	configPath := filepath.Join(projectPath, ".devcontainer", "devcontainer.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config DevContainerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}

// SaveConfig 保存配置到文件
func SaveConfig(projectPath string, config *DevContainerConfig) error {
	configPath := filepath.Join(projectPath, ".devcontainer", "devcontainer.json")

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	return nil
}

// SetEnvVariable 设置环境变量
func SetEnvVariable(projectPath string, key string, value string) error {
	config, err := LoadConfig(projectPath)
	if err != nil {
		return err
	}

	if config.ContainerEnv == nil {
		config.ContainerEnv = make(map[string]string)
	}

	config.ContainerEnv[key] = value

	return SaveConfig(projectPath, config)
}

// AddVSCodeExtension 添加 VS Code 扩展
func AddVSCodeExtension(projectPath string, extensionID string) error {
	config, err := LoadConfig(projectPath)
	if err != nil {
		return err
	}

	if config.Customizations == nil {
		config.Customizations = make(map[string]interface{})
	}

	vscodeConfig, ok := config.Customizations["vscode"].(map[string]interface{})
	if !ok {
		vscodeConfig = make(map[string]interface{})
		config.Customizations["vscode"] = vscodeConfig
	}

	extensions, ok := vscodeConfig["extensions"].([]interface{})
	if !ok {
		extensions = []interface{}{}
	}

	// 检查是否已存在
	for _, ext := range extensions {
		if extStr, ok := ext.(string); ok && extStr == extensionID {
			return fmt.Errorf("扩展 %s 已存在", extensionID)
		}
	}

	extensions = append(extensions, extensionID)
	vscodeConfig["extensions"] = extensions

	return SaveConfig(projectPath, config)
}
