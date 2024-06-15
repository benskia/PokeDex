package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Display the next 20 Pokemon locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 Pokemon locations.",
			callback:    commandMapb,
		},
	}
}

// NOTE: System Commands

func commandHelp() error {
	fmt.Print("\nAvailable commands:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf(" - Name: %v\n", cmd.name)
		fmt.Printf(" - Description: %v\n", cmd.description)
		fmt.Println()
	}
	return nil
}

func commandExit() error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}

// NOTE: Map Commands

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ApiResponse struct {
	Count   int        `json:"count"`
	Next    *string    `json:"next"`
	Prev    *string    `json:"previous"`
	Results []Location `json:"results"`
}

const endpoint string = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	Next *string
	Prev *string
}

var config = Config{
	Next: endpoint,
	Prev: nil,
}

func commandMap() error {
	fmt.Println("Getting list of locations...")

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error making GET request.")
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body.")
		return err
	}

	locations := ApiResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Struct.")
		return err
	}

	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("Name: %v\tURL: %v\n", location.Name, location.URL)
	}
	return nil
}

func commandMapb() error {
	return nil
}
