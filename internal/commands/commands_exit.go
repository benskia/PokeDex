package commands

import (
	"fmt"
	"os"

	"github.com/benskia/PokeDex/internal/cache"
)

func commandExit(_ *Config, _ *cache.Cache, _ string) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}
