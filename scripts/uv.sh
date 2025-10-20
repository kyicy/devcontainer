#!/usr/bin/env bash

curl -LsSf https://astral.sh/uv/install.sh | sh

mkdir -p $HOME/.config/uv \
    && echo '[[index]]' > $HOME/.config/uv/uv.toml \
    && echo 'url = "https://mirrors.ustc.edu.cn/pypi/simple"' >> $HOME/.config/uv/uv.toml \
    && echo 'default = true' >> $HOME/.config/uv/uv.toml
