package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// UserConfig 存储在用户主目录的默认配置
type UserConfig struct {
	GitUser     string `json:"git_user"`
	GitEmail    string `json:"git_email"`
	GithubToken string `json:"github_token"`
	GitBranch   string `json:"git_branch"`
	GithubProxy string `json:"github_proxy"`
	RemoteUser  string `json:"remote_user"`
	Workspace   string `json:"workspace"`
}

// GetUserConfigPath 获取用户配置文件路径
func GetUserConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("无法获取用户主目录: %w", err)
	}
	return filepath.Join(home, ".devinit.json"), nil
}

// LoadUserConfig 从用户主目录加载配置
func LoadUserConfig() (*UserConfig, error) {
	configPath, err := GetUserConfigPath()
	if err != nil {
		return nil, err
	}

	// 如果配置文件不存在,返回默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &UserConfig{
			GitBranch:   "master",
			GithubProxy: "http://host.docker.internal:7890",
			RemoteUser:  "admin",
			Workspace:   "/home/admin/gopath/src",
		}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config UserConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 填充默认值
	if config.GitBranch == "" {
		config.GitBranch = "master"
	}
	if config.GithubProxy == "" {
		config.GithubProxy = "http://host.docker.internal:7890"
	}
	if config.RemoteUser == "" {
		config.RemoteUser = "admin"
	}
	if config.Workspace == "" {
		config.Workspace = "/home/admin/gopath/src"
	}

	return &config, nil
}

// SaveUserConfig 保存配置到用户主目录
func SaveUserConfig(config *UserConfig) error {
	configPath, err := GetUserConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	return nil
}

// IsConfigExists 检查用户配置文件是否存在
func IsConfigExists() bool {
	configPath, err := GetUserConfigPath()
	if err != nil {
		return false
	}

	_, err = os.Stat(configPath)
	return err == nil
}
