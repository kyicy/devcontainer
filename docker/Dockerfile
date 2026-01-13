FROM debian:trixie

ENV DEBIAN_FRONTEND=noninteractive

RUN mv /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
COPY docker/aliyun.sources /etc/apt/sources.list.d/aliyun.sources

# Install base dependencies
RUN apt-get update && apt-get install -y -q --no-install-recommends apt-transport-https ca-certificates git zsh sudo \
    && rm -rf /var/lib/apt/lists/*


# Create ubuntu user with sudo privileges
RUN useradd -ms /bin/bash admin && \
    usermod -aG sudo admin
# New added for disable sudo password
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER admin
WORKDIR /home/admin
ENV HOME="/home/admin"

# oh my zsh
RUN git clone https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git && cd ohmyzsh/tools && REMOTE=https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git sh install.sh    
ENV SHELL=/usr/bin/zsh

# setup scripts
COPY docker/scripts $HOME/scripts

## nodejs
# Install nvm
ENV NVM_NODEJS_ORG_MIRROR="https://mirrors.ustc.edu.cn/node"

# Install rust
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
RUN mkdir -p $HOME/.cargo
COPY docker/cargo.toml $HOME/.cargo/config.toml

# golang env
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
ENV GVM_GO_BINARY_URL="https://mirrors.aliyun.com/golang/"
ENV GVM_GO_SOURCE_URL="https://gitee.com/mirrors/go"

# dotnet env
ENV PATH="$PATH:$HOME/.dotnet:$HOME/.dotnet/tools" 
ENV DOTNET_ROOT="$HOME/.dotnet"