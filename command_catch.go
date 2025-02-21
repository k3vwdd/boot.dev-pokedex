package main

import (
	"fmt"
	"math/rand"
)


type Pokemon struct {
   name string
}

func commandCatch(cfg *config, pokemonArg ...string) error {
    if len(pokemonArg) != 1 {
        return fmt.Errorf("catch command requires a pokemon")
    }

    pokemonName := pokemonArg[0]
    getPokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
        return fmt.Errorf("pokemon not found")
    }

    pokedex := make(map[string]Pokemon)
    caughtPokemon := Pokemon{
        name: getPokemon.Name,
    }

    threshold := 30
    baseExp := getPokemon.BaseExperience
    roll := rand.Intn(baseExp)

    fmt.Printf("Throwing a Pokeball at %v...\n",pokemonName)
    if roll < threshold {
        pokedex[pokemonName] = caughtPokemon
    fmt.Printf("%v was caught!\n",caughtPokemon.name)
    } else {
    fmt.Printf("%v escaped!\n",pokemonName)
    }
    return nil
}
