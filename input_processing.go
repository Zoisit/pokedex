package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main_loop() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{}
	first_areas := "https://pokeapi.co/api/v2/location-area/"
	conf.next =  &first_areas
	
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputs := cleanInput(scanner.Text())

		cmd, ok := getCommands()[inputs[0]]
		if ok {
			cmd.callback(&conf)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	inputs := strings.Fields(text)

	return inputs
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapBack,
		},
	}
}