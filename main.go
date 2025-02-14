package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	commands := map[string]cliCommand{
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

	for true {
		fmt.Print("Pokedex > ")
		more := scanner.Scan()
		if !more {
			fmt.Print("...")
		} else {
			text := strings.TrimSpace(scanner.Text())
			text = strings.ToLower(text)
			inputs := strings.Fields(text)

			cmd, ok := commands[inputs[0]]
			if ok {
				cmd.callback()
			} else {
				fmt.Printf("Unknown command\n")
			}
		}
	}
}


