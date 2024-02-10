FROM ubuntu:jammy

# Setup Environment
ENV GO_VERSION="1.21.7"
ENV ARCH=amd64
ENV OS=linux
ENV BINDIR=/usr/local/bin

# Install Go
RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.${OS}-${ARCH}.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.${OS}-${ARCH}.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.${OS}-${ARCH}.tar.gz"

# Configure Paths
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Install basic tools
RUN apt-get update
RUN apt-get install apt-transport-https ca-certificates wget gcc nano jq snapd zsh nodejs npm git gnupg file sudo build-essential procps -y

# Install NeonDB
RUN curl -sL https://github.com/neondatabase/neonctl/releases/latest/download/neonctl-linux -o neonctl
RUN chmod +x neonctl
RUN mv neonctl /usr/local/bin/neonctl

# Install IPFS
RUN wget "https://dist.ipfs.tech/kubo/v0.26.0/kubo_v0.26.0_${OS}-${ARCH}.tar.gz"
RUN tar -xvzf kubo_v0.26.0_${OS}-${ARCH}.tar.gz
RUN mv ./kubo/ipfs /usr/local/bin/ipfs

# Install Sonr
RUN curl -L https://sonr.sh | sh

# Add user
RUN chsh -s /bin/zsh
RUN useradd -m -s /bin/bash -G sudo coder
USER coder
