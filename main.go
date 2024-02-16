package main

import (
	"fmt"
	"os"

	"github.com/rt2013G/repl-pokedex/api"
)

func printCommands() {
	fmt.Println("-- list of commands: --")
	fmt.Println("	help - displays a list of commands")
	fmt.Println("	exit - exit the application")
	fmt.Println("	map - expore the next map page")
	fmt.Println("	mapb - expore the previous map page")
}

func main() {
	locConfig := api.LocConfig{}
	client := api.CreateHttpClient()

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
		case "map":
			_, err := client.LocationAreaRequest(&locConfig, true)
			if err != nil {
				fmt.Println("error")
				continue
			}
		case "mapb":
			_, err := client.LocationAreaRequest(&locConfig, false)
			if err != nil {
				fmt.Println("error")
				continue
			}
		default:
			fmt.Println("input not valid, use \"help\" for a list of commands")
		}
	}
}
