FROM ubuntu

RUN apt-get update
RUN apt-get install -y apt wget gcc nano jq snapd curl sudo gnupg file build-essential procps apt-transport-https ca-certificates

RUN curl -fsSL https://KByUT-ai1tBBANiAv48LZI3CxhN4imHcs@apt.fury.io/sonr/gpg.key | sudo dd of=/usr/share/keyrings/fury-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/fury-keyring.gpg
RUN echo "deb [trusted=yes] https://KByUT-ai1tBBANiAv48LZI3CxhN4imHcs@apt.fury.io/sonr/ /" > /etc/apt/sources.list.d/fury.list

# Install dev tools
RUN sudo apt update
RUN sudo apt install -y doppler gh sonrd zsh git kubo dendrite
RUN curl -fsSL https://bun.sh/install | bash

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

# Install NeonDB
RUN curl -sL https://github.com/neondatabase/neonctl/releases/latest/download/neonctl-linux -o neonctl
RUN chmod +x neonctl
RUN mv neonctl /usr/local/bin/neonctl

# Add user
RUN chsh -s /bin/zsh
RUN useradd -m -s /bin/bash -G sudo coder
USER coder
