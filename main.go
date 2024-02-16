package main

import (
	"fmt"
	"os"
)

func printCommands() {
	fmt.Println("-- list of commands: --")
	fmt.Println("	help - displays a list of commands")
	fmt.Println("	exit - exit the application")
}

func main() {
	for {
		fmt.Print("pokedex $ ")
		var input string
		fmt.Scanln(&input)

		if len(input) == 0 {
			continue
		}

		switch input {
		case "exit":
			os.Exit(0)
		case "help":
			printCommands()
		default:
			fmt.Println("input not valid, use \"help\" for a list of commands")
		}
	}
}
