package demo

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ditcraft/cli/config"
	"github.com/ditcraft/cli/git"
	"github.com/ditcraft/cli/helpers"
	"github.com/ditcraft/cli/smartcontracts/KNWToken"
	"github.com/ditcraft/cli/smartcontracts/KNWVoting"
	"github.com/ditcraft/cli/smartcontracts/ditDemoCoordinator"
	"github.com/ditcraft/cli/smartcontracts/ditToken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const correctETHAddressLength = 42

// ProposeCommit will start a new proposal on the ditContract of this repository
func ProposeCommit(_branch string, _branchHeadHash string) (string, int, error) {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return "", 0, err
		} else if !passedKYC {
			return "", 0, errors.New("You didn't pass the KYC yet")
		}
	}

	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return "", 0, err
	}

	repoHash := getHashOfString(repository)

	// Gathering the the knowledge-labels from the config
	knowledgeLabels := config.DitConfig.DemoRepositories[repository].KnowledgeLabels

	connection, err := getConnection()
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
	if err != nil {
		return "", 0, err
	}

	nextAddress, err := ditCoordinatorInstance.NextDitCoordinator(nil)
	if err != nil {
		return "", 0, err
	}
	if nextAddress != common.HexToAddress("0") {
		helpers.PrintLine("There was an update to the ditCraft smartcontracts. Please update your ditCLI in order to interact with them.", helpers.INFO)
		helpers.PrintLine("Go to: https://github.com/ditcraft/cli", helpers.INFO)
		os.Exit(0)
	}

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the ditToken to access it
	ditTokenInstance, err := getDitTokenInstance(connection)
	if err != nil {
		return "", 0, err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Prompting the user which knowledge-label he wants to use for this proposal
	answerProposalSummary := ""
	userInputString := "Summarize your proposed changes (max. 140 characters)"
	for len(answerProposalSummary) < 4 || len(answerProposalSummary) > 140 {
		answerProposalSummary = helpers.GetUserInput(userInputString)
	}

	// Prompting the user which knowledge-label he wants to use for this proposal
	answerKnowledgeLabel := 0
	userInputString = "Which Knowledge-Label suits these changes the most?"
	for i := range knowledgeLabels {
		userInputString += " (" + strconv.Itoa(i+1) + ") " + knowledgeLabels[i]
	}
	for answerKnowledgeLabel < 1 || answerKnowledgeLabel > len(knowledgeLabels) {
		answerKnowledgeLabel, _ = strconv.Atoi(helpers.GetUserInput(userInputString))
	}

	// Retrieving the xDIT balance of the user
	xDITBalance, err := ditTokenInstance.BalanceOf(nil, myAddress)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
	}

	// Formatting the xDIT balance to a human-readable format
	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(xDITBalance)), big.NewFloat(1000000000000000000))

	// Retrieving the xDIT balance of the user
	freeKNWBalance, err := KNWTokenInstance.FreeBalanceOfLabel(nil, myAddress, knowledgeLabels[answerKnowledgeLabel-1])
	if err != nil {
		return "", 0, errors.New("Failed to retrieve free KNW balance")
	}

	// Formatting the xDIT balance to a human-readable format
	floatKNWBalance := new(big.Float).Quo((new(big.Float).SetInt(freeKNWBalance)), big.NewFloat(1000000000000000000))

	// Prompting the user how much stake he wants to set for this proposal
	answerStake := "0"
	floatStakeParsed, _ := strconv.ParseFloat(answerStake, 64)
	floatStake := big.NewFloat(floatStakeParsed)

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f xDIT", floatBalance), helpers.INFO)
	userInputString = fmt.Sprintf("How much do you want to stake?")
	for floatStake.Cmp(big.NewFloat(0)) == 0 || floatStake.Cmp(floatBalance) != -1 {
		answerStake = helpers.GetUserInput(userInputString)
		floatStakeParsed, _ = strconv.ParseFloat(answerStake, 64)
		floatStake = big.NewFloat(floatStakeParsed)
	}

	// Prompting the user how much KNW he wants to use for this proposal
	answerKNW := "0"
	floatKNWParsed, _ := strconv.ParseFloat(answerKNW, 64)
	floatKNW := big.NewFloat(floatKNWParsed)
	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f KNW for the label '%s'", floatKNWBalance, knowledgeLabels[answerKnowledgeLabel-1]), helpers.INFO)
	if floatKNWBalance.Cmp(big.NewFloat(0)) == 1 {
		userInputString = fmt.Sprintf("How much do you want to use?")
		keepAsking := true
		for keepAsking {
			answerKNW = helpers.GetUserInput(userInputString)
			floatKNWParsed, _ = strconv.ParseFloat(answerKNW, 64)
			var ok bool
			floatKNW, ok = new(big.Float).SetString(answerKNW)
			if ok && floatKNW.Cmp(big.NewFloat(0)) >= 0 && floatKNW.Cmp(floatKNWBalance) <= 0 {
				keepAsking = false
			}
		}
	}

	// Prompting the user whether he is sure of this proposal and its details
	floatStakeString := fmt.Sprintf("%.2f", floatStakeParsed)
	floatKNWString := fmt.Sprintf("%.2f", floatKNWParsed)
	helpers.PrintLine("  Proposing the changes with the following settings:", helpers.INFO)
	helpers.PrintLine("  Proposal Description: '"+answerProposalSummary+"'", helpers.INFO)
	helpers.PrintLine("  Knowledge Label: "+knowledgeLabels[answerKnowledgeLabel-1], helpers.INFO)
	helpers.PrintLine("  Amount of KNW Tokens: "+floatKNWString+" KNW", helpers.INFO)
	helpers.PrintLine("  The following stake will automatically be deducted: "+floatStakeString+" xDIT", helpers.INFO)
	userIsSure := helpers.GetUserInputChoice("Is that correct?", "y", "n")
	if userIsSure == "n" {
		return "", 0, errors.New("Canceled proposal of changes due to users choice")
	}
	fmt.Println()

	// Setting the value of the transaction to be the selected stake
	weiFloatStake, _ := (new(big.Float).Mul(floatStake, big.NewFloat(1000000000))).Int64()
	intStake := (new(big.Int).Mul(big.NewInt(weiFloatStake), big.NewInt(1000000000)))

	weiFloatKNW, _ := (new(big.Float).Mul(floatKNW, big.NewFloat(1000000000))).Int64()
	intKNW := (new(big.Int).Mul(big.NewInt(weiFloatKNW), big.NewInt(1000000000)))

	approvedBalance, err := ditTokenInstance.Allowance(nil, myAddress, common.HexToAddress(config.DitConfig.DitCoordinator))
	if err != nil {
		return "", 0, errors.New("Failed to retrieve xDIT allowance")
	}

	if approvedBalance.Cmp(intStake) == -1 {
		helpers.PrintLine("Since xDIT is a token, we need to set an allowance first", helpers.INFO)
		// Crerating the transaction (basic values)
		auth, err := populateTx(connection)
		if err != nil {
			return "", 0, err
		}

		_, err = ditTokenInstance.Approve(auth, common.HexToAddress(config.DitConfig.DitCoordinator), xDITBalance)
		if err != nil {
			return "", 0, err
		}

		newAllowance := approvedBalance
		for newAllowance.Cmp(approvedBalance) == 0 {
			newAllowance, err = ditTokenInstance.Allowance(nil, myAddress, common.HexToAddress(config.DitConfig.DitCoordinator))
			if err != nil {
				return "", 0, errors.New("Failed to retrieve xDIT allowance")
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println()
		helpers.PrintLine("Now the actual proposal will be executed", helpers.INFO)
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return "", 0, err
	}

	// Retrieving the last/current proposalID of the ditContract
	// (This will increment after a proposal, so we can see when the proposal is live)
	lastProposalID, err := ditCoordinatorInstance.GetCurrentProposalID(nil, repoHash)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve the current proposal id")
	}

	// Proposing the changes
	transaction, err := ditCoordinatorInstance.ProposeCommit(auth, repoHash, big.NewInt(int64(answerKnowledgeLabel-1)), intKNW, big.NewInt(int64(60)), big.NewInt(int64(60)), intStake)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return "", 0, errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return "", 0, err
	}

	// Waiting for the proposals transaction to be mined
	helpers.Printf("Waiting for proposal transaction to be mined", helpers.INFO)
	newProposalID := lastProposalID
	waitingFor := 0
	for newProposalID.Cmp(lastProposalID) == 0 {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the current proposal ID every 5 seconds
		newProposalID, err = ditCoordinatorInstance.GetCurrentProposalID(nil, repoHash)
		if err != nil {
			return "", 0, errors.New("Failed to retrieve the current proposal id")
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 180 {
			fmt.Printf("\n")
			helpers.PrintLine("Waiting for over 3 minutes, maybe the transaction or the network failed?", 1)
			helpers.PrintLine("Check at: https://blockscout.com/poa/dai/tx/"+transaction.Hash().Hex(), 1)
			os.Exit(0)
		}
	}
	fmt.Printf("\n")

	// Gathering the information of the new proposal from the ditContract (including the times to commit and reveal)
	newVote, err := gatherProposalInfo(connection, ditCoordinatorInstance, repoHash, newProposalID.Int64())
	if err != nil {
		return "", 0, err
	}

	newVote.BranchHash = _branchHeadHash

	// Adding the new vote to the config object
	config.DitConfig.DemoRepositories[repository].ActiveVotes[newProposalID.String()] = &newVote

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return "", 0, err
	}

	// Conwerting the stake and used KNW count into a float so that it's human-readable
	intEth, ok := new(big.Int).SetString(newVote.NumTokens, 10)
	if !ok {
		return "", 0, errors.New("Failed to parse big.int from string")
	}
	floatETH := new(big.Float).Quo((new(big.Float).SetInt(intEth)), big.NewFloat(1000000000000000000))

	intKNW, ok = new(big.Int).SetString(newVote.NumKNW, 10)
	if !ok {
		return "", 0, errors.New("Failed to parse big.int from string")
	}
	floatKNW = new(big.Float).Quo((new(big.Float).SetInt(intKNW)), big.NewFloat(1000000000000000000))

	// Formatting the time of the commit and reveal phase into a readable format
	timeReveal := time.Unix(int64(newVote.RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	timeCommit := time.Unix(int64(newVote.CommitEnd), 0)
	timeCommitString := timeCommit.Format("15:04:05 on 2006/01/02")

	// Returning the response to the user, we don't want to print this right here and now
	// since we are not sure whether the actual git commands will succeed.
	// So it will be printed afterwards in the main routine
	var responseString string
	responseString += "---------------------------"
	responseString += "\nSuccessfully proposed merge to master. Vote on proposal started with ID " + newProposalID.String() + ""
	responseString += fmt.Sprintf("\nYou staked %.2f %s and used %.2f KNW", floatETH, config.DitConfig.Currency, floatKNW)
	responseString += "\nThe vote will end at " + timeRevealString
	if config.DitConfig.DemoRepositories[repository].Provider == "github" {
		if _branch == "" {
			_branch = "dit_proposal_" + newProposalID.String()
		}
		responseString += "\nYour proposed branch is at https://" + repository + "/tree/" + _branch
	}
	responseString += "\n---------------------------"

	if config.DitConfig.DemoModeActive {
		fmt.Println()
		helpers.PrintLine("Since this is the demo mode, five auto-validators will automatically vote on your proposed changes. You will see the outcome after the vote ended.", 3)
		helpers.PrintLine("All validators now have time to vote on your proposed changes until "+timeCommitString, 3)
		helpers.PrintLine("To finalize the vote, execute '"+helpers.ColorizeCommand("finalize "+newProposalID.String())+"' after "+timeRevealString, 3)
		fmt.Println()
	}
	return responseString, int(newProposalID.Int64()), nil
}

// SearchForHashInVotes will search for a vote on a given branch hash and return the outcome
func SearchForHashInVotes(_branchHash string) (int, bool, bool, string, error) {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return 0, false, false, "", err
		} else if !passedKYC {
			return 0, false, false, "", errors.New("You didn't pass the KYC yet")
		}
	}

	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return 0, false, false, "", err
	}

	if config.DitConfig.DemoRepositories[repository] == nil {
		return 0, false, false, "", errors.New("Repository hasn't been initialized")
	}

	connection, err := getConnection()
	if err != nil {
		return 0, false, false, repository, err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return 0, false, false, repository, err
	}

	// Searching for the corresponding vote in the votes stored in the config
	var activeVote *config.ActiveVote
	var voteID int
	for key, value := range config.DitConfig.DemoRepositories[repository].ActiveVotes {
		if value.BranchHash == _branchHash {
			activeVote = value
			voteID, _ = strconv.Atoi(key)
			break
		}
	}
	if activeVote == nil {
		return 0, false, false, repository, nil
	}

	// Verifying whether the vote has ended
	voteEnded, err := KNWVotingInstance.VoteEnded(nil, big.NewInt(int64(activeVote.KNWVoteID)))
	if err != nil {
		return voteID, false, false, repository, errors.New("Failed to retrieve vote ending status")
	}

	// If it is now active it hasn't started yet or it's over
	if !voteEnded {
		return voteID, false, false, repository, nil
	}

	// Verifying whether the user has revealed his vote on this proposal
	votePassed, err := KNWVotingInstance.IsPassed(nil, big.NewInt(int64(activeVote.KNWVoteID)))
	if err != nil {
		return voteID, true, false, repository, errors.New("Failed to retrieve opening status")
	}

	// If this is the case, the user already revealed his vote
	if !votePassed {
		return voteID, true, false, repository, nil
	}

	return voteID, true, true, repository, nil
}

// CheckForKYC will return whether a user has already passed the KYC
func CheckForKYC() (bool, error) {
	connection, err := getConnection()
	if err != nil {
		return false, err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the KNWToken to access it
	ditCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
	if err != nil {
		return false, err
	}

	passedKYC, err := ditCoordinatorInstance.PassedKYC(nil, myAddress)
	if err != nil {
		return false, err
	}

	if !passedKYC {
		return false, nil
	}

	config.DitConfig.PassedKYC = true

	err = config.Save()
	if err != nil {
		return false, err
	}

	return true, nil
}

// populateTX will set the necessary values for a ethereum transaction
// amount of gas, gasprice, nonce, sign this with the private key
func populateTx(_connection *ethclient.Client) (*bind.TransactOpts, error) {
	// Retrieve the decrypted private key through a password prompt
	var err error
	privateKeyString, err := config.GetPrivateKey(true)
	if err != nil {
		return nil, err
	}

	// Converting the private key string into a private key object
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, errors.New("Failed to convert ethereum private-key")
	}

	// Retrieving the current pending nonce of our address
	pendingNonce, err := _connection.PendingNonceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address))
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}
	// Retrieving the current non-pending nonce of our address
	nonpendingNonce, err := _connection.NonceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address), nil)
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}

	// Edge-Case for slow nodes
	nonce := pendingNonce
	if nonpendingNonce > pendingNonce {
		nonce = nonpendingNonce
	}

	// Setting the values into the transaction-options object
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000)
	auth.GasPrice = big.NewInt(1000000000)

	return auth, nil
}

func getDitDemoCoordinatorInstance(_connection *ethclient.Client, _ditDemoCoordinatorAddress ...string) (*ditDemoCoordinator.DitDemoCoordinator, error) {
	var ditDemoCoordinatorAddressString string

	if len(_ditDemoCoordinatorAddress) > 0 {
		ditDemoCoordinatorAddressString = _ditDemoCoordinatorAddress[0]
	} else {
		if len(config.DitConfig.DitCoordinator) != correctETHAddressLength {
			return nil, errors.New("Invalid ditDemoCoordinator address, please do '" + helpers.ColorizeCommand("update") + "' first")
		}
		ditDemoCoordinatorAddressString = config.DitConfig.DitCoordinator
	}

	// Convertig the hex-string-formatted address into an address object
	ditDemoCoordinatorAddress := common.HexToAddress(ditDemoCoordinatorAddressString)

	// Create a new instance of the ditDemoCoordinator to access it
	ditDemoCoordinatorInstance, err := ditDemoCoordinator.NewDitDemoCoordinator(ditDemoCoordinatorAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditDemoCoordinator at provided address")
	}

	return ditDemoCoordinatorInstance, nil
}

func getKNWTokenInstance(_connection *ethclient.Client) (*KNWToken.KNWToken, error) {
	KNWTokenAddress := common.HexToAddress(config.DitConfig.KNWToken)

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := KNWToken.NewKNWToken(KNWTokenAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find KNWToken at provided address")
	}

	return KNWTokenInstance, nil
}

func getKNWVotingInstance(_connection *ethclient.Client) (*KNWVoting.KNWVoting, error) {
	KNWVotingAddress := common.HexToAddress(config.DitConfig.KNWVoting)

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := KNWVoting.NewKNWVoting(KNWVotingAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find KNWVoting at provided address")
	}

	return KNWVotingInstance, nil
}

// getConnection will return a connection to the ethereum blockchain
func getConnection() (*ethclient.Client, error) {
	for i := 0; i < len(config.EthereumNodes); i++ {
		connection, err := ethclient.Dial(config.EthereumNodes[i])
		if err != nil {
			if i == len(config.EthereumNodes)-1 {
				return nil, errors.New("Failed to connect to the ethereum network" + strconv.Itoa(i))
			}
		} else {
			networkID, err := connection.NetworkID(context.Background())
			if err != nil {
				if i == len(config.EthereumNodes)-1 {
					return nil, errors.New("Failed to connect to the ethereum network" + strconv.Itoa(i))
				}
			} else {
				if networkID.Cmp(big.NewInt(100)) == 0 {
					return connection, nil
				}
			}
		}
	}
	return nil, errors.New("Failed to connect to the ethereum network")
}

func getHashOfString(_string string) [32]byte {
	repoHash32 := [32]byte{}
	copy(repoHash32[:], crypto.Keccak256([]byte(_string))[:])

	return repoHash32
}

// gatherProposalInfo will retrieve necessary information about a proposal from the ditContract
func gatherProposalInfo(_connection *ethclient.Client, _ditCoordinatorInstance *ditDemoCoordinator.DitDemoCoordinator, _repoHash [32]byte, _proposalID int64) (config.ActiveVote, error) {
	// Retrieve the proposal object from the ditContract
	var newVote config.ActiveVote
	proposal, err := _ditCoordinatorInstance.ProposalsOfRepository(nil, _repoHash, big.NewInt(_proposalID))
	if err != nil {
		return newVote, errors.New("Failed to retrieve the new proposal")
	}

	// Store the information about this vote in an object
	newVote.KNWVoteID = int(proposal.KNWVoteID.Int64())
	newVote.KnowledgeLabel = proposal.KnowledgeLabel

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(_connection)
	if err != nil {
		return newVote, err
	}

	// Retrieving the information about this vote from the KNWVOting contract itself
	KNWPoll, err := KNWVotingInstance.Votes(nil, big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve vote information")
	}

	ethereumAddress := config.DitConfig.EthereumKeys.Address

	newVote.CommitEnd = int(KNWPoll.CommitEndDate.Int64())
	newVote.RevealEnd = int(KNWPoll.OpenEndDate.Int64())

	// Retrieving the amount of xDIT a user has staked for this vote
	numTokens, err := KNWVotingInstance.GetGrossStake(nil, big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve grossStake")
	}

	// Retrieving the number of votes a user has for this vote
	numVotes, err := KNWVotingInstance.GetAmountOfVotes(nil, common.HexToAddress(ethereumAddress), big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve numVotes")
	}

	// Retrieving the number of KNW tokens a user has staked for this vote
	numKNW, err := KNWVotingInstance.GetUsedKNW(nil, common.HexToAddress(ethereumAddress), big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve numKNW")
	}

	newVote.NumTokens = numTokens.String()
	newVote.NumVotes = numVotes.String()
	newVote.NumKNW = numKNW.String()

	return newVote, nil
}

func getDitTokenInstance(_connection *ethclient.Client) (*ditToken.MintableERC20, error) {
	ditTokenAddress := common.HexToAddress(config.DitConfig.DitToken)

	// Create a new instance of the KNWToken contract to access it
	ditTokenInstance, err := ditToken.NewMintableERC20(ditTokenAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditToken at provided address")
	}

	return ditTokenInstance, nil
}
