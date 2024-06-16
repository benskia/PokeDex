package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

		err := cmd.Callback()
		if err != nil {
			fmt.Printf("Error executing command %v : %v", cmd.Name, err)
		}
	}
}
