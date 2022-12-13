#!/bin/bash

pathadd() {
    newelement=${1%/}
    if [ -d "$1" ] && ! echo $PATH | grep -E -q "(^|:)$newelement($|:)" ; then
        if [ "$2" = "after" ] ; then
            PATH="$PATH:$newelement"
        else
            PATH="$newelement:$PATH"
        fi
    fi
}

pathrm() {
    PATH="$(echo $PATH | sed -e "s;\(^\|:\)${1%/}\(:\|\$\);\1\2;g" -e 's;^:\|:$;;g' -e 's;::;:;g')"
}

if ! command -v ufw &> /dev/null
then
    echo "Installing ufw..."
    apt install ufw
fi

if ! command -v go &> /dev/null
then
    echo "Installing go1.18.9..."
    curl -OL https://golang.org/dl/go1.18.9.linux-amd64.tar.gz
    sudo tar -C /usr/local -xvf go1.18.9.linux-amd64.tar.gz
    rm -rf go1.18.9.linux-amd64.tar.gz
    pathadd "/usr/local/go/bin"
    pathadd "/root/go/bin" after
    export PATH
fi

if ! command -v ignite &> /dev/null
then
    echo "Installing Ignite..."
    curl https://get.ignite.com/cli! | bash
fi
