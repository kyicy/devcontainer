package generator

// DevContainerConfig 定义了 devcontainer 的配置
type DevContainerConfig struct {
	ProjectName string
	DockerImage string
	RemoteUser  string
	Extensions  []string
}

// DevContainerJSON devcontainer.json 结构
type DevContainerJSON struct {
	Name              string                 `json:"name"`
	DockerComposeFile string                 `json:"dockerComposeFile"`
	Service           string                 `json:"service"`
	WorkspaceFolder   string                 `json:"workspaceFolder"`
	ContainerEnv      map[string]string      `json:"containerEnv"`
	RemoteUser        string                 `json:"remoteUser"`
	Customizations    map[string]interface{} `json:"customizations"`
}
