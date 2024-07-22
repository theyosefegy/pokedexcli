package main

import (
	"time"

	pokeapi "github.com/theyosefegy/pokedexcli/Internal/PokeApi"
	commands "github.com/theyosefegy/pokedexcli/commands"
)

func main() {

	cfg := commands.Config{
		PokeClient:     pokeapi.NewClient(time.Hour),
		CaughtPokemons: map[string]pokeapi.Pokemon{},
	}

	startRepl(&cfg)
}
