package commands

import (
	"fmt"
	"os"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
)

type Config struct {
	Next *string
	Prev *string
}

func NewConfig() Config {
	next := new(string)
	*next = api.Endpoint + "?offset=0&limit=20"
	return Config{
		Next: next,
		Prev: nil,
	}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, *cache.Cache, string) error
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
		"explore": {
			Name:        "explore",
			Description: "Display which Pokemon can be found at a location.",
			Callback:    commandExplore,
		},
	}
}
