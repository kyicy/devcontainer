FROM debian:latest

ENV DEBIAN_FRONTEND=noninteractive

# Install base dependencies
RUN apt-get update  \
    && apt-get install -y -q --no-install-recommends sudo zsh curl git \
    && rm -rf /var/lib/apt/lists/*

# Create admin user with sudo privileges (no password required)
RUN useradd -ms /bin/zsh admin && \
    usermod -aG sudo admin && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Copy setup scripts before switching user
RUN mkdir -p /var/scripts
COPY docker/scripts /var/scripts
RUN chown -R admin:admin /var/scripts && \
    chmod -R +x /var/scripts/*.sh

USER admin

CMD ["sleep", "infinity"]