#!/usr/bin/env bash

# Check if SDKMAN is already installed
if [ -d "$HOME/.sdkman" ]; then
    echo "✓ SDKMAN is already installed at $HOME/.sdkman"
    echo "  Skipping SDKMAN installation."
else
    echo "Installing SDKMAN (Software Development Kit Manager)..."
    curl -s "https://get.sdkman.io" | bash
    echo "✓ SDKMAN installation completed"
fi