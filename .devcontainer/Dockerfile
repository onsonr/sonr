# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.191.0/containers/go/.devcontainer/base.Dockerfile
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: lts/*, 16, 14, 12, 10
ARG NODE_VERSION="16"
RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} && nvm use ${NODE_VERSION} && npm install -g nextron" 2>&1
