package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
    name        string
    description string
    callback    func(*Config) error
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

func getCommands() map[string]cliCommand {

    return map[string]cliCommand{
        "exit": {
            name:   "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "map": {
            name: "map",
            description: "Displays the names of the next 20 location areas in the Pokemon world",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "Displays the names of the previous 20 location areas in the Pokemon world",
            callback: commandMapB,
        },
    }
}

func startRepl() {
    cfg := &Config{} // global config to update for pagination
    fmt.Print("Pokedex > ")
    userInput := bufio.NewScanner(os.Stdin)
    for userInput.Scan() {
        if err := userInput.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
        }
        words := cleanInput(userInput.Text())
        if len(words) == 0 {
            fmt.Print("Pokedex > ")
            continue
        }
        commandName := words[0]
        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg) // creates a new config on each call, no memory of previous map command, create a global one to update
            if err != nil {
                fmt.Println(err)
            }
            fmt.Print("Pokedex > ")
            continue
        } else {
            fmt.Println("uknown command")
            continue
        }
    }

}
