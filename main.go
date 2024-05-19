package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joshua-lucas/pokedexcli/internal/repl"
)

var commands map[string]repl.CliCommand

func init() {
	commands = repl.GetCommands()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {

		input := repl.SanitizeInput(scanner.Text())

		for key, value := range commands {

			if input == key {
				value.Callback()
			}
		}

		fmt.Print("\nPokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
