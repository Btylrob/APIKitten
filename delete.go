package main

import (
	"fmt"
)

func deleteAPIKeys() {

	ClearTerm()

	var password string
	fmt.Print(boldWhite + "Enter encryption password: ")
	fmt.Scanln(&password)

	store, err := loadKeys()
	if err != nil {
		fmt.Println(boldPink+"Error loading keys:", err)
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

		fmt.Printf(boldPink+"Is this the key you want to delete? %s (y/n): ", decryptedKey)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			updatedKeys = append(updatedKeys, encryptedKey)
		} else {
			keyDeleted = true
		}
	}

	if !keyDeleted {
		fmt.Println(boldPink + "No API key was deleted.")
		return
	}

	store.EncryptedKeys = updatedKeys
	err = saveKeys(store)
	if err != nil {
		fmt.Println(boldPink+"Error saving updated keys:", err)
		return
	}

	fmt.Printf(boldPink+"API key deleted successfully. %s", cat)
}
