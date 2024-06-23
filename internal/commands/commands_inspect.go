package commands

import (
	"fmt"
)

func commandInspect(config *Config, pokemon string) error {
	if pokemon == "" {
		fmt.Print("\nNeed a Pokemon name to look it up in your Pokedex.\n\n")
		return nil
	}
	p, ok := config.Pokedex.Dex[pokemon]
	if !ok {
		fmt.Printf("\n%v hasn't been caught and recorded in your Pokedex yet.\n\n", pokemon)
		return nil
	}
	p.PrintDetails()
	return nil
}
