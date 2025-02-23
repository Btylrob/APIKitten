package main

import (
	"fmt"
	"os"
)

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
