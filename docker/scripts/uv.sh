#!/usr/bin/env bash

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
