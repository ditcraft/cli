package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"strconv"
	"strings"

	"github.com/ditcraft/cli/helpers"
	ditLog "github.com/ditcraft/cli/log"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ssh/terminal"
)

// DitConfig is exported since it needs to be accessed from other packages all the time
var DitConfig ditConfig

// Version of the config, will be incremented after every ditCLI update that modified the config file
// or the smart contracts in a way that an update is necessaray
var Version = 5

// EthereumNodes is an array of rpc nodes that are used. First one is the primary, if one fails,
// the next one is used
var EthereumNodes = []string{"https://node.ditcraft.io", "https://dai.poa.network"}

// MainnetNodes contains the infura address to the eth mainnet
var MainnetNodes = []string{"https://mainnet.infura.io/v3/e0c6c62366d14f509033c919f2c72767"}

// KyberNetworkProxy address that will be used for the ETH<->DAI swap
var KyberNetworkProxy = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"

type ditConfig struct {
	DitCoordinator   string                 `json:"dit_coordinator"`
	KNWVoting        string                 `json:"knw_voting"`
	KNWToken         string                 `json:"knw_token"`
	DitToken         string                 `json:"dit_token,omitempty"`
	Currency         string                 `json:"currency"`
	PassedKYC        bool                   `json:"passed_kyc"`
	DemoModeActive   bool                   `json:"demo_mode_active"`
	Version          int                    `json:"version"`
	APIKey           string                 `json:"api_key"`
	EthereumKeys     ethereumKeys           `json:"ethereum_keys"`
	LiveRepositories map[string]*Repository `json:"live_repositories"`
	DemoRepositories map[string]*Repository `json:"demo_repositories"`
}

type ethereumKeys struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

// Repository struct, exported since its used in the ethereum package for new repositories
type Repository struct {
	Provider        string                 `json:"provider"`
	KnowledgeLabels []KnowledgeLabel       `json:"knowledge_labels"`
	ActiveVotes     map[string]*ActiveVote `json:"active_votes"`
}

// KnowledgeLabel struct, exported since its used in the ethereum package for new repositories
type KnowledgeLabel struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

// ActiveVote struct, exported since its used in the ethereum package for new votes
type ActiveVote struct {
	KNWVoteID      int            `json:"knw_vote_id"`
	KnowledgeLabel KnowledgeLabel `json:"knowledge_label"`
	BranchHash     string         `json:"branch_hash"`
	NewHeadHash    string         `json:"new_head_hash"`
	Choice         int            `json:"choice"`
	Salt           string         `json:"salt"`
	NumTokens      string         `json:"num_tokens"`
	NumVotes       string         `json:"num_votes"`
	NumKNW         string         `json:"num_knw"`
	CommitEnd      int            `json:"commit_end"`
	RevealEnd      int            `json:"reveal_end"`
	Resolved       bool           `json:"resolved"`
}

// GetPrivateKey will prompt the user for his password and return the decrypted ethereum private key
func GetPrivateKey(_forTransaction bool) (string, error) {
	// Prompting the user
	if _forTransaction {
		helpers.PrintLine("This action requires to send a transaction to the ethereum blockchain.", helpers.INFO)
	}

	// Converting the encrypted private key from hex to bytes
	encPrivateKey, err := hex.DecodeString(DitConfig.EthereumKeys.PrivateKey)
	if err != nil {
		return "", errors.New("Failed to decode private key from the config")
	}

	helpers.Printf("Please provide your password to unlock your ethereum account: ", helpers.INFO)
	ditLog.AddToLog("\n")
	iteration := 0
	for iteration <= 2 {
		password, err := terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err != nil {
			if iteration < 2 {
				iteration++
				helpers.Printf("There was an error during the password input, please try again: ", helpers.INFO)
			} else {
				return "", errors.New("There was an error during the password input - exiting after three attempts")
			}
		} else {
			decryptedPrivateKey, err := decrypt(encPrivateKey, string(password))
			if err != nil {
				if iteration < 2 {
					iteration++
					helpers.Printf("Failed to decrypt the encrypted private key - wrong password? Please try again: ", helpers.INFO)
					ditLog.AddToLog("\n")
				} else {
					return "", errors.New("There was an error during the private key decryption, probably due to a wrong password")
				}
			} else {
				if len(DitConfig.APIKey) == 0 {
					_ = createAPIKey(string(decryptedPrivateKey), true)
				}
				return string(decryptedPrivateKey), nil
			}
		}
	}
	return "", errors.New("Failed to decrypt the encrypted private key")
}

func createAPIKey(_privateKey string, _save bool) error {
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		return errors.New("Failed to convert ethereum private-key")
	}
	hash := crypto.Keccak256Hash([]byte("api"))
	apiRequestSignature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return errors.New("Failed to sign api request token with ethereum key")
	}

	DitConfig.APIKey = hex.EncodeToString(apiRequestSignature)

	if _save {
		err = Save()
		if err != nil {
			return err
		}
	}

	return nil
}

// Load will load the config and set it to the exported variable "DitConfig"
func Load() error {
	configFile, err := getRawConfig()
	if err != nil {
		return err
	}

	// Parsing the json into a public object
	err = json.Unmarshal(configFile, &DitConfig)
	if err != nil {
		errorMessage := "Failed to unmarshal JSON of config file"
		position := strings.Index(string(configFile), "version")
		if position != -1 {
			versionNumberConfig, err := strconv.Atoi(strings.TrimSpace(strings.Split(string(configFile)[position+9:], ",")[0]))
			if err != nil {
				errorMessage += " - Your config file seems to be corrupted"
			} else {
				if versionNumberConfig < Version {
					errorMessage = "Your config file is outdated, please call '" + helpers.ColorizeCommand("update") + "' to fix this"
				} else if versionNumberConfig > Version {
					errorMessage = "Your ditCLI version is older than your config file, please update your ditCLI"
				}
			}
		} else {
			errorMessage = "Your config file seems to be outdated, please call '" + helpers.ColorizeCommand("update") + "' to fix this"
		}
		return errors.New(errorMessage)
	}

	// If the config is valid, it will contain an ethereum address with a length of 42
	if len(DitConfig.EthereumKeys.Address) != 42 {
		return errors.New("Invalid config file")
	}

	return nil
}

func getRawConfig() ([]byte, error) {
	// Retrieve the home directory of the user
	usr, err := user.Current()
	if err != nil {
		return nil, errors.New("Failed to retrieve home-directory of user")
	}

	// Reading the config file
	configFile, err := ioutil.ReadFile(usr.HomeDir + "/.ditconfig")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, errors.New("Config file not found - please use '" + helpers.ColorizeCommand("setup") + "'")
		}
		return nil, errors.New("Failed to load config file")
	}

	return configFile, nil
}

// Create will create a new config file
func Create(_demoMode bool) error {
	continueCreating := true
	if strings.Contains(DitConfig.DitCoordinator, "0x") {
		helpers.PrintLine("This action will re-initialize the ditCLI, resulting in a loss of your current Ethereum keys.", helpers.WARN)
		helpers.PrintLine("Note: If you want to switch between live and demo mode, use '"+helpers.ColorizeCommand("mode")+"'", helpers.WARN)
		answer := helpers.GetUserInputChoice("Are you sure that you want to proceed?", "y", "n")
		if answer == "n" {
			continueCreating = false
		} else {
			fmt.Println()
		}
	}
	if continueCreating {
		helpers.PrintLine("Initializing the ditCLI...", helpers.INFO)
		DitConfig = ditConfig{}
		DitConfig.DemoModeActive = _demoMode

		if !DitConfig.DemoModeActive {
			helpers.PrintLine("You are initializing the ditCLI in live mode, you will be staking real xDAI.", helpers.WARN)
			helpers.PrintLine("If you just want to play around with dit, you can also use the demo mode with '"+helpers.ColorizeCommand("setup --demo")+"'", helpers.INFO)
			fmt.Println()
		} else {
			helpers.PrintLine("You are initializing the ditCLI in demo mode, feel free to play around with it!", helpers.INFO)
			helpers.PrintLine("If you want to switch to the live mode later on, you can do so with '"+helpers.ColorizeCommand("mode live")+"'", helpers.INFO)
			fmt.Println()
		}
		// Prompting the user for his choice on the ethereum key generation/importing
		answerPrivateKeySelection := helpers.GetUserInputChoice("You can either (a) sample a new ethereum private-key or (b) provide your own one", "a", "b")

		// Sample new ethereum Keys
		if answerPrivateKeySelection == "a" {
			address, privateKey, err := sampleEthereumKeys()
			if err != nil {
				return err
			}
			DitConfig.EthereumKeys.PrivateKey = privateKey
			DitConfig.EthereumKeys.Address = address
		} else {
			// Import existing ones, prompting the user for input
			answerPrivateKeyInput := helpers.GetUserInput("Please provide a hex-formatted ethereum private-key")
			ditLog.RemoveLastLine("Please provide a hex-formatted ethereum private-key:")
			if len(answerPrivateKeyInput) == 64 || len(answerPrivateKeyInput) == 66 {
				// Remove possible "0x" at the beginning
				if strings.Contains(answerPrivateKeyInput, "0x") && len(answerPrivateKeyInput) == 66 {
					answerPrivateKeyInput = answerPrivateKeyInput[2:]
				}
				// Import the ethereum private key
				address, privateKey, err := importEthereumKey(answerPrivateKeyInput)
				if err != nil {
					return err
				}

				DitConfig.EthereumKeys.PrivateKey = privateKey
				DitConfig.EthereumKeys.Address = address
			} else {
				return errors.New("Invalid ethereum private-key")
			}
		}

		err := createAPIKey(DitConfig.EthereumKeys.PrivateKey, false)
		if err != nil {
			return err
		}

		// Prompting the user to set a password for the private keys encryption
		var password []byte
		keepAsking := true
		for keepAsking {
			helpers.Printf("Please provide a password to encrypt your private key: ", helpers.INFO)
			var err error
			password, err = terminal.ReadPassword(0)
			fmt.Printf("\n")
			if err != nil {
				return errors.New("Failed to retrieve password")
			}

			// Repeating the password to make sure that there are no typos
			helpers.Printf("Please repeat your password: ", helpers.INFO)
			passwordAgain, err := terminal.ReadPassword(0)
			fmt.Printf("\n")
			if err != nil {
				return errors.New("Failed to retrieve password")
			}

			// If passwords don't match or are empty
			if string(passwordAgain) != string(password) {
				helpers.PrintLine("Passwords didn't match - try again!", helpers.WARN)
			} else if len(password) == 0 {
				helpers.PrintLine("Password can't be empty - try again!", helpers.WARN)
			} else {
				// Stop if nothing of the above is true
				keepAsking = false
			}
		}

		// Encrypt the private keys with the password
		encryptedPrivateKey, err := encrypt([]byte(DitConfig.EthereumKeys.PrivateKey), string(password))
		if err != nil {
			return errors.New("Failed to encrypt ethereum private-key")
		}

		DitConfig.EthereumKeys.PrivateKey = hex.EncodeToString(encryptedPrivateKey)
		DitConfig.Version = Version
		DitConfig.LiveRepositories = make(map[string]*Repository)
		DitConfig.DemoRepositories = make(map[string]*Repository)

		// Write the config to the file
		err = Save()
		if err != nil {
			return err
		}

		helpers.PrintLine("Initialization successfull", helpers.INFO)
		helpers.PrintLine("Your Ethereum Address is: "+DitConfig.EthereumKeys.Address, helpers.INFO)
		return nil
	}

	return errors.New("Cancelling setup due to users choice")
}

// Update will migrate the current key-pair after a ditCLI update that
func Update(_liveDitCoordinator string, _demoDitCoordinator string) (bool, error) {
	didUpdate := false
	// Retrieving the config file
	configFile, err := getRawConfig()
	if err != nil {
		return false, err
	}

	// Parsing the config to detect parsing errors
	err = json.Unmarshal(configFile, &DitConfig)
	if err != nil && strings.Contains(err.Error(), "cannot unmarshal") {
		for _, repository := range DitConfig.LiveRepositories {
			for _, vote := range repository.ActiveVotes {
				// If there was a parsing error and a live vote is still not resolved, the user is warned before updating
				if !vote.Resolved {
					helpers.PrintLine("You have ongoing or unfinalized votes that might be unresolvable with the new ditCLI version after an update!", helpers.WARN)
					answer := helpers.GetUserInputChoice("Are you sure that you want to proceed?", "y", "n")
					if answer == "n" {
						return false, errors.New("Cancelling update due to users choice")
					}
					break
				}
			}
		}
	}

	if DitConfig.Version < Version {
		helpers.PrintLine("Updating the ditCLIs' config...", helpers.INFO)
		didUpdate = true
		newDitConfig := ditConfig{}
		newDitConfig.DemoModeActive = DitConfig.DemoModeActive
		newDitConfig.DitCoordinator = DitConfig.DitCoordinator
		newDitConfig.EthereumKeys = DitConfig.EthereumKeys
		newDitConfig.LiveRepositories = make(map[string]*Repository)
		newDitConfig.DemoRepositories = make(map[string]*Repository)
		newDitConfig.Version = Version

		DitConfig = newDitConfig
		// Write the config to the file
		err := Save()
		if err != nil {
			return false, err
		}
	}
	if (DitConfig.DemoModeActive && DitConfig.DitCoordinator != _demoDitCoordinator) || (!DitConfig.DemoModeActive && DitConfig.DitCoordinator != _liveDitCoordinator) {
		helpers.PrintLine("Updating the ditCoordinator...", helpers.INFO)
		didUpdate = true
		if !DitConfig.DemoModeActive {
			DitConfig.DitCoordinator = _liveDitCoordinator
		} else {
			DitConfig.DitCoordinator = _demoDitCoordinator
		}
	}
	if didUpdate {
		helpers.PrintLine("Update of the config was successful", helpers.INFO)
		return true, nil
	}
	helpers.PrintLine("Your ditCLI is already up to date", helpers.INFO)
	return false, nil
}

// Save will write the current config object to the file
func Save() error {
	// Convert the config object to JSON
	jsonBytes, err := json.MarshalIndent(DitConfig, "", "    ")
	if err != nil {
		return errors.New("Failed to marshal JSON of config file")
	}

	// Retrieve the home directory of the user
	usr, err := user.Current()
	if err != nil {
		return errors.New("Failed to retrieve home-directory of user")
	}

	// Write the above into the config file
	err = ioutil.WriteFile(usr.HomeDir+"/.ditconfig", jsonBytes, 0644)
	if err != nil {
		return errors.New("Failed to write config file")
	}

	return nil
}

// importEthereumKey will return the private key and the address of an imported private key
func importEthereumKey(privateKey string) (string, string, error) {
	helpers.PrintLine("Importing ethereum key...", helpers.INFO)

	// Converting the private key string into a private key object
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", "", errors.New("Failed to import ethereum keys")
	}

	// Calculating the address based on the privat key object
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()

	// Converting the private key to string
	privateKey = hex.EncodeToString(key.D.Bytes())

	return address, privateKey, err
}

func sampleEthereumKeys() (string, string, error) {
	helpers.PrintLine("Sampling ethereum key...", helpers.INFO)

	// Sampling a new private key
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", errors.New("Failed to generate ethereum keys")
	}

	// Calculating the address based on the privat key object
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()

	// Converting the private key to string
	privateKey := hex.EncodeToString(key.D.Bytes())

	return address, privateKey, err
}

// from: https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// from: https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, err
}

func createHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))[0:32]
}
