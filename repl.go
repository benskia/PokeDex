package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
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
	fmt.Println("Helping...")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting...")
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("PokeDex > ")
	scanner.Scan()
	cmd, ok := getCommands()[scanner.Text()]
	if !ok {
		fmt.Println("Invalid command.")
		return
	}
	cmd.callback()
}

func cleanInput(str string) []string {
	lowerCaseStr := strings.ToLower(str)
	words := strings.Fields(lowerCaseStr)
	return words
}
