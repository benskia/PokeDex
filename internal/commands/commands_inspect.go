package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

func commandInspect(_ *Config, _ *cache.Cache, pokemon string, pokedex *dex.Pokedex) error {
	if pokemon == "" {
		fmt.Println("Need a Pokemon name to look it up in your Pokedex.")
		return nil
	}
	p, ok := pokedex.Dex[pokemon]
	if !ok {
		fmt.Printf("%v hasn't been caught and recorded in your Pokedex yet.", pokemon)
		return nil
	}
	fmt.Println()
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Println()
	return nil
}
