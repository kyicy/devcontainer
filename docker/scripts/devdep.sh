#!/usr/bin/env bash

set -e

echo "🚀 Starting development environment setup..."

# Fix ownership of /home/admin directory
sudo chown -R admin:admin /home/admin

# Install base development dependencies
echo ""
echo "📦 Installing base development dependencies..."
sudo apt-get update

sudo apt-get install -y -q --no-install-recommends \
    build-essential \
    mercurial make binutils bison gcc bsdmainutils \
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

# Configure proxy functions in .zshrc
ZSHRC="$HOME/.zshrc"
PROXY_MARKER="# Proxy functions - managed by devdep.sh"

if grep -q "$PROXY_MARKER" "$ZSHRC" 2>/dev/null; then
    echo "✓ Proxy functions already configured in .zshrc"
    echo "  Skipping proxy functions configuration."
else
    echo ""
    echo "🔧 Adding proxy functions to .zshrc..."
    cat >> "$ZSHRC" << 'EOF'

# Proxy functions - managed by devdep.sh
set_proxy() {
    export https_proxy=http://host.docker.internal:7890 http_proxy=http://host.docker.internal:7890 all_proxy=socks5://host.docker.internal:7890
    echo "✓ Proxy set: https_proxy=http://host.docker.internal:7890"
}

unset_proxy() {
    unset https_proxy
    unset http_proxy
    unset all_proxy
    echo "✓ Proxy unset"
}
EOF
    echo "✓ Proxy functions added to .zshrc"
    echo "  Usage: set_proxy    - enable proxy"
    echo "        unset_proxy  - disable proxy"
fi


# Configure proxy functions in .zshrc
MISE_MARKER="# mise - managed by devdep.sh"
if ! grep -q "$MISE_MARKER" "$ZSHRC" 2>/dev/null; then
else
    
    cat >> "$ZSHRC" << 'EOF'

# mise - managed by devdep.sh
eval "$(/home/admin/.local/bin/mise activate zsh)"
alias x="mise x --"

EOF
fi


# Final ownership fix
sudo chown -R admin:admin /home/admin

echo ""
echo "✅ Development environment setup completed!"
echo ""