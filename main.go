package main

import (
	"time"
	"github.com/k3vwdd/boot.dev-pokedex/internal/pokeapi"
	"github.com/k3vwdd/boot.dev-pokedex/internal/pokecache"
)



func main() {
    cache := pokecache.NewCache(5 * time.Second)
    pokeClient := pokeapi.NewClient(5 * time.Second, cache)
    cfg := &config{
        pokeapiClient: pokeClient,
    }
    startRepl(cfg)
}
