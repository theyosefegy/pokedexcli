package commands

import (
	pokeapi "github.com/theyosefegy/pokedexcli/Internal/PokeApi"
)

type Config struct {
	PokeClient          pokeapi.Client
	CaughtPokemons      map[string]pokeapi.Pokemon
	NextLocationAreaURL *string
	PreLocationAreaURL  *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
	Aliases     []string
}

func handleError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// A 'GetCommands' function that returns a map (string -> command structure)
func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Sends this message.",
			Callback:    Help,
			Aliases:     []string{"h"},
		},
		"map": {
			Name:        "map",
			Description: "Shows the next available maps.",
			Callback:    Map,
			Aliases:     []string{"m"},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous available maps.",
			Callback:    Mapb,
			Aliases:     []string{"mb"},
		},
		"explore": {
			Name:        "explore",
			Description: "Displays all pokemons in the area.",
			Callback:    Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a pokemon.",
			Callback:    Catch,
		},
		"mypokemons": {
			Name:        "mypokemons",
			Description: "Displays all your pokemons.",
			Callback:    MyPokemons,
			Aliases:     []string{"pokemons", "myp"},
		},
		"inspect": {
			Name:        "inspect",
			Description: "View information about a caught pokemon",
			Callback:    Inspect,
		},
		"exit": {
			Name:        "exit",
			Description: "Turns off the Pokedex.",
			Callback:    Exit,
			Aliases:     []string{"kill", "close"},
		},
	}
}
