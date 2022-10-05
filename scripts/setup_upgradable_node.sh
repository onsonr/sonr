#!/bin/bash
echo "Setting Up Base Directory"
echo "Enter Base Folder Name"
read -r dir
echo "Enter Sonrd Binary Path (Hint: use whereis sonrd)"
read -r path
export DAEMON_HOME="$PWD"/$dir/daemon_home
export DAEMON_NAME=sonrd
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
mkdir "$dir"
# shellcheck disable=SC2164
cd "$dir"
mkdir daemon_home
mkdir daemon_home/data
mkdir daemon_home/config
# shellcheck disable=SC2164
cd daemon_home
# Installing in case its not installed
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
cosmovisor init "$path"
echo "Done"