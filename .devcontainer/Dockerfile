FROM ubuntu

RUN apt-get update
RUN apt-get install -y apt wget gcc nano jq snapd curl sudo gnupg file build-essential procps apt-transport-https ca-certificates

RUN curl -sLf --retry 3 --tlsv1.2 --proto "=https" 'https://packages.doppler.com/public/cli/gpg.DE2A7741A397C129.key'  | sudo gpg --dearmor -o /usr/share/keyrings/doppler-archive-keyring.gpg
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
RUN curl -fsSL https://KByUT-ai1tBBANiAv48LZI3CxhN4imHcs@apt.fury.io/sonr/gpg.key | sudo dd of=/usr/share/keyrings/fury-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/fury-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/doppler-archive-keyring.gpg] https://packages.doppler.com/public/cli/deb/debian any-version main" | sudo tee /etc/apt/sources.list.d/doppler-cli.list
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
RUN echo "deb [trusted=yes] https://KByUT-ai1tBBANiAv48LZI3CxhN4imHcs@apt.fury.io/sonr/ /" > /etc/apt/sources.list.d/fury.list

# Install dev tools
RUN sudo apt update
RUN sudo apt install -y doppler gh sonrd zsh nodejs npm git

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
