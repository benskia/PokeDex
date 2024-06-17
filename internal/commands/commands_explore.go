package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
)

func commandExplore(config *Config, cache *cache.Cache, area string) error {
	if area == "" {
		fmt.Println("Need an area name to explore (explore canalave-city-area).")
		return nil
	}
	endpoint := new(string)
	*endpoint = api.Endpoint + area
	fmt.Printf("Querying location-areas for %v...", area)
	locations, err := api.RequestAreaDetails(endpoint, cache)
	if err != nil {
		fmt.Println(err)
		return err
	}
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
