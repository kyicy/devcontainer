package generator

import (
	"encoding/json"
	"fmt"
	"os"
)

// generateDevcontainerJSON 生成 devcontainer.json
func generateDevcontainerJSON(path string, config *DevContainerConfig, serviceName string) error {
	containerEnv := map[string]string{}

	devContainer := DevContainerJSON{
		Name:              config.ProjectName + " Dev Container",
		DockerComposeFile: "docker-compose.yml",
		Service:           serviceName,
		WorkspaceFolder:   "/home/admin",
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
      - my_code:/home/admin
    # 容器启动命令 - 先修复权限再启动
    command: sleep infinity

# 定义 volume
volumes:
  my_code:
    external: true
`, serviceName, config.DockerImage)

	filePath := path + "/docker-compose.yml"
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("写入 docker-compose.yml 失败: %w", err)
	}

	return nil
}
