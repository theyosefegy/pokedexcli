package commands

import pokeapi "github.com/theyosefegy/pokedexcli/Internal/PokeApi"

type Config struct {
	PokeClient          pokeapi.Client
	NextLocationAreaURL *string
	PreLocationAreaURL  *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

// A 'GetCommands' function that returns a map (string -> command structure)
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
		"map": {
			Name:        "map",
			Description: "Shows the next available maps.",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous available maps.",
			Callback:    Mapb,
		},
	}
}
