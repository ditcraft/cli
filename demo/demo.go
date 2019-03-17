package demo

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ditcraft/client/config"
	"github.com/ditcraft/client/git"
	"github.com/ditcraft/client/helpers"
	"github.com/ditcraft/client/smartcontracts/KNWToken"
	"github.com/ditcraft/client/smartcontracts/KNWVoting"
	"github.com/ditcraft/client/smartcontracts/ditCoordinator"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const correctETHAddressLength = 42

var demoVoterAddresses = []string{
	"0x0000000000000000000000000000000000000000",
	"0x0000000000000000000000000000000000000000",
	"0x0000000000000000000000000000000000000000",
}
var demoVoterPrivateKeys = []string{
	"0000000000000000000000000000000000000000000000000000000000000000",
	"0000000000000000000000000000000000000000000000000000000000000000",
	"0000000000000000000000000000000000000000000000000000000000000000",
}

// Vote will cast the demo accounts' votes on a proposal
func Vote(_proposalID string) error {
	helpers.PrintLine("Starting demo voters to simulate participants", 3)
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		randomSelection := rand.Intn(2)
		randomSalt := rand.Intn(123456789)
		err := _executeVote(_proposalID, strconv.Itoa(randomSelection), strconv.Itoa(randomSalt), i)
		if err != nil {
			helpers.PrintLine("Error during vote commiting of demo voter "+strconv.Itoa(i), 2)
		}
	}
	return nil
}

func _executeVote(_proposalID string, _choice string, _salt string, _demoVoter int) error {
	// Converting the stdin string input of the user into Ints
	proposalID, _ := strconv.Atoi(_proposalID)
	choice, _ := strconv.Atoi(_choice)
	salt, _ := strconv.Atoi(_salt)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	repoHash := getHashOfString(config.DitConfig.Repositories[ditContractIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into address object
	myAddress := common.HexToAddress(demoVoterAddresses[_demoVoter])

	// Create a new instance of the ditContract to access it
	ditCoordingatorInstance, err := getDitCoordinatorInstance(connection)
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

	// Retrieving the proposal object from the ditContract

	proposal, err := ditCoordingatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
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

	var auth *bind.TransactOpts
	// Crerating the transaction (basic values)
	auth, err = populateTx(connection, demoVoterPrivateKeys[_demoVoter], demoVoterAddresses[_demoVoter])
	if err != nil {
		return err
	}

	// Setting the value of the transaction to be the default stake
	auth.Value = requiredStake

	// Voting on the proposal
	_, err = ditCoordingatorInstance.VoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), voteHash)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return errors.New("Failed to commit the vote: " + err.Error())
	}

	// We will also store the users choice and salt, so that the user doesn't need to remember the salt
	// when he will reveal the vote later on
	if _demoVoter == 0 {
		config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices = make([]int, 3)
		config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoSalts = make([]int, 3)
	}
	config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices[_demoVoter] = choice
	config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoSalts[_demoVoter] = salt

	// Saving the config back to the file
	err = config.Save()
	if err != nil {
		return err
	}

	// Formatting the time of the commit and reveal phase into a readable format
	timeCommit := time.Unix(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].CommitEnd), 0)
	timeCommitString := timeCommit.Format("15:04:05 on 2006/01/02")
	timeReveal := time.Unix(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].RevealEnd), 0)
	timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

	helpers.PrintLine("Demo-Voter "+strconv.Itoa(_demoVoter)+" casted a vote on the proposed commit", 3)
	if _demoVoter == 2 {
		fmt.Println()
		helpers.PrintLine("With dit, votes are casted in a concealed manner through a commitment scheme.", 3)
		helpers.PrintLine("This means that votes have to be opened and thus revealed to the public once the commit-phase is over.", 3)
		helpers.PrintLine("Please open the concealed demo votes with '"+helpers.ColorizeCommand("demo_open "+strconv.Itoa(int(proposalID)))+"' between "+timeCommitString+" and "+timeRevealString, 3)
	}

	return nil
}

// Open will reveal the demo voters' conceal votes
func Open(_proposalID string, _demoVoter int) error {
	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return err
	}

	repoHash := getHashOfString(config.DitConfig.Repositories[ditContractIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(demoVoterAddresses[_demoVoter])

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

	// Searching for the corresponding vote in the votes stored in the config
	var voteIndex int
	for i := range config.DitConfig.Repositories[ditContractIndex].ActiveVotes {
		if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[i].ID == proposalID {
			voteIndex = i
			break
		}
	}

	if len(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices) != 3 || len(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoSalts) != 3 {
		return errors.New("No demo voters voted on this vote")
	}

	// Verifying whether the reveal period of this vote is active
	revealPeriodActive, err := KNWVotingInstance.RevealPeriodActive(nil, big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].KNWVoteID)))
	if err != nil {
		return errors.New("Failed to retrieve opening status")
	}

	// If it is now active it hasn't started yet or it's over
	if !revealPeriodActive {
		return errors.New("The opening phase of this vote is not active")
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
		return errors.New("Failed to retrieve opening status")
	}

	// If this is the case, the user already revealed his vote
	if oldDidReveal {
		return errors.New("You already opened your vote on this proposal")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection, demoVoterPrivateKeys[_demoVoter], demoVoterAddresses[_demoVoter])
	if err != nil {
		return err
	}

	// Gathering the original choice and the salt from the config
	choice := big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices[_demoVoter]))
	salt := big.NewInt(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoSalts[_demoVoter]))

	// Revealing the vote on the proposal
	_, err = ditCoordinatorInstance.OpenVoteOnProposal(auth, repoHash, big.NewInt(int64(proposalID)), choice, salt)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return errors.New("Failed to open the vote: " + err.Error())
	}

	var stringChoice string
	if config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices[_demoVoter] == 1 {
		stringChoice = "for"
	} else {
		stringChoice = "against"
	}
	helpers.PrintLine("Demo-Voter "+strconv.Itoa(_demoVoter)+" opened his vote commitment (voted "+stringChoice+" the proposal) ", 3)

	if _demoVoter == 2 {
		// Formatting the time of the reveal phase into a readable format
		timeReveal := time.Unix(int64(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].RevealEnd), 0)
		timeRevealString := timeReveal.Format("15:04:05 on 2006/01/02")

		helpers.PrintLine("Successfully opened all the concealed demo votes.", 3)
		fmt.Println("")
		helpers.PrintLine("After the vote ended, the vote has to be finalized. This includes", 3)
		helpers.PrintLine("calculating the outcome and distributing KNW tokens and ETH stakes.", 3)
		helpers.PrintLine("To do so when the vote is over, execute '"+helpers.ColorizeCommand("finalize "+_proposalID)+"' after "+timeRevealString, 3)
	}
	return nil
}

// Finalize will finalize a demo voters' vote as it will trigger the calculation of the reward
// of this user including the ETH and KNW reward in case of a voting for the winning decision
// or the losing of ETH and KNW in case of a voting for the losing decision
// The first caller who executes this will also trigger the calculation whether the vote passed or not
func Finalize(_proposalID string, _demoVoter int) (bool, error) {
	// Converting the stdin string input of the user into an int
	proposalID, _ := strconv.Atoi(_proposalID)

	// Searching for this repositories object in the config
	ditContractIndex, err := searchForRepoInConfig()
	if err != nil {
		return false, err
	}

	repoHash := getHashOfString(config.DitConfig.Repositories[ditContractIndex].Name)

	connection, err := getConnection()
	if err != nil {
		return false, err
	}

	// Convertig the hex-string-formatted address into an address object
	myAddress := common.HexToAddress(demoVoterAddresses[_demoVoter])

	// Create a new instance of the ditContract to access it
	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection)
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

	if len(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoChoices) != 3 || len(config.DitConfig.Repositories[ditContractIndex].ActiveVotes[voteIndex].DemoSalts) != 3 {
		return false, errors.New("No demo voters voted on this vote")
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
	proposal, err := ditCoordinatorInstance.ProposalsOfRepository(nil, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		return false, errors.New("Failed to retrieve the new proposal")
	}

	// If not, we are not allowed to call this function (it would fail)
	if !didCommit && myAddress != proposal.Proposer {
		return false, errors.New("You didn't participate in this vote")
	}

	// Crerating the transaction (basic values)
	auth, err := populateTx(connection, demoVoterPrivateKeys[_demoVoter], demoVoterAddresses[_demoVoter])
	if err != nil {
		return false, err
	}

	// Resolving the vote
	_, err = ditCoordinatorInstance.FinalizeVote(auth, repoHash, big.NewInt(int64(proposalID)))
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return false, errors.New("Your account doesn't have enough ETH to pay for the transaction")
		}
		return false, errors.New("Failed to finalize the vote: " + err.Error())
	}

	helpers.PrintLine("Finalized vote for demo voter "+strconv.Itoa(_demoVoter), 3)

	return true, nil
}

// populateTX will set the necessary values for a ethereum transaction
// amount of gas, gasprice, nonce, sign this with the private key
func populateTx(_connection *ethclient.Client, _privateKey string, _address string) (*bind.TransactOpts, error) {
	// Converting the private key string into a private key object
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		return nil, errors.New("Failed to convert ethereum private-key")
	}

	// Retrieving the current pending nonce of our address
	pendingNonce, err := _connection.PendingNonceAt(context.Background(), common.HexToAddress(_address))
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}
	// Retrieving the current non-pending nonce of our address
	nonpendingNonce, err := _connection.NonceAt(context.Background(), common.HexToAddress(_address), nil)
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

	// Minimum gas price is 10 gwei for now, which works best for rinkeby
	// Will be changed later on
	defaultGasPrice := big.NewInt(10000000000)
	if gasPrice.Cmp(defaultGasPrice) != 1 {
		gasPrice = defaultGasPrice
	}
	// Setting the values into the transaction-options object
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func getDitCoordinatorInstance(_connection *ethclient.Client, _ditCoordinatorAddress ...string) (*ditCoordinator.DitCoordinator, error) {
	var ditCoordinatorAddressString string

	if len(_ditCoordinatorAddress) > 0 {
		ditCoordinatorAddressString = _ditCoordinatorAddress[0]
	} else {
		if len(config.DitConfig.DitCoordinator) != correctETHAddressLength {
			return nil, errors.New("Invalid ditCoordinator address, please do '" + helpers.ColorizeCommand("set_coordinator") + "' first")
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
	connection, err := ethclient.Dial("https://rinkeby.infura.io/v3/e0c6c62366d14f509033c919f2c72767")
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

func getHashOfString(_string string) [32]byte {
	repoHash32 := [32]byte{}
	copy(repoHash32[:], crypto.Keccak256([]byte(_string))[:])

	return repoHash32
}
