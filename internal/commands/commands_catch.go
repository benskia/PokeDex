package commands

import (
	"fmt"
	"math/rand"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

const maxCatchRate int = 255

func commandCatch(_ *Config, cache *cache.Cache, pokemon string, pokedex *dex.Pokedex) error {
	if pokemon == "" {
		fmt.Println("Need a Pokemon name to try catching it.")
		return nil
	}
	if _, ok := pokedex.Dex[pokemon]; ok {
		fmt.Printf("%v has already been captured and recorded in your Pokedex.\n", pokemon)
		return nil
	}
	endpoint := new(string)
	*endpoint = api.PokemonSpeciesEndpoint + pokemon
	pokemonDetails, err := api.RequestPokemonDetails(endpoint, cache)
	if err != nil {
		return err
	}
	if rand.Intn(maxCatchRate) <= pokemonDetails.CaptureRate {
		fmt.Printf("Captured %v!\n", pokemon)
		pokedex.Dex[pokemon] = dex.Pokemon{Name: pokemonDetails.Name}
	} else {
		fmt.Printf("Failed to capture %v!\n", pokemon)
	}
	return nil
}
