#!/usr/bin/env bash

set -e

# Check if NVM is already installed
if [ -d "$HOME/.nvm" ]; then
    echo "✓ NVM is already installed at $HOME/.nvm"
    echo "  Skipping NVM installation."
else
    echo "Installing NVM (Node Version Manager)..."
    curl https://gitee.com/mirrors/nvm/raw/master/install.sh | bash
    echo "✓ NVM installation completed"
fi

# Load nvm
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

# Check if Node.js LTS is already installed
if command -v node &> /dev/null; then
    echo "✓ Node.js is already installed (node version: $(node --version))"
    echo "  Skipping Node.js LTS installation."
else
    echo "Installing Node.js LTS..."
    # Install latest LTS version
    nvm install --lts
    # Set as default version
    nvm alias default lts/*
    nvm use default
    echo "✓ Node.js LTS installation completed"
fi

# Configure npm registry mirror
NPM_REGISTRY=$(npm config get registry)
if [ "$NPM_REGISTRY" = "https://registry.npmmirror.com" ]; then
    echo "✓ npm registry is already configured to use npmmirror"
    echo "  Skipping npm configuration."
else
    echo "Configuring npm to use Chinese mirror..."
    npm config set registry https://registry.npmmirror.com
    echo "✓ npm registry configured to: https://registry.npmmirror.com"
fi

# Check if Claude Code is already installed
if command -v claude &> /dev/null; then
    echo "✓ Claude Code is already installed"
    echo "  Skipping Claude Code installation."
else
    echo "Installing Claude Code..."
    npm install -g @anthropic-ai/claude-code
    echo "✓ Claude Code installation completed"
fi

echo ""
echo "Summary:"
echo "  ✓ Node.js: $(node --version)"
echo "  ✓ npm registry: $(npm config get registry)"
if command -v claude &> /dev/null; then
    echo "  ✓ Claude Code: installed"
fi