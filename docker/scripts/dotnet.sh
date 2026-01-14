#!/usr/bin/env bash

# Check if .NET is already installed
if command -v dotnet &> /dev/null; then
    echo "✓ .NET is already installed (dotnet version: $(dotnet --version))"
    echo "  Skipping .NET installation."
else
    echo "Installing .NET SDK (STS channel)..."
    curl -s https://builds.dotnet.microsoft.com/dotnet/scripts/v1/dotnet-install.sh | bash -s -- --channel STS
    echo "✓ .NET installation completed"
fi
