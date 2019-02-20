# dit Client
The dit client is a **d**istributed g**it**-client demonstrating a possible decentralized approach on how git and open-source software development might look like when we take a blockchain-based approach. Along the lines of the idea of DAOs (Decentralized Autonomous Organizations) we came up with the idea to bring demoncracy and decentralization into the world of git!

## Installation
A more detailed tutorial on how to build the dit client for yourself will follow in the near future. For now the easiest way is to make use of the released pre-compiled binaries that are provided.

### MacOS
* Download the latest `dit-macos-amd64` from the release page
* Open your terminal (you will need it for dit anyways) and go into the folder where you downloaded the dit client
* Move it to /usr/local/bin by executing the following
    * `mv dit-macos-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Linux
* Download the latest `dit-linux-amd64` from the release page
* Open your terminal (you will need it for dit anyways) and go into the folder where you downloaded the dit client
* Move it to /usr/local/bin by executing the following
    * `mv dit-linux-* /usr/local/bin/dit`
    * `chmod a+x /usr/local/bin/dit`
* Now you can use it by calling `dit` in your terminal

### Windows
Please note that the windows version hasn't been thoroughly tested, although it *should* work. I don't offer support for it at the moment.

## Usage
To use the dit client you can follow this guideline:
* Run `dit setup` to start the setup of dit
    * You can select to import an exisiting ethereum private key, otherwise dit will sample a new one for you
* Go into a repository that you want to work with, to do so you can either:
    * `dit clone <REPOSITORY_URL>` to pull the repository onto your machine and automatically initialize it as a dit repository or
    * `dit init` in an exisiting repository to initialize it as a dit repository
* Now you can use dit in this repository, to get some help with the possible commands just type `dit help`