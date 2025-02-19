package pokeapi

import (
    "net/http"
    "time"
    "github.com/k3vwdd/boot.dev-pokedex/internal/pokecache"
)

type Client struct {
    httpClient  http.Client
    cache   *pokecache.Cache
}

// Construct a new http Client
// Remember functions with a capital are exported functions
func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        cache: cache,
    }
}
