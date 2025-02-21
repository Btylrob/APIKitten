package main

import "fmt"

func StoreAPI() {

	ClearTerm()

	var apiKey, password string

	fmt.Print(boldWhite + "Enter API key: ")
	fmt.Scanln(&apiKey)
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
