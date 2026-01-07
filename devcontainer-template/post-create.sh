#!/usr/bin/env bash

set -e

echo "🚀 开始配置开发环境..."

# 检查代码是否已经存在
if [ -z "$(ls -A /workspace 2>/dev/null)" ]; then
    echo "📁 Workspace 为空，准备加载代码..."

    # 方式1: 从 Git 仓库克隆（推荐）
    if [ -n "$GIT_REPO_URL" ]; then
        echo "📥 正在从 Git 仓库克隆代码..."
        git clone "$GIT_REPO_URL" /workspace
        if [ -n "$GIT_BRANCH" ]; then
            echo "🌿 切换到分支: $GIT_BRANCH"
            git -C /workspace checkout "$GIT_BRANCH"
        fi
    else
        # 方式2: 提示用户手动操作
        echo "⚠️  未检测到 GIT_REPO_URL 环境变量"
        echo ""
        echo "请选择以下方式之一来加载代码："
        echo ""
        echo "方式 1: 从 Git 克隆"
        echo "  git clone <your-repo-url> /workspace"
        echo ""
        echo "方式 2: 复制本地文件（如果在宿主机操作）"
        echo "  docker cp /path/to/local/code devcontainer-app:/workspace/"
        echo ""
        echo "方式 3: 挂载现有备份（如果有备份）"
        echo "  docker run --rm -v devcontainer-workspace:/data -v \$(pwd):/backup debian:trixie \\"
        echo "    tar xzf /backup/workspace-backup.tar.gz -C /data"
        echo ""
        echo "方式 4: 重新创建容器并绑定本地目录"
        echo "  修改 docker-compose.yml，将 volume 改为："
        echo "    volumes:"
        echo "      - .:/workspace"
    fi
else
    echo "✅ Workspace 已有代码，跳过加载步骤"
fi

# 检查是否已经安装过开发依赖
if [ ! -f "$HOME/.devcontainer-initialized" ]; then
    echo "📦 首次初始化，安装系统依赖..."
    bash ~/scripts/devdep.sh
    touch "$HOME/.devcontainer-initialized"
fi

# 检查项目是否有特定的依赖配置文件
if [ -f "/workspace/.devcontainer-dependencies" ]; then
    echo "📋 检测到项目依赖配置，正在安装..."
    source /workspace/.devcontainer-dependencies
else
    echo "💡 提示: 创建 /workspace/.devcontainer-dependencies 文件来定义项目需要的语言环境"
fi

echo "✅ 开发环境配置完成！"
