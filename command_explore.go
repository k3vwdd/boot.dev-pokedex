package main


import (
    "fmt"
)


func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return fmt.Errorf("explore command requires exactly one argument")
    }
    locationName := args[0]
    exploreLocation, err := cfg.pokeapiClient.ExploreLocationArea(locationName)
    if err != nil {
        return fmt.Errorf("location not found, Try again with a valid location name")
    }

    fmt.Printf("Exploring...%v\n", locationName)
    fmt.Println("Found Pokemon:")
    for _, pokemon := range exploreLocation.PokemonEncounters {
        fmt.Println("- ", pokemon.Pokemon.Name)
    }

    return nil
}
