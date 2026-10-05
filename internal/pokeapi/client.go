package pokeapi

import (
	"net/http"
	"time"

	"github.com/JacksonS25/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(5 * time.Minute),
	}
}
