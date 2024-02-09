FROM ubuntu

RUN apt-get update
RUN apt-get install vim sudo curl -y
COPY ./scripts /scripts
RUN sh /scripts/goinstall.sh


RUN useradd -m -s /bin/bash -G sudo coder
USER coder

# Let's add a fancy prompt!
RUN echo "PS1='🐳 \[\033[1;36m\]\h \[\033[1;34m\]\W\[\033[0;35m\] \[\033[1;36m\]# \[\033[0m\]'" > ~/.bashrc
