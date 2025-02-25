package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/Zoisit/pokedex/internal/pokeapi"
)

func main_loop() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{}
	first_areas := "https://pokeapi.co/api/v2/location-area/"
	conf.next =  &first_areas
	conf.pokedex = make(map[string]pokeapi.PokemonInfo)
	
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputs := cleanInput(scanner.Text())

		cmd, ok := getCommands()[inputs[0]]
		if ok {
			conf.input = nil
			if len(inputs) > 1 {
				conf.input = &inputs[1]
			}
			err := cmd.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
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
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemon at the given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the given Pokémon. A succesful catch will add the Pokémon to the Pokédex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows information of the given Pokémon in the Pokédex.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the names of all Pokémon in the Pokédex. (All caught Pokémon.)",
			callback:    commandPokedex,
		},
	}
}