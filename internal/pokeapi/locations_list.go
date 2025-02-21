package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c* Client) GetPokemon(pokemonName string) (Pokemon, error) {
    url := baseUrl + "/pokemon/" + pokemonName

    cachedData, found := c.cache.Get(url)
    if found {
        cachedPokemonCaught := Pokemon{}
        err := json.Unmarshal(cachedData, &cachedPokemonCaught)
        if err != nil {
            return Pokemon{}, err
        }
        return cachedPokemonCaught, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return Pokemon{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }

    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return Pokemon{}, err
    }

    pokemonResponseData := Pokemon{}
    err = json.Unmarshal(data, &pokemonResponseData)
    if err != nil {
        return Pokemon{}, err
    }

    c.cache.Add(url, data)

    return pokemonResponseData, nil

}


func (c* Client) ExploreLocationArea(locationName string) (RespPokemonFromALocationArea, error) {
    url := baseUrl + "/location-area/" + locationName

    cachedData, found := c.cache.Get(url)
    if found {
        cachedLocationArea := RespPokemonFromALocationArea{}
        err := json.Unmarshal(cachedData, &cachedLocationArea)
        if err != nil {
            return RespPokemonFromALocationArea{}, err
        }
        return cachedLocationArea, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespPokemonFromALocationArea{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespPokemonFromALocationArea{}, err
    }

    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespPokemonFromALocationArea{}, err
    }

    locationsAreaResonse := RespPokemonFromALocationArea{}
    err = json.Unmarshal(data, &locationsAreaResonse)
    if err != nil {
        return RespPokemonFromALocationArea{}, err
    }

    c.cache.Add(url, data)

    return locationsAreaResonse, nil

}


func (c* Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {

    url := baseUrl + "/location-area"
    if pageUrl != nil {
        url = *pageUrl
    }

    cachedData, found := c.cache.Get(url)
    if found {
        fmt.Println("Serving from cache")
        cachedLocations := RespShallowLocations{}
        err := json.Unmarshal(cachedData, &cachedLocations)
        if err != nil {
            return RespShallowLocations{}, err
        }
        return cachedLocations, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespShallowLocations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowLocations{}, err
    }

    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespShallowLocations{}, err
    }

    locationsResponse := RespShallowLocations{}
    err = json.Unmarshal(data, &locationsResponse)
    if err != nil {
        return RespShallowLocations{}, err
    }
    c.cache.Add(url, data)

    return locationsResponse, nil
}
