package commands

import (
	"fmt"
	"time"

	"github.com/benskia/PokeDex/internal/pokeapi"
)

func commandMapb(config *Config, args ...string) error {
	if config.Prev == nil {
		fmt.Println("Already at the first page of locations.")
		return nil
	}
	if len(args) > 0 && args[0] == "report" {
		start := time.Now()
		defer reportTimeDelta(start)
	}
	locations, err := pokeapi.RequestLocationAreas(config.Prev, config.Cache)
	if err != nil {
		return err
	}
	config.Next = locations.Next
	config.Prev = locations.Prev
	if len(locations.Results) == 0 {
		fmt.Println("No location data received.")
		return nil
	}
	fmt.Println()
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()
	return nil
}
