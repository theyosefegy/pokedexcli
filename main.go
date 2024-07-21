package main

import (
	pokeapi "github.com/theyosefegy/pokedexcli/Internal/PokeApi"
	commands "github.com/theyosefegy/pokedexcli/commands"
)


func main() {

	cfg := commands.Config{
		PokeClient: pokeapi.NewClient(),
		NextLocationAreaURL: nil,
		PreLocationAreaURL: nil,
	}

	startRepl(&cfg)
}
