package ethereum

import (
	"context"
	"dit/config"
	"dit/git"
	"dit/helpers"
	"dit/smartcontracts/KNWToken"
	"dit/smartcontracts/KNWVoting"
	"dit/smartcontracts/ditContract"
	"dit/smartcontracts/ditCoordinator"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

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
	ditCoordinatorInstance, err := getditCoordinatorInstance(connection, _ditCoordinatorAddressString)
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

	// Setting the retrieved addresses to the config object
	config.DitConfig.DitCoordinator = _ditCoordinatorAddressString
	config.DitConfig.KNWVoting = KNWVotingAddress.Hex()
	config.DitConfig.KNWToken = KNWTokenAddress.Hex()

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

	// Searching through the repositories that are in our config
	for i := 0; i < len(config.DitConfig.Repositories); i++ {
		// If the repository was found in our config it has already been initialized
		// In that case we don't need to proceed further
		if config.DitConfig.Repositories[i].Name == repository {
			if len(_optionalRepository) > 0 {
				return nil
			}
			fmt.Println("Repository is already initialized")
			os.Exit(0)
		}
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the ditCoordinator to access it
	ditCoordinatorInstance, err := getditCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the address of the ditContract that was deployed for this repository
	ditContractAddress, err := ditCoordinatorInstance.GetRepository(nil, repository)
	if err != nil {
		return err
	}

	// If the address is zero, there is no contract deployed
	if ditContractAddress == common.HexToAddress("0") {
		// Prompt the user whether he wants to deploy one
		answer := helpers.GetUserInputChoice("There is no ditContract for this repository yet, do you want to deploy one?", "y", "n")
		if answer == "y" {
			// If yes: deploy the ditContract
			ditContractAddress, err = deployDitContract(ditCoordinatorInstance, repository)
			if err != nil {
				return err
			}
		} else {
			// If not: exit
			fmt.Println("No ditContract deployed - repository can't be used with dit")
			os.Exit(0)
		}
	}

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, 0, ditContractAddress.Hex())
	if err != nil {
		return err
	}

	// Retrieving the knowledge-labels of this ditContract
	var knowledgeLabels []string
	for i := 0; i < 3; i++ {
		contractKnowledgeLabels, err := ditContractInstance.KnowledgeLabels(nil, big.NewInt(int64(i)))
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
	newRepository.ContractAddress = ditContractAddress.Hex()
	newRepository.KnowledgeLabels = knowledgeLabels
	newRepository.ActiveVotes = make([]config.ActiveVote, 0)
	config.DitConfig.Repositories = append(config.DitConfig.Repositories, newRepository)

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return nil
	}

	return nil
}

// ProposeCommit will start a new proposal on the ditContract of this repository
func ProposeCommit(_commitMessage string) (string, int, error) {
	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return "", 0, err
	}

	// Gathering the the knowledge-labels from the config
	knowledgeLabels := config.DitConfig.Repositories[ditContractIndex].KnowledgeLabels

	connection, err := getConnection()
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, ditContractIndex)
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

	// Retrieving the ETH balance of the user
	ethBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve ETH balance")
	}

	// Formatting the ETH balance to a human-readable format
	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(ethBalance)), big.NewFloat(1000000000000000000))

	// Prompting the user how much stake he wants to set for this proposal
	answerStake := "0"
	floatStakeParsed, _ := strconv.ParseFloat(answerStake, 64)
	floatStake := big.NewFloat(floatStakeParsed)

	userInputString = fmt.Sprintf("You have a balance of %f ETH\nHow much do you want to stake?", floatBalance)
	for floatStake.Cmp(big.NewFloat(0)) == 0 || floatStake.Cmp(floatBalance) != -1 {
		answerStake = helpers.GetUserInput(userInputString)
		floatStakeParsed, _ = strconv.ParseFloat(answerStake, 64)
		floatStake = big.NewFloat(floatStakeParsed)
	}

	// Prompting the user whether he is sure of this proposal and its details
	floatStakeString := fmt.Sprintf("%f", floatStakeParsed)
	userIsSure := helpers.GetUserInputChoice("Proposing the commit with the following settings: \n"+
		" Commit Message: "+_commitMessage+"\n"+
		" Knowledge Label: "+knowledgeLabels[answerKnowledgeLabel-1]+"\n"+
		" The following stake with automatically be deducted: "+floatStakeString+" ETH\n"+
		"Is that correct?", "y", "n")
	if userIsSure == "n" {
		return "", 0, errors.New("Canceled proposal of commit due to users choice")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return "", 0, err
	}

	// Setting the value of the transaction to be the selected stake
	weiFloatStake, _ := (new(big.Float).Mul(floatStake, big.NewFloat(1000000000000000000))).Int64()
	intStake := big.NewInt(weiFloatStake)
	auth.Value = intStake

	// Retrieving the last/current proposalID of the ditContract
	// (This will increment after a proposal, so we can see when the proposal is live)
	lastProposalID, err := ditContractInstance.CurrentProposalID(nil)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve the current proposal id")
	}

	// Proposing the commit
	transaction, err := ditContractInstance.ProposeCommit(auth, big.NewInt(int64(answerKnowledgeLabel-1)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return "", 0, errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return "", 0, err
	}

	// Waiting for the proposals transaction to be mined
	fmt.Printf("Waiting for commit proposal transaction to be mined")
	newProposalID := lastProposalID
	waitingFor := 0
	for newProposalID.Cmp(lastProposalID) == 0 {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the current proposal ID every 5 seconds
		newProposalID, err = ditContractInstance.CurrentProposalID(nil)
		if err != nil {
			return "", 0, errors.New("Failed to retrieve the current proposal id")
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 120 {
			fmt.Println("\nWaiting for over 2 minutes, maybe the transaction or the network failed?")
			fmt.Println("Check at: https://rinkeby.etherscan.io/tx/" + transaction.Hash().Hex())
			os.Exit(0)
		}
	}
	fmt.Printf("\n")

	// Gathering the information of the new proposal from the ditContract (including the times to commit and reveal)
	newVote, err := gatherProposalInfo(connection, ditContractInstance, ditContractIndex, newProposalID.Int64())
	if err != nil {
		return "", 0, err
	}

	// Adding the new vote to the config object
	config.DitConfig.Repositories[ditContractIndex].ActiveVotes = append(config.DitConfig.Repositories[ditContractIndex].ActiveVotes, newVote)

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return "", 0, err
	}

	// Conwerting the stake and used KNW count into a float so that it's human-readable
	floatETH := new(big.Float).Quo((new(big.Float).SetInt(big.NewInt(int64(newVote.NumTokens)))), big.NewFloat(1000000000000000000))
	floatKNW := new(big.Float).Quo((new(big.Float).SetInt(big.NewInt(int64(newVote.NumKNW)))), big.NewFloat(1000000000000000000))

	// Formatting the time of the commit and reveal phase into a readable format
	timeReveal := time.Unix(int64(newVote.RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	// Returning the response to the user, we don't want to print this right here and now
	// since we are not sure whether the actual git commands will succeed.
	// So it will be printed afterwards in the main routine
	var responseString string
	responseString += "---------------------------"
	responseString += "\nSuccessfully proposed commit. Vote on proposal started with ID " + strconv.Itoa(int(newVote.ID)) + ""
	responseString += fmt.Sprintf("\nYou staked %f ETH and used %f KNW.", floatETH, floatKNW)
	responseString += "\nThe vote will end at " + timeRevealString
	responseString += "\nYour commit is at https://github.com/" + config.DitConfig.Repositories[ditContractIndex].Name + "/tree/dit_proposal_" + strconv.Itoa(int(newVote.ID))
	responseString += "\n---------------------------"

	return responseString, int(newProposalID.Int64()), nil
}

// Vote will cast a users vote on a proposal
func Vote(_proposalID string, _choice string, _salt string) error {
	// Converting the stdin string input of the user into Ints
	proposalID, _ := strconv.Atoi(_proposalID)
	choice, _ := strconv.Atoi(_choice)
	salt, _ := strconv.Atoi(_salt)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, ditContractIndex)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the proposal object from the ditContract
	proposal, err := ditContractInstance.Proposals(nil, big.NewInt(int64(proposalID)))
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

	fmt.Println("Voting on this proposal will automatically deduct the required stake from you account.")
	fmt.Printf("Required stake: %f\n", floatStake)
	fmt.Println("All participants of the vote will counter-stake the proposer.")
	fmt.Println("You will receive the remaining stake back, no matter how the vote ends.")

	answer := helpers.GetUserInputChoice("Is this okay for you?", "y", "n")
	if answer == "n" {
		// If not: exit
		fmt.Println("No vote executed due to users choice")
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
		return errors.New("Failed to retrieve reveal status")
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

	// Voting on the proposal
	transaction, err := ditContractInstance.VoteOnProposal(auth, big.NewInt(int64(proposalID)), voteHash)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return errors.New("Failed to commit the vote: " + err.Error())
	}

	// Waiting for the voting transaction to be mined
	fmt.Printf("Waiting for voting transaction to be mined")
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
		if waitingFor > 120 {
			fmt.Println("\nWaiting for over 2 minutes, maybe the transaction or the network failed?")
			fmt.Println("Check at: https://rinkeby.etherscan.io/tx/" + transaction.Hash().Hex())
			os.Exit(0)
		}
	}
	fmt.Printf("\n")

	// Gathering the information of the proposal from the ditContract
	newVote, err := gatherProposalInfo(connection, ditContractInstance, ditContractIndex, int64(proposalID))
	if err != nil {
		return err
	}

	// We will also store the users choice and salt, so that the user doesn't need to remember the salt
	// when he will reveal the vote later on
	newVote.Choice = choice
	newVote.Salt = salt

	// Adding the new vote to the config object
	config.DitConfig.Repositories[ditContractIndex].ActiveVotes = append(config.DitConfig.Repositories[ditContractIndex].ActiveVotes, newVote)

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return err
	}

	if choice == 1 {
		fmt.Println("You successfully voted in favor of the proposal")
	} else {
		fmt.Println("You successfully voted against the proposal")
	}

	// Formatting the time of the commit and reveal phase into a readable format
	timeCommit := time.Unix(int64(newVote.CommitEnd), 0)
	timeCommitString := timeCommit.Format("15:04:05 on 2006/01/02")
	timeReveal := time.Unix(int64(newVote.RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	// Letting the user know when and how he has to reveal the vote
	fmt.Println("Please reveal your vote with 'dit reveal_vote " + strconv.Itoa(int(proposalID)) + "' between " + timeCommitString + " and " + timeRevealString)

	return nil
}

// Reveal will reveal a vote (that was commited through a hash) to the public during the reveal phase
func Reveal(_proposalID string) error {
	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, ditContractIndex)
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
	for i := range config.DitConfig.Repositories[ditContractIndex].ActiveVotes {
		if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].ID == proposalID {
			voteIndex = i
			break
		}
	}

	// Verifying whether the reveal period of this vote is active
	revealPeriodActive, err := KNWVotingInstance.RevealPeriodActive(nil, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve reveal status")
	}

	// If it is now active it hasn't started yet or it's over
	if !revealPeriodActive {
		return errors.New("The reveal phase of this vote is not active")
	}

	// Verifying whether the user has commited a vote on this proposal
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve commit status")
	}

	// If this is not the case the user never participated in this proposal through a vote
	if !didCommit {
		return errors.New("You didn't vote on this proposal")
	}

	// Verifying whether the user has revealed his vote on this proposal
	oldDidReveal, err := KNWVotingInstance.DidReveal(nil, myAddress, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve reveal status")
	}

	// If this is the case, the user already revealed his vote
	if oldDidReveal {
		return errors.New("You already revealed your vote on this proposal")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	// Gathering the original choice and the salt from the config
	choice := big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].Choice))
	salt := big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].Salt))

	// Revealing the vote on the proposal
	transaction, err := ditContractInstance.RevealVoteOnProposal(auth, big.NewInt(int64(proposalID)), choice, salt)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return errors.New("Failed to reveal the vote: " + err.Error())
	}

	// Waiting for the reveal transaction to be mined
	fmt.Printf("Waiting for reveal transaction to be mined")
	waitingFor := 0
	newDidReveal := oldDidReveal
	for newDidReveal == oldDidReveal {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the reveal status of the user every 5 seconds
		newDidReveal, err = KNWVotingInstance.DidReveal(nil, myAddress, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
		if err != nil {
			return errors.New("Failed to retrieve reveal status")
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 120 {
			fmt.Println("\nWaiting for over 2 minutes, maybe the transaction or the network failed?")
			fmt.Println("Check at: https://rinkeby.etherscan.io/tx/" + transaction.Hash().Hex())
			os.Exit(0)
		}
	}
	fmt.Println("\nSuccessfully revealed your vote")

	return nil
}

// Resolve will finalize a vote as it will trigger the calculation of the reward of this user
// including the ETH and KNW reward in case of a voting for the winning decision
// or the losing of ETH and KNW in case of a voting for the losing decision
// The first caller who executes this will also trigger the calculation whether the vote passed or not
func Resolve(_proposalID string) (bool, error) {
	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return false, err
	}

	connection, err := getConnection()
	if err != nil {
		return false, err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, ditContractIndex)
	if err != nil {
		return false, err
	}

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return false, err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return false, err
	}

	// Searching for the corresponding vote in the votes stored in the config
	var voteIndex int
	for i := range config.DitConfig.Repositories[ditContractIndex].ActiveVotes {
		if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].ID == proposalID {
			voteIndex = i
			break
		}
	}

	// If this user already called the Resolve function for this vote it's not possible anymore
	if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].Resolved {
		return false, errors.New("You already resolved this vote")
	}

	// Verifying whether the vote has already ended
	pollEnded, err := KNWVotingInstance.PollEnded(nil, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, errors.New("Failed to retrieve vote status")
	}

	// If not, we can't resolve it
	if !pollEnded {
		return false, errors.New("The vote hasn't ended yet")
	}

	// Verifying whether the user is a participant of this vote
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, errors.New("Failed to retrieve commit status")
	}

	// Retrieve the selected proposal obkect
	proposal, err := ditContractInstance.Proposals(nil, big.NewInt(int64(proposalID)))
	if err != nil {
		return false, errors.New("Failed to retrieve the new proposal")
	}

	// If not, we are not allowed to call this function (it would fail)
	if !didCommit && myAddress != proposal.Proposer {
		return false, errors.New("You didn't participate in this vote")
	}

	// Saving the old ETH balance
	oldEthBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return false, errors.New("Failed to retrieve ETH balance")
	}

	// Saving the old KNW balance
	oldKNWBalance, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	if err != nil {
		return false, errors.New("Failed to retrieve KNW balance")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return false, err
	}

	// Resolving the vote
	transaction, err := ditContractInstance.ResolveVote(auth, big.NewInt(int64(proposalID)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return false, errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return false, errors.New("Failed to resolve the vote: " + err.Error())
	}

	// Waiting for the resolve transaction to be mined
	fmt.Printf("Waiting for resolve transaction to be mined")
	waitingFor := 0
	newEthBalance := oldEthBalance
	for oldEthBalance.Cmp(newEthBalance) == 0 {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the balance of the user every 5 seconds, if it changed, a transaction was executed
		newEthBalance, err = connection.BalanceAt(context.Background(), myAddress, nil)
		if err != nil {
			return false, errors.New("Failed to retrieve reveal status")
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 120 {
			fmt.Println("\nWaiting for over 2 minutes, maybe the transaction or the network failed?")
			fmt.Println("Check at: https://rinkeby.etherscan.io/tx/" + transaction.Hash().Hex())
			os.Exit(0)
		}
	}

	// Saving the new KNW balance after resolving the vote
	newKNWBalance, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	if err != nil {
		return false, errors.New("Failed to retrieve KNW balance")
	}

	// Retrieving the outcome of the vote
	pollPassed, err := KNWVotingInstance.IsPassed(nil, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return false, errors.New("Failed to retrieve vote outcome")
	}

	// Show the user how the vote ended
	if pollPassed {
		fmt.Println("\nSuccessfully resolve the vote - it passed")
	} else {
		fmt.Println("\nSuccessfully resolve the vote - it didn't pass")
	}

	// If the user got some ETH as a reward, this will be shown to the user
	if newEthBalance.Cmp(oldEthBalance) > 0 {
		difference := newEthBalance.Sub(newEthBalance, oldEthBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		fmt.Printf("You gained %f ETH\n", floatDifference)
	}

	// Also shwoing the user how man KNW tokens he got/lost
	if oldKNWBalance.Cmp(newKNWBalance) < 0 {
		difference := newKNWBalance.Sub(newKNWBalance, oldKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		fmt.Printf("You earned %f KNW Tokens for the knowledge label '%s'\n", floatDifference, config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	} else if oldKNWBalance.Cmp(newKNWBalance) > 0 {
		difference := oldKNWBalance.Sub(oldKNWBalance, newKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		fmt.Printf("You lost %f KNW Tokens for the knowledge label '%s'\n", floatDifference, config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KnowledgeLabel)
	}

	// Saving the resolved-status in the config
	config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].Resolved = true

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return false, nil
	}

	return pollPassed, nil
}

// GetVoteInfo will print information about a vote
func GetVoteInfo(_proposalID ...int) error {
	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := getDitContractInstance(connection, ditContractIndex)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the current proposalID
	currentProposalIDBigInt, err := ditContractInstance.CurrentProposalID(nil)
	if err != nil {
		return errors.New("Failed to retrieve the current proposal id")
	}

	// Converting the current proposal id of the ditContract to an int
	// if it is zero, there are no votes
	currentProposalID := int(currentProposalIDBigInt.Int64())
	if currentProposalID == 0 {
		fmt.Println("There are no votes for this repository")
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
	proposal, err := ditContractInstance.Proposals(nil, big.NewInt(int64(proposalID)))
	if err != nil {
		return errors.New("Failed to retrieve the new proposal")
	}

	// Retrieving information about this vote from the KNWVoting contract itself
	KNWPoll, err := KNWVotingInstance.PollMap(nil, proposal.KNWVoteID)
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
	timeReveal := time.Unix(KNWPoll.RevealEndDate.Int64(), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	// Printing the information about this vote
	fmt.Println("---------------------------")
	fmt.Println("Proposal ID: " + strconv.Itoa(proposalID))
	fmt.Println("URL: https://github.com/" + config.DitConfig.Repositories[ditContractIndex].Name + "/tree/dit_proposal_" + strconv.Itoa(proposalID))
	fmt.Println("Proposer: " + proposal.Proposer.Hex())
	fmt.Println("Knowledge-Label: " + proposal.KnowledgeLabel)

	// If the user participated in this vote, the choice and stake/KNW are also printed
	for i := range config.DitConfig.Repositories[ditContractIndex].ActiveVotes {
		if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].ID == proposalID {
			floatETH := new(big.Float).Quo((new(big.Float).SetInt(big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].NumTokens)))), big.NewFloat(1000000000000000000))
			floatKNW := new(big.Float).Quo((new(big.Float).SetInt(big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].NumKNW)))), big.NewFloat(1000000000000000000))
			fmt.Println("Your choice: " + strconv.Itoa(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].Choice))
			fmt.Printf("You staked %f ETH and used %f KNW\n", floatETH, floatKNW)
			break
		} else {
			fmt.Printf("Required (gross) stake: %f ETH\n", floatGrossStake)
		}
	}
	fmt.Printf("Current net stake: %f ETH\n", floatNetStake)
	fmt.Println("Commit phase end: " + timeCommitString)
	fmt.Println("Reveal phase end: " + timeRevealString)
	fmt.Println("Resolved? " + strconv.FormatBool(proposal.IsResolved))
	fmt.Println("Passed? " + strconv.FormatBool(proposal.ProposalAccepted))
	fmt.Println("---------------------------")

	return nil
}

// GetBalances will print the ETH and KNW balances
func GetBalances() error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	if len(config.DitConfig.KNWToken) != correctETHAddressLength {
		return errors.New("Invalid KNWToken address, please do 'dit set_coordinator' first")
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

	// Retrieving the ETH balance of the user
	ethBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return errors.New("Failed to retrieve ETH balance")
	}
	fmt.Println("Balances for address " + config.DitConfig.EthereumKeys.Address + ":")

	// Formatting the ETH balance to a human-readable format
	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(ethBalance)), big.NewFloat(1000000000000000000))
	fmt.Printf("ETH-Balance: %f ETH\n", floatBalance)

	// Printing each KNW balance
	fmt.Println("KNW-Balances:")
	if len(labelsOfAddress) == 0 {
		fmt.Println("- No labels with any balance found")
	} else {
		for i := 0; i < len(labelsOfAddress); i++ {
			// Retrieve each KNW balance
			balanceOfLabel, err := KNWTokenInstance.BalanceOfLabel(nil, myAddress, labelsOfAddress[i])
			if err != nil {
				return err
			}
			// Formatting each KNW balance to a human-readable format
			floatBalance := new(big.Float).Quo((new(big.Float).SetInt(balanceOfLabel)), big.NewFloat(1000000000000000000))
			fmt.Printf(" - "+labelsOfAddress[i]+": %f KNW\n", floatBalance)
		}
	}

	return nil
}

// gatherProposalInfo will retrieve necessary information about a proposal from the ditContract
func gatherProposalInfo(_connection *ethclient.Client, _ditContractInstance *ditContract.DitContract, _repositoryIndex int64, _proposalID int64) (config.ActiveVote, error) {
	// Retrieve the proposal object from the ditContract
	var newVote config.ActiveVote
	proposal, err := _ditContractInstance.Proposals(nil, big.NewInt(_proposalID))
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
	KNWPoll, err := KNWVotingInstance.PollMap(nil, big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve vote information")
	}

	newVote.CommitEnd = int(KNWPoll.CommitEndDate.Int64())
	newVote.RevealEnd = int(KNWPoll.RevealEndDate.Int64())

	// Retrieving the amount of ETH a user has staked for this vote
	numTokens, err := KNWVotingInstance.GetGrossStake(nil, big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve grossStake")
	}

	// Retrieving the number of votes a user has for this vote
	numVotes, err := KNWVotingInstance.GetNumVotes(nil, common.HexToAddress(config.DitConfig.EthereumKeys.Address), big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve numVotes")
	}

	// Retrieving the number of KNW tokens a user has staked for this vote
	numKNW, err := KNWVotingInstance.GetNumKNW(nil, common.HexToAddress(config.DitConfig.EthereumKeys.Address), big.NewInt(int64(newVote.KNWVoteID)))
	if err != nil {
		return newVote, errors.New("Failed to retrieve numKNW")
	}

	newVote.NumTokens = int(numTokens.Int64())
	newVote.NumVotes = int(numVotes.Int64())
	newVote.NumKNW = int(numKNW.Int64())

	return newVote, nil
}

// deployDitContract will deploy a new ditContract for a repository through the ditCoordinator
func deployDitContract(_ditCoordinatorInstance *ditCoordinator.DitCoordinator, _repository string) (common.Address, error) {
	connection, err := getConnection()
	if err != nil {
		return common.HexToAddress("0"), err
	}

	// Setting the voting setting for this repository
	var voteSettings [3]*big.Int

	// Majority (in percent)
	voteSettings[0] = big.NewInt(50)

	// KNW minting method (0 = regular)
	voteSettings[1] = big.NewInt(0)

	// KNW burning method (0 = square-root based, 1 = divide by 2 each time, 2 = proportial to the winning percentage)
	voteSettings[2] = big.NewInt(0)

	// Prompting the user to provide 1 to 3 knowledge-labels for this repository
	fmt.Println("Please provide knowledge labels that will be used for this repository:")
	fmt.Println("Note: At least one has to be provided, press ENTER to skip")
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
		return common.HexToAddress("0"), err
	}

	// Initializing the repository = deploying a new ditContract
	transaction, err := _ditCoordinatorInstance.InitRepository(auth, _repository, knowledgeLabels[0], knowledgeLabels[1], knowledgeLabels[2], voteSettings)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return common.HexToAddress("0"), errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return common.HexToAddress("0"), err
	}

	// Waiting for the deployment transaction to be mined
	fmt.Printf("Waiting for deployment transaction to be mined")
	ditContractAddress := common.HexToAddress("0")
	waitingFor := 0
	for ditContractAddress == common.HexToAddress("0") {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the repository status every 5 seconds
		ditContractAddress, err = _ditCoordinatorInstance.GetRepository(nil, _repository)
		if err != nil {
			return common.HexToAddress("0"), err
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 120 {
			fmt.Println("\nWaiting for over 2 minutes, maybe the transaction or the network failed?")
			fmt.Println("Check at: https://rinkeby.etherscan.io/tx/" + transaction.Hash().Hex())
			os.Exit(0)
		}
	}
	fmt.Printf("\n")

	return ditContractAddress, nil
}

// populateTX will set the necessary values for a ethereum transaction
// amount of gas, gasprice, nonce, sign this with the private key
func populateTx(_connection *ethclient.Client) (*bind.TransactOpts, error) {
	// Retrieve the decrypted private key through a password prompt
	privateKeyString, err := config.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	// Converting the private key string into a private key object
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, errors.New("Failed to convert ethereum private-key")
	}

	// Retrieving the current nonce of our address
	nonce, err := _connection.PendingNonceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address))
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
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
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func getDitContractInstance(_connection *ethclient.Client, _ditContractIndex int64, _ditContractAddress ...string) (*ditContract.DitContract, error) {
	var ditContractAddressString string

	if len(_ditContractAddress) > 0 {
		ditContractAddressString = _ditContractAddress[0]
	} else {
		// Gathering the ditContracts address and the knowledge-labels from the config
		ditContractAddressString = config.DitConfig.Repositories[_ditContractIndex].ContractAddress
	}

	// Convertig the hex-string-formatted address into an address object
	ditContractAddress := common.HexToAddress(ditContractAddressString)

	// Create a new instance of the ditContract to access it
	ditContractInstance, err := ditContract.NewDitContract(ditContractAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditCoordinator at provided address")
	}
	return ditContractInstance, nil
}

func getditCoordinatorInstance(_connection *ethclient.Client, _ditCoordinatorAddress ...string) (*ditCoordinator.DitCoordinator, error) {
	var ditCoordinatorAddressString string

	if len(_ditCoordinatorAddress) > 0 {
		ditCoordinatorAddressString = _ditCoordinatorAddress[0]
	} else {
		if len(config.DitConfig.DitCoordinator) != correctETHAddressLength {
			return nil, errors.New("Invalid ditCoordinator address, please do 'dit set_coordinator' first")
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
	// Connecting to rinkeby via infura
	connection, err := ethclient.Dial("https://rinkeby.infura.io/<INFURA-KEY-HERE>")
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

	// Search for the repo in our config
	for i := range config.DitConfig.Repositories {
		if config.DitConfig.Repositories[i].Name == repository {
			// Return the index of this repository if it was found
			return int64(i), nil
		}
	}

	// Return an error if nothing was found
	return 0, errors.New("Repository hasn't been initialized")
}
