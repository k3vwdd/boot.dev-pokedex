package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

)

func commandMapB(c *Config) error {
    var url string
    if c.Previous == "" {
        url = "https://pokeapi.co/api/v2/location-area"
    } else {
        url = c.Previous
    }
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return err
    }

    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
        return err
    }

    locations := Config{}
    err = json.Unmarshal(data, &locations)
    if err != nil {
        fmt.Println(err)
        return err
    }

    c.Results = locations.Results
    c.Next = locations.Next
    for _, val := range c.Results {
        if len(c.Results) == 0 {
            fmt.Println("No more Pokemon")
        }
        fmt.Println(val.Name)
    }
    //fmt.Println(c.Next)
    // you use decoder as an either or vs io.ReadAll, Not as both.
    // we're doing io.ReadAll then Unmarshal
    // use this below for larger JSON data streams
    //decoder := json.NewDecoder(res.Body)
    //err = decoder.Decode(&locations)
    //if err != nil {
    //    fmt.Println(err)
    //}
    return nil

}



