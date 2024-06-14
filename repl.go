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
		"map": {
			name:        "map",
			description: "Display all Pokemon location names.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous Pokemon locations.",
			callback:    commandMapb,
		},
	}
}

func commandHelp() error {
	fmt.Print("\nAvailable PokeDex commands:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf(" - Name: %v\n", cmd.name)
		fmt.Printf(" - Description: %v\n", cmd.description)
		fmt.Println()
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}

func startRepl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("PokeDex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		cmd, ok := getCommands()[cleanedInput[0]]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}
		cmd.callback()
	}
}

func cleanInput(str string) []string {
	lowerCaseStr := strings.ToLower(str)
	words := strings.Fields(lowerCaseStr)
	return words
}
