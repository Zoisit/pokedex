package main

import (
	"fmt"
	"os"
	"github.com/Zoisit/pokedex/internal/pokeapi"
	"math/rand"
)


type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) (error)
}

type config struct {
	next     *string 
	previous *string  
	input *string
	pokedex map[string]pokeapi.PokemonInfo
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	commands := getCommands()
	
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for c := range commands {
		fmt.Printf("%s: %s\n", commands[c].name, commands[c].description)
	}

	return nil
}

func commandMap(conf *config) error {
	if conf.next == nil {
		fmt.Println("You're on the last page of the location areas.")
	} else {
		location_areas, err := pokeapi.GetLocationAreas(*conf.next)

		if err != nil {
			return fmt.Errorf("There was an error retrieving the location areas: %v", err)
		}

		conf.next = location_areas.Next
		conf.previous = location_areas.Previous

		for _, la := range location_areas.Results {
			fmt.Println(la.Name)
		}	
	}

	return nil
}

func commandMapBack(conf *config) error {
	if conf.previous == nil {
		fmt.Println("You're on the first page of the location areas.")
	} else {
		location_areas, err := pokeapi.GetLocationAreas(*conf.previous)

		if err != nil {
			return fmt.Errorf("There was an error retrieving the location areas: %v", err)
		}

		conf.next = location_areas.Next
		conf.previous = location_areas.Previous

		for _, la := range location_areas.Results {
			fmt.Println(la.Name)
		}	
	}

	return nil
}

func commandExplore(conf *config) error {
	if conf.input == nil {
		return fmt.Errorf("Please specify a location to explore. Locations are provided by the 'map' and 'mapb' commands.")
	}
	location, err := pokeapi.GetLocationInfo(*conf.input)

	if err != nil {
		return fmt.Errorf("There was an error retrieving the information for location %v: %v", *conf.input, err)
	}

	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found.")
	} else {
		fmt.Println("Found Pokemon:")
		for _, pe := range location.PokemonEncounters {
			fmt.Println(" - " + pe.Pokemon.Name)
		}
	}	

	return nil
}

func commandCatch(conf *config) error {
	if conf.input == nil {
		return fmt.Errorf("Please specify a Pokemon to catch. Pokémon can be seen with the explore command.")
	}

	pokemon, err := pokeapi.GetPokemonInfo(*conf.input)

	if err != nil {
		return fmt.Errorf("There was an error retrieving the information for Pokémon %v: %v", *conf.input, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	//research: Blissey gives the highest amount of base exp, according to the api it's 635 (608 accodding to research), Sunkern is lowest with 36 - TODO: query all pokemon once and save the data 
	chance := float32(rand.Intn(pokemon.BaseExperience)) / float32(pokemon.BaseExperience) //TODO: might as well do completely random
	if chance * 10 >= 5 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped...\n", pokemon.Name)
	}

	return nil
}

func commandInspect(conf *config) error {
	if conf.input == nil {
		return fmt.Errorf("Please specify a Pokémon in your Pokédex to inspect.")
	}

	pokemon, ok := conf.pokedex[*conf.input] 
	if !ok {
		return fmt.Errorf("You have not caught a %s", *conf.input)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(conf *config) error {
	fmt.Println("Your Pokedex:")
	
	for k, _ := range conf.pokedex {
		fmt.Printf("- %s\n", k)
	}

	if len(conf.pokedex) == 0 {
		fmt.Println("You did not ctach any Pokémon yet.")
	}

	return nil
}