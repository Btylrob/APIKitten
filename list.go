package main

import (
	"fmt"
)

func listEncryptedKeys() {

	ClearTerm()

	store, _ := loadKeys()
	fmt.Println(boldWhite + "Stored Encrypted API Keys:")
	for _, key := range store.EncryptedKeys {
		fmt.Println(key)
	}
}
