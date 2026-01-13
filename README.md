# DevContainer 开发环境

一个为国内开发者优化的 DevContainer 解决方案，提供开箱即用的多语言开发环境。

## 📦 项目组成

本项目包含两个部分：

### 1. DevContainer Docker 镜像
基于 Debian Trixie 的开发容器镜像，预配置了完整的开发工具链。

### 2. devinit CLI 工具
用于快速初始化项目 DevContainer 配置的命令行工具。

---

## 🐳 Docker 镜像

### 特性

- 🐧 **基础镜像**: Debian Trixie
- 🇨🇳 **国内优化**: 预配置阿里云、清华大学等国内镜像源
- 👤 **用户配置**: 预创建 `admin` 用户，配置免密 sudo
- 🐚 **Shell 环境**: Oh My Zsh（清华大学镜像）
- 🔧 **多语言支持**: Node.js、Go、Rust、Python、.NET、Java

### 支持的开发环境

| 语言/工具 | 安装脚本 | 镜像源 |
|-----------|----------|--------|
| **Node.js** | `nvm.sh` | 中科大镜像 |
| **Go** | `gvm.sh` | 阿里云镜像 |
| **Rust** | `rustup.sh` | rsproxy 国内镜像 |
| **Python** | `uv.sh` | - |
| **.NET** | `dotnet.sh` | - |
| **Java** | `sdkman.sh` | - |

### 构建镜像

```bash
docker build -t ghcr.io/kyicy/devcontainer:latest .
```

### 使用镜像

在项目的 `.devcontainer/devcontainer.json` 中引用：

```json
{
  "image": "ghcr.io/kyicy/devcontainer:latest"
}
```

---

## 🚀 devinit CLI 工具

### 安装

```bash
cd devinit
go build -o devinit
sudo mv devinit /usr/local/bin/
```

### 功能

#### 1. 初始化项目

```bash
# 交互式模式
devinit init

# 非交互模式
devinit init --name myproject \
  --git-user "Your Name" \
  --git-email "you@example.com" \
  --non-interactive
```

**生成的文件结构**：

```
.devcontainer/
├── devcontainer.json          # DevContainer 配置
├── docker-compose.yml         # Docker Compose 配置
└── mapping/
    ├── .cam/                  # Claude 配置映射
    ├── .claude/               # Claude 数据映射
    ├── devcontainer-dependencies  # 项目依赖安装脚本
    ├── post-create.sh         # 容器创建后执行脚本
    └── .zsh_history           # Zsh 历史记录映射
```

#### 2. 管理配置

```bash
# 查看配置
devinit config view

# 设置环境变量
devinit config set-env NODE_ENV production

# 添加 VS Code 扩展
devinit config add-extension golang.go
```

### 命令参数

#### init 命令

| 参数 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--name` | `-n` | - | 项目名称 |
| `--workspace` | `-w` | `/home/admin/gopath/src` | 工作目录 |
| `--user` | `-u` | `admin` | 容器用户 |
| `--git-email` | - | - | Git 邮箱 |
| `--git-user` | - | - | Git 用户名 |
| `--github-token` | - | - | GitHub Token |
| `--git-branch` | - | `master` | Git 默认分支 |
| `--github-proxy` | - | `http://host.docker.internal:7890` | GitHub 代理 |
| `--non-interactive` | `-y` | `false` | 非交互模式 |

---

## 📁 项目结构

```
.
├── Dockerfile              # Docker 镜像定义
├── aliyun.sources          # 阿里云 APT 源配置
├── cargo.toml              # Cargo 国内镜像配置
├── scripts/                # 开发工具安装脚本
│   ├── nvm.sh             # Node.js (nvm) 安装
│   ├── gvm.sh             # Go (gvm) 安装
│   ├── rustup.sh          # Rust 安装
│   ├── uv.sh              # Python (uv) 安装
│   ├── dotnet.sh          # .NET 安装
│   ├── sdkman.sh          # Java (SDKMAN) 安装
│   └── devdep.sh          # 系统依赖安装
└── devinit/               # CLI 工具
    ├── main.go            # 入口文件
    ├── cmd/               # CLI 命令定义
    │   ├── root.go        # 根命令
    │   ├── init.go        # 初始化命令
    │   └── config.go      # 配置管理命令
    └── pkg/               # 核心逻辑包
        ├── config/        # 配置读写
        ├── generator/     # 文件生成器
        └── util/          # 工具函数
```

---

## 🌟 特色功能

### 1. 国内镜像优化

所有工具都配置了国内镜像源，确保快速下载：

- **APT**: 阿里云镜像
- **npm**: npmmirror 镜像
- **Go**: 阿里云 goproxy
- **Rust**: rsproxy 镜像
- **Oh My Zsh**: 清华大学镜像

### 2. GitHub 代理支持

自动配置 GitHub 代理，解决访问问题：

```bash
# 默认代理配置
git config --global http.https://github.com.proxy http://host.docker.internal:7890
```

### 3. GitHub 认证

支持通过环境变量配置 GitHub Token：

```bash
devinit init --github-token your_token_here
```

### 4. 开箱即用

容器启动后自动执行：

- 修复文件权限
- 配置 Git 用户信息
- 安装系统依赖（首次）
- 加载项目依赖配置

---

## 📝 使用示例

### 初始化一个 Go 项目

```bash
# 1. 使用 devinit 初始化
devinit init \
  --name my-go-project \
  --git-user "Your Name" \
  --git-email "you@example.com"

# 2. 添加 Go 相关扩展
devinit config add-extension golang.go
devinit config add-extension eamodio.gitlens

# 3. 在 VS Code 中重新打开容器
# 按 F1 -> "Dev Containers: Rebuild Container"
```

### 初始化一个全栈项目

```bash
devinit init \
  --name fullstack-app \
  --git-user "Your Name" \
  --git-email "you@example.com"

# 编辑 .devcontainer/mapping/devcontainer-dependencies
# 取消注释所需的开发环境脚本
```

---

## 🔧 配置文件说明

### devcontainer.json

主要配置文件，定义容器行为：

```json
{
  "name": "Project Dev Container",
  "dockerComposeFile": "docker-compose.yml",
  "service": "project_dev",
  "workspaceFolder": "/home/admin/gopath/src/project",
  "postCreateCommand": "bash $HOME/scripts/post-create.sh",
  "remoteUser": "admin",
  "customizations": {
    "vscode": {
      "extensions": ["golang.go", "eamodio.gitlens"]
    }
  }
}
```

### docker-compose.yml

Docker Compose 配置，定义服务、卷、网络：

```yaml
services:
  project_dev:
    image: ghcr.io/kyicy/devcontainer:latest
    volumes:
      - project_code:/home/admin/gopath
      - ./mapping/.cam:/home/admin/.cam
      # ... 更多映射
```

### devcontainer-dependencies

项目特定的依赖安装脚本：

```bash
#!/usr/bin/env bash
set -e

echo "🔧 安装项目所需的开发环境..."

# === 前端开发 ===
bash ~/scripts/nvm.sh

# === 后端开发 (Go) ===
bash ~/scripts/gvm.sh

echo "✅ 项目依赖安装完成"
```

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

---

## 📄 许可证

[LICENSE](LICENSE)
