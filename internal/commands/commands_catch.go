package commands

import (
	"fmt"
	"math/rand"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
)

const maxCatchRate int = 255

func commandCatch(_ *Config, cache *cache.Cache, pokemon string) error {
	if pokemon == "" {
		fmt.Println("Need a Pokemon name to try catching it.")
		return nil
	}
	endpoint := new(string)
	*endpoint = api.PokemonSpeciesEndpoint + pokemon
	pokemonDetails, err := api.RequestPokemonDetails(endpoint, cache)
	if err != nil {
		return err
	}
	if rand.Intn(maxCatchRate) <= pokemonDetails.CaptureRate {
		fmt.Println("Captured!")
	} else {
		fmt.Println("Failed capture!")
	}
	return nil
}
