package commands

import (
	"fmt"
	"os"
)

func commandExit(_ *Config, _ string) error {
	fmt.Println("Shutting down PokeDex...")
	os.Exit(0)
	return nil
}
