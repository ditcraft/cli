package ethereum

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
	"github.com/ditcraft/cli/git"
	"github.com/ditcraft/cli/helpers"
	"github.com/ditcraft/cli/smartcontracts/KNWToken"
	"github.com/ditcraft/cli/smartcontracts/KNWVoting"
	KyberNetwork "github.com/ditcraft/cli/smartcontracts/KyberNetworkProxy"
	"github.com/ditcraft/cli/smartcontracts/ditCoordinator"
	"github.com/ditcraft/cli/smartcontracts/ditDemoCoordinator"
	"github.com/ditcraft/cli/smartcontracts/ditToken"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora"
	"github.com/schollz/closestmatch"
)

const correctETHAddressLength = 42

// // GetDaiPrice will return the price in dai
// func GetDaiPrice() error {
// 	connection, err := getConnectionMainnet()
// 	if err != nil {
// 		return err
// 	}

// 	KyberInstance, err := getKyberInstance(connection)

// 	priceStruct, err := KyberInstance.GetExpectedRate(nil, common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"), big.NewInt(1*10^18))
// 	if err != nil {
// 		return err
// 	}

// 	floatRate := new(big.Float).Quo((new(big.Float).SetInt(priceStruct.ExpectedRate)), big.NewFloat(1000000000000000000))

// 	fmt.Println(fmt.Sprintf("Current exchange rate for 1 ETH: %.2f DAI", floatRate))
// 	fmt.Println("These rates are directly sourced from the KyberNetwork smart contracts.")
// 	fmt.Println("For more details visit: https://kyberswap.com")
// 	return nil
// }

// // SwapETHtoXDai will swap all ETH on mainnet to xdai on the xdai chain
// func SwapETHtoXDai() error {
// 	connectionMainnet, err := getConnectionMainnet()
// 	if err != nil {
// 		return err
// 	}

// 	connectionXDai, err := getConnection()
// 	if err != nil {
// 		return err
// 	}

// 	balanceMainnet, err := connectionMainnet.BalanceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address), nil)
// 	if err != nil {
// 		return err
// 	}

// 	KyberInstance, err := getKyberInstance(connectionMainnet)

// 	priceStruct, err := KyberInstance.GetExpectedRate(nil, common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"), big.NewInt(1*10^18))
// 	if err != nil {
// 		return err
// 	}

// 	minRate := new(big.Int).Div(new(big.Int).Mul(priceStruct.ExpectedRate, big.NewInt(97)), big.NewInt(100))

// 	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(balanceMainnet)), big.NewFloat(1000000000000000000))
// 	floatRate := new(big.Float).Quo((new(big.Float).SetInt(priceStruct.ExpectedRate)), big.NewFloat(1000000000000000000))

// 	floatValueDai, _ := new(big.Float).Mul(floatBalance, floatRate).Float64()
// 	if floatValueDai < 0.01 {
// 		return errors.New("The amount of ETH you have on the Ethereum Mainnet has to equivalted to a value greater than 0.01 DAI")
// 	}

// 	response, err := http.Get("https://ethgasstation.info/json/ethgasAPI.json")
// 	if err != nil {
// 		return err
// 	}

// 	data, _ := ioutil.ReadAll(response.Body)

// 	position := strings.Index(string(data), "average")
// 	if position == -1 {
// 		return errors.New("Failed to retrieve the current gas price on Mainnet")
// 	}
// 	gasPricePart := string(data)[position+10:]
// 	gasPriceString := strings.Split(gasPricePart, ",")[0]
// 	gasPrice, err := strconv.ParseFloat(gasPriceString, 64)
// 	if err != nil {
// 		return err
// 	}
// 	gasPrice = (gasPrice * 1.1) / 10

// 	helpers.PrintLine("Starting the deposit of xDai from ETH onto your account:", helpers.INFO)
// 	helpers.PrintLine("---------------------------", helpers.INFO)
// 	helpers.PrintLine(fmt.Sprintf("The current exchange rate is 1 ETH = %.2f DAI", floatRate), helpers.INFO)
// 	helpers.PrintLine(fmt.Sprintf("You will be converting %.2f ETH to ~%.2f xDAI", floatBalance, floatValueDai), helpers.INFO)
// 	helpers.PrintLine(fmt.Sprintf("The current gas price is %.2f GWei", gasPrice), helpers.INFO)
// 	helpers.PrintLine("---------------------------", helpers.INFO)

// 	daiInstance, err := getDAIInstance(connectionMainnet)
// 	if err != nil {
// 		return err
// 	}

// 	daiBalanceOld, err := daiInstance.BalanceOf(nil, common.HexToAddress(config.DitConfig.EthereumKeys.Address))

// 	privateKeyString, err := config.GetPrivateKey(true)
// 	if err != nil {
// 		return err
// 	}

// 	// Crerating the transaction (basic values)
// 	auth, err := populateTx(connectionMainnet, privateKeyString)
// 	if err != nil {
// 		return err
// 	}

// 	auth.GasPrice = new(big.Int).Mul(big.NewInt(int64(gasPrice*100)), big.NewInt(10000000))
// 	auth.GasLimit = uint64(375000)
// 	auth.Value = new(big.Int).Div(balanceMainnet, big.NewInt(100))

// 	tx, err := KyberInstance.SwapEtherToToken(auth, common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359"), minRate)
// 	if err != nil {
// 		return err
// 	}

// 	daiBalanceNew := daiBalanceOld
// 	helpers.PrintLine("\nConverting your Ether to DAI via KyberNetwork...", helpers.INFO)
// 	// Waiting for the proposals transaction to be mined
// 	helpers.Printf("Waiting for KyberNetwork transaction to be mined", helpers.INFO)
// 	waitingFor := 0
// 	for daiBalanceOld.Cmp(daiBalanceNew) == 0 {
// 		waitingFor += 20
// 		time.Sleep(20 * time.Second)
// 		fmt.Printf(".")
// 		// Checking the current dai balance every 5 seconds
// 		daiBalanceNew, err = daiInstance.BalanceOf(nil, common.HexToAddress(config.DitConfig.EthereumKeys.Address))
// 		if err != nil {
// 			return errors.New("Failed to retrieve the current dai balance")
// 		}
// 		// If we are waiting for more than 30 minutes, the transaction might have failed
// 		if waitingFor > 60*30 {
// 			fmt.Printf("\n")
// 			helpers.PrintLine("Waiting for over 30 minutes, maybe the transaction or the network failed?", 1)
// 			helpers.PrintLine("Check at: https://etherscan.com/tx/"+tx.Hash().Hex(), 1)
// 			helpers.ExitAndLog(0)
// 		}
// 	}
// 	fmt.Printf("\n")

// 	xdaiBalanceOld, err := connectionXDai.BalanceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address), nil)
// 	if err != nil {
// 		return err
// 	}

// 	response, err = http.Get("https://ethgasstation.info/json/ethgasAPI.json")
// 	if err != nil {
// 		return err
// 	}

// 	data, _ = ioutil.ReadAll(response.Body)

// 	position = strings.Index(string(data), "average")
// 	if position == -1 {
// 		return errors.New("Failed to retrieve the current gas price on Mainnet")
// 	}
// 	gasPricePart = string(data)[position+10:]
// 	gasPriceString = strings.Split(gasPricePart, ",")[0]
// 	gasPrice, err = strconv.ParseFloat(gasPriceString, 64)
// 	if err != nil {
// 		return err
// 	}
// 	gasPrice = (gasPrice * 1.1) / 10

// 	// Crerating the transaction (basic values)
// 	auth, err = populateTx(connectionMainnet, privateKeyString)
// 	if err != nil {
// 		return err
// 	}

// 	auth.GasPrice = new(big.Int).Mul(big.NewInt(int64(gasPrice*100)), big.NewInt(10000000))
// 	auth.GasLimit = uint64(50000)

// 	tx, err = daiInstance.Transfer(auth, common.HexToAddress("0x4aa42145Aa6Ebf72e164C9bBC74fbD3788045016"), daiBalanceNew)
// 	if err != nil {
// 		return err
// 	}

// 	xdaiBalanceNew := xdaiBalanceOld
// 	helpers.PrintLine("\nTransferring your DAI from Ethereum Mainnet to xDAI on the POA Network...", helpers.INFO)
// 	// Waiting for the proposals transaction to be mined
// 	helpers.Printf("Waiting for xDAI-Bridge transaction to be mined", helpers.INFO)
// 	waitingFor = 0
// 	for xdaiBalanceOld.Cmp(xdaiBalanceNew) == 0 {
// 		waitingFor += 20
// 		time.Sleep(20 * time.Second)
// 		fmt.Printf(".")
// 		// Checking the current dai balance every 5 seconds
// 		xdaiBalanceNew, err = connectionXDai.BalanceAt(context.Background(), common.HexToAddress(config.DitConfig.EthereumKeys.Address), nil)
// 		if err != nil {
// 			return errors.New("Failed to retrieve the current xdai balance")
// 		}
// 		// If we are waiting for more than 30 minutes, the transaction might have failed
// 		if waitingFor > 60*30 {
// 			fmt.Printf("\n")
// 			helpers.PrintLine("Waiting for over 30 minutes, maybe the transaction or the network failed?", 1)
// 			helpers.PrintLine("Check at: https://etherscan.com/tx/"+tx.Hash().Hex(), 1)
// 			helpers.ExitAndLog(0)
// 		}
// 	}
// 	fmt.Printf("\n")

// 	return nil
// }

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

	// If in demo mode, also retrieve the xDIT token address
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
		config.DitConfig.Currency = "xDIT"
	} else {
		config.DitConfig.DitToken = ""
		config.DitConfig.Currency = "xDAI"
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

	var repositoryMap map[string]*config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryMap = config.DitConfig.DemoRepositories
	} else {
		repositoryMap = config.DitConfig.LiveRepositories
	}

	// Checking the repositories that are in our config
	if repositoryMap[repository] != nil {
		if len(_optionalRepository) > 0 {
			return nil
		}
		helpers.PrintLine("Repository is already initialized", helpers.INFO)
		helpers.ExitAndLog(0)
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

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return err
	}

	nextAddress, err := ditCoordinatorInstance.NextDitCoordinator(nil)
	if err != nil {
		return err
	}
	if nextAddress != common.HexToAddress("0") {
		helpers.PrintLine("There was an update to the ditCraft smartcontracts. Please update your ditCLI in order to interact with them.", helpers.INFO)
		helpers.PrintLine(fmt.Sprintf("Please execute %s to update the ditCLI", aurora.Green("bash <(curl https://get.ditcraft.io -Ls)")), helpers.INFO)
		helpers.PrintLine("Or go to: https://github.com/ditcraft/cli", helpers.INFO)
		helpers.ExitAndLog(0)
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
			err = initDitRepository(ditCoordinatorInstance, repoHash, repository)
			if err != nil {
				return err
			}
		} else {
			// If not: exit
			helpers.PrintLine("Initialization cancelled - repository can't be used with dit until this is done", 1)
			helpers.PrintLine("Execute '"+helpers.ColorizeCommand("init")+"' in the directory of the repository to do so", 1)
			helpers.ExitAndLog(0)
		}
	}

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
	repositoryMap[repository] = &newRepository
	if config.DitConfig.DemoModeActive {
		config.DitConfig.DemoRepositories = repositoryMap
	} else {
		config.DitConfig.LiveRepositories = repositoryMap
	}

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return nil
	}

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

	if config.DitConfig.LiveRepositories[repository] == nil {
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
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return 0, false, false, repository, err
	}

	// Searching for the corresponding vote in the votes stored in the config
	var activeVote *config.ActiveVote
	var voteID int
	for key, value := range config.DitConfig.LiveRepositories[repository].ActiveVotes {
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

	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, GetHashOfString(repository), big.NewInt(int64(voteID)))
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

	repoHash := GetHashOfString(repository)

	// Gathering the the knowledge-labels from the config
	knowledgeLabels := config.DitConfig.LiveRepositories[repository].KnowledgeLabels

	connection, err := getConnection()
	if err != nil {
		return "", 0, err
	}

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
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

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Prompting the user to summarize the changes included in this proposal
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

	// Retrieving the xDAI balance of the user
	xDAIBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
	}

	// Formatting the xDAI balance to a human-readable format
	floatBalance := new(big.Float).Quo((new(big.Float).SetInt(xDAIBalance)), big.NewFloat(1000000000000000000))

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

	helpers.PrintLine(fmt.Sprintf("You have a balance of %.2f xDAI", floatBalance), helpers.INFO)
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
	helpers.PrintLine("  The following stake will automatically be deducted: "+floatStakeString+"xDAI", helpers.INFO)
	userIsSure := helpers.GetUserInputChoice("Is that correct?", "y", "n")
	if userIsSure == "n" {
		return "", 0, errors.New("Canceled proposal of changes due to users choice")
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

	secondFloatLength, _ := (new(big.Float).Mul(floatLength, big.NewFloat(60))).Int64()
	intLength := big.NewInt(secondFloatLength)

	// Retrieving the last/current proposalID of the ditContract
	// (This will increment after a proposal, so we can see when the proposal is live)
	lastProposalID, err := ditCoordinatorInstance.GetCurrentProposalID(nil, repoHash)
	if err != nil {
		return "", 0, errors.New("Failed to retrieve the current proposal id")
	}

	// Proposing the commit
	transaction, err := ditCoordinatorInstance.ProposeCommit(auth, repoHash, answerProposalSummary, proposalIdentifier, big.NewInt(int64(knowledgeLabels[answerKnowledgeLabel-1].ID)), intKNW, intLength)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return "", 0, errors.New("Your account doesn't have enough xDAI to pay for the transaction")
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
	config.DitConfig.LiveRepositories[repository].ActiveVotes[newProposalID.String()] = &newVote

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return "", 0, err
	}

	// Conwerting the stake and used KNW count into a float so that it's human-readable
	intxDAI, ok := new(big.Int).SetString(newVote.NumTokens, 10)
	if !ok {
		return "", 0, errors.New("Failed to parse big.int from string")
	}
	floatxDAI := new(big.Float).Quo((new(big.Float).SetInt(intxDAI)), big.NewFloat(1000000000000000000))

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
	responseString += "\nSuccessfully proposed merge to master. Vote on proposal started with ID " + newProposalID.String() + ""
	responseString += fmt.Sprintf("\nYou staked %.2f %s and used %f KNW.", floatxDAI, config.DitConfig.Currency, floatKNW)
	responseString += "\nThe vote will end at " + timeRevealString
	if config.DitConfig.LiveRepositories[repository].Provider == "github" {
		if _branch == "" {
			_branch = "dit_proposal_" + newProposalID.String()
		}
		responseString += "\nYour proposed branch is at https://" + repository + "/tree/" + _branch
	}
	responseString += "\nTo finalize the vote, execute '" + helpers.ColorizeCommand("finalize "+newProposalID.String()) + "' after " + timeRevealString
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

	repoHash := GetHashOfString(repository)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
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

	// Setting the value of the transaction to be the default stake
	auth.Value = requiredStake

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

	if config.DitConfig.LiveRepositories[repository] == nil {
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
		config.DitConfig.LiveRepositories[repository] = &newRepository
	}

	// Adding the new vote to the config object
	config.DitConfig.LiveRepositories[repository].ActiveVotes[_proposalID] = &newVote

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

	// Letting the user know when and how he has to reveal the vote
	helpers.PrintLine("With dit, votes are casted in a concealed manner through a commitment scheme.", helpers.INFO)
	helpers.PrintLine("This means that votes have to be revealed to the public once the commit-phase is over.", helpers.INFO)
	helpers.PrintLine("Please open your vote with '"+helpers.ColorizeCommand("open "+strconv.Itoa(int(proposalID)))+"' between "+timeCommitString+" and "+timeRevealString, helpers.INFO)

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

	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return err
	}

	var repositoryMap map[string]*config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryMap = config.DitConfig.DemoRepositories
	} else {
		repositoryMap = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repository)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	if repositoryMap[repository].ActiveVotes[_proposalID] == nil {
		return errors.New("You didn't participate in this vote")
	}

	// Verifying whether the reveal period of this vote is active
	revealPeriodActive, err := KNWVotingInstance.OpenPeriodActive(nil, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve opening status")
	}

	// If it is now active it hasn't started yet or it's over
	if !revealPeriodActive {
		return errors.New("The opening phase of this vote is not active")
	}

	// Verifying whether the user has commited a vote on this proposal
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve commit status")
	}

	// If this is not the case the user never participated in this proposal through a vote
	if !didCommit {
		return errors.New("You didn't vote on this proposal")
	}

	// Verifying whether the user has revealed his vote on this proposal
	oldDidReveal, err := KNWVotingInstance.DidOpen(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
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
	choice := big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].Choice))
	salt, ok := new(big.Int).SetString(repositoryMap[repository].ActiveVotes[_proposalID].Salt, 10)
	if !ok {
		return errors.New("Error during salt recovery: " + err.Error())
	}

	// Revealing the vote on the proposal
	transaction, err := ditCoordinatorInstance.OpenVoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), choice, salt)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDAI to pay for the transaction")
		}
		return errors.New("Failed to open the vote: " + err.Error())
	}

	// Waiting for the reveal transaction to be mined
	helpers.Printf("Waiting for opening transaction to be mined", helpers.INFO)
	waitingFor := 0
	newDidReveal := oldDidReveal
	for newDidReveal == oldDidReveal {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the reveal status of the user every 5 seconds
		newDidReveal, err = KNWVotingInstance.DidOpen(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
		if err != nil {
			return errors.New("Failed to retrieve opening status")
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

	// Formatting the time of the reveal phase into a readable format
	timeReveal := time.Unix(int64(repositoryMap[repository].ActiveVotes[_proposalID].RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	helpers.PrintLine("Successfully opened your vote", helpers.INFO)
	helpers.PrintLine("To finalize it when the vote is over, execute '"+helpers.ColorizeCommand("finalize "+_proposalID)+"' after "+timeRevealString, helpers.INFO)

	return nil
}

// Finalize will finalize a vote as it will trigger the calculation of the reward of this user
// including the xDAI and KNW reward in case of a voting for the winning decision
// or the losing of xDAI and KNW in case of a voting for the losing decision
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

	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return false, false, err
	}

	var repositoryMap map[string]*config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryMap = config.DitConfig.DemoRepositories
	} else {
		repositoryMap = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repository)

	connection, err := getConnection()
	if err != nil {
		return false, false, err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
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

	if repositoryMap[repository].ActiveVotes[_proposalID] == nil {
		return false, false, errors.New("You didn't participate in this vote")
	}

	// If this user already called the Resolve function for this vote it's not possible anymore
	if repositoryMap[repository].ActiveVotes[_proposalID].Resolved {
		return false, false, errors.New("You already finalized this vote")
	}

	// Verifying whether the vote has already ended
	pollEnded, err := KNWVotingInstance.VoteEnded(nil, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve vote status")
	}

	// If not, we can't resolve it
	if !pollEnded {
		return false, false, errors.New("The vote hasn't ended yet")
	}

	// Verifying whether the user is a participant of this vote
	didCommit, err := KNWVotingInstance.DidCommit(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve commit status")
	}

	// Retrieve the selected proposal obkect
	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve the new proposal")
	}

	// Indicates whether the called was the proposer or not
	isProposer := (proposal.Proposer == myAddress)

	// If not, we are not allowed to call this function (it would fail)
	if !didCommit && myAddress != proposal.Proposer {
		return false, false, errors.New("You didn't participate in this vote")
	}

	// Saving the old xDAI balance
	oldxDAIBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
	if err != nil {
		return false, false, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
	}

	// Saving the old KNW balance
	oldKNWBalance, err := KNWTokenInstance.BalanceOfID(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KnowledgeLabel.ID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve KNW balance")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return false, false, err
	}

	// Resolving the vote
	transaction, err := ditCoordinatorInstance.FinalizeVote(auth, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return false, false, errors.New("Your account doesn't have enough xDAI to pay for the transaction")
		}
		return false, false, errors.New("Failed to finalize the vote: " + err.Error())
	}

	// Waiting for the resolve transaction to be mined
	helpers.Printf("Waiting for finalizing transaction to be mined", helpers.INFO)
	waitingFor := 0
	newxDAIBalance := oldxDAIBalance
	for oldxDAIBalance.Cmp(newxDAIBalance) == 0 {
		waitingFor += 5
		time.Sleep(5 * time.Second)
		fmt.Printf(".")
		// Checking the balance of the user every 5 seconds, if it changed, a transaction was executed
		newxDAIBalance, err = connection.BalanceAt(context.Background(), myAddress, nil)
		if err != nil {
			return false, false, errors.New("Failed to retrieve opening status")
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

	// Saving the new KNW balance after resolving the vote
	newKNWBalance, err := KNWTokenInstance.BalanceOfID(nil, myAddress, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KnowledgeLabel.ID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve KNW balance")
	}

	// Retrieving information about this vote from the KNWVoting contract itself
	KNWPoll, err := KNWVotingInstance.Votes(nil, proposal.KNWVoteID)
	if err != nil {
		return false, false, errors.New("Failed to retrieve vote information")
	}
	totalVoters := int(new(big.Int).Add(KNWPoll.ParticipantsAgainst, KNWPoll.ParticipantsFor).Int64())
	winningPercentage := int(KNWPoll.WinningPercentage.Int64())

	// Retrieving the outcome of the vote
	pollPassed, err := KNWVotingInstance.IsPassed(nil, big.NewInt(int64(repositoryMap[repository].ActiveVotes[_proposalID].KNWVoteID)))
	if err != nil {
		return false, false, errors.New("Failed to retrieve vote outcome")
	}

	fmt.Printf("\n")
	// Show the user how the vote ended
	if KNWPoll.WinningPercentage.Cmp(big.NewInt(50)) == 1 {
		if pollPassed {
			helpers.PrintLine("Successfully finalized the vote - it passed with "+strconv.Itoa(winningPercentage)+"%% approval amongst "+strconv.Itoa(totalVoters)+" validators", helpers.INFO)
			if proposal.Proposer == common.HexToAddress(config.DitConfig.EthereumKeys.Address) {
				helpers.PrintLine("You received your stake back and will share the opposing voters slashed stakes", helpers.INFO)
			}
		} else {
			helpers.PrintLine("Successfully finalized the vote - it didn't pass with "+strconv.Itoa(winningPercentage)+"%% disapproval amongst "+strconv.Itoa(totalVoters)+" validators", helpers.INFO)
			if proposal.Proposer == common.HexToAddress(config.DitConfig.EthereumKeys.Address) {
				helpers.PrintLine("Your stake was slashed and will be distributed amongst the voters", helpers.INFO)
			}
		}
	} else if KNWPoll.WinningPercentage.Cmp(big.NewInt(50)) == 0 {
		helpers.PrintLine("Successfully finalized the vote - it ended in a draw between "+strconv.Itoa(totalVoters)+" validators and didn't get accepted", helpers.INFO)
		helpers.PrintLine("You received your stake back", helpers.INFO)
	} else {
		helpers.PrintLine("Successfully finalized the vote - no one voted on it and it didn't get accepted", helpers.INFO)
		helpers.PrintLine("You received your stake back", helpers.INFO)
	}

	// If the user got some xDAI as a reward, this will be shown to the user
	if newxDAIBalance.Cmp(oldxDAIBalance) > 0 {
		difference := newxDAIBalance.Sub(newxDAIBalance, oldxDAIBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You gained %.2f "+config.DitConfig.Currency+"\n", floatDifference), helpers.INFO)
	}

	// Also shwoing the user how man KNW tokens he got/lost
	if oldKNWBalance.Cmp(newKNWBalance) < 0 {
		difference := newKNWBalance.Sub(newKNWBalance, oldKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You earned %.2f KNW Tokens for the knowledge label '%s'", floatDifference, repositoryMap[repository].ActiveVotes[_proposalID].KnowledgeLabel.Label), helpers.INFO)
	} else if oldKNWBalance.Cmp(newKNWBalance) > 0 {
		difference := oldKNWBalance.Sub(oldKNWBalance, newKNWBalance)
		floatDifference := new(big.Float).Quo((new(big.Float).SetInt64(difference.Int64())), big.NewFloat(1000000000000000000))
		helpers.PrintLine(fmt.Sprintf("You lost %.2f KNW Tokens for the knowledge label '%s'", floatDifference, repositoryMap[repository].ActiveVotes[_proposalID].KnowledgeLabel.Label), helpers.INFO)
	}

	// Saving the resolved-status in the config
	repositoryMap[repository].ActiveVotes[_proposalID].Resolved = true

	if config.DitConfig.DemoModeActive {
		config.DitConfig.DemoRepositories = repositoryMap
	} else {
		config.DitConfig.LiveRepositories = repositoryMap
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
	// Retrieve the name of the repo we are in from git
	repository, err := git.GetRepository()
	if err != nil {
		return err
	}

	var repositoryMap map[string]*config.Repository
	if config.DitConfig.DemoModeActive {
		repositoryMap = config.DitConfig.DemoRepositories
	} else {
		repositoryMap = config.DitConfig.LiveRepositories
	}

	repoHash := GetHashOfString(repository)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		return err
	}

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return err
	}

	// Retrieving the current proposalID
	currentProposalIDBigInt, err := ditCoordinatorInstance.GetCurrentProposalID(nil, repoHash)
	if err != nil {
		return errors.New("Failed to retrieve the current proposal id")
	}

	// Converting the current proposal id of the ditContract to an int
	// if it is zero, there are no votes
	currentProposalID := int(currentProposalIDBigInt.Int64())
	if currentProposalID == 0 {
		helpers.PrintLine("There are no votes for this repository", 1)
		helpers.ExitAndLog(0)
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
	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return errors.New("Failed to retrieve the new proposal")
	}

	// Retrieving information about this vote from the KNWVoting contract itself
	KNWPoll, err := KNWVotingInstance.Votes(nil, proposal.KNWVoteID)
	if err != nil {
		return errors.New("Failed to retrieve vote information")
	}

	knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, proposal.KnowledgeID)
	if err != nil {
		return errors.New("Failed to retrieve knowledge label")
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
	helpers.PrintLine("---------------------------", helpers.INFO)
	helpers.PrintLine("Proposal ID: "+strconv.Itoa(proposalID), helpers.INFO)
	helpers.PrintLine("Description: "+proposal.Description, helpers.INFO)
	identifierParts := strings.Split(proposal.Identifier, ":")
	if len(identifierParts) == 1 {
		helpers.PrintLine("Branch: dit_proposal_"+strconv.Itoa(proposalID), helpers.INFO)
		helpers.PrintLine("URL: https://"+repository+"/tree/"+identifierParts[0], helpers.INFO)
	} else {
		helpers.PrintLine("Branch: "+identifierParts[0], helpers.INFO)
		helpers.PrintLine("URL: https://"+repository+"/tree/"+identifierParts[1], helpers.INFO)
	}
	helpers.PrintLine("Proposer: "+proposal.Proposer.Hex(), helpers.INFO)
	helpers.PrintLine("Knowledge-Label: "+knowledgeLabel, helpers.INFO)

	// If the user participated in this vote, the choice and stake/KNW are also printed
	if repositoryMap[repository] != nil && repositoryMap[repository].ActiveVotes[strconv.Itoa(proposalID)] != nil {
		// TODO: FIX nil pointer crash
		intxDAI, ok := new(big.Int).SetString(repositoryMap[repository].ActiveVotes[strconv.Itoa(proposalID)].NumTokens, 10)
		if !ok {
			return errors.New("Failed to parse big.int from string")
		}
		floatxDAI := new(big.Float).Quo((new(big.Float).SetInt(intxDAI)), big.NewFloat(1000000000000000000))
		intKNW, ok := new(big.Int).SetString(repositoryMap[repository].ActiveVotes[strconv.Itoa(proposalID)].NumKNW, 10)
		if !ok {
			return errors.New("Failed to parse big.int from string")
		}
		floatKNW := new(big.Float).Quo((new(big.Float).SetInt(intKNW)), big.NewFloat(1000000000000000000))
		fmt.Println()
		if proposal.Proposer.Hex() != config.DitConfig.EthereumKeys.Address {
			helpers.PrintLine("Your choice: "+strconv.Itoa(repositoryMap[repository].ActiveVotes[strconv.Itoa(proposalID)].Choice), helpers.INFO)
		}
		helpers.PrintLine(fmt.Sprintf("You staked %.2f %s and used %f KNW", floatxDAI, config.DitConfig.Currency, floatKNW), helpers.INFO)
	}
	helpers.PrintLine(fmt.Sprintf("Required (gross) stake: %.2f "+config.DitConfig.Currency, floatGrossStake), helpers.INFO)
	helpers.PrintLine(fmt.Sprintf("Current net stake: %.2f "+config.DitConfig.Currency, floatNetStake), helpers.INFO)
	fmt.Println()
	helpers.PrintLine("Vote phase end: "+timeCommitString, helpers.INFO)
	helpers.PrintLine("Opening phase end: "+timeRevealString, helpers.INFO)
	helpers.PrintLine("Resolved? "+strconv.FormatBool(proposal.IsFinalized), helpers.INFO)
	if proposal.IsFinalized {
		winningPercentage := int(KNWPoll.WinningPercentage.Int64())
		if KNWPoll.WinningPercentage.Cmp(big.NewInt(50)) == 1 {
			if proposal.ProposalAccepted {
				helpers.PrintLine("Passed? Accepted with "+strconv.Itoa(winningPercentage)+"%% approval", helpers.INFO)
			} else {
				helpers.PrintLine("Passed? Rejected with "+strconv.Itoa(winningPercentage)+"%% approval", helpers.INFO)
			}
		} else if KNWPoll.WinningPercentage.Cmp(big.NewInt(50)) == 0 {
			helpers.PrintLine("Passed? Ended in a draw", helpers.INFO)
		} else {
			helpers.PrintLine("Passed? No validators voted on this proposal - it didn't get accepted", helpers.INFO)
		}
	}
	helpers.PrintLine("---------------------------", helpers.INFO)

	return nil
}

// GetBalances will print the xDAI and KNW balances
func GetBalances() (float64, []string, []float64, error) {
	connection, err := getConnection()
	if err != nil {
		return 0, nil, nil, err
	}

	if len(config.DitConfig.KNWToken) != correctETHAddressLength {
		return 0, nil, nil, errors.New("Invalid KNWToken address, please do '" + helpers.ColorizeCommand("set_coordinator") + "' first")
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(config.DitConfig.EthereumKeys.Address)

	// Create a new instance of the KNWToken to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return 0, nil, nil, err
	}

	// Retrieving the number of knowledge ids that exist
	amountOfIDs, err := KNWTokenInstance.AmountOfIDs(nil)
	if err != nil {
		return 0, nil, nil, errors.New("Failed to retrieve knowledge label count of address")
	}

	var bigFloatBalance *big.Float
	if !config.DitConfig.DemoModeActive {
		// Retrieving the xDAI balance of the user
		xDAIBalance, err := connection.BalanceAt(context.Background(), myAddress, nil)
		if err != nil {
			return 0, nil, nil, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
		}
		helpers.PrintLine("Balances for address "+config.DitConfig.EthereumKeys.Address+":", helpers.INFO)

		// Formatting the xDAI balance to a human-readable format
		bigFloatBalance = new(big.Float).Quo((new(big.Float).SetInt(xDAIBalance)), big.NewFloat(1000000000000000000))
	} else {
		// Create a new instance of the ditToken to access it
		ditTokenInstance, err := getDitTokenInstance(connection)
		if err != nil {
			return 0, nil, nil, err
		}
		// Retrieving the xDIT balance of the user
		xDITBalance, err := ditTokenInstance.BalanceOf(nil, myAddress)
		if err != nil {
			return 0, nil, nil, errors.New("Failed to retrieve " + config.DitConfig.Currency + " balance")
		}

		// Formatting the xDAI balance to a human-readable format
		bigFloatBalance = new(big.Float).Quo((new(big.Float).SetInt(xDITBalance)), big.NewFloat(1000000000000000000))
	}

	floatBalance, _ := bigFloatBalance.Float64()

	// Retrieving each KNW balance
	var labelArray []string
	var balanceArray []float64
	for i := 0; i < int(amountOfIDs.Int64()); i++ {
		// Retrieve each KNW balance
		balanceOfID, err := KNWTokenInstance.BalanceOfID(nil, myAddress, big.NewInt(int64(i)))
		if err != nil {
			return floatBalance, nil, nil, err
		}
		if balanceOfID.Cmp(big.NewInt(10^16)) >= 0 {
			knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, big.NewInt(int64(i)))
			if err != nil {
				return floatBalance, nil, nil, err
			}
			// Formatting each KNW balance to a human-readable format
			floatBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(balanceOfID)), big.NewFloat(1000000000000000000))).Float64()
			labelArray = append(labelArray, knowledgeLabel)
			balanceArray = append(balanceArray, floatBalance)
		}
	}

	return floatBalance, labelArray, balanceArray, nil
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

	// Retrieving the amount of xDAI a user has staked for this vote
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
func initDitRepository(_ditCoordinatorInstance *ditCoordinator.DitCoordinator, _repoHash [32]byte, _name string) error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		return err
	}

	amountOfIDs, err := KNWTokenInstance.AmountOfIDs(nil)
	if err != nil {
		return err
	}

	var allowedKnowledgeLabels []string
	labelToID := make(map[string]*big.Int)
	for i := 0; i < int(amountOfIDs.Int64()); i++ {
		knowledgeLabel, err := KNWTokenInstance.LabelOfID(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		labelComponents := strings.Split(knowledgeLabel, "/")
		allowedKnowledgeLabels = append(allowedKnowledgeLabels, labelComponents[1])
		labelToID[labelComponents[1]] = big.NewInt(int64(i))
	}

	neededMajority := big.NewInt(50)

	// Prompting the user to provide knowledge-labels for this repository
	helpers.PrintLine("Please provide knowledge labels that will be used for this repository:", helpers.INFO)
	helpers.PrintLine("Note: At least one has to be provided, press ENTER to finish", helpers.INFO)
	var knowledgeIDs []*big.Int
	for len(knowledgeIDs) < 10 && len(allowedKnowledgeLabels) > 0 {
		newLabel := helpers.GetUserInput("Knowledge Label " + strconv.Itoa(len(knowledgeIDs)+1))
		foundLabel := false
		for _, element := range allowedKnowledgeLabels {
			if element == newLabel {
				foundLabel = true
				break
			}
		}
		if !foundLabel {
			closestMatcher := closestmatch.New(allowedKnowledgeLabels, []int{2, 3, 4, 5})
			proposedLabel := closestMatcher.Closest(newLabel)
			if len(proposedLabel) > 0 {
				choice := helpers.GetUserInputChoice("Closest match found was '"+proposedLabel+"', do you want to use that?", "y", "n")
				if choice == "y" {
					newLabel = proposedLabel
					foundLabel = true
				} else {
					fmt.Println("Please use a knowledge-label from this list: " + strings.Join(allowedKnowledgeLabels, ", "))
				}
			} else if len(newLabel) != 0 {
				fmt.Println("Please use a knowledge-label from this list: " + strings.Join(allowedKnowledgeLabels, ", "))
			}
		}
		if foundLabel && len(newLabel) > 0 {
			knowledgeIDs = append(knowledgeIDs, labelToID[newLabel])
			delete(labelToID, newLabel)
			for i := range allowedKnowledgeLabels {
				if allowedKnowledgeLabels[i] == newLabel {
					allowedKnowledgeLabels[i] = allowedKnowledgeLabels[len(allowedKnowledgeLabels)-1]
					allowedKnowledgeLabels = allowedKnowledgeLabels[:len(allowedKnowledgeLabels)-1]
					break
				}
			}
		}
		if len(newLabel) == 0 && len(knowledgeIDs) > 0 {
			break
		}
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	// Initializing the repository = deploying a new ditContract
	transaction, err := _ditCoordinatorInstance.InitRepository(auth, _name, knowledgeIDs, neededMajority)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough xDAI to pay for the transaction")
		}
		return err
	}

	// Waiting for the deployment transaction to be mined
	helpers.Printf("Waiting for initialization transaction to be mined", helpers.INFO)
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
			helpers.ExitAndLog(0)
		}
	}
	fmt.Printf("\n")

	return nil
}

// populateTX will set the necessary values for a ethereum transaction
// amount of gas, gasprice, nonce, sign this with the private key
func populateTx(_connection *ethclient.Client, _privatKey ...string) (*bind.TransactOpts, error) {
	// Retrieve the decrypted private key through a password prompt
	var err error
	var privateKeyString string
	if len(_privatKey) != 1 {
		privateKeyString, err = config.GetPrivateKey(true)
		if err != nil {
			return nil, err
		}
	} else {
		privateKeyString = _privatKey[0]
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

func getKyberInstance(_connection *ethclient.Client) (*KyberNetwork.KyberNetwork, error) {
	KyberNetworkProxyAddress := common.HexToAddress(config.KyberNetworkProxy)

	// Create a new instance of the KyberNetworkProxy contract to access it
	KyberInstance, err := KyberNetwork.NewKyberNetwork(KyberNetworkProxyAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find KyberNetworkProxy at provided address")
	}

	return KyberInstance, nil
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

// getConnection will return a connection to the ethereum blockchain
func getConnectionMainnet() (*ethclient.Client, error) {
	for i := 0; i < len(config.MainnetNodes); i++ {
		connection, err := ethclient.Dial(config.MainnetNodes[i])
		if err != nil {
			if i == len(config.MainnetNodes)-1 {
				return nil, errors.New("Failed to connect to the ethereum network" + strconv.Itoa(i))
			}
		} else {
			networkID, err := connection.NetworkID(context.Background())
			if err != nil {
				if i == len(config.MainnetNodes)-1 {
					return nil, errors.New("Failed to connect to the ethereum network" + strconv.Itoa(i))
				}
			} else {
				if networkID.Cmp(big.NewInt(1)) == 0 {
					return connection, nil
				}
			}
		}
	}
	return nil, errors.New("Failed to connect to the ethereum network")
}

// GetHashOfString takes a string a returns it as keccak256
func GetHashOfString(_string string) [32]byte {
	repoHash32 := [32]byte{}
	copy(repoHash32[:], crypto.Keccak256([]byte(_string))[:])

	return repoHash32
}

func getDAIInstance(_connection *ethclient.Client) (*ditToken.MintableERC20, error) {
	daiAddress := common.HexToAddress("0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359")

	// Create a new instance of the KNWToken contract to access it
	daiInstance, err := ditToken.NewMintableERC20(daiAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find DAI contract at provided address")
	}

	return daiInstance, nil
}
