package main

import (
	"bufio"
	"fmt"
	"os"
)

func StoreAPI() {

	ClearTerm()

	var apiKey, password, use string

	fmt.Print(boldWhite + "Enter API key: ")
	reader := bufio.NewReader(os.Stdin)
	apiKey, _ = reader.ReadString('\n')
	apiKey = apiKey[:len(apiKey)-1]

	fmt.Print(boldWhite + "Enter Use: ")
	use, _ = reader.ReadString('\n')
	use = use[:len(use)-1]

	apiKey = apiKey + " " + use

	fmt.Print(boldWhite + "Enter encryption password: ")
	fmt.Scanln(&password)

	encryptedKey, err := encrypt(apiKey, password)
	if err != nil {
		fmt.Println(boldPink+"Error encrypting key:", err)
		return
	}

	store, _ := loadKeys()
	store.EncryptedKeys = append(store.EncryptedKeys, encryptedKey)
	err = saveKeys(store)
	if err != nil {
		fmt.Println("Error saving key:", err)
		return
	}

	fmt.Printf(boldWhite+"API key stored securely.%s", cat)
}
