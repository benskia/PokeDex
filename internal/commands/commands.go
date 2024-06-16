package commands

import (
	"fmt"
	"os"

	"github.com/benskia/PokeDex/internal/api"
)

type Config struct {
	Next *string
	Prev *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 Pokemon locations.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 Pokemon locations.",
			Callback:    commandMapb,
		},
	}
}

func commandHelp(config *Config) error {
	fmt.Print("\nAvailable commands:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - Name: %v\n", cmd.Name)
		fmt.Printf(" - Description: %v\n", cmd.Description)
		fmt.Println()
	}
	return nil
}

func commandExit(config *Config) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	fmt.Println("Querying next list of locations...")
	locations, err := api.RequestLocationAreas()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Caching Next and Prev endpoints...")
	config.Next = locations.Next
	config.Prev = locations.Prev

	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("Name: %v\tURL: %v\n", location.Name, location.URL)
	}
	return nil
}

func commandMapb(config *Config) error {
	return nil
}
