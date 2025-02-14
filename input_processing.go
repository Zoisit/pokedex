package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main_loop() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputs := cleanInput(scanner.Text())

		cmd, ok := getCommands()[inputs[0]]
		if ok {
			cmd.callback()
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
	}
}