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

## 脚本说明

### `/var/scripts/set_up_system.sh` - 系统级配置

安装系统级别的开发工具和环境：

- 📦 **基础开发工具**: `build-essential`, `cmake`, `make`, `git`
- 🔧 **CLI 工具**: `fzf`, `fd-find`, `jq`, `curl`, `wget`, `zsh`
- 🌐 **网络工具**: `netcat-openbsd`, `openssh-client`
- 🐚 **oh-my-zsh**: 使用清华镜像加速安装

### `/var/scripts/set_up_user.sh` - 用户级配置

配置用户级别的开发环境和工具链：

- 🔐 **权限修复**: 修复 `/home/admin` 目录的所有权
- 🌐 **代理配置**: 添加 `set_proxy`/`unset_proxy` 函数用于代理切换
- 📊 **环境变量**: 配置 Rust 和 Go 的国内镜像源
- ⚙️ **mise**: 统一的开发工具版本管理器（支持别名 `x`）
- 🦀 **Rust**: 通过 rustup 安装，配置 rsproxy 国内镜像
- 🐍 **uv**: Python 包管理器，配置 USTC 镜像源

## 配置说明

- **镜像**: `ghcr.io/kyicy/devcontainer:latest`
- **工作目录**: `/home/admin`
- **用户**: `admin`
- **数据持久化**: 通过 Docker volume 实现
