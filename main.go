package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func cleanInput(text string) []string {
    return strings.Fields(text)
}

func main() {
    fmt.Print("Pokedex > ")
    userInput := bufio.NewScanner(os.Stdin)
    for userInput.Scan() {
        words := strings.Fields(strings.ToLower(userInput.Text()))
        var command string
        if len(words) > 0 {
            command = words[0]
            fmt.Print("Your command was: ", command)
            continue
        }
        fmt.Print("Pokedex > ")
    }
    if err := userInput.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
    }
}
