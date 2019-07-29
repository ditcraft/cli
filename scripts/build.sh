#! /bin/bash

if [ "$1" == "all"  ]
    then
        echo "Building all"
        rsync -a -e ssh --exclude='binaries/*' * $GOREMOTEXGO:~/go/src/github.com/ditcraft/cli
        ssh $GOREMOTEXGO 'cd ~/go/src/github.com/ditcraft/cli; export GOPATH=~/go; /usr/local/go/bin/go get -d ./...; ~/go/bin/xgo -go 1.12.x --targets=linux/amd64 .; mv cli-linux-amd64 binaries/ditCLI-linux-amd64'
        # ssh $GOREMOTEXGO 'cd ~/go/src/github.com/ditcraft/cli; export GOPATH=~/go; /usr/local/go/bin/go get -d ./...; ~/go/bin/xgo -go 1.12.x --targets=windows-10/* .; mv cli-windows-10-amd64.exe binaries/ditCLI-windows-amd64.exe'
        rsync -e ssh $GOREMOTEXGO:~/go/src/github.com/ditcraft/cli/binaries/* binaries/
        ssh $GOREMOTEXGO 'cd ~/go/src/github.com/ditcraft/cli; rm -rf *'
    else
        echo "Building only macos"
fi
go build -o binaries/ditCLI-macos-amd64 $GOPATH/src/github.com/ditcraft/cli/
cp -rf binaries/ditCLI-macos-amd64 /usr/local/bin/dit
chmod a+x /usr/local/bin/dit