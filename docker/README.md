# Docker 镜像

此目录包含 DevContainer Docker 镜像的所有相关文件。

## 📁 文件说明

- **Dockerfile** - 主镜像定义文件
- **aliyun.sources** - 阿里云 APT 源配置（用于加速软件包下载）
- **cargo.toml** - Cargo 国内镜像配置（用于加速 Rust 依赖下载）
- **.dockerignore** - Docker 构建时的排除文件列表
- **scripts/** - 开发工具安装脚本目录
  - `nvm.sh` - Node.js (nvm) 安装
  - `gvm.sh` - Go (gvm) 安装
  - `rustup.sh` - Rust 安装
  - `uv.sh` - Python (uv) 安装
  - `dotnet.sh` - .NET 安装
  - `sdkman.sh` - Java (SDKMAN) 安装
  - `devdep.sh` - 系统依赖安装

## 🔨 构建镜像

从项目根目录构建：

```bash
docker build -t ghcr.io/kyicy/devcontainer:latest -f docker/Dockerfile .
```

## 📦 镜像特性

- 基于 Debian Trixie
- 预配置国内镜像源（阿里云、清华大学等）
- 预创建 `admin` 用户，配置免密 sudo
- 预装 Oh My Zsh
- 支持多语言开发环境（Node.js、Go、Rust、Python、.NET、Java）

## 🚀 CI/CD

GitHub Workflow 会在以下情况自动构建镜像：

1. 推送到 `main` 分支且修改了 `docker/` 目录
2. 创建版本标签（如 `v1.0.0`）
3. 手动触发 Workflow

详见：[.github/workflows/docker.yml](../.github/workflows/docker.yml)
