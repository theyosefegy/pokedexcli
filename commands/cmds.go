package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func()
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Sends this message.",
			Callback:    Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Turns off the Pokedex.",
			Callback:    Exit,
		},
	}
}
