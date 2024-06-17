package commands

import (
	"fmt"
	"math/rand"

	"github.com/benskia/PokeDex/internal/api"
	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

// This is based off Rattata's base_experience of 51 and capture rate of ~100% (255 / 255)
const maxCatchValue int = 50

func commandCatch(_ *Config, cache *cache.Cache, pokemon string, pokedex *dex.Pokedex) error {
	if pokemon == "" {
		fmt.Print("\nNeed a Pokemon name to try catching it.\n\n")
		return nil
	}
	if _, ok := pokedex.Dex[pokemon]; ok {
		fmt.Printf("\n%v has already been captured and recorded in your Pokedex.\n\n", pokemon)
		return nil
	}
	fmt.Printf("\nThrowing a Pokeball at %v...\n", pokemon)
	endpoint := new(string)
	*endpoint = api.PokemonEndpoint + pokemon
	pd, err := api.RequestPokemonDetails(endpoint, cache)
	if err != nil {
		return err
	}
	if rand.Intn(pd.BaseExperience) <= maxCatchValue {
		numTypes := len(pd.Types)
		pokemonTypes := make([]string, numTypes)
		for i, t := range pd.Types {
			pokemonTypes[i] = t.Type.Name
		}
		pokedex.Dex[pokemon] = dex.Pokemon{
			Name:   pd.Name,
			Height: pd.Height,
			Weight: pd.Weight,
			Stats: struct {
				HP             int
				Attack         int
				Defense        int
				SpecialAttack  int
				SpecialDefense int
				Speed          int
			}{
				HP:             pd.Stats[0].BaseStat,
				Attack:         pd.Stats[1].BaseStat,
				Defense:        pd.Stats[2].BaseStat,
				SpecialAttack:  pd.Stats[3].BaseStat,
				SpecialDefense: pd.Stats[4].BaseStat,
				Speed:          pd.Stats[5].BaseStat,
			},
			Types: pokemonTypes,
		}
		fmt.Printf("\nCaptured %v!\n\n", pokemon)
	} else {
		fmt.Printf("\nFailed to capture %v!\n\n", pokemon)
	}
	return nil
}
