package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"dit/helpers"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ssh/terminal"
)

var DemoMode = true

type ditConfig struct {
	DitCoordinator string         `json:"dit_coordinator"`
	KNWVoting      string         `json:"knw_voting"`
	KNWToken       string         `json:"knw_token"`
	EthereumKeys   ethereumKeys   `json:"ethereum_keys"`
	DemoKeys       []ethereumKeys `json:"demo_keys,omitempty"`
	Repositories   []Repository   `json:"repositories"`
}

type ethereumKeys struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

// Repository struct, exported since its used in the ethereum package for new repositories
type Repository struct {
	Name            string       `json:"name"`
	ContractAddress string       `json:"contract_address"`
	KnowledgeLabels []string     `json:"knowledge_labels"`
	ActiveVotes     []ActiveVote `json:"active_votes"`
}

// ActiveVote struct, exported since its used in the ethereum package for new votes
type ActiveVote struct {
	DemoVoter      int    `json:"demo_voter"`
	ID             int    `json:"id"`
	KNWVoteID      int    `json:"knw_vote_id"`
	KnowledgeLabel string `json:"knowledge_label"`
	Choice         int    `json:"choice"`
	Salt           int    `json:"salt"`
	NumTokens      int    `json:"num_tokens"`
	NumVotes       int    `json:"num_votes"`
	NumKNW         int    `json:"num_knw"`
	CommitEnd      int    `json:"commit_end"`
	RevealEnd      int    `json:"reveal_end"`
	Resolved       bool   `json:"resolved"`
}

// DitConfig is exported since it needs to be accessed from other packages all the time
var DitConfig ditConfig

var demoPassword = "alig49239kjjhsdf653"

// GetPrivateKey will prompt the user for his password and return the decrypted ethereum private key
func GetPrivateKey() (string, error) {
	// Prompting the user
	helpers.PrintLine("This action requires to send a transaction to the ethereum blockchain.", 0)
	helpers.Printf("Please provide your password to unlock your ethereum account: ", 0)
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		return "", errors.New("Failed to retrieve password")
	}

	// Converting the encrypted private key from hex to bytes
	encPrivateKey, err := hex.DecodeString(DitConfig.EthereumKeys.PrivateKey)
	if err != nil {
		return "", errors.New("Failed to decode private key from config")
	}

	// Decrypting the private key
	decryptedPrivateKey, err := decrypt(encPrivateKey, string(password))
	if err != nil {
		return "", errors.New("Failed to decrypt the encrypted private key - wrong password?")
	}

	return string(decryptedPrivateKey), nil
}

// GetDemoPrivateKeys will prompt the user for his password and return the decrypted ethereum private key
func GetDemoPrivateKeys(_num int) (string, error) {
	// Converting the encrypted private key from hex to bytes
	encPrivateKey, err := hex.DecodeString(DitConfig.DemoKeys[_num].PrivateKey)
	if err != nil {
		return "", errors.New("Failed to decode demo private keys from config")
	}
	// Decrypting the private key
	decryptedPrivateKey, err := decrypt(encPrivateKey, demoPassword)
	if err != nil {
		return "", errors.New("Failed to decrypt the encrypted demo private keys")
	}

	return string(decryptedPrivateKey), nil
}

// Load will load the config and set it to the exported variable "DitConfig"
func Load() error {
	// Retrieve the home directory of the user
	usr, err := user.Current()
	if err != nil {
		return errors.New("Failed to retrieve home-directory of user")
	}

	// Reading the config file
	configFile, err := ioutil.ReadFile(usr.HomeDir + "/.ditconfig")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return errors.New("Config file not found - please use 'dit setup'")
		} else {
			return errors.New("Failed to load config file")
		}
	}

	// Parsing the json into a public object
	err = json.Unmarshal(configFile, &DitConfig)
	if err != nil {
		return errors.New("Failed to unmarshal JSON of config file")
	}

	// If the config is valid, it will contain an ethereum address with a length of 42
	if len(DitConfig.EthereumKeys.Address) != 42 {
		return errors.New("Invalid config file")
	}

	return nil
}

// Create will create a new config file
func Create() error {
	helpers.PrintLine("Initializing the ditClient...", 0)

	if !DemoMode {
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
	} else {
		helpers.PrintLine("Pre-funded private key was chosen due to demo mode being active", 3)
		DitConfig.EthereumKeys.PrivateKey = "0EFF81044DE7940865C6ED0039D9604FD9C7348E43168945ABCA5340120181CF"
		DitConfig.EthereumKeys.Address = "0x1b0eFa0068Fd403D293048F8758aC153DB252A11"

		DitConfig.DemoKeys = make([]ethereumKeys, 3)

		// Encrypt the private keys with the password
		encryptedPrivateKey1, _ := encrypt([]byte("B7D74145BBDD4ED8F6A2467A95C584E8C9AFFA2AFC956427C1517DA24C05C1A7"), demoPassword)
		encryptedPrivateKey2, _ := encrypt([]byte("7F8480337E77D8B241942F069DBDD58F0A7CCC9F9FAF62456F1964BC4FAAB15B"), demoPassword)
		encryptedPrivateKey3, _ := encrypt([]byte("E909852AEB18AD19ABCA2AB5CAC121FBC9BF323317312C8E8B642E348F7DE420"), demoPassword)

		DitConfig.DemoKeys[0].PrivateKey = hex.EncodeToString(encryptedPrivateKey1)
		DitConfig.DemoKeys[0].Address = "0x1e9f642f0bCa054A743fB117E91aa05E9b842F98"
		DitConfig.DemoKeys[1].PrivateKey = hex.EncodeToString(encryptedPrivateKey2)
		DitConfig.DemoKeys[1].Address = "0xCa5b5488DFB2e6E37a6B572D7Fb5E686806a1503"
		DitConfig.DemoKeys[2].PrivateKey = hex.EncodeToString(encryptedPrivateKey3)
		DitConfig.DemoKeys[2].Address = "0x125Ba426110FBc83980cc716a63ed4B6229Ec0c0"

	}

	// Prompting the user to set a password for the private keys encryption
	var password []byte
	keepAsking := true
	for keepAsking {
		helpers.Printf("Please provide a password to encrypt your private key: ", 0)
		var err error
		password, err = terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err != nil {
			return errors.New("Failed to retrieve password")
		}

		// Repeating the password to make sure that there are no typos
		helpers.Printf("Please repeat your password: ", 0)
		passwordAgain, err := terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err != nil {
			return errors.New("Failed to retrieve password")
		}

		// If passwords don't match or are empty
		if string(passwordAgain) != string(password) {
			helpers.PrintLine("Passwords didn't match - try again!", 1)
		} else if len(password) == 0 {
			helpers.PrintLine("Password can't be empty - try again!", 1)
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
	DitConfig.Repositories = make([]Repository, 0)

	// Write the config to the file
	err = Save()
	if err != nil {
		return err
	}

	helpers.PrintLine("Initialization successfull", 0)

	return nil
}

// Save will write the current config object to the file
func Save() error {
	// Convert the config object to JSON
	jsonBytes, err := json.Marshal(DitConfig)
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
	helpers.PrintLine("Importing ethereum key...", 0)

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
	helpers.PrintLine("Sampling ethereum key...", 0)

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
