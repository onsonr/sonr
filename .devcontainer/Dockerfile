FROM jetpackio/devbox:latest

# Installing your devbox project
WORKDIR /code
USER root:root
RUN mkdir -p /code && chown ${DEVBOX_USER}:${DEVBOX_USER} /code
USER ${DEVBOX_USER}:${DEVBOX_USER}

# Copy devbox configuration files
COPY --chown=${DEVBOX_USER}:${DEVBOX_USER} devbox.json devbox.json
COPY --chown=${DEVBOX_USER}:${DEVBOX_USER} devbox.lock devbox.lock

# Install additional packages
RUN devbox add \
    neovim \
    yazi \
    rust \
    lazygit \
    github-cli \
    rclone \
    bat \
    eza \
    python3 \
    zsh

# Install aider-chat using pip
RUN devbox run -- pip install aider-chat

# Install mods using cargo
RUN devbox run -- cargo install mods

RUN devbox run -- echo "Installed Packages."

RUN devbox shellenv --init-hook >> ~/.profile

# Set shell to zsh
ENV SHELL=/bin/zsh
