package commands

import (
	"fmt"
	"os"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/dex"
)

func commandExit(_ *Config, _ *cache.Cache, _ string, _ *dex.Pokedex) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}
