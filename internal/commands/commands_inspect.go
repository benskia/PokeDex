package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

func commandInspect(_ *Config, _ *cache.Cache, pokemon string, pokedex *dex.Pokedex) error {
	if pokemon == "" {
		fmt.Print("\nNeed a Pokemon name to look it up in your Pokedex.\n\n")
		return nil
	}
	p, ok := pokedex.Dex[pokemon]
	if !ok {
		fmt.Printf("\n%v hasn't been caught and recorded in your Pokedex yet.\n\n", pokemon)
		return nil
	}
	p.PrintDetails()
	return nil
}
