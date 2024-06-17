package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

func commandPokedex(_ *Config, _ *cache.Cache, _ string, pokedex *dex.Pokedex) error {
	fmt.Println("\nYour Pokedex:")
	for name := range pokedex.Dex {
		fmt.Printf("  - %v\n", name)
	}
	fmt.Println()
	return nil
}
