package main

import (
	"fmt"
    "errors"
)

func commandMapF(cfg *config, args ...string) error {
    locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResponse.Next
    cfg.previousLocationsURL = locationsResponse.Previous

    for _, location := range locationsResponse.Results {
        fmt.Println(location.Name)
    }
    return nil
}

func commandMapB(cfg *config, args ...string) error {
    if cfg.previousLocationsURL == nil {
        return errors.New("You're on the first page")
    }

    locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResponse.Next
    cfg.previousLocationsURL = locationsResponse.Previous

    for _, location := range locationsResponse.Results {
        fmt.Println(location.Name)
    }
    return nil
}




