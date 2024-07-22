package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func Catch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemonData, err := cfg.PokeClient.GetPokemon(pokemonName)
	handleError(err)

	const threshold = 50
	randNum := rand.Intn(pokemonData.BaseExperience)

	fmt.Println(pokemonData.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonData.Name)
	}

	cfg.CaughtPokemons[pokemonName] = pokemonData
	
	fmt.Printf("%s was caught successfuly!\n", pokemonData.Name)

	return nil
}
