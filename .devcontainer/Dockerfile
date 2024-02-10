FROM ubuntu

# Install dev tools
RUN apt-get update
RUN apt-get install -y wget gcc nano jq snapd zsh nodejs npm curl sudo gnupg file build-essential procps

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
RUN apt-get install apt-transport-https ca-certificates zsh nodejs npm git -y

# Install Doppler/GH
RUN curl -sLf --retry 3 --tlsv1.2 --proto "=https" 'https://packages.doppler.com/public/cli/gpg.DE2A7741A397C129.key'  | sudo gpg --dearmor -o /usr/share/keyrings/doppler-archive-keyring.gpg
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
RUN curl -fsSL https://apt.sonr.build/gpg.key | sudo dd of=/usr/share/keyrings/fury-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/fury-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/doppler-archive-keyring.gpg] https://packages.doppler.com/public/cli/deb/debian any-version main" | sudo tee /etc/apt/sources.list.d/doppler-cli.list
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
RUN echo "deb https://apt.sonr.build/ * *" > /etc/apt/sources.list.d/fury.list
RUN apt-get install gh doppler sonrd

# Install NeonDB
RUN curl -sL https://github.com/neondatabase/neonctl/releases/latest/download/neonctl-linux -o neonctl
RUN chmod +x neonctl
RUN mv neonctl /usr/local/bin/neonctl

# Install IPFS
RUN wget "https://dist.ipfs.tech/kubo/v0.26.0/kubo_v0.26.0_${OS}-${ARCH}.tar.gz"
RUN tar -xvzf kubo_v0.26.0_${OS}-${ARCH}.tar.gz
RUN mv ./kubo/ipfs /usr/local/bin/ipfs

# Add user
RUN chsh -s /bin/zsh
RUN useradd -m -s /bin/bash -G sudo coder
USER coder
