package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ditcraft/client/config"
	"github.com/ditcraft/client/demo"
	"github.com/ditcraft/client/ethereum"
	"github.com/ditcraft/client/git"
	"github.com/ditcraft/client/helpers"
)

// The current dit coordinator address
var defaultDitCoodinator = "0x60F01B8F86Aa3D1a61d1E1730B49BaeE09D8d72c"

func main() {
	var err error

	args := os.Args[1:]

	// Loading the config
	err = config.Load()

	// If no arguments are provided print the usage information
	if len(args) == 0 {
		printUsage()
	}

	// Converting everything to lowercase, so that a command will not fail if something
	// is capitalized by accident
	command := strings.ToLower(args[0])

	// If there is an error during the config load and the user is not
	// in the process of setting dit up: print an error
	if err != nil && command != "setup" {
		helpers.PrintLine(err.Error(), 2)
		os.Exit(0)
	}

	if err == nil && config.DitConfig.DitCoordinator != defaultDitCoodinator && command != "setup" {
		helpers.PrintLine("You are using an old version of the deployed ditCoordinator contract", 1)
		helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("setup")+"'", 1)
	}

	switch command {
	case "setup":
		// Create a new config
		if len(args) == 2 && strings.Contains(args[1], "demo") {
			err = config.Create(true)
		} else {
			err = config.Create(false)
		}

		if err == nil {
			// Load the new config
			config.Load()
			// Set the DitCoordinator
			err = ethereum.SetDitCoordinator(defaultDitCoodinator)
			if err == nil {
				fmt.Println()
				helpers.PrintLine("ditCoordinator automatically set to the current deployed one at "+defaultDitCoodinator, 0)
				helpers.PrintLine("If you wish to change this use '"+helpers.ColorizeCommand("set_coordinator <DIT_COORDINATOR_ADDRESS>")+"'", 0)
			}
		}
		break
	case "set_coordinator":
		checkIfExists(args, 1, "the address of the ditCoordinator contract")
		// Set the DitCoordinator to a provided address
		err = ethereum.SetDitCoordinator(args[1])
		if err == nil {
			helpers.PrintLine("ditCoordinator successfully set", 0)
		}
		break
	case "get_balance":
		// Retrieve the balances in ETH and KNW
		err = ethereum.GetBalances()
		break
	case "get_address":
		// Return the ETH address
		helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, 0)
		helpers.PrintLine("Etherscan-URL: https://rinkeby.etherscan.io/address/"+config.DitConfig.EthereumKeys.Address, 0)
		break
	case "clone":
		checkIfExists(args, 1, "a URL to a repository")
		var repositoryName string
		// Clone a repository and return the name of it
		repositoryName, err = git.Clone(args[1])
		if err == nil {
			// Initialize this new repository as a ditRepository
			err = ethereum.InitDitRepository(repositoryName)
			if err == nil {
				helpers.PrintLine("ditRepository successfully cloned and initialized", 0)
			}
		}
		break
	case "init":
		// Initialize the current repository as a ditRepository
		err = ethereum.InitDitRepository()
		if err == nil {
			helpers.PrintLine("ditRepository successfully initiated", 0)
		}
		break
	case "commit", "propose_commit", "demo_commit":
		checkIfExists(args, 1, "a commit message")
		var voteDetails string
		var proposalID int

		err = git.CheckForChanges()
		if err == nil {
			// Propose a new commit
			voteDetails, proposalID, err = ethereum.ProposeCommit(args[1])
			if err == nil {
				// Push the commit into a proposal branch
				err = git.Commit(proposalID, args[1])
				if err == nil {
					// Print the details of the vote
					stringLines := strings.Split(voteDetails, "\n")
					for i := range stringLines {
						helpers.PrintLine(stringLines[i], 0)
					}
				}
			}
		}
		break
	case "vote_info":
		// Prints the details of a vote
		if len(args) == 2 {
			voteID, _ := strconv.Atoi(args[1])
			err = ethereum.GetVoteInfo(voteID)
		} else {
			err = ethereum.GetVoteInfo()
		}
		break
	case "vote":
		checkIfExists(args, 1, "a proposal ID - dit vote <PROPOSAL_ID> <CHOICE> <SALT>")
		checkIfExists(args, 2, "your choice (1/0) - dit vote <PROPOSAL_ID> <CHOICE> <SALT>")
		checkIfExists(args, 3, "a salt (a non-zero number) - dit vote <PROPOSAL_ID> <CHOICE> <SALT>")
		// Votes on a proposal
		err = ethereum.Vote(args[1], args[2], args[3])
		break
	case "demo_vote":
		checkIfExists(args, 1, "a proposal ID - dit demo_vote <PROPOSAL_ID>")
		// Demo-Votes on a proposal
		err = demo.Vote(args[1])
		break
	case "open", "open_vote", "reveal_vote", "reveal":
		checkIfExists(args, 1, "a proposal ID")
		// Reveals the vote on a proposal
		err = ethereum.Open(args[1])
		break
	case "demo_open", "demo_open_vote", "demo_reveal_vote", "demo_reveal":
		checkIfExists(args, 1, "a proposal ID")
		// Reveals the votes of the demo voters on a proposal
		for i := 0; i < 3; i++ {
			err = demo.Open(args[1], i)
		}
		break
	case "finalize", "finalize_vote":
		checkIfExists(args, 1, "a proposal ID")
		var pollPassed bool
		// Resolves a proposal
		pollPassed, err = ethereum.Finalize(args[1])
		if err == nil {
			fmt.Println()
			if pollPassed {
				// Merges the proposal branch into master after a successful vote
				err = git.Merge(args[1])
				if err == nil {
					helpers.PrintLine("Successfully merged dit proposal "+args[1]+" into the master branch", 0)
				}
			} else {
				// Deletes the proposal branch after an unsuccessful vote
				err = git.DeleteBranch(args[1])
				if err == nil {
					helpers.PrintLine("Removed the dit proposal "+args[1]+" from the repository", 0)
				}
			}
		}

		break
	default:
		printUsage()
		break
	}
	if err != nil {
		if len(err.Error()) > 0 {
			helpers.PrintLine(err.Error(), 2)
		}
	}
}

// Will check whether a provided argument exists and tell the user if its missing
func checkIfExists(_arguments []string, _index int, _description string) {
	if len(_arguments) < _index+1 {
		helpers.PrintLine("Please provide "+_description, 2)
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("--------- dit client v0.1 --------")
	if config.DitConfig.DemoModeActive {
		fmt.Println("--------- demo mode active -------")
	}
	fmt.Println("------------- General ------------")
	fmt.Println(" - dit setup\t\t\t\t\tCreates or imports the ethereum keys and creates a config")
	fmt.Println(" - dit set_coordinator <ADDRESS>\t\tSaves the ditCoordinator address and retrieves the")
	fmt.Println("\t\t\t\t\t\tKNWToken and KNWVoting addresses")
	fmt.Println("")
	fmt.Println("------------- Ethereum ------------")
	fmt.Println(" - dit get_address\t\t\t\tReturns ethereum address of the account")
	fmt.Println(" - dit get_balance\t\t\t\tReturns the ETH and KNW balance of the account")
	fmt.Println("")
	fmt.Println("----------- Repositories ----------")
	fmt.Println(" - dit clone <REPOSITORY_URL>\t\t\tClones a repository (GitHub only) and automatically")
	fmt.Println("\t\t\t\t\t\tcalls 'dit init' afterwards")
	fmt.Println(" - dit init\t\t\t\t\tRetrieves the address of the ditContract for the repository")
	fmt.Println("\t\t\t\t\t\tyou are using (GitHub only)")
	fmt.Println(" - dit commit <COMMIT_MESSAGE>—\t\t\tProposes a new commit with the specified")
	fmt.Println("\t\t\t\t\t\tcommit message (This will start a vote)")
	fmt.Println("")
	fmt.Println("------------- Voting --------------")
	fmt.Println(" - dit vote <PROPOSAL_ID> <CHOICE> <SALT>\tCasts a concealed vote on a proposed commit")
	fmt.Println(" - dit open <PROPOSAL_ID>\t\t\tOpens and reveals the vote commitment on a proposed commit")
	fmt.Println(" - dit finalize <PROPOSAL_ID>\t\t\tFinalizes the vote on proposed commit and claims the tokens")
	os.Exit(0)
}
