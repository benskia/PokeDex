package commands

import (
	"fmt"
	"os"
)

func commandExit(config *Config, args ...string) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}
