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

#// ! ||----------------------------------------------------------------------------------||
#// ! ||                                    IPFS Script                                   ||
#// ! ||----------------------------------------------------------------------------------||

setup_config() {
	ipfs config --json Gateway '{
        "HTTPHeaders": {
            "Access-Control-Allow-Origin": [
                "*"
            ],
        },
        "RootRedirect": "",
        "Writable": false,
        "PathPrefixes": [
            "/blog",
            "/refs"
        ],
        "APICommands": [],
        "NoFetch": false,
        "NoDNSLink": false,
        "PublicGateways": {
            "dweb.link": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": true
            },
            "gateway.ipfs.io": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": false
            },
            "ipfs.io": {
                "NoDNSLink": false,
                "Paths": [
                    "/ipfs",
                    "/ipns",
                    "/api"
                ],
                "UseSubdomains": false
            }
        }
    }'
}

#// ! ||----------------------------------------------------------------------------------||
#// ! ||                                    Root Menu                                   ||
#// ! ||----------------------------------------------------------------------------------||

PS3='| [sonr/rails/ipfs.sh] → Select Option: '
rootOptions=("Install" "Setup Config" "Status" "Quit")
select opt in "${rootOptions[@]}"; do
    case $opt in
        "Install")
			pkg_install ipfs
	    # optionally call a function or run some code here
            ;;
        "Setup Config")
			setup_config
	    # optionally call a function or run some code here
            ;;
        "Status")
            service_status ipfs
	    # optionally call a function or run some code here
	    break
            ;;
	"Quit")
	    echo "User requested exit"
	    exit
	    ;;
        *) echo "invalid option $REPLY";;
    esac
done
