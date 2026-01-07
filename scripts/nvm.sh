#!/usr/bin/env bash

set -e

# 安装 nvm
curl https://gitee.com/mirrors/nvm/raw/master/install.sh | bash

# 加载 nvm
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

# 安装最新的 LTS 版本
nvm install --lts

# 设置为默认版本
nvm alias default lts/*
nvm use default

# 配置 npm 镜像源
npm config set registry https://registry.npmmirror.com

# 安装 Claude Code
npm install -g @anthropic-ai/claude-code

echo "✅ Node.js LTS 安装完成并设为默认"
echo "✅ npm 镜像已设置为: https://registry.npmmirror.com"
echo "✅ Claude Code 已安装"