package commands

import (
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
)

func commandHelp(_ *Config, _ *cache.Cache, _ string) error {
	fmt.Print("\nAvailable commands:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - Name: %v\n", cmd.Name)
		fmt.Printf(" - Description: %v\n", cmd.Description)
		fmt.Println()
	}
	return nil
}
