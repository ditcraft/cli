#!/bin/bash

RELEASE="stable"
BASE_URL_STABLE="https://ditcraft.io/dl/stable/"
BASE_URL_PRERELEASE="https://ditcraft.io/dl/prerelease/"
FILE=""
DOWNLOAD_URL=""
INSTALL_SUCCESSFUL=true
IS_UPDATE=false

check_os() {
	if [ "$(uname)" = "Linux" ] ; then
		FILE="ditCLI-linux-amd64"
	elif [ "$(uname)" = "Darwin" ] ; then
		FILE="ditCLI-macos-amd64"
	else
		echo "Unknown operating system"
	fi
}

check_update() {
	which dit &> /dev/null 
	if [[ $? -ne 0 ]] ; then
		echo "Installing ditCLI $RELEASE version..."
	else
		IS_UPDATE=true
		echo "Updating ditCLI $RELEASE version..."
	fi
}

get_download_url() {
	if [ "$RELEASE" = "prerelease" ]; then
		DOWNLOAD_URL="$BASE_URL_PRERELEASE$FILE"
	else
		DOWNLOAD_URL="$BASE_URL_STABLE$FILE"
	fi
}

install() {
    TMPDIR=$(mktemp -d) && cd $TMPDIR
    curl -Ss -O $DOWNLOAD_URL

	if [ "$(uname)" = "Linux" ] ; then
	  sudo cp $TMPDIR/$FILE /usr/bin/dit && sudo chmod +x /usr/bin/dit
	fi

	if [ "$(uname)" = "Darwin" ] ; then
	  sudo cp $TMPDIR/$FILE /usr/local/bin/dit && sudo chmod +x /usr/local/bin/dit
	fi

	if [[ $? -ne 0 ]] ; then
		INSTALL_SUCCESSFUL=false
	fi
}

help() {
	echo "Usage is: -r --release [ stable / prerelease ]"
	echo "Default is: stable"
}

cleanup() {
	rm $TMPDIR/*
	rmdir $TMPDIR
}

# Check whether curl is installed
which curl &> /dev/null 
if [[ $? -ne 0 ]] ; then
	echo 'This script needs "curl" to work - please install it and retry!'
	exit 1
fi


while [ "$1" != "" ]; do
	case $1 in
	-r | --release )           shift
		RELEASE=$1
		;;
	* )  	help
		exit 1
	esac
	shift
done

check_os
check_update
get_download_url
install
cleanup

if [ $INSTALL_SUCCESSFUL = true ] ; then
	if [ $IS_UPDATE = true ] ; then
		echo "ditCLI successfully updated - you are ready to go!"
	else
		echo "ditCLI successfully installed - you are ready to go!"
		echo "If it's your first time using the ditCLI feel free to take a look at 'dit help'"
	fi
else
	if [ $IS_UPDATE = true ] ; then
		echo "*ERROR* ditCLI update failed. Please refer to the error above."
	else
		echo "*ERROR* ditCLI installation failed. Please refer to the error above."
	fi
fi