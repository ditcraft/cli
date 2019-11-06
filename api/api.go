package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ditcraft/cli/helpers"
	ditLog "github.com/ditcraft/cli/log"
	"github.com/logrusorgru/aurora"

	"github.com/ditcraft/cli/config"
)

// user struct
type user struct {
	DitAddress          string                  `json:"dit_address" bson:"dit_address"`
	AuthorizedAddresses authorizedAddress       `json:"authorized_adresses" bson:"authorized_adresses"`
	TwitterID           string                  `json:"twitter_id" bson:"twitter_id"`
	GitHubID            string                  `json:"github_id" bson:"github_id"`
	GitHubToken         string                  `json:"github_token" bson:"github_token"`
	MainAccount         string                  `json:"main_account" bson:"main_account"`
	XDAIBalance         float64                 `json:"xdai_balance" bson:"xdai_balance"`
	XDITBalance         float64                 `json:"xdit_balance" bson:"xdit_balance"`
	KNWTokensLive       []knwLabels             `json:"knw_tokens_live" bson:"knw_tokens_live"`
	ProposalsLive       []proposalDetailsUser   `json:"proposals_live" bson:"proposals_live"`
	RepositoriesLive    []repositoryDetailsUser `json:"repositories_live" bson:"repositories_live"`
	KNWTokensDemo       []knwLabels             `json:"knw_tokens_demo" bson:"knw_tokens_demo"`
	ProposalsDemo       []proposalDetailsUser   `json:"proposals_demo" bson:"proposals_demo"`
	RepositoriesDemo    []repositoryDetailsUser `json:"repositories_demo" bson:"repositories_demo"`
}

type authorizedAddress struct {
	DitCLI      string `json:"dit_cli" bson:"dit_cli"`
	DitExplorer string `json:"dit_explorer" bson:"dit_explorer"`
	Alice       string `json:"alice" bson:"alice"`
}

type knwLabels struct {
	Label      string  `json:"label" bson:"label"`
	Balance    float64 `json:"balance" bson:"balance"`
	RawBalance string  `json:"raw_balance" bson:"raw_balance"`
}

type proposalDetailsUser struct {
	KNWVoteID     int       `json:"knw_vote_id" bson:"knw_vote_id"`
	VoteDate      time.Time `json:"vote_date" bson:"vote_date"`
	UsedKNW       float64   `json:"used_knw" bson:"used_knw"`
	Stake         float64   `json:"stake" bson:"stake"`
	IsProposer    bool      `json:"is_proposer" bson:"is_proposer"`
	Opened        bool      `json:"opened" bson:"opened"`
	Finalized     bool      `json:"finalized" bson:"finalized"`
	VotedRight    bool      `json:"voted_right" bson:"voted_right"`
	KNWDifference float64   `json:"knw_difference" bson:"knw_difference"`
}

type repositoryDetailsUser struct {
	Hash                string    `json:"hash" bson:"hash"`
	Name                string    `json:"name" bson:"name"`
	Provider            string    `json:"provider" bson:"provider"`
	URL                 string    `json:"url" bson:"url"`
	LastActivityDate    time.Time `json:"last_activity_date" bson:"last_activity_date"`
	EarnedKNW           float64   `json:"earned_knw" bson:"earned_knw"`
	AmountOfProposals   int       `json:"amount_of_proposals" bson:"amount_of_proposals"`
	AmountOfValidations int       `json:"amount_of_validations" bson:"amount_of_validations"`
}

type requestProposal struct {
	APIKey            string `json:"api_key" bson:"api_key"`
	Mode              string `json:"mode" bson:"mode"`
	KNWVoteID         int    `json:"knw_vote_id" bson:"knw_vote_id"`
	UserAddress       string `json:"user_address" bson:"user_address"`
	UserTwitterHandle string `json:"user_twitter_handle" bson:"user_twitter_handle"`
	UserIsProposer    bool   `json:"user_is_proposer" bson:"user_is_proposer"`
	UserIsValidator   bool   `json:"user_is_validator" bson:"user_is_validator"`
	SinceDate         string `json:"since_date" bson:"since_date"`
	Amount            int    `json:"amount" bson:"amount"`
	OnlyFinalized     bool   `json:"only_finalized" bson:"only_finalized"`
	OnlyActive        bool   `json:"only_active" bson:"only_active"`
	RepositoryName    string `json:"repository_name" bson:"repository_name"`
	RepositoryHash    string `json:"repository_hash" bson:"repository_hash"`
}

type requestRepository struct {
	APIKey            string `json:"api_key" bson:"api_key"`
	Mode              string `json:"mode" bson:"mode"`
	Name              string `json:"name" bson:"name"`
	Hash              string `json:"hash" bson:"hash"`
	Provider          string `json:"provider" bson:"provider"`
	KnowledgeLabel    string `json:"knowledge_label" bson:"knowledge_label"`
	UserAddress       string `json:"user_address" bson:"user_address"`
	UserTwitterHandle string `json:"user_twitter_handle" bson:"user_twitter_handle"`
	UserIsProposer    bool   `json:"user_is_proposer" bson:"user_is_proposer"`
	UserIsValidator   bool   `json:"user_is_validator" bson:"user_is_validator"`
	OnlyActive        bool   `json:"only_active" bson:"only_active"`
	Amount            int    `json:"amount" bson:"amount"`
}

type requestUser struct {
	APIKey        string `json:"api_key" bson:"api_key"`
	Mode          string `json:"mode" bson:"mode"`
	Address       string `json:"address" bson:"address"`
	TwitterHandle string `json:"twitter_handle" bson:"twitter_handle"`
	TwitterID     string `json:"twitter_id" bson:"twitter_id"`
}

type sendLog struct {
	APIKey  string `json:"api_key" bson:"api_key"`
	Address string `json:"address" bson:"address"`
	Version string `json:"version" bson:"version"`
	LogData string `json:"log_data" bson:"log_data"`
}

type responseStatus struct {
	Status string `json:"status" bson:"status"`
}

// GetBalances will return the xDAI/xDIT and KNW-token balances
func GetBalances() (float64, []string, []float64, error) {
	var request requestUser
	request.APIKey = config.DitConfig.APIKey
	request.Address = config.DitConfig.EthereumKeys.Address
	if config.DitConfig.DemoModeActive {
		request.Mode = "demo"
	} else {
		request.Mode = "live"
	}

	jsonValue, _ := json.Marshal(request)
	response, err := http.Post("https://server.ditcraft.io/api/user", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return 0, nil, nil, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	var userObject user
	err = json.Unmarshal(data, &userObject)
	if err != nil {
		return 0, nil, nil, err
	}

	var labelArray []string
	var balanceArray []float64
	var tokenObject []knwLabels
	if config.DitConfig.DemoModeActive {
		tokenObject = userObject.KNWTokensDemo
	} else {
		tokenObject = userObject.KNWTokensLive
	}
	for _, knwToken := range tokenObject {
		if knwToken.Balance >= 0.01 {
			labelArray = append(labelArray, knwToken.Label)
			balanceArray = append(balanceArray, knwToken.Balance)
		}
	}
	if config.DitConfig.DemoModeActive {
		return userObject.XDITBalance, labelArray, balanceArray, err
	}
	return userObject.XDAIBalance, labelArray, balanceArray, err
}

// UploadLog func
func UploadLog(_version string) error {
	helpers.PrintLine("This will upload your ditCLI's logs onto our servers.", helpers.INFO)
	helpers.PrintLine("Sensitive information like your private keys are "+fmt.Sprintf("%s", aurora.Bold("**not**"))+" included in these logs!", helpers.INFO)
	choice := helpers.GetUserInputChoice("Do you wish to proceed?", "y", "n")
	if choice != "y" {
		return errors.New("Canceled upload of logs due to users choice")
	}

	var sendLog sendLog
	sendLog.APIKey = config.DitConfig.APIKey
	sendLog.Address = config.DitConfig.EthereumKeys.Address
	sendLog.Version = _version
	logData, err := ditLog.GetRawLog()
	if err != nil {
		return err
	}
	sendLog.LogData = string(logData)

	jsonValue, _ := json.Marshal(sendLog)
	response, err := http.Post("https://server.ditcraft.io/api/uploadlog", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(response.Body)
	var status responseStatus
	err = json.Unmarshal(data, &status)
	if err != nil {
		return err
	}

	if len(status.Status) < 5 {
		return errors.New("There was an error uploading the logs - please try again")
	}

	helpers.PrintLine("Successfully uploaded your logs under the id "+fmt.Sprintf("%s", aurora.Bold("#"+status.Status)), helpers.INFO)
	helpers.PrintLine("Please include this id when referring to your logs so that we can easily identify them.", helpers.INFO)
	return nil
}
