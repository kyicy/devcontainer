FROM debian:bookworm-slim

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
    curl \
    git \
    libssl-dev \
    wget \
    zip unzip \
    tar \
    zsh \
    && rm -rf /var/lib/apt/lists/*


# oh my zsh
RUN git clone https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git && cd ohmyzsh/tools && REMOTE=https://mirrors.tuna.tsinghua.edu.cn/git/ohmyzsh.git sh install.sh    

ENV SHELL=/usr/bin/zsh


## nodejs
ENV NVM_NODEJS_ORG_MIRROR="https://mirrors.ustc.edu.cn/node"
ENV NODE_VERSION="24.7.0"

# Install nvm with node and npm
RUN curl https://gitee.com/mirrors/nvm/raw/master/install.sh | bash \
    && . /root/.nvm/nvm.sh \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use default \
    && npm config set registry https://registry.npmmirror.com

# Install uv
RUN curl -LsSf https://astral.sh/uv/install.sh | sh
RUN mkdir -p /root/.config/uv \
    && echo '[[index]]' > /root/.config/uv/uv.toml \
    && echo 'url = "https://mirrors.ustc.edu.cn/pypi/simple"' >> /root/.config/uv/uv.toml \
    && echo 'default = true' >> /root/.config/uv/uv.toml

# Install rust
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
RUN curl --proto '=https' --tlsv1.2 -sSf https://rsproxy.cn/rustup-init.sh | bash -s -- -y
RUN mkdir -p /root/.cargo && \
    cat > /root/.cargo/config.toml << EOF
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
ENV GOLANG_VERSION=1.25.1
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
ENV GOPROXY="https://goproxy.cn,direct"
RUN curl -fSL "https://golang.google.cn/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz


RUN curl -s "https://get.sdkman.io" | bash
RUN /bin/bash -c "source /root/.sdkman/bin/sdkman-init.sh; sdk version; sdk install kotlin; sdk install java 25-open"
