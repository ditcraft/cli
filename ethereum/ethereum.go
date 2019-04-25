package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ditcraft/client/config"
	"github.com/ditcraft/client/git"
	"github.com/ditcraft/client/helpers"
	"github.com/ditcraft/client/smartcontracts/KNWToken"
	"github.com/ditcraft/client/smartcontracts/KNWVoting"
	"github.com/ditcraft/client/smartcontracts/ditCoordinator"
	"github.com/ditcraft/client/smartcontracts/ditDemoCoordinator"
	"github.com/ditcraft/client/smartcontracts/ditToken"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const correctETHAddressLength = 42

// SetDitCoordinator will write the dit coordinators address to the config and
// retrieve the addresses of the KNWToken and KNWVoting contracts
func SetDitCoordinator(_ditCoordinatorAddressString string) error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the ditCoordinator to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection, _ditCoordinatorAddressString)
	if err != nil {
		return err
	}

	// Retrieving the KNWVoting contracts address from the ditCoordinator
	KNWVotingAddress, err := ditCoordinatorInstance.KNWVotingAddress(nil)
	if err != nil {
		return errors.New("Failed to retrieve address of the KNWVoting contract")
	}

	// Check whether the retrieved address is valid (i.e. not zero)
	if KNWVotingAddress == common.HexToAddress("0") {
		return errors.New("Received an invalid address of the KNWVoting contract")
	}

	// Retrieving the KNWToken contracts address from the ditCoordinator
	KNWTokenAddress, err := ditCoordinatorInstance.KNWTokenAddress(nil)
	if err != nil {
		return errors.New("Failed to retrieve address of the KNWToken contract")
	}

	// Check whether the retrieved address is valid (i.e. not zero)
	if KNWTokenAddress == common.HexToAddress("0") {
		return errors.New("Received an invalid address of the KNWToken contract")
	}

	// If in demo mode, also retrieve the xDit token address
	if config.DitConfig.DemoModeActive {
		// Create a new instance of the ditDemoCoordinator to access it
		ditDemoCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection, _ditCoordinatorAddressString)
		if err != nil {
			return err
		}

		// Retrieving the ditToken contracts address from the ditDemoCoordinator
		ditTokenAddress, err := ditDemoCoordinatorInstance.XDitTokenAddress(nil)
		if err != nil {
			return errors.New("Failed to retrieve address of the KNWToken contract")
		}
		config.DitConfig.DitToken = ditTokenAddress.Hex()
		config.DitConfig.Currency = "xDit"
	} else {
		config.DitConfig.DitToken = ""
		config.DitConfig.Currency = "xDai"
	}

	// Setting the retrieved addresses to the config object
	config.DitConfig.DitCoordinator = _ditCoordinatorAddressString
	config.DitConfig.KNWVoting = KNWVotingAddress.Hex()
	config.DitConfig.KNWToken = KNWTokenAddress.Hex()
	config.DitConfig.PassedKYC = false

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return nil
	}

	return nil
}

// InitDitRepository will initialize a new repository and store its ditContracts address
// together with additional information into the config
// When there is no ditContract for this repository, the user will be prompted to deploy one
func InitDitRepository(_optionalRepository ...string) error {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return err
		} else if !passedKYC {
			return errors.New("You didn't pass the KYC yet")
		}
	}

	var repository string

	// If a repository name is already provided with the call of this function we will use it...
	if len(_optionalRepository) > 0 {
		repository = _optionalRepository[0]
	} else {
		// ...otherwise we will retrieve the name of this repository
		var err error
		repository, err = git.GetRepository()
		if err != nil {
			return err
		}
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	// Searching through the repositories that are in our config
	for i := 0; i < len(repositoryArray); i++ {
		// If the repository was found in our config it has already been initialized
		// In that case we don't need to proceed further
		if repositoryArray[i].Name == repository {
			if len(_optionalRepository) > 0 {
				return nil
			}
			helpers.PrintLine("Repository is already initialized", 0)
			os.Exit(0)
		}
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	repoHash := GetHashOfString(repository)

	// Create a new instance of the ditCoordinator to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the address of the ditContract that was deployed for this repository
	isInitialized, err := ditCoordinatorInstance.RepositoryIsInitialized(nil, repoHash)
	if err != nil {
		return err
	}

	// If the address is zero, there is no contract deployed
	if !isInitialized {
		// Prompt the user whether he wants to deploy one
		answer := helpers.GetUserInputChoice("This repository hasn't been initialized yet, do you want to initialize it?", "y", "n")
		if answer == "y" {
			// If yes: deploy the ditContract
			err = initDitRepository(ditCoordinatorInstance, repoHash)
			if err != nil {
				return err
			}
		} else {
			// If not: exit
			helpers.PrintLine("No ditContract deployed - repository can't be used with dit", 1)
			os.Exit(0)
		}
	}

	// Retrieving the knowledge-labels of this ditContract
	var knowledgeLabels []string
	for i := 0; i < 3; i++ {
		contractKnowledgeLabels, err := ditCoordinatorInstance.GetKnowledgeLabels(nil, repoHash, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		if len(contractKnowledgeLabels) > 0 {
			knowledgeLabels = append(knowledgeLabels, contractKnowledgeLabels)
		}
	}

	// Inserting this repositories details into the cobfig
	var newRepository config.Repository
	newRepository.Name = repository
	if strings.Contains(repository, "github.com") {
		newRepository.Provider = "github"
	}
	newRepository.KnowledgeLabels = knowledgeLabels
	newRepository.ActiveVotes = make([]config.ActiveVote, 0)
	repositoryArray = append(repositoryArray, newRepository)
	if config.DitConfig.DemoModeActive {
		config.DitConfig.DemoRepositories = repositoryArray
	} else {
		config.DitConfig.LiveRepositories = repositoryArray
	}

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return nil
	}

	return nil
}

// ProposeCommit will start a new proposal on the ditContract of this repository
func ProposeCommit(_commitMessage string) (string, int, error) {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return "", 0, err
		} else if !passedKYC {
			return "", 0, errors.New("You didn't pass the KYC yet")
		}
	}

	// Searching for this repositories object in the config
	repoIndex, err := searchForRepoInConfig()
	if err != nil {
		return "", 0, err
	}

	repoHash := GetHashOfString(config.DitConfig.LiveRepositories[repoIndex].Name)

	// Gathering the the knowledge-labels from the config
	knowledgeLabels := config.DitConfig.LiveRepositories[repoIndex].KnowledgeLabels

	connection, err := getConnection()
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return "", 0, err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Prompting the user which knowledge-label he wants to use for this proposal
	answerKnowledgeLabel := 0
	userInputString := "Which Knowledge-Label suits this commit most?"
	for i := range knowledgeLabels {
		userInputString += " (" + strconv.Itoa(i+1) + ") " + knowledgeLabels[i]
	}
	for answerKnowledgeLabel < 1 || answerKnowledgeLabel > len(knowledgeLabels) {
		answerKnowledgeLabel, _ = strconv.Atoi(helpers.GetUserInput(userInputString))
	}

	// Retrieving the xDai balance of the user
	xDaiBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
	}

	// Formatting the xDai balance to a human-readable format
	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(xDaiBalance)), big.NewFloat(1000000000000000000))

	// Retrieving the xDit balance of the user
	freeKNWBalance, err := KNWTokenInstance.FreeBalanceOfLabel(nil, myAddress, knowledgeLabels[answerKnowledgeLabel-1])
	if err != nil {
		return "", 0, errors.New("Failed to retrieve free KNW balance")
	}

	// Formatting the xDit balance to a human-readable format
	floatKNWBalance := new(big.Float).Quo((new(big.Float).SetInt(freeKNWBalance)), big.NewFloat(1000000000000000000))

	// Prompting the user how much stake he wants to set for this proposal
	answerStake := "0"
	floatStakeParsed, _ := strconv.ParseFloat(answerStake, 64)
	floatStake := big.NewFloat(floatStakeParsed)

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f xDai", floatBalance), 0)
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

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f KNW for the label '%s'", floatKNWBalance, knowledgeLabels[answerKnowledgeLabel-1]), 0)
	if floatKNWBalance.Cmp(big.NewFloat(0)) == 1 {
		userInputString = fmt.Sprintf("How much do you want to use?")
		keepAsking := true
		for keepAsking {
			answerKNW = helpers.GetUserInput(userInputString)
			floatKNWParsed, _ = strconv.ParseFloat(answerKNW, 64)
			floatKNW = big.NewFloat(floatKNWParsed)
			if floatKNW.Cmp(big.NewFloat(0)) >= 0 && floatKNW.Cmp(floatKNWBalance) <= 0 {
				keepAsking = false
			}
		}
	}

	// Prompting the user whether he is sure of this proposal and its details
	floatStakeString := fmt.Sprintf("%.2f", floatStakeParsed)
	helpers.PrintLine("  Proposing the commit with the following settings:", 0)
	helpers.PrintLine("  Commit Message: "+_commitMessage+"", 0)
	helpers.PrintLine("  Knowledge Label: "+knowledgeLabels[answerKnowledgeLabel-1], 0)
	helpers.PrintLine("  The following stake with automatically be deducted: "+floatStakeString+"xDai", 0)
	userIsSure := helpers.GetUserInputChoice("Is that correct?", "y", "n")
	if userIsSure == "n" {
		return "", 0, errors.New("Canceled proposal of commit due to users choice")
	}
	fmt.Println()
	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return "", 0, err
	}

	// Setting the value of the transaction to be the selected stake
	weiFloatStake, _ := (new(big.Float).Mul(floatStake, big.NewFloat(1000000000))).Int64()
	intStake := (new(big.Int).Mul(big.NewInt(weiFloatStake), big.NewInt(1000000000)))
	auth.Value = intStake

	weiFloatKNW, _ := (new(big.Float).Mul(floatKNW, big.NewFloat(1000000000))).Int64()
	intKNW := (new(big.Int).Mul(big.NewInt(weiFloatKNW), big.NewInt(1000000000)))

	// Retrieving the last/current proposalID of the ditContract
	// (This will increment after a proposal, so we can see when the proposal is live)
	lastProposalID, err := ditCoordinatorInstance.GetCurrentProposalID(nil, repoHash)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve the current proposal id")
	}

	// Proposing the commit
	transaction, err := ditCoordinatorInstance.ProposeCommit(auth, repoHash, big.NewInt(int64(answerKnowledgeLabel-1)), intKNW, big.NewInt(int64(180)), big.NewInt(int64(180)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return "", 0, errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return "", 0, err
	}

	// Waiting for the proposals transaction to be mined
	helpers.Printf("Waiting for commit proposal transaction to be mined", 0)
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

	// Adding the new vote to the config object
	config.DitConfig.LiveRepositories[repoIndex].ActiveVotes = append(config.DitConfig.LiveRepositories[repoIndex].ActiveVotes, newVote)

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return "", 0, err
	}

	// Conwerting the stake and used KNW count into a float so that it's human-readable
	intxDai, ok := new(big.Int).SetString(newVote.NumTokens, 10)
	if !ok {
		return "", 0, errors.New("Failed to parse big.int from string")
	}
	floatxDai := new(big.Float).Quo((new(big.Float).SetInt(intxDai)), big.NewFloat(1000000000000000000))

	intKNW, ok = new(big.Int).SetString(newVote.NumKNW, 10)
	if !ok {
		return "", 0, errors.New("Failed to parse big.int from string")
	}
	floatKNW = new(big.Float).Quo((new(big.Float).SetInt(intKNW)), big.NewFloat(1000000000000000000))

	// Formatting the time of the commit and reveal phase into a readable format
	timeReveal := time.Unix(int64(newVote.RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	// Returning the response to the user, we don't want to print this right here and now
	// since we are not sure whether the actual git commands will succeed.
	// So it will be printed afterwards in the main routine
	var responseString string
	responseString += "---------------------------"
	responseString += "\nSuccessfully proposed commit. Vote on proposal started with ID " + strconv.Itoa(int(newVote.ID)) + ""
	responseString += fmt.Sprintf("\nYou staked %.2f %s and used %f KNW.", floatxDai, config.DitConfig.Currency, floatKNW)
	responseString += "\nThe vote will end at " + timeRevealString
	if config.DitConfig.LiveRepositories[repoIndex].Provider == "github" {
		responseString += "\nYour commit is at https://" + config.DitConfig.LiveRepositories[repoIndex].Name + "/tree/dit_proposal_" + strconv.Itoa(int(newVote.ID))
	}
	responseString += "\n---------------------------"

	return responseString, int(newProposalID.Int64()), nil
}

// Vote will cast a users vote on a proposal
func Vote(_proposalID string, _choice string, _salt string) error {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return err
		} else if !passedKYC {
			return errors.New("You didn't pass the KYC yet")
		}
	}

	// Converting the stdin string input of the user into Ints
	proposalID, _ := strconv.Atoi(_proposalID)
	choice, _ := strconv.Atoi(_choice)
	salt, _ := strconv.Atoi(_salt)

	// Searching for this repositories object in the config
	repoIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repositoryArray[repoIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCooordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the proposal object from the ditContract
	proposal, err := ditCooordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return errors.New("Failed to retrieve proposal")
	}

	// Verifiying whether the proposal is valid (if KNWVoteID is zero its not valid or non existent)
	if proposal.KNWVoteID.Int64() == 0 {
		return errors.New("Invalid proposalID")
	}

	if proposal.Proposer == myAddress {
		return errors.New("You can't vote on your own proposal")
	}

	// Retrieving the default stake from the ditContract
	requiredStake, err := KNWVotingInstance.GetGrossStake(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve the required stake of the vote")
	}

	floatStake := new(big.Float).Quo((new(big.Float).SetInt(requiredStake)), big.NewFloat(1000000000000000000))

	// Retrieving the xDit balance of the user
	freeKNWBalance, err := KNWTokenInstance.FreeBalanceOfLabel(nil, myAddress, proposal.KnowledgeLabel)
	if err != nil {
		return errors.New("Failed to retrieve free KNW balance")
	}

	// Formatting the xDit balance to a human-readable format
	floatKNWBalance := new(big.Float).Quo((new(big.Float).SetInt(freeKNWBalance)), big.NewFloat(1000000000000000000))

	// Prompting the user how much KNW he wants to use for this proposal
	answerKNW := "0"
	floatKNWParsed, _ := strconv.ParseFloat(answerKNW, 64)
	floatKNW := big.NewFloat(floatKNWParsed)

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f KNW for the label '%s'", floatKNWBalance, proposal.KnowledgeLabel), 0)
	if floatKNWBalance.Cmp(big.NewFloat(0)) == 1 {
		userInputString := fmt.Sprintf("How much do you want to use?")
		for floatKNW.Cmp(big.NewFloat(0)) == -1 || floatStake.Cmp(floatKNWBalance) != -1 {
			answerKNW = helpers.GetUserInput(userInputString)
			floatKNWParsed, _ = strconv.ParseFloat(answerKNW, 64)
			floatKNW = big.NewFloat(floatKNWParsed)
		}
	}

	// Prompting the user whether he is sure of this vote and its details
	floatKNWString := fmt.Sprintf("%.2f", floatKNWParsed)

	helpers.PrintLine("Voting on the commit with the following settings:", 0)
	helpers.PrintLine("Knowledge Label: "+proposal.KnowledgeLabel, 0)
	helpers.PrintLine("Amount of KNW Tokens: "+floatKNWString+" KNW", 0)
	fmt.Println()
	helpers.PrintLine("Voting on this proposal will automatically deduct the required stake from you account.", 0)
	helpers.PrintLine(fmt.Sprintf("Required stake: %.2f", floatStake), 0)
	helpers.PrintLine("All participants of the vote will counter-stake the proposer.", 0)
	helpers.PrintLine("You will receive the remaining stake back, no matter how the vote ends.", 0)

	answer := helpers.GetUserInputChoice("Is this okay for you?", "y", "n")
	if answer == "n" {
		// If not: exit
		helpers.PrintLine("No vote executed due to users choice", 1)
		os.Exit(0)
	}

	// In order to create a valid abi-encoded hash of the vote choice and salt
	// we need to create an abi object
	uint256Type, _ := abi.NewType("uint256")
	arguments := abi.Arguments{
		{
			Type: uint256Type,
		},
		{
			Type: uint256Type,
		},
	}

	// We will now put pack this abi object into a bytearray
	bytes, _ := arguments.Pack(
		big.NewInt(int64(choice)),
		big.NewInt(int64(salt)),
	)

	// And finally hash this bytearray with keccak256, resulting in the votehash
	voteHash := crypto.Keccak256Hash(bytes)

	// Verifying whether the commit period of this vote is active
	commitPeriodActive, err := KNWVotingInstance.CommitPeriodActive(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve opening status")
	}

	// If it is now active it's probably over
	if !commitPeriodActive {
		return errors.New("The commit phase of this vote has ended")
	}

	// Verifying whether the user has already commited a vote on this proposal
	oldDidCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve commit status")
	}

	// If this is the case, the user can not vote again
	if oldDidCommit {
		return errors.New("You already voted on this proposal")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	// Setting the value of the transaction to be the default stake
	auth.Value = requiredStake

	weiFloatKNW, _ := (new(big.Float).Mul(floatKNW, big.NewFloat(1000000000000000000))).Int64()
	intKNW := big.NewInt(weiFloatKNW)

	// Voting on the proposal
	transaction, err := ditCooordinatorInstance.VoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), voteHash, intKNW)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return errors.New("Failed to commit the vote: " + err.Error())
	}

	// Waiting for the voting transaction to be mined
	helpers.Printf("Waiting for voting transaction to be mined", 0)
	waitingFor := 0
	newDidCommit := oldDidCommit
	for newDidCommit == oldDidCommit {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the commit status of the user every 5 seconds
		newDidCommit, err = KNWVotingInstance.DidCommit(nil, myAddress, proposal.KNWVoteID)
		if err != nil {
			return errors.New("Failed to retrieve commit status")
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

	// Gathering the information of the proposal from the ditContract
	newVote, err := gatherProposalInfo(connection, ditCooordinatorInstance, repoHash, int64(proposalID))
	if err != nil {
		return err
	}

	// We will also store the users choice and salt, so that the user doesn't need to remember the salt
	// when he will reveal the vote later on
	newVote.Choice = choice
	newVote.Salt = salt

	// Adding the new vote to the config object
	repositoryArray[repoIndex].ActiveVotes = append(repositoryArray[repoIndex].ActiveVotes, newVote)
	if config.DitConfig.DemoModeActive {
		config.DitConfig.DemoRepositories = repositoryArray
	} else {
		config.DitConfig.LiveRepositories = repositoryArray
	}

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return err
	}

	// Formatting the time of the commit and reveal phase into a readable format
	timeCommit := time.Unix(int64(newVote.CommitEnd), 0)
	timeCommitString := timeCommit.Format("15:04:05 on 2006/01/02")
	timeReveal := time.Unix(int64(newVote.RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	if choice == 1 {
		helpers.PrintLine("You successfully voted in favor of the proposal", 0)
	} else {
		helpers.PrintLine("You successfully voted against the proposal", 0)
	}

	// Letting the user know when and how he has to reveal the vote
	helpers.PrintLine("With dit, votes are casted in a concealed manner through a commitment scheme.", 0)
	helpers.PrintLine("This means that votes have to be revealed to the public once the commit-phase is over.", 0)
	helpers.PrintLine("Please open your vote with '"+helpers.ColorizeCommand("open "+strconv.Itoa(int(proposalID)))+"' between "+timeCommitString+" and "+timeRevealString, 0)

	return nil
}

// Open will open and thus reveal a vote (that was commited through a hash) to the public during the opening phase
func Open(_proposalID string) error {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return err
		} else if !passedKYC {
			return errors.New("You didn't pass the KYC yet")
		}
	}

	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	repoIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repositoryArray[repoIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCooordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Searching for the corresponding vote in the votes stored in the config
	var voteIndex int
	for i := range repositoryArray[repoIndex].ActiveVotes {
		if repositoryArray[repoIndex].ActiveVotes[i].ID == proposalID {
			voteIndex = i
			break
		}
	}

	// Verifying whether the reveal period of this vote is active
	revealPeriodActive, err := KNWVotingInstance.OpenPeriodActive(nil, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve opening status")
	}

	// If it is now active it hasn't started yet or it's over
	if !revealPeriodActive {
		return errors.New("The opening phase of this vote is not active")
	}

	// Verifying whether the user has commited a vote on this proposal
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve commit status")
	}

	// If this is not the case the user never participated in this proposal through a vote
	if !didCommit {
		return errors.New("You didn't vote on this proposal")
	}

	// Verifying whether the user has revealed his vote on this proposal
	oldDidReveal, err := KNWVotingInstance.DidOpen(nil, myAddress, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve opening status")
	}

	// If this is the case, the user already revealed his vote
	if oldDidReveal {
		return errors.New("You already opened your vote on this proposal")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	// Gathering the original choice and the salt from the config
	choice := big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].Choice))
	salt := big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].Salt))

	// Revealing the vote on the proposal
	transaction, err := ditCooordinatorInstance.OpenVoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), choice, salt)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return errors.New("Failed to open the vote: " + err.Error())
	}

	// Waiting for the reveal transaction to be mined
	helpers.Printf("Waiting for opening transaction to be mined", 0)
	waitingFor := 0
	newDidReveal := oldDidReveal
	for newDidReveal == oldDidReveal {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the reveal status of the user every 5 seconds
		newDidReveal, err = KNWVotingInstance.DidOpen(nil, myAddress, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
		if err != nil {
			return errors.New("Failed to retrieve opening status")
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

	// Formatting the time of the reveal phase into a readable format
	timeReveal := time.Unix(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	helpers.PrintLine("Successfully opened your vote", 0)
	helpers.PrintLine("To finalize it when the vote is over, execute '"+helpers.ColorizeCommand("finalize "+_proposalID)+"' after "+timeRevealString, 3)

	return nil
}

// Finalize will finalize a vote as it will trigger the calculation of the reward of this user
// including the xDai and KNW reward in case of a voting for the winning decision
// or the losing of xDai and KNW in case of a voting for the losing decision
// The first caller who executes this will also trigger the calculation whether the vote passed or not
func Finalize(_proposalID string) (bool, bool, error) {
	if !config.DitConfig.PassedKYC {
		passedKYC, err := CheckForKYC()
		if err != nil {
			return false, false, err
		} else if !passedKYC {
			return false, false, errors.New("You didn't pass the KYC yet")
		}
	}

	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	repoIndex, err := searchForRepoInConfig()
	if err != nil {
		return false, false, err
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repositoryArray[repoIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return false, false, err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCooordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return false, false, err
	}

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return false, false, err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return false, false, err
	}

	// Searching for the corresponding vote in the votes stored in the config
	var voteIndex int
	for i := range repositoryArray[repoIndex].ActiveVotes {
		if repositoryArray[repoIndex].ActiveVotes[i].ID == proposalID {
			voteIndex = i
			break
		}
	}

	// If this user already called the Resolve function for this vote it's not possible anymore
	if repositoryArray[repoIndex].ActiveVotes[voteIndex].Resolved {
		return false, false, errors.New("You already finalized this vote")
	}

	// Verifying whether the vote has already ended
	pollEnded, err := KNWVotingInstance.VoteEnded(nil, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve vote status")
	}

	// If not, we can't resolve it
	if !pollEnded {
		return false, false, errors.New("The vote hasn't ended yet")
	}

	// Verifying whether the user is a participant of this vote
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve commit status")
	}

	// Retrieve the selected proposal obkect
	proposal, err := ditCooordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve the new proposal")
	}

	// Indicates whether the called was the proposer or not
	isProposer := (proposal.Proposer == myAddress)

	// If not, we are not allowed to call this function (it would fail)
	if !didCommit && myAddress != proposal.Proposer {
		return false, false, errors.New("You didn't participate in this vote")
	}

	// Saving the old xDai balance
	oldxDaiBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return false, false, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
	}

	// Saving the old KNW balance
	oldKNWBalance, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, repositoryArray[repoIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	if err != nil {
		return false, false, errors.New("Failed to retrieve KNW balance")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return false, false, err
	}

	// Resolving the vote
	transaction, err := ditCooordinatorInstance.FinalizeVote(auth, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return false, false, errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return false, false, errors.New("Failed to finalize the vote: " + err.Error())
	}

	// Waiting for the resolve transaction to be mined
	helpers.Printf("Waiting for finalizing transaction to be mined", 0)
	waitingFor := 0
	newxDaiBalance := oldxDaiBalance
	for oldxDaiBalance.Cmp(newxDaiBalance) == 0 {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the balance of the user every 5 seconds, if it changed, a transaction was executed
		newxDaiBalance, err = connection.BalanceAt(context.Background(), myAddress, nil)
		if err != nil {
			return false, false, errors.New("Failed to retrieve opening status")
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

	// Saving the new KNW balance after resolving the vote
	newKNWBalance, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, repositoryArray[repoIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	if err != nil {
		return false, false, errors.New("Failed to retrieve KNW balance")
	}

	// Retrieving the outcome of the vote
	pollPassed, err := KNWVotingInstance.IsPassed(nil, big.NewInt(int64(repositoryArray[repoIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve vote outcome")
	}

	fmt.Printf("\n")
	// Show the user how the vote ended
	if pollPassed {
		helpers.PrintLine("Successfully finalized the vote - it passed", 0)
		if proposal.Proposer == common.HexToAddress(config.DitConfig.EthereumKeys.Address) {
			helpers.PrintLine("You received your stake back and will share the opposing voters slashes stakes", 0)
		}
	} else {
		helpers.PrintLine("Successfully finalized the vote - it didn't pass", 0)
		if proposal.Proposer == common.HexToAddress(config.DitConfig.EthereumKeys.Address) {
			helpers.PrintLine("Your stake was slashed and will be distributed amongst the voters", 0)
		}
	}

	// If the user got some xDai as a reward, this will be shown to the user
	if newxDaiBalance.Cmp(oldxDaiBalance) > 0 {
		difference := newxDaiBalance.Sub(newxDaiBalance, oldxDaiBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You gained %.2f "+config.DitConfig.Currency+"\n", floatDifference), 0)
	}

	// Also shwoing the user how man KNW tokens he got/lost
	if oldKNWBalance.Cmp(newKNWBalance) < 0 {
		difference := newKNWBalance.Sub(newKNWBalance, oldKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You earned %.2f KNW Tokens for the knowledge label '%s'", floatDifference, repositoryArray[repoIndex].ActiveVotes[voteIndex].KnowledgeLabel), 0)
	} else if oldKNWBalance.Cmp(newKNWBalance) > 0 {
		difference := oldKNWBalance.Sub(oldKNWBalance, newKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You lost %.2f KNW Tokens for the knowledge label '%s'", floatDifference, repositoryArray[repoIndex].ActiveVotes[voteIndex].KnowledgeLabel), 0)
	}

	// Saving the resolved-status in the config
	repositoryArray[repoIndex].ActiveVotes[voteIndex].Resolved = true

	if config.DitConfig.DemoModeActive {
		config.DitConfig.DemoRepositories = repositoryArray
	} else {
		config.DitConfig.LiveRepositories = repositoryArray
	}

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return false, false, err
	}

	return pollPassed, isProposer, nil
}

// GetVoteInfo will print information about a vote
func GetVoteInfo(_proposalID ...int) error {
	// Searching for this repositories object in the config
	repoIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repositoryArray[repoIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the ditContract to access it
	ditCooordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the current proposalID
	currentProposalIDBigInt, err := ditCooordinatorInstance.GetCurrentProposalID(nil, repoHash)
	if err != nil {
		return errors.New("Failed to retrieve the current proposal id")
	}

	// Converting the current proposal id of the ditContract to an int
	// if it is zero, there are no votes
	currentProposalID := int(currentProposalIDBigInt.Int64())
	if currentProposalID == 0 {
		helpers.PrintLine("There are no votes for this repository", 1)
		os.Exit(0)
	}

	// If a proposalID is specified in the function call, use that
	var proposalID int
	if len(_proposalID) > 0 {
		proposalID = _proposalID[0]
		// If not, let the user select on that is between 1 and the current proposal ID
	} else {
		for proposalID < 1 || proposalID > currentProposalID {
			proposalID, _ = strconv.Atoi(helpers.GetUserInput("Select a proposal ID between 1 and " + strconv.Itoa(currentProposalID)))
		}
	}

	// Retrieve the selected proposal obkect
	proposal, err := ditCooordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return errors.New("Failed to retrieve the new proposal")
	}

	// Retrieving information about this vote from the KNWVoting contract itself
	KNWPoll, err := KNWVotingInstance.Votes(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve vote information")
	}

	// Retrieving grossStake from this vote
	grossStake, err := KNWVotingInstance.GetGrossStake(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve the gross stake")
	}
	floatGrossStake := new(big.Float).Quo((new(big.Float).SetInt(grossStake)), big.NewFloat(1000000000000000000))

	// Retrieving netStake from this vote
	netStake, err := KNWVotingInstance.GetNetStake(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve the net stake")
	}
	floatNetStake := new(big.Float).Quo((new(big.Float).SetInt(netStake)), big.NewFloat(1000000000000000000))

	// Formatting the commit and reveal time to a human-readable format
	timeCommit := time.Unix(KNWPoll.CommitEndDate.Int64(), 0)
	timeCommitString := timeCommit.Format("15:04:05 on 2006/01/02")
	timeReveal := time.Unix(KNWPoll.OpenEndDate.Int64(), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	// Printing the information about this vote
	helpers.PrintLine("---------------------------", 0)
	helpers.PrintLine("Proposal ID: "+strconv.Itoa(proposalID), 0)
	helpers.PrintLine("URL: https://github.com/"+repositoryArray[repoIndex].Name+"/tree/dit_proposal_"+strconv.Itoa(proposalID), 0)
	helpers.PrintLine("Proposer: "+proposal.Proposer.Hex(), 0)
	helpers.PrintLine("Knowledge-Label: "+proposal.KnowledgeLabel, 0)

	// If the user participated in this vote, the choice and stake/KNW are also printed
	for i := range repositoryArray[repoIndex].ActiveVotes {
		if repositoryArray[repoIndex].ActiveVotes[i].ID == proposalID {
			intxDai, ok := new(big.Int).SetString(repositoryArray[repoIndex].ActiveVotes[i].NumTokens, 10)
			if !ok {
				return errors.New("Failed to parse big.int from string")
			}
			floatxDai := new(big.Float).Quo((new(big.Float).SetInt(intxDai)), big.NewFloat(1000000000000000000))
			intKNW, ok := new(big.Int).SetString(repositoryArray[repoIndex].ActiveVotes[i].NumKNW, 10)
			if !ok {
				return errors.New("Failed to parse big.int from string")
			}
			floatKNW := new(big.Float).Quo((new(big.Float).SetInt(intKNW)), big.NewFloat(1000000000000000000))
			if proposal.Proposer.Hex() != config.DitConfig.EthereumKeys.Address {
				helpers.PrintLine("Your choice: "+strconv.Itoa(repositoryArray[repoIndex].ActiveVotes[i].Choice), 0)
			}
			helpers.PrintLine(fmt.Sprintf("You staked %.2f %s and used %f KNW\n", floatxDai, config.DitConfig.Currency, floatKNW), 0)
			break
		} else {
			helpers.PrintLine(fmt.Sprintf("Required (gross) stake: %.2f "+config.DitConfig.Currency+"\n", floatGrossStake), 0)
		}
	}
	helpers.PrintLine(fmt.Sprintf("Current net stake: %.2f "+config.DitConfig.Currency+"\n", floatNetStake), 0)
	helpers.PrintLine("Vote phase end: "+timeCommitString, 0)
	helpers.PrintLine("Opening phase end: "+timeRevealString, 0)
	helpers.PrintLine("Resolved? "+strconv.FormatBool(proposal.IsFinalized), 0)
	helpers.PrintLine("Passed? "+strconv.FormatBool(proposal.ProposalAccepted), 0)
	helpers.PrintLine("---------------------------", 0)

	return nil
}

// GetBalances will print the xDai and KNW balances
func GetBalances() error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	if len(config.DitConfig.KNWToken) != correctETHAddressLength {
		return errors.New("Invalid KNWToken address, please do '" + helpers.ColorizeCommand("set_coordinator") + "' first")
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the number of knowledge-labels where the user has a balance in
	labelCount, err := KNWTokenInstance.LabelCountOfAddress(nil, myAddress)
	if err != nil {
		return errors.New("Failed to retrieve knowledge label count of address")
	}

	// Converting this count to an int
	labelCountInt := labelCount.Int64()

	// Retrieving each of the labels
	var labelsOfAddress []string
	for i := 1; i <= int(labelCountInt); i++ {
		newLabel, err := KNWTokenInstance.LabelOfAddress(nil, myAddress, big.NewInt(int64(i)))
		if err != nil {
			return errors.New("Failed to retrieve knowledge label count of address")
		}
		if len(newLabel) > 0 {
			labelsOfAddress = append(labelsOfAddress, newLabel)
		}
	}

	var floatBalance *big.Float
	if !config.DitConfig.DemoModeActive {
		// Retrieving the xDai balance of the user
		xDaiBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
		if err != nil {
			return errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
		}
		helpers.PrintLine("Balances for address "+config.DitConfig.EthereumKeys.Address+":", 0)

		// Formatting the xDai balance to a human-readable format
		floatBalance = new(big.Float).Quo((new(big.Float).SetInt(xDaiBalance)), big.NewFloat(1000000000000000000))
	} else {
		// Create a new instance of the ditToken to access it
		ditTokenInstance, err := getDitTokenInstance(connection)
		if err != nil {
			return err
		}

		// Retrieving the xDit balance of the user
		xDitBalance, err := ditTokenInstance.BalanceOf(nil, myAddress)
		if err != nil {
			return errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
		}
		helpers.PrintLine("Balances for address "+config.DitConfig.EthereumKeys.Address+":", 0)

		// Formatting the xDai balance to a human-readable format
		floatBalance = new(big.Float).Quo((new(big.Float).SetInt(xDitBalance)), big.NewFloat(1000000000000000000))
	}
	helpers.PrintLine(fmt.Sprintf(config.DitConfig.Currency+"-Balance: %.2f "+config.DitConfig.Currency, floatBalance), 0)
	helpers.PrintLine("", 0)

	// Printing each KNW balance
	helpers.PrintLine("KNW-Balances:", 0)
	if len(labelsOfAddress) == 0 {
		helpers.PrintLine("- No labels with any balance found", 0)
	} else {
		for i := 0; i < len(labelsOfAddress); i++ {
			// Retrieve each KNW balance
			balanceOfLabel, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, labelsOfAddress[i])
			if err != nil {
				return err
			}
			// Formatting each KNW balance to a human-readable format
			floatBalance := new(big.Float).Quo((new(big.Float).SetInt(balanceOfLabel)), big.NewFloat(1000000000000000000))
			helpers.PrintLine(fmt.Sprintf(" - "+labelsOfAddress[i]+": %.2f KNW", floatBalance), 0)
		}
	}

	return nil
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
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
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

// gatherProposalInfo will retrieve necessary information about a proposal from the ditContract
func gatherProposalInfo(_connection *ethclient.Client, _ditCoordinatorInstance *ditCoordinator.DitCoordinator, _repoHash [32]byte, _proposalID int64) (config.ActiveVote, error) {
	// Retrieve the proposal object from the ditContract
	var newVote config.ActiveVote
	proposal, err := _ditCoordinatorInstance.ProposalsOfRepository(nil, _repoHash, big.NewInt(_proposalID))
	if err != nil {
		return newVote, errors.New("Failed to retrieve the new proposal")
	}

	// Store the information about this vote in an object
	newVote.ID = int(_proposalID)
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

	// Retrieving the amount of xDai a user has staked for this vote
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

// TODO
func initDitRepository(_ditCoordinatorInstance *ditCoordinator.DitCoordinator, _repoHash [32]byte) error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	neededMajority := big.NewInt(50)

	// Prompting the user to provide 1 to 3 knowledge-labels for this repository
	helpers.PrintLine("Please provide knowledge labels that will be used for this repository:", 0)
	helpers.PrintLine("Note: At least one has to be provided, press ENTER to skip", 0)
	var knowledgeLabels []string
	for len(knowledgeLabels) < 3 {
		newLabel := helpers.GetUserInput("Knowledge Label " + strconv.Itoa(len(knowledgeLabels)+1))
		if len(newLabel) > 0 {
			knowledgeLabels = append(knowledgeLabels, newLabel)
		}
		if len(newLabel) == 0 && len(knowledgeLabels) > 0 {
			if len(knowledgeLabels) < 3 {
				for i := 0; i <= 3-len(knowledgeLabels); i++ {
					knowledgeLabels = append(knowledgeLabels, "")
				}
			}
			break
		}
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	// Initializing the repository = deploying a new ditContract
	transaction, err := _ditCoordinatorInstance.InitRepository(auth, _repoHash, knowledgeLabels[0], knowledgeLabels[1], knowledgeLabels[2], neededMajority)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDai to pay for the transaction")
		}
		return err
	}

	// Waiting for the deployment transaction to be mined
	helpers.Printf("Waiting for initialization transaction to be mined", 0)
	boolTrue := true
	isInitialized := false
	waitingFor := 0
	for isInitialized != boolTrue {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the repository status every 5 seconds
		isInitialized, err = _ditCoordinatorInstance.RepositoryIsInitialized(nil, _repoHash)
		if err != nil {
			return err
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

	return nil
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

	// Retrieving the suggested gasprice by the network
	gasPrice, err := _connection.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("Failed to retrieve the gas-price for ethereum transaction")
	}

	// Setting the values into the transaction-options object
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(0)
	// auth.GasLimit = uint64(1000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func getDitCoordinatorInstance(_connection *ethclient.Client, _ditCoordinatorAddress ...string) (*ditCoordinator.DitCoordinator, error) {
	var ditCoordinatorAddressString string

	if len(_ditCoordinatorAddress) > 0 {
		ditCoordinatorAddressString = _ditCoordinatorAddress[0]
	} else {
		if len(config.DitConfig.DitCoordinator) != correctETHAddressLength {
			return nil, errors.New("Invalid ditCoordinator address, please do '" + helpers.ColorizeCommand("update") + "' first")
		}
		ditCoordinatorAddressString = config.DitConfig.DitCoordinator
	}

	// Convertig the hex-string-formatted address into an address object
	ditCoordinatorAddress := common.HexToAddress(ditCoordinatorAddressString)

	// Create a new instance of the ditCoordinator to access it
	ditCoordinatorInstance, err := ditCoordinator.NewDitCoordinator(ditCoordinatorAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditCoordinator at provided address")
	}

	return ditCoordinatorInstance, nil
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

func getDitTokenInstance(_connection *ethclient.Client) (*ditToken.MintableERC20, error) {
	ditTokenAddress := common.HexToAddress(config.DitConfig.DitToken)

	// Create a new instance of the KNWToken contract to access it
	ditTokenInstance, err := ditToken.NewMintableERC20(ditTokenAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditToken at provided address")
	}

	return ditTokenInstance, nil
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
	connection, err := ethclient.Dial("https://sokol.poa.network")
	if err != nil {
		return nil, errors.New("Failed to connect to the ethereum network")
	}

	return connection, nil
}

// searchForRepoInConfig will search for the current repo in the config
func searchForRepoInConfig() (int64, error) {
	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return 0, err
	}

	var repositoryArray []config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryArray = config.DitConfig.DemoRepositories
	} else {
		repositoryArray = config.DitConfig.LiveRepositories
	}

	// Search for the repo in our config
	for i := range repositoryArray {
		if repositoryArray[i].Name == repository {
			// Return the index of this repository if it was found
			return int64(i), nil
		}
	}

	// Return an error if nothing was found
	return 0, errors.New("Repository hasn't been initialized")
}

// GetHashOfString takes a string a returns it as keccak256
func GetHashOfString(_string string) [32]byte {
	repoHash32 := [32]byte{}
	copy(repoHash32[:], crypto.Keccak256([]byte(_string))[:])

	return repoHash32
}
