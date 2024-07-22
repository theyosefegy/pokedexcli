package commands

import (
	"fmt"
)

func Help(cfg *Config, args ...string) error {
	availCommands := GetCommands()

	info := ""

	for _, cmd := range availCommands {
		info += fmt.Sprintf("\t- %s: %s\n", cmd.Name, cmd.Description)
	}

	helpText := `
--------------------------------------------
    Welcome to the Pokedex help command!	
--------------------------------------------
    Here's your available commands:		
--------------------------------------------
%s
--------------------------------------------
`

	fmt.Printf(helpText, info)
	return nil
}
