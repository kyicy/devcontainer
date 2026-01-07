#!/usr/bin/env bash

set -e

sudo apt-get update

sudo apt-get install -y -q --no-install-recommends \
    build-essential \
    curl mercurial make binutils bison gcc bsdmainutils \
    libssl-dev \
    wget \
    zip unzip \
    tar \
    libicu-dev
