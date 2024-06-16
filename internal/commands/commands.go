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
	if config.Next == nil {
		fmt.Println("Already at the last page of locations.")
		return nil
	}
	fmt.Println("Querying next list of locations...")
	locations, err := api.RequestLocationAreas(config.Next)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Caching Next and Prev endpoints...")
	config.Next = locations.Next
	config.Prev = locations.Prev

	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("Name: %v\t\tURL: %v\n", location.Name, location.URL)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.Prev == nil {
		fmt.Println("Already at the first page of locations.")
		return nil
	}

	fmt.Println("Querying previous list of locations...")
	locations, err := api.RequestLocationAreas(config.Prev)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Caching Next and Prev endpoints...")
	config.Next = locations.Next
	config.Prev = locations.Prev

	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("Name: %v\t\tURL: %v\n", location.Name, location.URL)
	}
	return nil
}
