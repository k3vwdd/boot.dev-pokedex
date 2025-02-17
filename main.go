package main

import (
	"time"

	"github.com/k3vwdd/boot.dev-pokedex/internal/pokeapi"
)



func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second)
    cfg := &config{
        pokeapiClient: pokeClient,
    }
    startRepl(cfg)
}
