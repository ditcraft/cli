#!/bin/bash

RELEASE="stable"
BASE_URL_STABLE="https://ditcraft.io/dl/stable/"
BASE_URL_PRERELEASE="https://ditcraft.io/dl/prerelease/"
FILE=""
DOWNLOAD_URL=""

check_os() {
	if [ "$(uname)" = "Linux" ] ; then
        FILE="ditCLI-linux-amd64"
	elif [ "$(uname)" = "Darwin" ] ; then
		FILE="ditCLI-macos-amd64"
	else
		echo "Unknown operating system"
	fi
}

get_package() {
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
}

help() {
	echo "Usage is: -r --release [ stable / prerelease ]"
    echo "Default is: stable"
}

cleanup() {
  rm $TMPDIR/*
  rmdir $TMPDIR
}

## MAIN ##

## curl installed? 
which curl &> /dev/null 
if [[ $? -ne 0 ]] ; then
    echo '"curl" binary not found, please install and retry'
    exit 1
fi
##

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

echo "Installing ditCLI $RELEASE version..."
check_os
get_package
install
cleanup