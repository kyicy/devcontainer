#!/usr/bin/env bash

set -e

# Install base development dependencies
echo ""
echo "📦 Installing base development dependencies..."
sudo apt-get update

sudo apt-get install -y -q \
    # Build tools
    build-essential cmake make git \
    # CLI utilities
    fzf fd-find jq curl wget zsh \
    # System & Security
    apt-transport-https ca-certificates \
    # Network tools
    netcat-openbsd openssh-client

echo "✓ Base development dependencies installed"

# Install oh-my-zsh with shallow clone to reduce download size
if [ -d "$HOME/.oh-my-zsh" ]; then
    echo "✓ oh-my-zsh is already installed at $HOME/.oh-my-zsh"
    echo "  Skipping oh-my-zsh installation."
else
    echo "Installing oh-my-zsh..."
    cd /home/admin
    git clone --depth 1 https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git && \
        cd ohmyzsh/tools && \
        REMOTE=https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git sh install.sh
    echo "✓ oh-my-zsh installation completed"
fi
