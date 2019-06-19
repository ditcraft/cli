#! /bin/bash

if [ "$1" == "all"  ]
    then
        echo "Building all"
        ~/go/bin/xgo --targets=linux/amd64 $GOPATH/src/github.com/ditcraft/cli/
        mv cli-linux-amd64 binaries/ditCLI-linux-amd64
        ~/go/bin/xgo --targets=windows-10/amd64 $GOPATH/src/github.com/ditcraft/cli/
        mv cli-windows-10-amd64.exe binaries/ditCLI-windows-amd64.exe
    else
        echo "Building only macos"
fi
go build -o binaries/ditCLI-macos-amd64 $GOPATH/src/github.com/ditcraft/cli/
cp -rf binaries/ditCLI-macos-amd64 /usr/local/bin/dit
chmod a+x /usr/local/bin/dit