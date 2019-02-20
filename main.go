package main

import (
	"dit/config"
	"dit/ethereum"
	"dit/git"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The current dit coordinator address
var defaultDitCoodinator = "0x808d9B0E0b36DCd34b64113F4c2AabadDc15743f"

func main() {
	args := os.Args[1:]

	// If no arguments are provided print the usage information
	if len(args) == 0 {
		printUsage()
	}

	var err error

	// Loading the config
	err = config.Load()

	// Converting everything to lowercase, so that a command will not fail if something
	// is capitalized by accident
	command := strings.ToLower(args[0])

	// If there is an error during the config load and the user is not
	// in the process of setting dit up: print an error
	if err != nil && command != "setup" {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	switch command {
	case "setup":
		// Create a new config
		err = config.Create()
		if err == nil {
			// Load the new config
			config.Load()
			// Set the DitCoordinator
			err = ethereum.SetDitCoordinator(defaultDitCoodinator)
			if err == nil {
				fmt.Println("\nditCoordinator automatically set to the current deployed one at " + defaultDitCoodinator)
				fmt.Println("If you wish to change this use 'dit set_coordinator <DIT_COORDINATOR_ADDRESS>")
			}
		}
		break
	case "set_coordinator":
		checkIfExists(args, 1, "the address of the ditCoordinator contract")
		// Set the DitCoordinator to a provided address
		err = ethereum.SetDitCoordinator(args[1])
		if err == nil {
			fmt.Println("ditCoordinator successfully set")
		}
		break
	case "get_balance":
		// Retrieve the balances in ETH and KNW
		err = ethereum.GetBalances()
		break
	case "get_address":
		// Return the ETH address
		fmt.Println("Ethereum Address: " + config.DitConfig.EthereumKeys.Address)
		fmt.Println("Etherscan-URL: https://rinkeby.etherscan.io/address/" + config.DitConfig.EthereumKeys.Address)
		break
	case "clone":
		var repositoryName string
		// Clone a repository and return the name of it
		repositoryName, err = git.Clone(args[1])
		if err == nil {
			// Initialize this new repository as a ditRepository
			err = ethereum.InitDitRepository(repositoryName)
			if err == nil {
				fmt.Println("ditRepository successfully cloned and initialized")
			}
		}
		break
	case "init":
		// Initialize the current repository as a ditRepository
		err = ethereum.InitDitRepository()
		if err == nil {
			fmt.Println("ditRepository successfully initiated")
		}
		break
	case "propose_commit":
		checkIfExists(args, 1, "a commit message")
		var voteDetails string
		var proposalID int
		// Propose a new commit
		voteDetails, proposalID, err = ethereum.ProposeCommit(args[1])
		if err == nil {
			// Push the commit into a proposal branch
			err = git.Commit(proposalID, args[1])
			if err == nil {
				// Print the details of the vote
				fmt.Println(voteDetails)
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
	case "reveal_vote", "reveal":
		checkIfExists(args, 1, "a proposal ID")
		// Reveals the vote on a proposal
		err = ethereum.Reveal(args[1])
		break
	case "resolve_vote", "resolve":
		checkIfExists(args, 1, "a proposal ID")
		var pollPassed bool
		// Resolves a proposal
		pollPassed, err = ethereum.Resolve(args[1])
		if err == nil {
			if pollPassed {
				// Merges the proposal branch into master after a successful vote
				err = git.Merge(args[1])
				if err == nil {
					fmt.Println("Successfully merged dit proposal " + args[1] + " into the master branch")
				}
			} else {
				// Deletes the proposal branch after an unsuccessful vote
				err = git.DeleteBranch(args[1])
				if err == nil {
					fmt.Println("Removed the dit proposal " + args[1] + " from the repository")
				}
			}
		}

		break
	default:
		printUsage()
		break
	}
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}

// Will check whether a provided argument exists and tell the user if its missing
func checkIfExists(_arguments []string, _index int, _description string) {
	if len(_arguments) < _index+1 {
		fmt.Println("Error: Please provide " + _description)
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("\n-------- dit client v0.1 -------")
	fmt.Println("------------- General ------------")
	fmt.Println(" - dit setup\t\t\t\t\tCreates or imports the ethereum keys and creates a config")
	fmt.Println("\t\t\t\t\t\tWill also set the ditCoordinator address to the currently deployed one")
	fmt.Println(" - dit set_coordinator <ADDRESS>\t\tSaves the ditCoordinator address and retrieves the KNWToken and KNWVoting addresses")
	fmt.Println("\n------------- Ethereum ------------")
	fmt.Println(" - dit get_address\t\t\t\tReturns ethereum address of the account")
	fmt.Println(" - dit get_balance\t\t\t\tReturns the ETH and KNW balance of the account")
	fmt.Println("\n----------- Repositories ----------")
	fmt.Println(" - dit clone <REPOSITORY_URL>\t\t\tClones a repository (GitHub only) and automatically calls 'dit init' afterwards")
	fmt.Println(" - dit init\t\t\t\t\tRetrieves the address of the ditContract for the repository you are using (GitHub only)")
	fmt.Println(" - dit propose_commit <COMMIT_MESSAGE>\t\tProposes a new commit with the specified commit message (This will start a vote)")
	fmt.Println(" - dit vote <PROPOSAL_ID> <CHOICE> <SALT>\tVotes on a proposal")
	fmt.Println(" - dit reveal_vote <PROPOSAL_ID>\t\tReveals a vote on a proposal")
	fmt.Println(" - dit resolve_vote <PROPOSAL_ID>\t\tResolves a vote on a proposal")
	fmt.Println("")
	os.Exit(0)
}
