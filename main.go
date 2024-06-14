package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("Help executed...")
	return nil
}

func commandExit() error {
	fmt.Println("Exit executed...")
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokeDex > ")
		scanner.Scan()
		cmd, ok := getCommandMap()[scanner.Text()]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}
		cmd.callback()
	}
}
