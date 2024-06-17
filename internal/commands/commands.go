package commands

import (
	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

type Config struct {
	Next *string
	Prev *string
}

func NewConfig() Config {
	next := new(string)
	*next = api.LocationAreaEndpoint + "?offset=0&limit=20"
	return Config{
		Next: next,
		Prev: nil,
	}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, *cache.Cache, string, *dex.Pokedex) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Display a list of available commands.",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex CLI.",
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
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon, by name.",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Display a Pokemon's details, by name.",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display the names of all the Pokemon you've caught.",
			Callback:    commandPokedex,
		},
	}
}
