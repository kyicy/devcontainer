FROM debian:trixie

ENV DEBIAN_FRONTEND=noninteractive

# Use Aliyun mirror for faster downloads in China
RUN mv /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
COPY docker/aliyun.sources /etc/apt/sources.list.d/aliyun.sources

# Install base dependencies
RUN apt-get update && \
    apt-get install -y -q --no-install-recommends \
        apt-transport-https \
        ca-certificates \
        git \
        zsh \
        sudo && \
    rm -rf /var/lib/apt/lists/*

# Create admin user with sudo privileges (no password required)
RUN useradd -ms /bin/bash admin && \
    usermod -aG sudo admin && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER admin
WORKDIR /home/admin

# Set up development environment variables
ENV SHELL=/usr/bin/zsh \
    NVM_NODEJS_ORG_MIRROR="https://mirrors.ustc.edu.cn/node" \
    RUSTUP_DIST_SERVER="https://rsproxy.cn" \
    RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup" \
    GOPROXY="https://mirrors.aliyun.com/goproxy/,direct" \
    GVM_GO_BINARY_URL="https://mirrors.aliyun.com/golang/" \
    GVM_GO_SOURCE_URL="https://gitee.com/mirrors/go" \
    PATH="$PATH:/home/admin/.dotnet:/home/admin/.dotnet/tools" \
    DOTNET_ROOT="/home/admin/.dotnet"

# Install oh-my-zsh with shallow clone to reduce download size
RUN git clone --depth 1 https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git && \
    cd ohmyzsh/tools && \
    REMOTE=https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git sh install.sh && \
    cd /home/admin && \
    rm -rf ohmyzsh

# Install Rust configuration
RUN mkdir -p /home/admin/.cargo
COPY docker/cargo.toml /home/admin/.cargo/config.toml

# Copy setup scripts (keep this at the end for better layer caching)
COPY docker/scripts /home/admin/scripts