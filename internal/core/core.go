package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/benskia/PokeDex/internal/cache"
	"github.com/benskia/PokeDex/internal/commands"
)

func cleanInput(str string) []string {
	if len(str) == 0 {
		return nil
	}
	lowerCaseStr := strings.ToLower(str)
	words := strings.Fields(lowerCaseStr)
	return words
}

func StartRepl() {
	fmt.Print("\nWelcome to the PokeDex!\n\n")
	config := commands.NewConfig()
	cache := cache.NewCache(time.Minute * 5)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("PokeDex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if cleanedInput == nil {
			continue
		}

		cmd, ok := commands.GetCommands()[cleanedInput[0]]
		if !ok {
			fmt.Println("Invalid command. Try 'help' for a list of commands.")
			continue
		}

		area := ""
		if len(cleanedInput) > 1 {
			area = cleanedInput[1]
		}

		err := cmd.Callback(&config, cache, area)
		if err != nil {
			fmt.Printf("Error executing command %v : %v", cmd.Name, err)
		}
	}
}
