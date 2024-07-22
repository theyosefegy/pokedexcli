package commands

import (
	"fmt"
	"strings"
)

func Help(cfg *Config, args ...string) error {
	availCommands := GetCommands()

	var info strings.Builder

	for _, cmd := range availCommands {
		info.WriteString(fmt.Sprintf("\t- %s: %s\n", cmd.Name, cmd.Description))

		if len(cmd.Aliases) > 0 {
			info.WriteString(fmt.Sprintf("\t  Aliases: %s\n", strings.Join(cmd.Aliases, ", ")))
		}
	}

	helpText := `
----------------------------------------------------
    	Welcome to the Pokedex help command!	
----------------------------------------------------
    	Here are your available commands:		
----------------------------------------------------
%s
----------------------------------------------------
`

	fmt.Printf(helpText, info.String())
	return nil
}
