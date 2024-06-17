package commands

import (
	"fmt"
	"os"

	"github.com/benskia/PokeDex/internal/cache"
)

func commandExit(config *Config, cache *cache.Cache, area string) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}
