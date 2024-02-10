FROM ubuntu:jammy

# Install basic tools
RUN apt-get update
RUN apt-get install apt-transport-https ca-certificates curl git gnupg file sudo build-essential procps -y

# Add repositories
RUN curl -sLf --retry 3 --tlsv1.2 --proto "=https" 'https://packages.doppler.com/public/cli/gpg.DE2A7741A397C129.key' | sudo gpg --dearmor -o /usr/share/keyrings/doppler-archive-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/doppler-archive-keyring.gpg] https://packages.doppler.com/public/cli/deb/debian any-version main" | sudo tee /etc/apt/sources.list.d/doppler-cli.list
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null

# Install dev tools
RUN apt-get update
RUN apt-get install -y wget gcc doppler nano snapd zsh gh nodejs npm

# Install NeonDB
RUN curl -sL https://github.com/neondatabase/neonctl/releases/latest/download/neonctl-linux -o neonctl
RUN chmod +x neonctl
RUN mv neonctl /usr/local/bin/neonctl

# Install Go
ENV GO_VERSION="1.21.7"
RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Install Node Tools
RUN npm install -g @go-task/cli
RUN npm install -g @bufbuild/buf

# Install Dep (IPFS)
RUN wget "https://dist.ipfs.tech/kubo/v0.26.0/kubo_v0.26.0_linux-amd64.tar.gz"
RUN tar -xvzf kubo_v0.26.0_linux-amd64.tar.gz
RUN mv ./kubo/ipfs /usr/local/bin/ipfs

# Add user
RUN chsh -s /bin/zsh
RUN useradd -m -s /bin/bash -G sudo coder
USER coder

# Configure ZSH
RUN mkdir -p "$HOME/.zsh"
RUN git clone https://github.com/sindresorhus/pure.git "$HOME/.zsh/pure"
RUN echo "fpath+=$HOME/.zsh/pure" >> "$HOME/.zshrc"
RUN echo "autoload -U promptinit; promptinit" >> "$HOME/.zshrc"
RUN echo "prompt pure" >> "$HOME/.zshrc"
