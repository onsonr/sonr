#!bin/bash

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Utilities                                   ||
#// ! ||--------------------------------------------------------------------------------||

# Prompts a confirmation message and yes no options for the user.
confirm() {
	echo $1
	read -r -p "${1:-Are you sure? [y/N]} " response
	case "$response" in
		[yY][eE][sS]|[yY])
			true
			;;
		*)
			false
			;;
	esac
}

# Runs Apt get and installs a package if it is not already installed. - LINUX ONLY
apt_install() {
	if [ "$(uname)" == "Linux" ]; then
		if ! dpkg -s $1 >/dev/null 2>&1; then
			sudo apt-get install $1
		fi
		echo "$1 is already installed."
	else
		echo "This function can only be run on Linux."
		exit 1
	fi
}

# Runs brew and installs a package if it is not already installed. - MAC ONLY
brew_install() {
	if [ "$(uname)" == "Darwin" ]; then
		if ! brew ls --versions $1 >/dev/null 2>&1; then
			brew install $1
		fi
		echo "$1 is already installed."
	else
		echo "This function can only be run on Mac."
		exit 1
	fi
}

# Pkg install uses apt for linux and brew for mac.
pkg_install() {
	if [ "$(uname)" == "Linux" ]; then
		apt__install $1
	elif [ "$(uname)" == "Darwin" ]; then
		brew_install $1
	else
		echo "This function can only be run on Linux or Mac."
		exit 1
	fi
}

# Sh install runs a remote shell script if it is not already installed.
sh_install() {
	# check if the package is already installed
	if ! command -v $1 &> /dev/null; then
		# if not installed, run the install script
		confirm "Installing $1 requires sudo privileges." && \
			curl -fsSL $2 -o install-$1.sh && \
			sudo sh install-$1.sh
	fi
	echo "$1 is already installed."
}

# Downloads a file from a GitHub repository's latest release.
download_release_file() {
    REPO=$1
    FILE=$2
    OUTPUT=$3
    wget https://github.com/$REPO/releases/latest/download/$FILE -O $OUTPUT
}

# Downloads the latest tagged release tarball from a GitHub repository for current OS and architecture.
download_release_tarball() {
    REPO=$1
    BINARY=$2
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m | tr '[:upper:]' '[:lower:]' | sed 's/x86_64/amd64/')
    wget https://github.com/$REPO/releases/latest/download/$BINARY-$OS-$ARCH.tar.gz
    sudo tar -xvf $BINARY-$OS-$ARCH.tar.gz -C /usr/local/bin
    rm $BINARY-$OS-$ARCH.tar.gz
}

# Downloads the latest tagged release tarball from a GitHub repository for current OS and architecture.
download_release_binary() {
    REPO=$1
    BINARY=$2
    OS=$3
    ARCH=$4
    wget https://github.com/$REPO/releases/latest/download/$BINARY-$OS-$ARCH -O /usr/local/bin/$BINARY
    chmod +x /usr/local/bin/$BINARY
}

# Installs earthly for windows, linux, and mac.
earthly_install() {
	if [ "$(uname)" == "Linux" ]; then
		download_release_binary earthly/earthly earthly linux amd64
		earthly bootstrap --with-autocomplete
	elif [ "$(uname)" == "Darwin" ]; then
		brew_install earthly
		earthly bootstrap
	elif [ "$(uname)" == "Windows" ]; then
		download_release_binary earthly/earthly earthly linux amd64
		earthly bootstrap --with-autocomplete
	else
		echo "This function can only be run on Linux, Mac, or Windows."
		exit 1
	fi
}

# Checks the status of a systemd service. - LINUX ONLY
service_status() {
    if [ "$(uname)" == "Linux" ]; then
        sudo systemctl status $1
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

# Restarts a systemd service if it is enabled. - LINUX ONLY
restart_service() {
    if [ "$(uname)" == "Linux" ]; then
        sudo systemctl daemon-reload
        if systemctl is-enabled --quiet $1; then
            sudo systemctl stop $1
            sudo systemctl restart $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

# Starts a systemd service if it is not already enabled. - LINUX ONLY
start_service() {
    if [ "$(uname)" == "Linux" ]; then
        if ! systemctl is-enabled --quiet $1; then
            sudo systemctl enable $1
            sudo systemctl start $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

# Stops a systemd service if it is enabled. - LINUX ONLY
stop_service() {
    if [ "$(uname)" == "Linux" ]; then
        if systemctl is-enabled --quiet $1; then
            sudo systemctl stop $1
        fi
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}


# // ! ||--------------------------------------------------------------------------------||
# // ! ||                              Install Dependencies                              ||
# // ! ||--------------------------------------------------------------------------------||

install() {
    download_release_tarball sonr-io/sonr sonrd
    download_release_tarball sonr-io/IceFireDB icefirekv
    download_release_tarball sonr-io/IceFireDB icefiresql
    mkdir -p /var/lib/icefiresql
    download_release_file sonr-io/IceFireDB config.sqlite.yaml /var/lib/icefiresql/config.yml
}

register_services() {
    if [ "$(uname)" == "Linux" ]; then
        register_icefirekv_service
        register_icefiresql_service
        register_sonrd_service
    else
        echo "This function can only be run on Linux."
        exit 1
    fi
}

upgrade() {
    download_release_tarball sonr-io/sonr sonrd
    download_release_tarball sonr-io/IceFireDB icefirekv
    download_release_tarball sonr-io/IceFireDB icefiresql
    restart_service icefirekv
    restart_service icefiresql
    restart_service sonrd
}

#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Startup                                     ||
#// ! ||--------------------------------------------------------------------------------||

OPTIONS=("Install Sonr and Deps" "Register Systemd Service" "Upgrade Installation" "Check Service Status" "Exit")
PS3='Please enter your choice: '

select CHOICE in "${OPTIONS[@]}"
do
  echo ''
  case $CHOICE in
    "Install Sonr and Deps")
        install
        break
        ;;
    "Register Systemd Service")
        register_services
        break
        ;;
    "Upgrade Installation")
        upgrade
        break
        ;;
    "Check Service Status")
        sudo systemctl status sonrd
        break
        ;;
    "Exit")
        echo "Exiting..."
        exit 1
        ;;
    *) echo "invalid option $REPLY";;
  esac
done
