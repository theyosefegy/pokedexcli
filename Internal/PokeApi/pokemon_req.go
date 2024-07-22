package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)

	if ok {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	// Creating a new HTTP request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// The 'Do' method sends a HTTP request and returns HTTP response.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Status Code
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("connection lost, status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
