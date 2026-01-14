#!/usr/bin/env bash

set -e

echo "🚀 Starting development environment setup..."

# Fix ownership of /home/admin directory
sudo chown -R admin:admin /home/admin

# Configure Git user information
if git config --global user.email > /dev/null 2>&1; then
    echo "✓ Git user already configured"
    echo "  Skipping Git configuration."
else
    echo ""
    echo "🔧 Git User Configuration"
    echo "========================="
    read -p "Enter your Git email: " git_email
    read -p "Enter your Git name: " git_name

    if [ -n "$git_email" ] && [ -n "$git_name" ]; then
        git config --global user.email "$git_email"
        git config --global user.name "$git_name"
        echo "✓ Git user configuration completed"
    else
        echo "⚠️  Skipping Git configuration (empty input)"
    fi
fi

# Configure GitHub proxy
read -p "Enter GitHub proxy URL (press Enter to skip): " github_proxy
if [ -n "$github_proxy" ]; then
    echo "Configuring GitHub proxy to $github_proxy..."
    git config --global http.https://github.com.proxy "$github_proxy"
    echo "✓ GitHub proxy configured"
else
    echo "⚠️  Skipping GitHub proxy configuration"
fi

# Configure GitHub authentication
if [ -f ~/.git-credentials ]; then
    echo "✓ GitHub authentication already configured"
    echo "  Skipping GitHub authentication setup."
else
    echo ""
    echo "🔐 GitHub Authentication Configuration"
    echo "====================================="
    echo "It's recommended to configure a GitHub Token for Git operations"
    echo ""
    echo "To get a GitHub Token:"
    echo "1. Visit https://github.com/settings/tokens"
    echo "2. Click 'Generate new token (classic)'"
    echo "3. Select appropriate permissions (at minimum, repo scope)"
    echo "4. Generate and copy the token"
    echo ""

    read -p "Configure GitHub Token now? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        while true; do
            read -p "Enter your GitHub Token: " -s github_token
            echo
            if [ -n "$github_token" ]; then
                # Configure credential helper
                git config --global credential.helper store

                # Save token
                echo "https://oauth2:${github_token}@github.com" > ~/.git-credentials
                chmod 600 ~/.git-credentials

                echo "✓ GitHub authentication configured"

                # Test authentication
                echo "🧪 Testing GitHub authentication..."
                if git ls-remote https://github.com/ >/dev/null 2>&1; then
                    echo "✓ GitHub authentication test successful"
                else
                    echo "⚠️  Warning: GitHub authentication test failed, please check your token"
                    rm -f ~/.git-credentials
                fi
                break
            else
                echo "❌ Token cannot be empty, please try again"
            fi
        done
    else
        echo "⚠️  Skipping GitHub Token configuration"
        echo "   You can configure it later by running:"
        echo "   git config --global credential.helper store"
        echo '   echo "https://TOKEN@github.com" > ~/.git-credentials'
    fi
fi

# Install base development dependencies
echo ""
echo "📦 Installing base development dependencies..."
sudo apt-get update

sudo apt-get install -y -q --no-install-recommends \
    build-essential \
    curl mercurial make binutils bison gcc bsdmainutils \
    libssl-dev \
    wget \
    zip unzip \
    tar \
    libicu-dev

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
        REMOTE=https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git sh install.sh && \
        cd /home/admin && \
        rm -rf ohmyzsh
    echo "✓ oh-my-zsh installation completed"
fi

# Check if project has specific dependencies configuration
if [ -f "$HOME/scripts/devcontainer-dependencies" ]; then
    echo ""
    echo "📋 Detected project dependencies configuration, installing..."
    source "$HOME/scripts/devcontainer-dependencies"
fi

# Final ownership fix
sudo chown -R admin:admin /home/admin

echo ""
echo "✅ Development environment setup completed!"
echo ""
echo "Summary:"
if git config --global user.email > /dev/null 2>&1; then
    echo "  ✓ Git configured: $(git config --global user.name) <$(git config --global user.email)>"
fi
echo "  ✓ Base dependencies installed"
echo "  ✓ oh-my-zsh installed"
if [ -f "$HOME/scripts/devcontainer-dependencies" ]; then
    echo "  ✓ Project dependencies installed"
fi
