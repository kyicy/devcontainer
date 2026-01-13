package generator

import (
	"encoding/json"
	"fmt"
	"os"
)

// generateDevcontainerJSON 生成 devcontainer.json
func generateDevcontainerJSON(path string, config *DevContainerConfig, serviceName string) error {
	containerEnv := map[string]string{
		"NODE_ENV":   "development",
		"GIT_EMAIL":  config.GitEmail,
		"GIT_USER":   config.GitUser,
		"GIT_BRANCH": config.GitBranch,
	}

	if config.GithubToken != "" {
		containerEnv["GITHUB_TOKEN"] = config.GithubToken
	}

	devContainer := DevContainerJSON{
		Name:              config.ProjectName + " Dev Container",
		DockerComposeFile: "docker-compose.yml",
		Service:           serviceName,
		WorkspaceFolder:   "/home/admin",
		PostCreateCommand: "bash $HOME/scripts/post-create.sh",
		ContainerEnv:      containerEnv,
		RemoteUser:        config.RemoteUser,
		Customizations: map[string]interface{}{
			"vscode": map[string]interface{}{
				"extensions": config.Extensions,
			},
		},
	}

	data, err := json.MarshalIndent(devContainer, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %w", err)
	}

	filePath := path + "/devcontainer.json"
	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

// generateDockerCompose 生成 docker-compose.yml
func generateDockerCompose(path string, config *DevContainerConfig, serviceName string) error {
	content := fmt.Sprintf(`services:
  %s:
    image: %s

    # 使用独立的 volume 存储代码
    volumes:
      # 方案1: 使用命名的 volume（推荐）
      - %s_code:/home/admin/gopath
      - ./mapping/.claude:/home/admin/.claude
      - ./mapping/devcontainer-dependencies:/home/admin/scripts/devcontainer-dependencies
      - ./mapping/.zsh_history:/home/admin/.zsh_history
      - ./mapping/post-create.sh:/home/admin/scripts/post-create.sh
    # 容器启动命令 - 先修复权限再启动
    command: sleep infinity

# 定义 volume
volumes:
  %s_code: {}
`, serviceName, config.DockerImage, config.ProjectName, config.ProjectName)

	filePath := path + "/docker-compose.yml"
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("写入 docker-compose.yml 失败: %w", err)
	}

	return nil
}

// generatePostCreateScript 生成 post-create.sh
func generatePostCreateScript(path string, config *DevContainerConfig) error {
	content := `#!/usr/bin/env bash

set -e

echo "🚀 开始配置开发环境..."

sudo chown -R admin:admin /home/admin

# 配置 Git 用户信息
git config --global user.email "$GIT_EMAIL"
git config --global user.name "$GIT_USER"

` + fmt.Sprintf(`# 配置 GitHub 代理
git config --global http.https://github.com.proxy %s

`, config.GithubProxy) + `# 配置 GitHub 认证（使用 GITHUB_TOKEN 环境变量）
if [ -n "$GITHUB_TOKEN" ]; then
    echo "🔐 配置 GitHub 认证..."
    git config --global credential.helper store
    # 使用 token 作为认证凭据
    echo "https://oauth2:${GITHUB_TOKEN}@github.com" > ~/.git-credentials
    chmod 600 ~/.git-credentials
    echo "✅ GitHub 认证配置完成"
else
    echo "⚠️  警告: 未设置 GITHUB_TOKEN 环境变量，Git 操作可能需要手动认证"
fi

# 检查是否已经安装过开发依赖
if [ ! -f "$HOME/.devcontainer-initialized" ]; then
    echo "📦 首次初始化，安装系统依赖..."
    bash ~/scripts/devdep.sh
    touch "$HOME/.devcontainer-initialized"
fi

# 检查项目是否有特定的依赖配置文件
if [ -f "$HOME/scripts/devcontainer-dependencies" ]; then
    echo "📋 检测到项目依赖配置，正在安装..."
    source /$HOME/scripts/devcontainer-dependencies
fi

sudo chown -R admin:admin /home/admin

echo "✅ 开发环境配置完成！"
`

	filePath := path + "/post-create.sh"
	if err := os.WriteFile(filePath, []byte(content), 0o755); err != nil {
		return fmt.Errorf("写入 post-create.sh 失败: %w", err)
	}

	return nil
}

// generateDependenciesFile 生成 devcontainer-dependencies
func generateDependenciesFile(path string) error {
	content := `#!/usr/bin/env bash
# 项目依赖配置
#
# 根据项目需要取消相应的注释

set -e

echo "🔧 安装项目所需的开发环境..."

# === 前端开发 ===
bash ~/scripts/nvm.sh

# === 后端开发 (Go) ===
# bash ~/scripts/gvm.sh

# === 后端开发 (Rust) ===
# bash ~/scripts/rustup.sh

# === Python 开发 ===
# bash ~/scripts/uv.sh

# === .NET 开发 ===
# bash ~/scripts/dotnet.sh

# === Java 开发 ===
# bash ~/scripts/sdkman.sh

echo "✅ 项目依赖安装完成"
`

	filePath := path + "/devcontainer-dependencies"
	if err := os.WriteFile(filePath, []byte(content), 0o755); err != nil {
		return fmt.Errorf("写入 devcontainer-dependencies 失败: %w", err)
	}

	return nil
}
