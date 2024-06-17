package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

func commandExplore(_ *Config, cache *cache.Cache, area string, _ *dex.Pokedex) error {
	if area == "" {
		fmt.Println("Need an area name to explore (explore canalave-city-area).")
		return nil
	}
	endpoint := new(string)
	*endpoint = api.LocationAreaEndpoint + area
	locationArea, err := api.RequestAreaDetails(endpoint, cache)
	if err != nil {
		return err
	}
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Printf("Found no Pokemon encounters in %v.\n", area)
		return nil
	}
	fmt.Println()
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
