package commands

import (
	"fmt"
	"time"

	"github.com/benskia/PokeDex/internal/pokeapi"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Print("\nNeed an area name to explore (explore canalave-city-area).\n\n")
		return nil
	}
	if len(args) > 1 && args[1] == "report" {
		start := time.Now()
		defer reportTimeDelta(start)
	}
	area := args[0]
	endpoint := new(string)
	*endpoint = pokeapi.LocationAreaEndpoint + area
	locationArea, err := pokeapi.RequestAreaDetails(endpoint, config.Cache)
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
