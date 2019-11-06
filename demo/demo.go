package demo

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ditcraft/cli/config"
	"github.com/ditcraft/cli/ethereum"
	"github.com/ditcraft/cli/git"
	"github.com/ditcraft/cli/helpers"
	"github.com/ditcraft/cli/smartcontracts/KNWToken"
	"github.com/ditcraft/cli/smartcontracts/KNWVoting"
	"github.com/ditcraft/cli/smartcontracts/ditDemoCoordinator"
	"github.com/ditcraft/cli/smartcontracts/ditToken"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora"
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
		helpers.PrintLine(fmt.Sprintf("Please execute %s to update the ditCLI", aurora.Green("bash <(curl https://get.ditcraft.io -Ls)")), helpers.INFO)
		helpers.PrintLine("Or go to: https://github.com/ditcraft/cli", helpers.INFO)
		helpers.ExitAndLog(0)
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

	proposalIdentifier := _branch + ":" + _branchHeadHash

	// Prompting the user which knowledge-label he wants to use for this proposal
	answerKnowledgeLabel := 0
	userInputString = "Which Knowledge-Label suits these changes the most?"
	for i := range knowledgeLabels {
		userInputString += " (" + strconv.Itoa(i+1) + ") " + knowledgeLabels[i].Label
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
	freeKNWBalance, err := KNWTokenInstance.FreeBalanceOfID(nil, myAddress, big.NewInt(int64(knowledgeLabels[answerKnowledgeLabel-1].ID)))
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
		answerStake = strings.Replace(answerStake, ",", ".", -1)
		floatStakeParsed, _ = strconv.ParseFloat(answerStake, 64)
		floatStake = big.NewFloat(floatStakeParsed)
	}

	// Prompting the user how much KNW he wants to use for this proposal
	answerKNW := "0"
	floatKNWParsed, _ := strconv.ParseFloat(answerKNW, 64)
	floatKNW := big.NewFloat(floatKNWParsed)

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f KNW for the label '%s'", floatKNWBalance, knowledgeLabels[answerKnowledgeLabel-1].Label), helpers.INFO)
	if floatKNWBalance.Cmp(big.NewFloat(0)) == 1 {
		userInputString = fmt.Sprintf("How much do you want to use?")
		keepAsking := true
		for keepAsking {
			answerKNW = helpers.GetUserInput(userInputString)
			answerKNW = strings.Replace(answerKNW, ",", ".", -1)
			floatKNWParsed, _ = strconv.ParseFloat(answerKNW, 64)
			var ok bool
			floatKNW, ok = new(big.Float).SetString(answerKNW)
			if ok && floatKNW.Cmp(big.NewFloat(0)) >= 0 && floatKNW.Cmp(floatKNWBalance) <= 0 {
				keepAsking = false
			}
		}
	}

	// Prompting the user how long the vote should last
	answerLength := "0"
	floatLengthParsed, _ := strconv.ParseFloat(answerLength, 64)
	floatLength := big.NewFloat(floatLengthParsed)
	userInputString = fmt.Sprintf("How long should the vote last in minutes?")
	keepAsking := true
	for keepAsking {
		answerLength = helpers.GetUserInput(userInputString)
		answerLength = strings.Replace(answerLength, ",", ".", -1)
		floatLengthParsed, _ = strconv.ParseFloat(answerLength, 64)
		var ok bool
		floatLength, ok = new(big.Float).SetString(answerLength)
		if ok && floatLength.Cmp(big.NewFloat(1)) >= 0 && floatLength.Cmp(big.NewFloat(7*24*60)) <= 0 {
			keepAsking = false
		}
	}

	// Prompting the user whether he is sure of this proposal and its details
	floatStakeString := fmt.Sprintf("%.2f", floatStakeParsed)
	floatKNWString := fmt.Sprintf("%.2f", floatKNWParsed)
	floatLengthString := fmt.Sprintf("%.1f", floatLengthParsed)

	helpers.PrintLine("  Proposing the changes with the following settings:", helpers.INFO)
	helpers.PrintLine("  Proposal Description: '"+answerProposalSummary+"'", helpers.INFO)
	helpers.PrintLine("  Knowledge Label: "+knowledgeLabels[answerKnowledgeLabel-1].Label, helpers.INFO)
	helpers.PrintLine("  Amount of KNW Tokens: "+floatKNWString+" KNW", helpers.INFO)
	helpers.PrintLine("  Length of Vote: "+floatLengthString+" Minutes", helpers.INFO)
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

	secondFloatLength, _ := (new(big.Float).Mul(floatLength, big.NewFloat(60))).Int64()
	intLength := big.NewInt(secondFloatLength)

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
	transaction, err := ditCoordinatorInstance.ProposeCommit(auth, repoHash, answerProposalSummary, proposalIdentifier, big.NewInt(int64(knowledgeLabels[answerKnowledgeLabel-1].ID)), intKNW, intLength, intStake)
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
			helpers.ExitAndLog(0)
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
	responseString += "\nTo finalize the vote, execute '" + helpers.ColorizeCommand("finalize "+newProposalID.String()) + "' after " + timeRevealString
	responseString += "\n---------------------------"

	if config.DitConfig.DemoModeActive {
		fmt.Println()
		helpers.PrintLine("Since this is the demo mode, five auto-validators will automatically vote on your proposed changes. You will see the outcome after the vote ended.", 3)
		helpers.PrintLine("All validators now have time to vote on your proposed changes until "+timeCommitString, 3)
		fmt.Println()
	}
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

	var salt *big.Int
	if len(_salt) > 0 {
		var ok bool
		salt, ok = new(big.Int).SetString(_salt, 10)
		if !ok {
			return errors.New("Invalid salt")
		}
	} else {
		max := new(big.Int)
		max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

		var err error
		salt, err = rand.Int(rand.Reader, max)
		if err != nil {
			return errors.New("Error during salt generation: " + err.Error())
		}
	}

	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return err
	}

	repoHash := ethereum.GetHashOfString(repository)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the ditToken to access it
	ditTokenInstance, err := getDitTokenInstance(connection)
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
	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return errors.New("Failed to retrieve proposal")
	}

	// Verifiying whether the proposal is valid (if KNWVoteID is zero its not valid or non existent)
	if proposal.KNWVoteID.Int64() == 0 {
		return errors.New("Invalid proposalID - the vote doesn't exist")
	}

	if proposal.Proposer == myAddress {
		return errors.New("You can't vote on your own proposal")
	}

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

	// Retrieving the default stake from the ditContract
	requiredStake, err := KNWVotingInstance.GetGrossStake(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve the required stake of the vote")
	}

	floatStake := new(big.Float).Quo((new(big.Float).SetInt(requiredStake)), big.NewFloat(1000000000000000000))

	// Retrieving the xDIT balance of the user
	freeKNWBalance, err := KNWTokenInstance.FreeBalanceOfID(nil, myAddress, proposal.KnowledgeID)
	if err != nil {
		return errors.New("Failed to retrieve free KNW balance")
	}

	// Retrieving the knowledge label
	knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, proposal.KnowledgeID)
	if err != nil {
		return errors.New("Failed to retrieve knowledge label")
	}

	// Formatting the xDIT balance to a human-readable format
	floatKNWBalance := new(big.Float).Quo((new(big.Float).SetInt(freeKNWBalance)), big.NewFloat(1000000000000000000))

	// Prompting the user how much KNW he wants to use for this proposal
	answerKNW := "0"
	floatKNWParsed, _ := strconv.ParseFloat(answerKNW, 64)
	floatKNW := big.NewFloat(floatKNWParsed)

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f KNW for the label '%s'", floatKNWBalance, knowledgeLabel), helpers.INFO)
	if floatKNWBalance.Cmp(big.NewFloat(0)) == 1 {
		userInputString := fmt.Sprintf("How much do you want to use?")
		keepAsking := true
		for keepAsking {
			answerKNW = helpers.GetUserInput(userInputString)
			answerKNW = strings.Replace(answerKNW, ",", ".", -1)
			floatKNWParsed, _ = strconv.ParseFloat(answerKNW, 64)
			var ok bool
			floatKNW, ok = new(big.Float).SetString(answerKNW)
			if ok && floatKNW.Cmp(big.NewFloat(0)) >= 0 && floatKNW.Cmp(floatKNWBalance) <= 0 {
				keepAsking = false
			}
		}
	}

	// Prompting the user whether he is sure of this vote and its details
	floatKNWString := fmt.Sprintf("%.2f", floatKNWParsed)

	helpers.PrintLine("Voting on the proposal with the following settings:", helpers.INFO)
	helpers.PrintLine("Knowledge Label: "+knowledgeLabel, helpers.INFO)
	helpers.PrintLine("Amount of KNW Tokens: "+floatKNWString+" KNW", helpers.INFO)
	fmt.Println()
	helpers.PrintLine("Voting on this proposal will automatically deduct the required stake from you account.", helpers.INFO)
	helpers.PrintLine(fmt.Sprintf("Required stake: %.2f", floatStake), helpers.INFO)
	helpers.PrintLine("All participants of the vote will counter-stake the proposer.", helpers.INFO)
	helpers.PrintLine("You will receive the remaining stake back, no matter how the vote ends.", helpers.INFO)

	answer := helpers.GetUserInputChoice("Is this okay for you?", "y", "n")
	if answer == "n" {
		// If not: exit
		helpers.PrintLine("No vote executed due to users choice", 1)
		helpers.ExitAndLog(0)
	}

	fmt.Println("")

	approvedBalance, err := ditTokenInstance.Allowance(nil, myAddress, common.HexToAddress(config.DitConfig.DitCoordinator))
	if err != nil {
		return errors.New("Failed to retrieve xDIT allowance")
	}

	if approvedBalance.Cmp(requiredStake) == -1 {
		helpers.PrintLine("Since xDIT is a token, we need to set an allowance first", helpers.INFO)

		// Retrieving the xDIT balance of the user
		xDITBalance, err := ditTokenInstance.BalanceOf(nil, myAddress)
		if err != nil {
			return errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
		}

		// Crerating the transaction (basic values)
		auth, err := populateTx(connection)
		if err != nil {
			return err
		}

		_, err = ditTokenInstance.Approve(auth, common.HexToAddress(config.DitConfig.DitCoordinator), xDITBalance)
		if err != nil {
			return err
		}

		newAllowance := approvedBalance
		for newAllowance.Cmp(approvedBalance) == 0 {
			newAllowance, err = ditTokenInstance.Allowance(nil, myAddress, common.HexToAddress(config.DitConfig.DitCoordinator))
			if err != nil {
				return errors.New("Failed to retrieve xDIT allowance")
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println()
		helpers.PrintLine("Now the actual vote will be executed", helpers.INFO)
	}

	// In order to create a valid abi-encoded hash of the vote choice and salt
	// we need to create an abi object
	uint256Type, _ := abi.NewType("uint256", nil)
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
		salt,
	)

	// And finally hash this bytearray with keccak256, resulting in the votehash
	voteHash := crypto.Keccak256Hash(bytes)

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	weiFloatKNW, _ := (new(big.Float).Mul(floatKNW, big.NewFloat(1000000000000000000))).Int64()
	intKNW := big.NewInt(weiFloatKNW)

	// Voting on the proposal
	transaction, err := ditCoordinatorInstance.VoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), voteHash, intKNW)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDAI to pay for the transaction")
		}
		return errors.New("Failed to commit the vote: " + err.Error())
	}

	// Waiting for the voting transaction to be mined
	helpers.Printf("Waiting for voting transaction to be mined", helpers.INFO)
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
			helpers.ExitAndLog(0)
		}
	}
	fmt.Printf("\n")

	// Gathering the information of the proposal from the ditContract
	newVote, err := gatherProposalInfo(connection, ditCoordinatorInstance, repoHash, int64(proposalID))
	if err != nil {
		return err
	}

	// We will also store the users choice and salt, so that the user doesn't need to remember the salt
	// when he will reveal the vote later on
	newVote.Choice = choice
	newVote.Salt = salt.String()

	if config.DitConfig.DemoRepositories[repository] == nil {
		// Retrieving the knowledge-labels of this ditContract
		var knowledgeLabels []config.KnowledgeLabel
		contractKnowledgeIDs, err := ditCoordinatorInstance.GetKnowledgeIDs(nil, repoHash)
		if err != nil {
			return err
		}
		for i := 0; i < len(contractKnowledgeIDs); i++ {
			knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, contractKnowledgeIDs[i])
			if err != nil {
				return err
			}
			if len(knowledgeLabel) > 0 {
				knowledgeLabels = append(knowledgeLabels, config.KnowledgeLabel{ID: int(contractKnowledgeIDs[i].Int64()), Label: knowledgeLabel})
			}
		}

		// Inserting this repositories details into the cobfig
		var newRepository config.Repository
		if strings.Contains(repository, "github.com") {
			newRepository.Provider = "github"
		}
		newRepository.KnowledgeLabels = knowledgeLabels
		newRepository.ActiveVotes = make(map[string]*config.ActiveVote)
		config.DitConfig.DemoRepositories[repository] = &newRepository
	}

	// Adding the new vote to the config object
	config.DitConfig.DemoRepositories[repository].ActiveVotes[_proposalID] = &newVote

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
		helpers.PrintLine("You successfully voted in favor of the proposal", helpers.INFO)
	} else {
		helpers.PrintLine("You successfully voted against the proposal", helpers.INFO)
	}

	fmt.Println("")

	// Letting the user know when and how he has to reveal the vote
	helpers.PrintLine("With dit, votes are casted in a concealed manner through a commitment scheme.", helpers.INFO)
	helpers.PrintLine("This means that votes have to be revealed to the public once the commit-phase is over.", helpers.INFO)
	helpers.PrintLine("Please open your vote with '"+helpers.ColorizeCommand("open "+strconv.Itoa(int(proposalID)))+"' between "+timeCommitString+" and "+timeRevealString, helpers.INFO)

	return nil
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

	// Create a new instance of the KNWToken to access it
	ditCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
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

	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, getHashOfString(repository), big.NewInt(int64(voteID)))
	if err != nil {
		return voteID, true, false, repository, errors.New("Failed to retrieve proposal status")
	}

	if int(proposal.KNWVoteID.Int64()) != activeVote.KNWVoteID {
		return voteID, true, false, repository, errors.New("Unable to retrieve correct vote from smart contract")
	}

	splitIdentifier := strings.Split(proposal.Identifier, ":")
	if _branchHash != splitIdentifier[len(splitIdentifier)-1] {
		return voteID, true, false, repository, errors.New("Head hash of proposal and branch don't match")
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

	if !config.DitConfig.PassedKYC {
		config.DitConfig.PassedKYC = true

		err = config.Save()
		if err != nil {
			return false, err
		}
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

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := getKNWTokenInstance(_connection)
	if err != nil {
		return newVote, err
	}

	knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, proposal.KnowledgeID)
	if err != nil {
		return newVote, err
	}

	// Store the information about this vote in an object
	newVote.KNWVoteID = int(proposal.KNWVoteID.Int64())
	newVote.KnowledgeLabel.ID = int(proposal.KnowledgeID.Int64())
	newVote.KnowledgeLabel.Label = knowledgeLabel

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
