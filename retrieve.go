package main

import (
	"fmt"
)

func retrieveAPIKeys() {

	ClearTerm()

	store, _ := loadKeys()

	var password string
	fmt.Print(boldWhite + "Enter decryption password: ")
	fmt.Scanln(&password)

	ClearTerm()
	fmt.Println(boldWhite + "Decrypted API Keys:")
	for _, encryptedKey := range store.EncryptedKeys {
		decryptedKey, err := decrypt(encryptedKey, password)
		if err != nil {
			fmt.Println(boldPink+"Locked", locked)
		} else {
			fmt.Println(decryptedKey)
		}
	}
}
