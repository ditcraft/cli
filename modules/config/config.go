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

type ditConfig struct {
	DitCoordinator string       `json:"dit_coordinator"`
	KNWVoting      string       `json:"knw_voting"`
	KNWToken       string       `json:"knw_token"`
	EthereumKeys   ethereumKeys `json:"ethereum_keys"`
	Repositories   []Repository `json:"repositories"`
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

// GetPrivateKey will prompt the user for his password and return the decrypted ethereum private key
func GetPrivateKey() (string, error) {
	// Prompting the user
	fmt.Printf("This action requires to send a transaction to the ethereum blockchain.\nPlease provide your password to unlock your ethereum account: ")
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		fmt.Println("Failed to retrieve password")
	}

	// Converting the encrypted private key from hex to bytes
	encPrivateKey, err := hex.DecodeString(DitConfig.EthereumKeys.PrivateKey)
	if err != nil {
		fmt.Println("Failed to decode private key from config")
	}

	// Decrypting the private key
	decryptedPrivateKey, err := decrypt(encPrivateKey, string(password))
	if err != nil {
		return "", errors.New("Failed to decrypt the encrypted private key - wrong password?")
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
	fmt.Println("Initializing the ditClient...")

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

	// Prompting the user to set a password for the private keys encryption
	var password []byte
	keepAsking := true
	for keepAsking {
		fmt.Printf("Please provide a password to encrypt your private key: ")
		var err error
		password, err = terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err != nil {
			return errors.New("Failed to retrieve password")
		}

		// Repeating the password to make sure that there are no typos
		fmt.Printf("Please repeat your password: ")
		passwordAgain, err := terminal.ReadPassword(0)
		fmt.Printf("\n")
		if err != nil {
			fmt.Println("Failed to retrieve password")
		}

		// If passwords don't match or are empty
		if string(passwordAgain) != string(password) {
			fmt.Println("Passwords didn't match - try again!")
		} else if len(password) == 0 {
			fmt.Println("Password can't be empty - try again!")
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

	fmt.Println("Initialization successfull")

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
	fmt.Println("Importing ethereum key...")

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
	fmt.Println("Sampling ethereum key...")

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
