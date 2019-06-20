#!/bin/bash

BASE_URL="https://ditcraft.io/dl/"
RELEASE=""
FILE=""
INSTALL_SUCCESSFUL=true
IS_UPDATE=false

determine_os() {
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

install() {
	DOWNLOAD_URL="$BASE_URL$RELEASE/$FILE"

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

######################################################
# Check whether curl is installed
which curl &> /dev/null 
if [[ $? -ne 0 ]] ; then
	echo 'This script needs "curl" to work - please install it and retry!'
	exit 1
fi

# Check command line inputs for release
if [ "$1" != "" -a  "$1" != "" ]; then
	if [ "$1" = "-r" ] || [ "$1" = "--release" ]; then
		if [ "$2" = "prerelease" ]; then
			RELEASE="prerelease"
		elif [ "$2" = "stable" ]; then
			RELEASE="stable"
		fi
	fi
elif [ "$1" != "" ]; then
	help
	exit 1
else
	RELEASE="stable"
fi

# Print help and exit if no valid choice was done
if [ "$RELEASE" = "" ]; then
	help
	exit 1
fi

determine_os
check_update
install
cleanup

# Print outcome
if [ $INSTALL_SUCCESSFUL = true ] ; then
	if [ $IS_UPDATE = true ] ; then
		echo "ditCLI successfully updated - you are ready to go!"
	else
		echo "ditCLI successfully installed - you are ready to go!"
		echo "If it's your first time using the ditCLI, feel free to take a look at 'dit help' to get started."
	fi
else
	if [ $IS_UPDATE = true ] ; then
		echo "*ERROR* ditCLI update failed. Please refer to the error above."
	else
		echo "*ERROR* ditCLI installation failed. Please refer to the error above."
	fi
fi
######################################################