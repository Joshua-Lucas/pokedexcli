package repl

import (
	"fmt"
	"os"
	"strings"

	"github.com/joshua-lucas/pokedexcli/internal/apis/pokeapi"
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
		"map": {
			Name:        "map",
			Description: "Displays next 20 location areas at a time in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous 20 location areas at a time in the Pokemon world",
			Callback:    commandMapb,
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

// Prints the next 20 map locations in the pokeman world
func commandMap() error {

	locationRes, _ := pokeapi.GetLocations(*GlobalConfig.next)

	// updated config struct with the returned next and prev url
	configUpdates := Config{
		next:     &locationRes.Next,
		previous: &locationRes.Previous,
	}

	UpdateConfig(GlobalConfig, configUpdates)

	// Loop through the location area results and print them out.
	for _, locationArea := range locationRes.Results {
		fmt.Println(locationArea.Name)
	}

	return nil

}

// Prints the previous 20 map locations
func commandMapb() error {

	if GlobalConfig.previous == nil || *GlobalConfig.previous == "" {
		fmt.Println("There are no previous locations. Run \033[3m map \033[0m command to see the next locations")
		return nil

	}

	locationRes, err := pokeapi.GetLocations(*GlobalConfig.previous)

	if err != nil {
		return err
	}
	// updated config struct with the returned next and prev url
	configUpdates := Config{
		next:     &locationRes.Next,
		previous: &locationRes.Previous,
	}

	UpdateConfig(GlobalConfig, configUpdates)

	// Loop through the location area results and print them out.
	for _, locationArea := range locationRes.Results {
		fmt.Println(locationArea.Name)
	}

	return nil

}
