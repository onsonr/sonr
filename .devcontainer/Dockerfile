FROM sonrhq/ubuntu-go:1.22.0


# update and install dependencies
RUN apt-get update && apt-get install -y \
	kubo \
	dendrite \
	unzip

RUN add-apt-repository ppa:maveonair/helix-editor
RUN apt update
RUN apt install helix

# Install NeonDB
RUN curl -sL https://github.com/neondatabase/neonctl/releases/latest/download/neonctl-linux -o neonctl
RUN chmod +x neonctl
RUN mv neonctl /usr/local/bin/neonctl

# Add user
RUN chsh -s /bin/zsh
RUN useradd -m -s /bin/bash -G sudo coder
USER coder

