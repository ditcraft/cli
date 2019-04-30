# dit Client
The ditCraft client, simply called **dit**, is a **d**ecentralized g**it**-client, empowering distributed and democratic governance for software projects.

## Installation
The easiest way to use dit is to make use of the released pre-compiled binaries that are provided. Alternatively you can build the client yourself, as described [here](#build-the-client-yourself).

### MacOS
* Download the latest `dit-macos-amd64` from the [release page](https://github.com/ditcraft/client/releases/latest)
* Open your terminal (you will need it anyways) and go into the folder where you downloaded the dit client
* Move it to /usr/local/bin by executing the following
    * `mv dit-macos-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Linux
* Download the latest `dit-linux-amd64` from the [release page](https://github.com/ditcraft/client/releases/latest)
* Open your terminal (you will need it anyways) and go into the folder where you downloaded the dit client
* Move it to /usr/local/bin by executing the following
    * `mv dit-linux-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Windows
Please note that the windows version hasn't been thoroughly tested, although it *should* work. We don't offer support for it at the moment, but will do so in the near future.

## Build the client yourself
If you want to build the client yourself feel free to do so by following this guide. The only prerequisite that you need is go.
* Run `go get github.com/ditcraft/client`
* Enter the directory of the client with `cd $GOPATH/src/github.com/ditcraft/client`
* Install the necessary dependencies with `go get -d ./...`
* Build the client with `go build -o dit main.go`
* For Linux/MacOS: Move `dit` to /usr/local/bin and grant execution permissions
    * `mv dit /usr/local/bin/dit && chmod a+x /ust/local/bin/dit`

**Please note**: If you intend to make use of the demo mode described below, you need to insert four Ethereum addresses and their corresponding private keys with funds on the Rinkeby Ethereum testnet into the code before building it. One (*the main one*) in `config.go` (*line 64/65*) and three (*the demo voters*) in `demo.go` (*lines 29-38*). The precompiled binaries already contain pre-funded ones, so only need to worry about this when building it yourself.

## Demo-Mode Usage
To test the dit client you can follow this guideline, showcasing you the features:

* Run `dit setup` to start the setup of dit
    * Follow the clients' directions, you'll need to do a quicky KYC (know-your-coder) via our automated Twitter bot
* Go into a repository that you want to work with, to do so you can either:
    * `dit clone <REPOSITORY_URL>` to pull the repository onto your machine and automatically initialize it as a dit repository or
    * `dit init` in an exisiting repository to initialize it as a dit repository
* Now you can use dit in this repository, to get some help with the possible commands just type `dit help`
* Create or change a file in the repository where you want to test the dit client
* Run `dit commit <COMMIT_MESSAGE>`
    * dit will automatically commit these changes to a new branch and show you all the information necessary
* Five demo validators will automatically vote on your proposal, so you'll directly see some action
    * Note: For testing and developing purposes the timeframe to cast and open the vote is 3 minutes each. In reality this would last some days or a week.
* Afterwards, ou need to wait until the opening phase is over to finalize the vote with `dit finalize <PROPOSAL_ID>`
    * Now you will see the outcome of the vote.
    * If the vote passed dit will automatically merge your commit into the master branch.
    * If the vote didn't pass dit will automatically remove your commit from the repository.
* Congratulations, you just participated in a decentralized version of git that is goverened through smart-contracts on the ethereum blockchain!