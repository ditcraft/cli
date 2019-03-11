# dit Client
The ditCraft client, simply called **dit**, is a **d**ecentralized g**it**-client, empowering distributed and democratic governance for software projects.

## Installation
A more detailed tutorial on how to build the dit client by yourself will follow in the near future. For now the easiest way is to make use of the released pre-compiled binaries that are provided.

### MacOS
* Download the latest `dit-macos-amd64` from the [release page](https://github.com/ditcraft/client/releases/latest)
* Open your terminal (you will need it anyways) and go into the folder where you downloaded the ditCraft client
* Move it to /usr/local/bin by executing the following
    * `mv dit-macos-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Linux
* Download the latest `dit-linux-amd64` from the [release page](https://github.com/ditcraft/client/releases/latest)
* Open your terminal (you will need it for dit anyways) and go into the folder where you downloaded the dit client
* Move it to /usr/local/bin by executing the following
    * `mv dit-linux-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Windows
Please note that the windows version hasn't been thoroughly tested, although it *should* work. We don't offer support for it at the moment, but will do so in the near future.

## Demo-Mode Usage
To test the dit client you can follow this guideline, showcasing you the features:

* Run `dit setup --demo` to start the setup of dit
    * The client will automatically use predefined privatekeys that can be used for the demo
* Go into a repository that you want to work with, to do so you can either:
    * `dit clone <REPOSITORY_URL>` to pull the repository onto your machine and automatically initialize it as a dit repository or
    * `dit init` in an exisiting repository to initialize it as a dit repository
* Now you can use dit in this repository, to get some help with the possible commands just type `dit help`
* Create or change a file in the repository where you want to test the dit client
* Run `dit commit <COMMIT_MESSAGE>`
    * dit will automatically commit these changes to a new branch and show you all the information necessary
* You can now let simulated demo voters vote on this proposal with `dit demo_vote <PROPOSA_ID>`
* When it's time to open the demo voters votes to the public, call `dit demo_open <PROPOSAL_ID>`
    * The timeframe to open your vote is displayed after calling `dit commit` or `dit demo_vote`
    * Note: For testing and developing purposes the timeframe to cast and open the vote is 3 minutes each. In reality this would last some days or a week. 
* After you opened the demo voters votes, you need to wait until the opening phase is over to finalize the vote with `dit finalize <PROPOSAL_ID>`
    * Now you will see the outcome of the vote.
    * If the vote passed dit will automatically merge your commit into the master branch.
    * If the vote didn't pass dit will automatically remove your commit from the repository.
* Congratulations, you just participated in a distributed version of git that is goverened through smart-contracts on the ethereum blockchain!