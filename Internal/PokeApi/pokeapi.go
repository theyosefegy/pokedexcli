package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/theyosefegy/pokedexcli/Internal/PokeCache"
)

const (
	baseURL string = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}


func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
