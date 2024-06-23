package commands

import (
	"fmt"
)

func commandPokedex(config *Config, _ string) error {
	fmt.Println("\nYour Pokedex:")
	for name := range config.Pokedex.Dex {
		fmt.Printf("  - %v\n", name)
	}
	fmt.Println()
	return nil
}
