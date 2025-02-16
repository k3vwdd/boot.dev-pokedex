package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

type cliCommand struct {
    name        string
    description string
    callback    func() error
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
    }
}

func startRepl() {
    fmt.Print("Pokedex > ")
    userInput := bufio.NewScanner(os.Stdin)
    for userInput.Scan() {
        if err := userInput.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
        }
        words := cleanInput(userInput.Text())
        if len(words) == 0 {
            continue
        }
        commandName := words[0]
        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("uknown command")
            continue
        }
    }
}
