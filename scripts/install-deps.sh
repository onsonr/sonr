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

# Installs earthly for windows, linux, and mac.
earthly_install() {
	if [ "$(uname)" == "Linux" ]; then
		download_release_binary earthly/earthly earthly linux amd64
		earthly bootstrap --with-autocomplete
	elif [ "$(uname)" == "Darwin" ]; then
		brew_install earthly
		earthly bootstrap --with-autocomplete
	elif [ "$(uname)" == "Windows" ]; then
		download_release_binary earthly/earthly earthly linux amd64
		earthly bootstrap --with-autocomplete
	else
		echo "This function can only be run on Linux, Mac, or Windows."
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

#// ! ||---------------------------------------------------------------------------||
#// ! ||                                    Main                                   ||
#// ! ||---------------------------------------------------------------------------||

install() {
	pkg_install "gum"
	sh_install "docker" "https://get.docker.com/"
	earthly_install
}


#// ! ||--------------------------------------------------------------------------------||
#// ! ||                                    Startup                                     ||
#// ! ||--------------------------------------------------------------------------------||

OPTIONS=("YES" "NO")
PS3='This script will install the necessary dependencies for running a Sonr Node. Would you like to continue?'

select CHOICE in "${OPTIONS[@]}"
do
  echo ''
  case $CHOICE in
    "YES")
        install
        break
        ;;
    "NO")
        echo "Exiting..."
        exit 1
        ;;
    *) echo "invalid option $REPLY";;
  esac
done
