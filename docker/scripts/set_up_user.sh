#!/usr/bin/env bash

set -e

echo "🚀 Starting development environment setup..."

# Fix ownership of /home/admin directory
sudo chown -R admin:admin /home/admin

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


# Configure development proxy environment variables in .zshrc
DEV_PROXY_MARKER="# Development proxy environment variables - managed by devdep.sh"
if grep -q "$DEV_PROXY_MARKER" "$ZSHRC" 2>/dev/null; then
    echo "✓ Development proxy environment variables already configured in .zshrc"
    echo "  Skipping development proxy configuration."
else
    echo ""
    echo "🔧 Adding development proxy environment variables to .zshrc..."
    cat >> "$ZSHRC" << 'EOF'

# Development proxy environment variables - managed by devdep.sh
export RUSTUP_DIST_SERVER="https://rsproxy.cn"
export RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
export GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
EOF
    echo "✓ Development proxy environment variables added to .zshrc"
fi

# Configure mise in .zshrc
MISE_MARKER="# mise - managed by devdep.sh"
if grep -q "$MISE_MARKER" "$ZSHRC" 2>/dev/null; then
    echo "  Skipping mise configuration."
else
    curl https://mise.run | sh
    cat >> "$ZSHRC" << 'EOF'

# mise - managed by devdep.sh
eval "$(/home/admin/.local/bin/mise activate zsh)"
alias x="mise x --"

EOF
fi

# Check if Rust is already installed
if command -v rustc &> /dev/null && command -v cargo &> /dev/null; then
    echo "✓ Rust is already installed (rustc version: $(rustc --version))"
    echo "  Skipping Rust installation."
else
    echo "Installing Rust via rustup..."
    curl --proto '=https' --tlsv1.2 -sSf https://rsproxy.cn/rustup-init.sh | sh
    echo "✓ Rust installation completed"
fi

# Configure Cargo to use Chinese mirrors
CARGO_CONFIG="$HOME/.cargo/config.toml"
if [ -f "$CARGO_CONFIG" ]; then
    echo "✓ Cargo config already exists at $CARGO_CONFIG"
    echo "  Skipping Cargo configuration."
else
    echo "Configuring Cargo to use Chinese mirrors..."
    mkdir -p $HOME/.cargo
    cat > "$CARGO_CONFIG" << 'EOF'
[source.crates-io]
replace-with = 'rsproxy-sparse'
[source.rsproxy]
registry = "https://rsproxy.cn/crates.io-index"
[source.rsproxy-sparse]
registry = "sparse+https://rsproxy.cn/index/"
[registries.rsproxy]
index = "https://rsproxy.cn/crates.io-index"
[net]
git-fetch-with-cli = true
EOF
    echo "✓ Cargo configuration completed"
fi


# Check if uv is already installed
if command -v uv &> /dev/null; then
    echo "✓ uv is already installed (uv version: $(uv --version))"
    echo "  Skipping uv installation."
else
    echo "Installing uv (Python package manager)..."
    curl -LsSf https://astral.sh/uv/install.sh | sh
    echo "✓ uv installation completed"
fi

# Configure uv to use USTC mirror
UV_CONFIG="$HOME/.config/uv/uv.toml"
if [ -f "$UV_CONFIG" ]; then
    echo "✓ uv config already exists at $UV_CONFIG"
    echo "  Skipping uv configuration."
else
    echo "Configuring uv to use USTC mirror..."
    mkdir -p $HOME/.config/uv
    cat > "$UV_CONFIG" << 'EOF'
[[index]]
url = "https://mirrors.ustc.edu.cn/pypi/simple"
default = true
EOF
    echo "✓ uv configuration completed"
fi


# Final ownership fix
sudo chown -R admin:admin /home/admin

echo ""
echo "✅ Development environment setup completed!"
echo ""