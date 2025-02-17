package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
)

// list locations

func (c* Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
    url := baseUrl + "/location-area"
    if pageUrl != nil {
        url = *pageUrl
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

    return locationsResponse, nil
}
