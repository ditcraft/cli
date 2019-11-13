package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ditcraft/cli/api"

	"github.com/ditcraft/cli/config"
	"github.com/ditcraft/cli/demo"
	"github.com/ditcraft/cli/ethereum"
	"github.com/ditcraft/cli/git"
	"github.com/ditcraft/cli/helpers"
	ditLog "github.com/ditcraft/cli/log"
	"github.com/ditcraft/cli/router"
	"github.com/logrusorgru/aurora"
)

// The current dit coordinator address
var liveDitCoodinator = "0x4BB184D45e8e3a6dC6684fda7110894cd6226f16"
var demoDitCoodinator = "0x9A65c773A216a4F4748B1D65C0fB039d9F2b223D"

// The current version
var version = "v0.3.1"

func main() {
	var err error

	args := os.Args[1:]

	// Loading the config
	err = config.Load()

	// Loading the logfile
	logErr := ditLog.Load()
	if logErr != nil {
		helpers.PrintLine(logErr.Error(), helpers.ERROR)
	}

	// If no arguments are provided print the usage information
	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	// Converting everything to lowercase, so that a command will not fail if something
	// is capitalized by accident
	command := strings.ToLower(args[0])

	logErr = ditLog.AddCommand("dit " + strings.Join(args, " "))
	if logErr != nil {
		if len(logErr.Error()) > 0 {
			helpers.PrintLine(logErr.Error(), helpers.ERROR)
		}
	}

	// If there is an error during the config load and the user is not
	// in the process of setting dit up: print an error
	if err != nil && command != "setup" && command != "update" {
		helpers.PrintLine(err.Error(), helpers.ERROR)
		helpers.ExitAndLog(0)
	}

	if err == nil && (command != "setup" && command != "update") {
		if config.DitConfig.Version < config.Version {
			helpers.PrintLine("Your config is not up to date, the ditCLI might not work correctly", helpers.WARN)
			helpers.PrintLine("To fix this call '"+helpers.ColorizeCommand("update")+"'", helpers.WARN)
		} else if config.DitConfig.Version > config.Version {
			helpers.PrintLine("Your config file was created with a newer version of the ditCLI", helpers.ERROR)
			helpers.PrintLine("To fix this update the ditCLI", helpers.ERROR)
			helpers.ExitAndLog(0)
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
					helpers.PrintLine("Visit the ditExplorer (https://explorer.ditcraft.io), login via GitHub or Twitter and follow the instructions.", helpers.INFO)
					helpers.PrintLine("Provide your ditCLI's address in the process: "+config.DitConfig.EthereumKeys.Address+"", helpers.INFO)
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
					helpers.PrintLine("It seems that the smart contracts have been updated, which means that you need to re-do the KYC.", helpers.INFO)
					helpers.PrintLine("Visit the ditExplorer (https://explorer.ditcraft.io), login via GitHub or Twitter and follow the instructions.", helpers.INFO)
					helpers.PrintLine("Provide your ditCLI's address in the process: "+config.DitConfig.EthereumKeys.Address+"", helpers.INFO)
				}
			}
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
			helpers.ExitAndLog(0)
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
				helpers.PrintLine("Visit the ditExplorer (https://explorer.ditcraft.io), login via GitHub or Twitter and follow the instructions.", helpers.INFO)
				helpers.PrintLine("Provide your ditCLI's address in the process: "+config.DitConfig.EthereumKeys.Address+"", helpers.INFO)
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
		err = router.GetBalances()
		break
	case "get_address", "address":
		// Return the ETH address
		helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, helpers.INFO)
		helpers.PrintLine("URL: https://explorer.ditcraft.io/address/"+config.DitConfig.EthereumKeys.Address, helpers.INFO)
		break
	case "export_keys", "export_key":
		helpers.PrintLine("Exporting your keys will print your ethereum private key in plain text", helpers.WARN)
		helpers.PrintLine("This might compromise your keys, please don't proceed if you are not sure about this", helpers.WARN)
		answer := helpers.GetUserInputChoice("Are you sure that you want to proceed?", "y", "n")
		if answer == "y" {
			// Return the ETH address and private key
			var privateKey string
			privateKey, err = config.GetPrivateKey(false)
			if len(privateKey) == 64 {
				helpers.PrintLine("Ethereum Address: "+config.DitConfig.EthereumKeys.Address, helpers.INFO)
				helpers.PrintLine("Ethereum Private-Key: "+privateKey, helpers.CONFIDENTIAL)
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
				err = git.SideBranchIsPushed(args[1])
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
							voteID, voteEnded, votePassed, repository, err = ethereum.SearchForHashInVotes(branchHash)
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
											if config.DitConfig.DemoModeActive {
												config.DitConfig.DemoRepositories[repository].ActiveVotes[strconv.Itoa(voteID)].NewHeadHash = newHeadHash
											} else {
												config.DitConfig.LiveRepositories[repository].ActiveVotes[strconv.Itoa(voteID)].NewHeadHash = newHeadHash
											}
											err = config.Save()
										}
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
						_, errRevParse := git.ExecuteCommandWithoutStdOut("rev-parse", "origin/master")
						changes, errTree := git.ExecuteCommandWithoutStdOut("ls-tree", "--name-only", "-r", masterHeadHash)
						if errRevParse != nil && errTree == nil {
							changes = strings.ToLower(changes)
							changes = strings.Replace(changes, ".md", "", 1)
							if changes == "readme" {
								// Allowing the push in this case since it's the first one and
								// only contains a README to ensure that the ditCLI works correctly
								_ = git.ExecuteCommand(args[:]...)
							}
						} else {
							helpers.PrintLine("You are not allowed to push changes to master without prior approval from the community", helpers.ERROR)
						}
						// TODO: offer to start a new vote on these changes?
					}
				}
			}
		}
		break
	case "checkout":
		var isOnMaster bool
		isOnMaster, err = git.IsOnMaster()
		if isOnMaster {
			var output string
			output, err = git.ExecuteCommandWithoutStdOut("status")
			if strings.Contains(output, "No commits yet") {
				helpers.PrintLine("It seems that the master branch is empty and you want to switch to another branch.", helpers.ERROR)
				helpers.PrintLine("The ditCLI depends on an exisiting master-branch with a valid state, please create a README first to initialize the master branch.", helpers.ERROR)
				helpers.ExitAndLog(0)
			}
		}
		_ = git.ExecuteCommand(args[:]...)
	case "add",
		"bisect",
		"branch",
		"commit",
		"config",
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
	case "upload_logs":
		err = api.UploadLog(version)
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
		checkIfExists(args, 1, "a proposal ID - dit vote <PROPOSAL_ID> <CHOICE> <SALT (optional)>")
		checkIfExists(args, 2, "your choice (accept/deny) - dit vote <PROPOSAL_ID> <CHOICE> <SALT (optional)>")
		salt := ""
		if len(args) > 3 {
			salt = args[3]
		}
		// Votes on a proposal
		switch args[2] {
		case "accept", "yes", "y", "yay", "1", "true":
			if config.DitConfig.DemoModeActive {
				err = demo.Vote(args[1], "1", salt)
			} else {
				err = ethereum.Vote(args[1], "1", salt)
			}
		case "deny", "decline", "no", "n", "nay", "0", "false":
			if config.DitConfig.DemoModeActive {
				err = demo.Vote(args[1], "0", salt)
			} else {
				err = ethereum.Vote(args[1], "0", salt)
			}
		}
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
					helpers.PrintLine("You are now allowed to integrate your changes into the master branch, e.g. through "+helpers.ColorizeCommand("merge <branch>"), helpers.INFO)
				} else {
					helpers.PrintLine(fmt.Sprintf("You are %s allowed to integrate your changes into the master branch", aurora.Bold(aurora.Red("NOT"))), helpers.INFO)
				}
			}
		}
		break
	// case "balance":
	// 	err = ethereum.GetDaiPrice()
	// 	break
	// case "wallet":
	// 	if args[1] == "deposit" {
	// 		err = ethereum.SwapETHtoXDai()
	// 	}
	// 	break
	case "-v", "--version", "version":
		helpers.PrintLine(version, helpers.INFO)
	default:
		printUsage()
		ditLog.AddToLog("Printed usage\n")
		helpers.ExitAndLog(0)
		break
	}

	if err != nil {
		if len(err.Error()) > 0 {
			helpers.PrintLine(err.Error(), helpers.ERROR)
		}
	}
	helpers.ExitAndLog(0)
}

// Will check whether a provided argument exists and tell the user if its missing
func checkIfExists(_arguments []string, _index int, _description string) {
	if len(_arguments) < _index+1 {
		helpers.PrintLine("Please provide "+_description, helpers.ERROR)
		helpers.ExitAndLog(0)
	}
}

func printUsage() {
	fmt.Println("---- ditCLI v0.3.1 prerelease ----")
	// fmt.Println("---------- ditCLI v0.3.1 ---------")
	if config.DitConfig.DemoModeActive {
		fmt.Println("--------- demo mode active -------")
	} else {
		fmt.Println("--------- live mode active -------")
	}
	fmt.Println("------------- General ------------")
	fmt.Println(" - dit setup\t\t\t\t\tCreates or imports the ethereum keys and creates a config")
	fmt.Println(" - dit update\t\t\t\t\tUpdates the ditCLIs config after an update of the ditCLI")
	fmt.Println(" - dit mode <MODE>\t\t\t\tSwitch between the modes of the ditCLI (live or demo)")
	fmt.Println(" - dit upload_logs\t\t\t\tUploads the ditCLI's logs onto our servers for troubleshooting")
	fmt.Println("\t\t\t\t\t\tor feedback")
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
	fmt.Println(" - dit merge <BRANCH>\t\t\t\tWill initiate a new vote when this command is executed in")
	fmt.Println("\t\t\t\t\t\tthe master branch to gain community approval")
	fmt.Println(" - dit vote <PROPOSAL_ID> <CHOICE>\t\tCasts a concealed vote on a proposal where <CHOICE> is")
	fmt.Println("\t\t\t\t\t\teither 'accept' or 'deny'. If a <SALT> is appended as a")
	fmt.Println("\t\t\t\t\t\tthird argument it will be used, otherwise the client will")
	fmt.Println("\t\t\t\t\t\tgenerate a random one to conceal the vote")
	fmt.Println(" - dit open <PROPOSAL_ID>\t\t\tOpens and reveals the vote commitment on a proposal")
	fmt.Println(" - dit finalize <PROPOSAL_ID>\t\t\tFinalizes the vote on a proposal and claims the tokens")
	fmt.Println("")
	fmt.Println("---------- Git Commands ----------")
	fmt.Println(" - dit <GIT_COMMAND>\t\t\t\tExample Usage: dit commit -m \"Add new feature\"")
	fmt.Println("")
	fmt.Println("   Avialable git commands within the ditCLI:")
	fmt.Println(" - add, bisect, branch, checkout, commit\tYou can use these git commands as you are usually using")
	fmt.Println(" - diff, fetch, grep, log, merge, pull\t\tthem with the regular git client. If you are not familiar")
	fmt.Println(" - push, rebase, reset, rm, show, stash\t\twith the usage of one of these commands, please refer to")
	fmt.Println(" - status, tag\t\t\t\t\tthe git clients' help by invoking 'git help' directly")
}
