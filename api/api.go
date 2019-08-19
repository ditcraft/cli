package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ditcraft/cli/config"
)

type user struct {
	Address       string                  `json:"address" bson:"address"`
	TwitterHandle string                  `json:"twitter_handle" bson:"twitter_handle"`
	TwitterID     string                  `json:"twitter_id" bson:"twitter_id"`
	XDAIBalance   float64                 `json:"xdai_balance" bson:"xdai_balance"`
	XDITBalance   float64                 `json:"xdit_balance" bson:"xdit_balance"`
	KNWTokens     []knwLabels             `json:"knw_tokens" bson:"knw_tokens"`
	Proposals     []proposalDetailsUser   `json:"proposals" bson:"proposals"`
	Repositories  []repositoryDetailsUser `json:"repositories" bson:"repositories"`
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
	for _, knwToken := range userObject.KNWTokens {
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
