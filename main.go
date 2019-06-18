package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ditcraft/cli/config"
	"github.com/ditcraft/cli/demo"
	"github.com/ditcraft/cli/ethereum"
	"github.com/ditcraft/cli/git"
	"github.com/ditcraft/cli/helpers"
	"github.com/logrusorgru/aurora"
)

// The current dit coordinator address
var liveDitCoodinator = "0x429e37f14462bdfca0f1168dae24f66f61e6b04c"
var demoDitCoodinator = "0x1dc6f1edd14b0b5d24305a0cfb6d4f0a5de3b4f6"
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
		helpers.PrintLine(err.Error(), helpers.ERROR)
		os.Exit(0)
	}

	if err == nil && (command != "setup" && command != "update") {
		if config.DitConfig.Version != config.Version {
			helpers.PrintLine("Your config is not up to date, the ditCLI might not work correctly", helpers.WARN)
			helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("update")+"'", helpers.WARN)
		} else if (!config.DitConfig.DemoModeActive && config.DitConfig.DitCoordinator != liveDitCoodinator) || (config.DitConfig.DemoModeActive && config.DitConfig.DitCoordinator != demoDitCoodinator) {
			helpers.PrintLine("You are using an old version of the deployed ditCoordinator contract", helpers.WARN)
			helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("update")+"'", helpers.WARN)
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
				helpers.PrintLine("ditCoordinator automatically set to the current deployed one", helpers.INFO)
				passedKYC, err := ethereum.CheckForKYC()
				if err == nil && !passedKYC {
					fmt.Println()
					helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", helpers.INFO)
					helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", helpers.INFO)
					helpers.PrintLine("@ditcraft I want to use dit, the decentralized git client! Please verify me "+config.DitConfig.EthereumKeys.Address+"", helpers.INFO)
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
					helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", helpers.INFO)
					helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", helpers.INFO)
					helpers.PrintLine("@ditcraft I want to try dit, the decentralized git. Please verify me "+config.DitConfig.EthereumKeys.Address+"!", helpers.INFO)
				}
			}
		}
		if err == nil {
			helpers.PrintLine("Update successful", helpers.INFO)
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
			helpers.PrintLine("ditCLI switched to the "+args[1]+" mode", helpers.INFO)
			if args[1] == "live" {
				helpers.PrintLine("You are now using the ditCLI in live mode, you will be staking real xDAI!", helpers.WARN)
			}
			passedKYC, err := ethereum.CheckForKYC()
			if err == nil && !passedKYC {
				fmt.Println()
				helpers.PrintLine("You didn't pass the KYC yet. Please do the KYC now:", helpers.INFO)
				helpers.PrintLine("Go to our Twitter @ditcraft and tweet the following at us:", helpers.INFO)
				helpers.PrintLine("@ditcraft I want to try dit, the decentralized git. Please verify me "+config.DitConfig.EthereumKeys.Address+"!", helpers.INFO)
			}
		}
	case "set_coordinator":
		checkIfExists(args, 1, "the address of the ditCoordinator contract")
		// Set the DitCoordinator to a provided address
		err = ethereum.SetDitCoordinator(args[1])
		if err == nil {
			helpers.PrintLine("ditCoordinator successfully set", helpers.INFO)
		}
		break
	case "get_balance":
		// Retrieve the balances in ETH and KNW
		err = ethereum.GetBalances()
		break
	case "get_address":
		// Return the ETH address
		helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, helpers.INFO)
		helpers.PrintLine("URL: https://explorer.ditcraft.io/address/"+config.DitConfig.EthereumKeys.Address, helpers.INFO)
		break
	case "export_keys":
		helpers.PrintLine("Exporting your keys will print your ethereum private key in plain text", helpers.WARN)
		helpers.PrintLine("This might compromise your keys, please don't proceed if you are not sure about this", helpers.WARN)
		answer := helpers.GetUserInputChoice("Are you sure that you want to proceed?", "y", "n")
		if answer == "y" {
			// Return the ETH address and private key
			var privateKey string
			privateKey, err = config.GetPrivateKey(false)
			if len(privateKey) == 64 {
				helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, helpers.INFO)
				helpers.PrintLine("Ethereum Private-Key: "+privateKey, helpers.INFO)
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
				helpers.PrintLine("ditRepository successfully cloned and initialized", helpers.INFO)
			}
		}
		break
	case "init":
		// Initialize the current repository as a ditRepository
		err = ethereum.InitDitRepository()
		if err == nil {
			helpers.PrintLine("ditRepository successfully initiated", helpers.INFO)
		}
		break
	case "merge", "rebase":
		var isOnMaster bool
		isOnMaster, err = git.IsOnMaster()
		if !isOnMaster {
			_ = git.ExecuteCommand(args[:]...)
		} else {
			checkIfExists(args, 1, "the branch to be "+command+"d")
			err = git.MasterIsClean()
			if err == nil {
				var branchHash string
				branchHash, err = git.GetHeadHashOfBranch(args[1])
				if err == nil {
					var repository, voteDetails string
					var voteID int
					var newVote, voteEnded, votePassed bool
					if config.DitConfig.DemoModeActive {
						voteID, voteEnded, votePassed, repository, err = demo.SearchForHashInVotes(branchHash)
						if err == nil {
							if voteID == 0 {
								helpers.PrintLine("You are trying to "+command+" into the master branch - this requires a vote.", helpers.INFO)
								voteDetails, voteID, err = demo.ProposeCommit(args[1], branchHash)
								newVote = true
							}
						}
					} else {
						voteID, voteEnded, votePassed, repository, err = demo.SearchForHashInVotes(branchHash)
						if err == nil {
							if voteID == 0 {
								helpers.PrintLine("You are trying to "+command+" into the master branch - this requires a vote.", helpers.INFO)
								voteDetails, voteID, err = ethereum.ProposeCommit(args[1], branchHash)
								newVote = true
							}
						}
					}
					if err == nil && voteID > 0 {
						if newVote {
							stringLines := strings.Split(voteDetails, "\n")
							for i := range stringLines {
								helpers.PrintLine(stringLines[i], helpers.INFO)
							}
						} else {
							if !voteEnded {
								helpers.PrintLine("The vote on this "+command+" hasn't ended yet - please wait!", helpers.ERROR)
							} else if voteEnded && !votePassed {
								helpers.PrintLine("The vote on this "+command+" hasn't passed - you are not allowed to "+command+" into master!", helpers.ERROR)
							} else if voteEnded && votePassed {
								err = git.ExecuteCommand(args[:]...)
								if err == nil {
									var newHeadHash string
									newHeadHash, err = git.GetHeadHashOfBranch("master")
									if err == nil {
										config.DitConfig.DemoRepositories[repository].ActiveVotes[strconv.Itoa(voteID)].NewHeadHash = newHeadHash
										err = config.Save()
									}
								}
							}
						}
					}
				}
			}
		}
		break
	case "push":
		var isOnMaster bool
		isOnMaster, err = git.IsOnMaster()
		if !isOnMaster {
			_ = git.ExecuteCommand(args[:]...)
		} else {
			var masterHeadHash string
			masterHeadHash, err = git.GetHeadHashOfBranch("master")
			if err == nil {
				var repository string
				repository, err = git.GetRepository()

				var repositoryMap map[string]*config.Repository
				if config.DitConfig.DemoModeActive {
					repositoryMap = config.DitConfig.DemoRepositories
				} else {
					repositoryMap = config.DitConfig.LiveRepositories
				}

				if err == nil && repositoryMap[repository] == nil {
					err = errors.New("Repository hasn't been initialized")
				}

				if err == nil {
					foundVote := false
					for _, value := range repositoryMap[repository].ActiveVotes {
						if value.NewHeadHash == masterHeadHash {
							foundVote = true
							break
						}
					}
					if foundVote {
						_ = git.ExecuteCommand(args[:]...)
					} else {
						helpers.PrintLine("You are not allowed to push changes to master without prior approval from the community", helpers.ERROR)
						// TODO: offer to start a new vote on these changes?
					}
				}
			}
		}
		break
	case "add",
		"bisect",
		"branch",
		"checkout",
		"commit",
		"diff",
		"fetch",
		"grep",
		"log",
		"pull",
		"reset",
		"rm",
		"show",
		"stash",
		"status",
		"tag":
		_ = git.ExecuteCommand(args[:]...)
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
					helpers.PrintLine("You are now allowed to integrate your changes into the master branch", helpers.INFO)
				} else {
					helpers.PrintLine(fmt.Sprintf("You are %s allowed to integrate your changes into the master branch", aurora.Bold(aurora.Red("NOT"))), helpers.INFO)
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
			helpers.PrintLine(err.Error(), helpers.ERROR)
		}
	}
}

// Will check whether a provided argument exists and tell the user if its missing
func checkIfExists(_arguments []string, _index int, _description string) {
	if len(_arguments) < _index+1 {
		helpers.PrintLine("Please provide "+_description, helpers.ERROR)
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("----- ditCLI v0.3 prerelease -----")
	if config.DitConfig.DemoModeActive {
		fmt.Println("--------- demo mode active -------")
	} else {
		fmt.Println("--------- live mode active -------")
	}
	fmt.Println("------------- General ------------")
	fmt.Println(" - dit setup\t\t\t\t\tCreates or imports the ethereum keys and creates a config")
	fmt.Println(" - dit update\t\t\t\t\tUpdates the ditCLIs config after an update of the ditCLI")
	fmt.Println(" - dit mode <MODE>\t\t\t\tSwitch between the modes of the ditCLI (live or demo)")
	fmt.Println("")
	fmt.Println("------------ Ethereum ------------")
	fmt.Println(" - dit get_address\t\t\t\tReturns ethereum address of the account")
	fmt.Println(" - dit get_balance\t\t\t\tReturns the " + config.DitConfig.Currency + " and KNW balance of the account")
	fmt.Println(" - dit export_keys\t\t\t\tReturns the unencrypted ethereum private key and address")
	fmt.Println("")
	fmt.Println("---------- Repositories ----------")
	fmt.Println(" - dit clone <REPOSITORY_URL>\t\t\tClones a repository and automatically calls 'dit init'")
	fmt.Println("\t\t\t\t\t\tafterwards")
	fmt.Println(" - dit init\t\t\t\t\tSets an already existing repository up for using it with dit")
	fmt.Println("")
	fmt.Println("------------- Voting -------------")
	fmt.Println(" - dit vote <PROPOSAL_ID> <CHOICE> <SALT>\tCasts a concealed vote on a proposed commit")
	fmt.Println(" - dit open <PROPOSAL_ID>\t\t\tOpens and reveals the vote commitment on a proposed commit")
	fmt.Println(" - dit finalize <PROPOSAL_ID>\t\t\tFinalizes the vote on proposed commit and claims the tokens")
	fmt.Println("")
	fmt.Println("---------- Git Commands ----------")
	fmt.Println(" - dit <GIT_COMMAND>\t\t\t\tExample Usage: dit commit -m \"Add new feature\"")
	fmt.Println("")
	fmt.Println("   Avialable git commands within the ditCLI:")
	fmt.Println(" - add, bisect, branch, checkout, commit\tYou can use these git commands as you are usually using them")
	fmt.Println(" - diff, fetch, grep, log, merge, pull\t\twith the regular git client. If you are not familiar with")
	fmt.Println(" - push, rebase, reset, rm, show, stash\t\tthe usage of one of these commands, please refer to the git ")
	fmt.Println(" - status, tag\t\t\t\t\tclients' help by invoking 'git help' directly")
	os.Exit(0)
}
