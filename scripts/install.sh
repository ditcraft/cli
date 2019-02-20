#! /bin/bash

mkdir -p $GOPATH/src/dit
mkdir -p $GOPATH/src/dit/client
mkdir -p $GOPATH/src/dit/helpers
mkdir -p $GOPATH/src/dit/git
mkdir -p $GOPATH/src/dit/config
mkdir -p $GOPATH/src/dit/ethereum
mkdir -p $GOPATH/src/dit/smartcontracts
mkdir -p $GOPATH/src/dit/smartcontracts/KNWToken
mkdir -p $GOPATH/src/dit/smartcontracts/KNWVoting
mkdir -p $GOPATH/src/dit/smartcontracts/ditCoordinator
mkdir -p $GOPATH/src/dit/smartcontracts/ditContract

cp -r modules/* $GOPATH/src/dit/
cp main.go $GOPATH/src/dit/client/client.go

go install $GOPATH/src/dit/helpers/helpers.go
go install $GOPATH/src/dit/git/git.go
go install $GOPATH/src/dit/config/config.go
go install $GOPATH/src/dit/smartcontracts/KNWToken/KNWToken.go
go install $GOPATH/src/dit/smartcontracts/KNWVoting/KNWVoting.go
go install $GOPATH/src/dit/smartcontracts/ditCoordinator/ditCoordinator.go
go install $GOPATH/src/dit/smartcontracts/ditContract/ditContract.go
go install $GOPATH/src/dit/ethereum/ethereum.go