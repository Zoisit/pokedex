package main

import (
	"fmt"
	"os"
)


type cliCommand struct {
	name        string
	description string
	callback    func() error
}



func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for c := range commands {
		fmt.Printf("%s: %s\n", commands[c].name, commands[c].description)
	}

	return nil
}