package config

const (
	defaultUrl = "amqp://guest:guest@localhost:5672/"
)

type Config struct {
	Url string `long:"url" short:"u" description:"AMQP endpoint to consume from"`
}

// DefaultConfig returns a config with default hardcoded values.
func DefaultConfig() *Config {
	return &Config{
		Url: defaultUrl,
	}
}
