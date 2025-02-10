package cli

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

var cat = "\U0001f431"

type KeyStore struct {
	EncryptedKeys []string `json:"encrypted_keys"`
}

const keyFile = "api_keys.json"

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

func storeAPIKey() {

	ClearTerm()

	var apiKey, password string

	fmt.Print("Enter API key: ")
	fmt.Scanln(&apiKey)
	fmt.Print("Enter encryption password: ")
	fmt.Scanln(&password)

	encryptedKey, err := encrypt(apiKey, password)
	if err != nil {
		fmt.Println("Error encrypting key:", err)
		return
	}

	store, _ := loadKeys()
	store.EncryptedKeys = append(store.EncryptedKeys, encryptedKey)
	err = saveKeys(store)
	if err != nil {
		fmt.Println("Error saving key:", err)
		return
	}

	fmt.Printf("API key stored securely.%s", cat)
}

func retrieveAPIKeys() {

	ClearTerm()

	var password string
	fmt.Print("Enter decryption password: ")
	fmt.Scanln(&password)

	store, _ := loadKeys()
	fmt.Println("Decrypted API Keys:")
	for _, encryptedKey := range store.EncryptedKeys {
		decryptedKey, err := decrypt(encryptedKey, password)
		if err != nil {
			fmt.Println("Error decrypting key or incorrect password.")
		} else {
			fmt.Println(decryptedKey)
		}
	}
}

func listEncryptedKeys() {

	ClearTerm()

	store, _ := loadKeys()
	fmt.Println("Stored Encrypted API Keys:")
	for _, key := range store.EncryptedKeys {
		fmt.Println(key)
	}
}

func deleteAPIKeys() {

	ClearTerm()

	var password string
	fmt.Print("Enter encryption password: ")
	fmt.Scanln(&password)

	store, err := loadKeys()
	if err != nil {
		fmt.Println("Error loading keys:", err)
		return
	}

	var updatedKeys []string
	keyDeleted := false

	// Decrypt and remove the matching key
	for _, encryptedKey := range store.EncryptedKeys {
		decryptedKey, err := decrypt(encryptedKey, password)
		if err != nil {
			updatedKeys = append(updatedKeys, encryptedKey)
			continue
		}

		fmt.Printf("Is this the key you want to delete? %s (y/n): ", decryptedKey)
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			updatedKeys = append(updatedKeys, encryptedKey)
		} else {
			keyDeleted = true
		}
	}

	if !keyDeleted {
		fmt.Println("No API key was deleted.")
		return
	}

	store.EncryptedKeys = updatedKeys
	err = saveKeys(store)
	if err != nil {
		fmt.Println("Error saving updated keys:", err)
		return
	}

	fmt.Printf("API key deleted successfully. %s", cat)
}

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

func help() {

	limeGreen := "\033[92m"
	resetColor := "\033[0m" // Reset to default terminal color

	ClearTerm()
	fmt.Println(limeGreen + `
	API Key Manager - Secure Storage & Retrieval

Usage:
  api_key_manager [command]

Commands:
  store       Encrypt and store a new API key
  retrieve    Decrypt and display stored API keys
  list        Show encrypted API keys
  help        Display this help menu
  exit        Quit the program

Examples:
  api_key_manager store
  api_key_manager retrieve
  api_key_manager list

Notes:
- API keys are encrypted using AES-256-GCM.
- You must provide the correct password to decrypt keys.
- Encrypted keys are stored in 'api_keys.json'.

For more information, visit: https://yourprojectdocs.example
	` + resetColor)
}

func Start() {

	fmt.Println(` 
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
	`)

	for {
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "-s", "--store":
			storeAPIKey()
		case "-r", "--retrieve":
			retrieveAPIKeys()
		case "-d", "--delete":
			deleteAPIKeys()
		case "-l", "--list":
			listEncryptedKeys()
		case "-h", "--help":
			help()
		case "c", "--close":
			os.Exit(0)
		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}
