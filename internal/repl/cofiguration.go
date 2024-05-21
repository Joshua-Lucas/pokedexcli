package repl

var GlobalConfig *Config

func init() {
	GlobalConfig = newConfig()
}

type Config struct {
	next     *string
	previous *string
}

// returns inital configuration struct.
func newConfig() *Config {
	initialUrl := "https://pokeapi.co/api/v2/location-area/"

	return &Config{
		next: &initialUrl,
	}
}

// Returns an updated configuration file
func UpdateConfig(current *Config, updates Config) *Config {

	if updates.next != nil {
		current.next = updates.next
	}

	if updates.previous != nil {
		current.previous = updates.previous
	}

	return current
}
