package main

import (
	"fmt"
)

func help() {

	ClearTerm()
	fmt.Println(boldPink + `
	API Kitten - Secure API Storage & Retrieval

Commands:
    -s, --store       Encrypt and store a new API key
  	-r, --retrieve    Decrypt and display stored API keys
	-d, --delete 	  Deletes stored API key
  	-l, --list        Show encrypted API keys
  	-h, --help        Display this help menu
	-v, --version     Display version

Notes:
- API keys are encrypted using AES-256-GCM.
- You must provide the correct password to decrypt keys.
- Encrypted keys are stored in 'api_keys.json'.

For more information, visit: https://github.com/Btylrob/APIKitten
	` + resetColor)
}
