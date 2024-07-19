package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theyosefegy/pokedexcli/commands"
)

func startRepl() {
	for {
		fmt.Print(" > ")
		
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		// return a "cleaned" slice of the user input.		
		cleaned := cleanInput(text)
		
		command := cleaned[0]
		availCommands := commands.GetCommands()


		// check if the command does actually exists.
		cmd, ok := availCommands[command]

		if !ok {
			fmt.Println("Invalid Command, please use the 'help' command :).")
			continue
		}

		// Calling the command function.
		cmd.Callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}