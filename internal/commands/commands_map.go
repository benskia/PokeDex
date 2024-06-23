package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
	"github.com/benskia/PokeDex/internal/pokeapi"
)

func commandMap(config *Config, cache *cache.Cache, _ string, _ *dex.Pokedex) error {
	if config.Next == nil {
		fmt.Println("Already at the last page of locations.")
		return nil
	}
	locations, err := pokeapi.RequestLocationAreas(config.Next, cache)
	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Prev = locations.Prev
	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	fmt.Println()
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()
	return nil
}
