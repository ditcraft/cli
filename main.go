package main

import (
	"errors"
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
var liveDitCoodinator = "0x73F2cdF96941B0f51282Fd21bab108df829C9c71"
var demoDitCoodinator = "0x0488Cea56d44E3BBEAD34D6Ffd2CBcD561d08b90"
var configVersion = "v0.2"

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
	if err != nil && command != "setup" && command != "update" {
		helpers.PrintLine(err.Error(), 2)
		os.Exit(0)
	}

	if err == nil && (command != "setup" && command != "update") {
		if config.DitConfig.Version != config.Version {
			helpers.PrintLine("Your config is not up to date, the client might not work correctly", 1)
			helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("update")+"'", 1)
		} else if (!config.DitConfig.DemoModeActive && config.DitConfig.DitCoordinator != liveDitCoodinator) || (config.DitConfig.DemoModeActive && config.DitConfig.DitCoordinator != demoDitCoodinator) {
			helpers.PrintLine("You are using an old version of the deployed ditCoordinator contract", 1)
			helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("update")+"'", 1)
		}
	}

	switch command {
	case "setup":
		var ditCoordinatorAddress string
		// Create a new config
		if len(args) == 2 && strings.Contains(args[1], "live") {
			err = config.Create(false)
			ditCoordinatorAddress = liveDitCoodinator
		} else {
			err = config.Create(true)
			ditCoordinatorAddress = demoDitCoodinator
		}

		if err == nil {
			// Load the new config
			config.Load()
			// Set the DitCoordinator
			err = ethereum.SetDitCoordinator(ditCoordinatorAddress)
			if err == nil {
				fmt.Println()
				helpers.PrintLine("ditCoordinator automatically set to the current deployed one", 0)
				passedKYC, err := ethereum.CheckForKYC()
				if err == nil && !passedKYC {
					fmt.Println()
					helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", 0)
					helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", 0)
					helpers.PrintLine("@ditcraft I want to use dit, the decentralized git client! Please verify me "+config.DitConfig.EthereumKeys.Address+"", 0)
				}
			}
		}
		break
	case "update":
		var updatedDitCoordinator bool
		updatedDitCoordinator, err = config.Update(liveDitCoodinator, demoDitCoodinator)
		if updatedDitCoordinator {
			err = ethereum.SetDitCoordinator(config.DitConfig.DitCoordinator)
			if !config.DitConfig.PassedKYC {
				var passedKYC bool
				passedKYC, err = ethereum.CheckForKYC()
				if err == nil && !passedKYC {
					fmt.Println()
					helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", 0)
					helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", 0)
					helpers.PrintLine("@ditcraft I want to try dit, the decentralized git. Please verify me "+config.DitConfig.EthereumKeys.Address+"!", 0)
				}
			}
		}
		if err == nil {
			helpers.PrintLine("Update successful", 0)
		}
		break
	case "mode":
		checkIfExists(args, 1, "the mode that you want to switch to (live/demo)")
		var ditCoordinatorAddress string
		if args[1] == "live" {
			ditCoordinatorAddress = liveDitCoodinator
			config.DitConfig.DemoModeActive = false
		} else if args[1] == "demo" {
			ditCoordinatorAddress = demoDitCoodinator
			config.DitConfig.DemoModeActive = true
		} else {
			printUsage()
			os.Exit(0)
		}
		err = ethereum.SetDitCoordinator(ditCoordinatorAddress)
		if err == nil {
			helpers.PrintLine("dit client switched to the "+args[1]+" mode", 0)
			if args[1] == "live" {
				helpers.PrintLine("You are now using the client in live mode, you will be staking real xDai!", 1)
			}
			passedKYC, err := ethereum.CheckForKYC()
			if err == nil && !passedKYC {
				fmt.Println()
				helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", 0)
				helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", 0)
				helpers.PrintLine("@ditcraft I want to try dit, the decentralized git. Please verify me "+config.DitConfig.EthereumKeys.Address+"!", 0)
			}
		}
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
		helpers.PrintLine("URL: https://blockscout.com/poa/dai/address/"+config.DitConfig.EthereumKeys.Address, 0)
		break
	case "export_keys":
		helpers.PrintLine("Exporting your keys will print your ethereum private key in plain text", 1)
		helpers.PrintLine("This might compromise your keys, please don't proceed if you are not sure about this", 1)
		answer := helpers.GetUserInputChoice("Are you sure that you want to proceed?", "y", "n")
		if answer == "y" {
			// Return the ETH address and private key
			var privateKey string
			privateKey, err = config.GetPrivateKey(false)
			if len(privateKey) == 64 {
				helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, 0)
				helpers.PrintLine("Ethereum Private-Key: "+privateKey, 0)
			}
		} else {
			err = errors.New("Cancelling key export due to users choice")
		}
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
		err = git.Validate()
		if err == nil {
			var voteDetails string
			var proposalID int

			err = git.CheckForChanges()
			if err == nil {
				// Propose a new commit
				if config.DitConfig.DemoModeActive {
					voteDetails, proposalID, err = demo.ProposeCommit(args[1])
				} else {
					voteDetails, proposalID, err = ethereum.ProposeCommit(args[1])
				}
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
	case "open", "open_vote", "reveal_vote", "reveal":
		checkIfExists(args, 1, "a proposal ID")
		// Reveals the vote on a proposal
		err = ethereum.Open(args[1])
		break
	case "finalize", "finalize_vote":
		checkIfExists(args, 1, "a proposal ID")
		err = git.Validate()
		if err == nil {
			var pollPassed bool
			var isProposer bool
			// Finalizes a proposal
			pollPassed, isProposer, err = ethereum.Finalize(args[1])
			if err == nil && isProposer {
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
	fmt.Println("--------- dit client v0.2 --------")
	if config.DitConfig.DemoModeActive {
		fmt.Println("--------- demo mode active -------")
	} else {
		fmt.Println("--------- live mode active -------")
	}
	fmt.Println("------------- General ------------")
	fmt.Println(" - dit setup\t\t\t\t\tCreates or imports the ethereum keys and creates a config")
	fmt.Println(" - dit update\t\t\t\t\tUpdates the clients config after an update of the dit client")
	fmt.Println(" - dit mode <MODE>\t\t\t\tSwitch between the modes of the client (live or demo)")
	fmt.Println("")
	fmt.Println("------------- Ethereum ------------")
	fmt.Println(" - dit get_address\t\t\t\tReturns ethereum address of the account")
	fmt.Println(" - dit get_balance\t\t\t\tReturns the " + config.DitConfig.Currency + " and KNW balance of the account")
	fmt.Println(" - dit export_keys\t\t\t\tReturns the unencrypted ethereum private key and address")
	fmt.Println("")
	fmt.Println("----------- Repositories ----------")
	fmt.Println(" - dit clone <REPOSITORY_URL>\t\t\tClones a repository (GitHub only) and automatically")
	fmt.Println("\t\t\t\t\t\tcalls 'dit init' afterwards")
	fmt.Println(" - dit init\t\t\t\t\tRetrieves the address of the ditContract for the repository")
	fmt.Println("\t\t\t\t\t\tyou are using (GitHub only)")
	fmt.Println(" - dit commit <COMMIT_MESSAGE>\t\t\tProposes a new commit with the specified")
	fmt.Println("\t\t\t\t\t\tcommit message (This will start a vote)")
	fmt.Println("")
	fmt.Println("------------- Voting --------------")
	fmt.Println(" - dit vote <PROPOSAL_ID> <CHOICE> <SALT>\tCasts a concealed vote on a proposed commit")
	fmt.Println(" - dit open <PROPOSAL_ID>\t\t\tOpens and reveals the vote commitment on a proposed commit")
	fmt.Println(" - dit finalize <PROPOSAL_ID>\t\t\tFinalizes the vote on proposed commit and claims the tokens")
	os.Exit(0)
}
