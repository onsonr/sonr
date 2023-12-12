#!bin/bash

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                 1. Installation                                ||
#// ! ||--------------------------------------------------------------------------------||

install_gum() {
	apt install git -y
    if ! dpkg -s gum >/dev/null 2>&1; then
        # Debian/Ubuntu
        sudo mkdir -p /etc/apt/keyrings
        curl -fsSL https://repo.charm.sh/apt/gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/charm.gpg
        echo "deb [signed-by=/etc/apt/keyrings/charm.gpg] https://repo.charm.sh/apt/ * *" | sudo tee /etc/apt/sources.list.d/charm.list
        sudo apt update && sudo apt install gum
    fi
}

install_docker() {
    if ! dpkg -s docker >/dev/null 2>&1; then
        curl -fsSL get.docker.com -o get-docker.sh && sudo sh get-docker.sh
	fi
}

install_earthly(){
    if ! dpkg -s earthly >/dev/null 2>&1; then
        sudo /bin/sh -c 'wget https://github.com/earthly/earthly/releases/latest/download/earthly-linux-amd64 -O /usr/local/bin/earthly && chmod +x /usr/local/bin/earthly && /usr/local/bin/earthly bootstrap --with-autocomplete'
    fi
}

download_deps(){
    install_gum
    gum spin --spinner line --title "Installing docker..." -- install_docker
    gum spin --spinner line --title "Installing earthly..." -- install_earthly
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                              2. Setup Environment                              ||
#// ! ||--------------------------------------------------------------------------------||

clone_repo(){
	git clone git@github.com:sonrhq/sonr.git
}


OPTIONS=("YES" "NO")
PS3='This script will install the necessary dependencies for running a Sonr Node. Would you like to continue?'

select CHOICE in "${OPTIONS[@]}"
do
  echo ''
  case $CHOICE in
    "YES")
        download_deps

        ;;
    "NO")
        echo "Exiting..."
        exit 1
        ;;
    *) echo "invalid option $REPLY";;
  esac
done
