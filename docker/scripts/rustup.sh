#!/usr/bin/env bash

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