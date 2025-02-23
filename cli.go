package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"encoding/base64"
	"encoding/json"

	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

// emojis / term-art / color code

const (
	cat    = "\U0001f431"
	locked = "\U0001F512"
)

const (
	boldPink   = "\033[1;95m"
	resetColor = "\033[0m" //default terminal color
	boldWhite  = "\033[1;97m"
)

// version
const vers = "apikitten version 0.0.1"

type KeyStore struct {
	EncryptedKeys []string `json:"encrypted_keys"`
}

// json file store
const keyFile = "api_keys.json"

// clear term
func ClearTerm() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// version
func version() {
	ClearTerm()
	fmt.Println(vers)
}

// encrypt and decrypt
func encrypt(text, password string) (string, error) {
	key := deriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encryptedText, password string) (string, error) {

	key := deriveKey(password)
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, ciphertext := data[:12], data[12:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// cmds retrieve / store / delete / help
func deriveKey(password string) []byte {
	key := make([]byte, 32)
	copy(key, password)
	return key
}

func saveKeys(store KeyStore) error {
	file, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(store)
}

func loadKeys() (KeyStore, error) {
	var store KeyStore
	file, err := os.Open(keyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return KeyStore{}, nil
		}
		return store, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&store)
	return store, err
}

// start terminal menu
func Start() {

	fmt.Println(boldPink + ` 
	 ________  ________  ___                                       
    |\   __  \|\   __  \|\  \                                      
    \ \  \|\  \ \  \|\  \ \  \                                     
     \ \   __  \ \   ____\ \  \                                    
      \ \  \ \  \ \  \___|\ \  \                                   
       \ \__\ \__\ \__\    \ \__\                                  
        \|__|\|__|\|__|     \|__|                                                                                                           
 ___  __    ___  _________  _________  _______   ________      
|\  \|\  \ |\  \|\___   ___\\___   ___\\  ___ \ |\   ___  \    
\ \  \/  /|\ \  \|___ \  \_\|___ \  \_\ \   __/|\ \  \\ \  \   
 \ \   ___  \ \  \   \ \  \     \ \  \ \ \  \_|/_\ \  \\ \  \  
  \ \  \\ \  \ \  \   \ \  \     \ \  \ \ \  \_|\ \ \  \\ \  \ 
   \ \__\\ \__\ \__\   \ \__\     \ \__\ \ \_______\ \__\\ \__\
    \|__| \|__|\|__|    \|__|      \|__|  \|_______|\|__| \|__|
	-CLI-based encryption tool for API keys and tokens.

USAGE: 
	A CLI locker for API keys and tokens using AES-256-GCM encryption.

AUTHORS(S):
	Bran Robinson <btylrob>

COMMANDS:
  	-s, --store       Encrypt and store a new API key
  	-r, --retrieve    Decrypt and display stored API keys
	-d, --delete 	  Deletes stored API key
  	-l, --list        Show encrypted API keys
  	-h, --help        Display this help menu
	-v, --version     Display version
	-b, --back 		  Goes back to main menu
	`)

	for {
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "-s", "--store":
			StoreAPI()
		case "-r", "--retrieve":
			retrieveAPIKeys()
		case "-d", "--delete":
			deleteAPIKeys()
		case "-l", "--list":
			listEncryptedKeys()
		case "-h", "--help":
			help()
		case "-v", "--version":
			version()
		case "-c", "--close":
			os.Exit(0)
		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}
