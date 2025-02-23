package main

import (
	"fmt"
	"os"
	"github.com/Zoisit/pokedex/internal/pokeapi"
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
		return fmt.Errorf("There was an error retrieving the information for location %v: %v", conf.input, err)
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