package pokeapi

import (
    "net/http"
    "time"
)

type Client struct {
    httpClient  http.Client
}

// Construct a new http Client
// Remember functions with a capital are exported functions
func NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}
