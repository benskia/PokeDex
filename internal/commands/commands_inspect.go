package commands

import (
	"fmt"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Print("\nNeed a Pokemon name to look it up in your Pokedex.\n\n")
		return nil
	}
	pokemon := args[0]
	p, ok := config.Pokedex.Dex[pokemon]
	if !ok {
		fmt.Printf("\n%v hasn't been caught and recorded in your Pokedex yet.\n\n", pokemon)
		return nil
	}
	p.PrintDetails()
	return nil
}
