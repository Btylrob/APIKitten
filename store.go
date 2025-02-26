package main

import (
	"bufio"
	"fmt"
	"os"
)

func StoreAPI() {

	ClearTerm()

	var apiKey, password string

	fmt.Print(boldWhite + "Enter API key: ")
	reader := bufio.NewReader(os.Stdin)
	apiKey, _ = reader.ReadString('\n')
	apiKey = apiKey[:len(apiKey)-1]

	fmt.Print(boldWhite + "Enter encryption password: ")
	password, _ = reader.ReadString('\n')
	password = password[:len(password)-1]

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
