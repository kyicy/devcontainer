# Devcontainer

快速初始化新项目的 devcontainer 配置，提供一致的开发环境。

## 使用教程

### 1. 创建数据卷

```sh
docker volume create my_code
```

### 2. 配置 devcontainer

在你的项目中创建 `.devcontainer` 目录，并添加以下文件：

#### `.devcontainer/devcontainer.json`

```json
{
  "name": "Your Project Name",
  "dockerComposeFile": "docker-compose.yml",
  "service": "dev",
  "workspaceFolder": "/home/admin/",
  "remoteUser": "admin"
}
```

#### `.devcontainer/docker-compose.yml`

```yaml
services:
  dev:
    image: ghcr.io/kyicy/devcontainer:latest
    # 使用独立的 volume 存储代码
    volumes:
      - my_code:/home/admin
    # 可选：自定义工作目录
    working_dir: /home/admin

# 定义 volume
volumes:
  my_code:
    external: true
```
