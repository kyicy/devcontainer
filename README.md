# devcontainer

> 一个专为国内开发者优化的多功能开发环境容器

## 简介

这是一个标准化的多语言开发环境容器，提供了开箱即用的开发工具链。项目针对中国开发者进行了深度优化，使用国内镜像源，大幅提升包管理和依赖下载速度。

## 特性

- 🚀 **开箱即用** - 预配置主流开发语言和工具
- 🇨🇳 **国内优化** - 全链路使用国内镜像源（阿里云、清华、中科大等）
- 🏗️ **多架构支持** - 支持 AMD64 和 ARM64 架构
- 👤 **用户友好** - 预配置 admin 用户，免密 sudo，集成 Oh My Zsh
- 🔄 **CI/CD 就绪** - 完整的 GitHub Actions 工作流

## 支持的语言和环境

| 语言/工具 | 版本管理器 | 说明 |
|----------|-----------|------|
| **Go** | GVM | Go 版本管理，使用阿里云镜像 |
| **Rust** | rustup | Rust 工具链，使用 rsproxy 镜像 |
| **Node.js** | NVM | Node 版本管理，使用中科大镜像 |
| **Python** | uv | 现代 Python 包管理器，使用中科大 PyPI 镜像 |
| **.NET** | - | .NET SDK (STS channel) |
| **Java** | SDKMAN | Java 工具链管理器 |

## 快速开始

### VS Code Dev Containers

1. 安装 [Dev Containers 扩展](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
2. 在项目根目录创建 `.devcontainer/devcontainer.json`：

```json
{
  "image": "ghcr.io/kyicy/devcontainer:latest",
  "remoteUser": "admin"
}
```

3. 按 `F1` 选择 `Dev Containers: Reopen in Container`

### GitHub Codespaces

直接使用此镜像作为 Codespaces 的基础镜像。

### Docker 直接使用

```bash
docker pull ghcr.io/kyicy/devcontainer:latest

docker run -it --rm \
  --cap-add=SYS_PTRACE \
  --security-opt seccomp=unconfined \
  -v $(pwd):/workspace \
  ghcr.io/kyicy/devcontainer:latest
```

### Docker Compose

```yaml
services:
  dev:
    image: ghcr.io/kyicy/devcontainer:latest
    volumes:
      - .:/workspace
    working_dir: /workspace
    command: /bin/zsh
```

## 镜像源配置

项目使用以下国内镜像源以提升访问速度：

- **APT 包管理**：阿里云 Debian 镜像
- **Go 模块**：阿里云 Go 代理 + Gitee 源
- **Rust crates**：rsproxy.cn (字节跳动)
- **Node.js/npm**：中科大镜像
- **Python PyPI**：中科大镜像

## 环境变量

容器预配置了以下环境变量：

```bash
# Go 代理
GOPROXY=https://goproxy.cn,direct
GO111MODULE=on

# Rust 镜像
RUSTUP_DIST_SERVER=https://rsproxy.cn/dist
RUSTUP_UPDATE_ROOT=https://rsproxy.cn/rustup

# Node.js 镜像
NVM_NODEJS_ORG_MIRROR=https://mirrors.ustc.edu.cn/npm/node-snapshot
```

## 默认用户

- **用户名**：admin
- **密码**：无（使用 SSH 密钥认证）
- **权限**：sudo 免密
- **Shell**：Zsh with Oh My Zsh

## 构建和发布

### 本地构建

```bash
# AMD64 架构
docker build -t devcontainer:latest .

# 多架构构建
docker buildx build --platform linux/amd64,linux/arm64 -t devcontainer:latest .
```

### 发布流程

项目使用 GitHub Actions 自动构建和发布：

- 推送到 `main` 分支：自动构建并打上 `latest` 标签
- 推送版本标签：自动构建并打上对应版本标签
- 镜像发布到：`ghcr.io/kyicy/devcontainer`

## 技术栈

- **基础镜像**：debian:trixie (testing)
- **Shell**：Zsh + Oh My Zsh
- **CI/CD**：GitHub Actions
- **镜像仓库**：GitHub Container Registry

## 许可证

[MIT License](LICENSE)

## 贡献

欢迎提交 Issue 和 Pull Request！