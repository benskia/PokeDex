package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
)

func commandMap(config *Config, cache *cache.Cache, _ string) error {
	if config.Next == nil {
		fmt.Println("Already at the last page of locations.")
		return nil
	}
	locations, err := api.RequestLocationAreas(config.Next, cache)
	if err != nil {
		fmt.Println(err)
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
