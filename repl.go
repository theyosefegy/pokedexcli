package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theyosefegy/pokedexcli/commands"
)

func startRepl(cfg *commands.Config) {
	for {
		fmt.Print(" > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		// Return a "cleaned" slice of the user input.
		cleaned := cleanInput(text)

		// A map that contains all available commands I created.
		availCommands := commands.GetCommands()

		command := cleaned[0]

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		// Check if the command exists, including aliases.
		cmd, isAvail := getCommand(availCommands, command)

		if !isAvail {
			fmt.Println("Invalid Command, please use the 'help' command :).")
			continue
		}

		// Calling the command function.
		err := cmd.Callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

// getCommand checks if a command or its alias exists in the available commands.
func getCommand(availCommands map[string]commands.CliCommand, input string) (commands.CliCommand, bool) {
	if cmd, exists := availCommands[input]; exists {
		return cmd, true
	}

	for _, cmd := range availCommands {
		for _, alias := range cmd.Aliases {
			if alias == input {
				return cmd, true
			}
		}
	}

	return commands.CliCommand{}, false
}
