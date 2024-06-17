package commands

import (
	"fmt"
	"math/rand"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
)

const MaxCatchRate int = 255

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
	captureRate := pokemonDetails.CaptureRate
	rand.
	return nil
}
