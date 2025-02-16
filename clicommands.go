package main

import (
	"fmt"
	"os"
)


type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) (error)
}

type config struct {
	next     *string `json:"next"`
	previous *string  `json:"previous"`
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	commands := getCommands()
	
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for c := range commands {
		fmt.Printf("%s: %s\n", commands[c].name, commands[c].description)
	}

	return nil
}

func commandMap(conf *config) error {
	if conf.next == nil {
		fmt.Println("You're on the last page of the location areas.")
	} else {
		location_areas, err := getLocationAreas(*conf.next)

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
		location_areas, err := getLocationAreas(*conf.previous)

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