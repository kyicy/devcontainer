FROM debian:bookworm-20250908

ENV DEBIAN_FRONTEND=noninteractive

RUN mv /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak

RUN cat > /etc/apt/sources.list.d/aliyun.sources << EOF
Types: deb
URIs: http://mirrors.aliyun.com/debian
Suites: bookworm bookworm-updates
Components: main non-free non-free-firmware contrib
Signed-By: /usr/share/keyrings/debian-archive-keyring.gpg

Types: deb
URIs: http://mirrors.aliyun.com/debian-security
Suites: bookworm-security
Components: main non-free non-free-firmware contrib
Signed-By: /usr/share/keyrings/debian-archive-keyring.gpg
EOF

# Install base dependencies
RUN apt-get update && apt-get install -y -q --no-install-recommends \
    apt-transport-https \
    build-essential \
    ca-certificates \
    curl git mercurial make binutils bison gcc bsdmainutils \
    libssl-dev \
    wget \
    zip unzip \
    tar \
    zsh \
    libicu-dev \
    sudo \
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

## nodejs
# Install nvm
ENV NVM_NODEJS_ORG_MIRROR="https://mirrors.ustc.edu.cn/node"
RUN curl https://gitee.com/mirrors/nvm/raw/master/install.sh | bash

# Install uv
RUN curl -LsSf https://astral.sh/uv/install.sh | sh
RUN mkdir -p $HOME/.config/uv \
    && echo '[[index]]' > $HOME/.config/uv/uv.toml \
    && echo 'url = "https://mirrors.ustc.edu.cn/pypi/simple"' >> $HOME/.config/uv/uv.toml \
    && echo 'default = true' >> $HOME/.config/uv/uv.toml

# Install rust
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
RUN mkdir -p $HOME/.cargo && \
    cat > $HOME/.cargo/config.toml << EOF
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

# Install golang
ENV GOPROXY="https://goproxy.cn,direct"
ENV GVM_GO_BINARY_URL="https://mirrors.aliyun.com/golang/"
ENV GVM_GO_SOURCE_URL="https://gitee.com/mirrors/go"
RUN curl -s -S -L "https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer" | bash

ENV PATH="$PATH:$HOME/.dotnet:$HOME/.dotnet/tools"
RUN curl -s https://builds.dotnet.microsoft.com/dotnet/scripts/v1/dotnet-install.sh | bash  -s -- --channel STS

RUN curl -s "https://get.sdkman.io" | bash


