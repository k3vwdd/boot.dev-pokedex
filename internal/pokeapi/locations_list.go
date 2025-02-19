package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


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
