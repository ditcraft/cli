# ditCLI
The ditCLI, or simply called **dit**, is a **d**ecentralized g**it**-client, empowering distributed and democratic governance for software projects.

## Installation
There are essentially three ways on how to install the ditCLI: through our easy one-line installer, through making use of the released pre-compiled binaries or by building the ditCLI yourself, as described [here](#build-the-ditcli-yourself).
The easiest way to use dit is to make use of the released pre-compiled binaries that are provided. Alternatively you can build the ditCLI yourself, as described [here](#build-the-ditcli-yourself).

### One-Line Installer
If you are a Linux or MacOS user, you can use our easy one-line install script to get right started. This script will also update your ditCLI if you are on an old version. Just copy the following line into your terminal and you are ready to go in just a few seconds:

```
bash <(curl https://get.ditcraft.io -Ls)
```

### Installing the binaries manually

#### Linux
* Download the latest `ditCLI-linux-amd64` from the [release page](https://github.com/ditcraft/cli/releases/latest)
* Open your terminal (you will need it anyways) and go into the folder where you downloaded the ditCLI
* Move it to /usr/local/bin by executing the following
    * `mv ditCLI-linux-* /usr/bin/dit`
    * `chmod a+x /usr/bin/dit`
* Now you can use it by calling `dit` in your terminal

#### MacOS
* Download the latest `ditCLI-macos-amd64` from the [release page](https://github.com/ditcraft/cli/releases/latest)
* Open your terminal (you will need it anyways) and go into the folder where you downloaded the ditCLI
* Move it to /usr/local/bin by executing the following
    * `mv ditCLI-macos-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

#### Windows
Please note that the windows version hasn't been thoroughly tested, although it *should* work. We don't offer support for it at the moment, but will do so in the near future.

## Build the ditCLI yourself
If you want to build the ditCLI yourself feel free to do so by following this guide. The only prerequisite that you need is go.
* Run `go get github.com/ditcraft/cli`
* Enter the directory of the ditCLI with `cd $GOPATH/src/github.com/ditcraft/cli`
* Install the necessary dependencies with `go get -d ./...`
* Build the ditCLI with `go build -o dit main.go`
* For Linux/MacOS: Move `dit` to /usr/local/bin and grant execution permissions
    * `mv dit /usr/local/bin/dit && chmod +x /usr/local/bin/dit`

## Demo-Mode Usage
To test the ditCLI you can follow this guideline, showcasing you the features:

* Run `dit setup` to start the setup of dit
    * Follow the ditCLIs' directions, you'll need to do a quick KYC (know-your-coder) via our automated Twitter bot
* Go into a repository that you want to work with, to do so you can either:
    * `dit clone <REPOSITORY_URL>` to pull the repository onto your machine and automatically initialize it as a dit repository or
    * `dit init` in an exisiting repository to initialize it as a dit repository
* Now you can use dit in this repository, to get some help with the possible commands just type `dit help`
* Create a second branch through `dit checkout -b <NEW_BRANCH_NAME>`, make some changes. Commit and push these changes in the side branch.
* Return to the master branch through `dit checkout master`
* Run `dit merge <NEW_BRANCH_NAME>`
    * As you are now in the process of changing the master branch, you will be asked to initiate a new vote - please do so!
* Five demo validators will automatically vote on your proposal, so you'll directly see some action
    * Note: For testing and developing purposes the timeframe to cast and open the vote is one minute each. In reality this would last some hours or days
* Afterwards, you need to wait until the opening phase is over to finalize the vote with `dit finalize <PROPOSAL_ID>`
    * Now you will see the outcome of the vote.
    * If the vote passed, you will be allowed to merge your changes into the master branch
    * If the vote didn't pass, you won't be allowed to do so
* Execute `dit merge <NEW_BRANCH_NAME>` once again after a successful vote, you will see that you are now allowed to do so
    * If your changes didn't pass the vote, revisit the new branch, do some more changes and re-initiate a vote
* After the merge went through, you can now push these changes to the remote repository via `dit push`
* Congratulations, you just participated in a decentralized version of git that is goverened through smart-contracts on the ethereum blockchain!
