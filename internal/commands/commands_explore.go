package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
	"github.com/benskia/PokeDex/internal/pokeapi"
)

func commandExplore(_ *Config, cache *cache.Cache, area string, _ *dex.Pokedex) error {
	if area == "" {
		fmt.Print("\nNeed an area name to explore (explore canalave-city-area).\n\n")
		return nil
	}
	endpoint := new(string)
	*endpoint = pokeapi.LocationAreaEndpoint + area
	locationArea, err := pokeapi.RequestAreaDetails(endpoint, cache)
	if err != nil {
		return err
	}
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Printf("\nFound no Pokemon encounters in %v.\n\n", area)
		return nil
	}
	fmt.Println()
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
