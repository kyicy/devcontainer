#!/usr/bin/env bash

# Check if GVM is already installed
if [ -d "$HOME/.gvm" ]; then
    echo "✓ GVM is already installed at $HOME/.gvm"
    echo "  Skipping GVM installation."
else
    echo "Installing GVM (Go Version Manager)..."
    curl -s -S -L "https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer" | bash
    echo "✓ GVM installation completed"
fi