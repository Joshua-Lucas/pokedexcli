package repl

import (
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

// returns the available commands for the repl
func GetCommands() map[string]CliCommand {

	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

// Returns an error there is an issue getting the help document.
func commandHelp() error {

	commands := GetCommands()

	fmt.Print("\nWelcome to the Pokedex!\n")

	fmt.Print("\nUsage:\n\n")

	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}

	return nil
}

// Returns and error if it cannot exit the program.
func commandExit() error {

	os.Exit(0)

	return nil
}

func SanitizeInput(input string) string {
	trimmedInput := strings.Trim(input, " ")

	return strings.ToLower(trimmedInput)
}
