package configs

// Config holds the configuration values
type Config struct {
	ServerAddress string
}

// LoadConfig loads the configuration from file or environment variables
func LoadConfig() (*Config, error) {
	// Placeholder for actual config loading
	return &Config{
		ServerAddress: ":8080",
	}, nil
}
