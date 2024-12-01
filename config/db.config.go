package config

import "github.com/spf13/viper"

// Config struct for application settings
type Config struct {
	AppName     string `json:"APP_NAME"`
	Environment string `json:"ENVIRONMENT"`
	DatabaseURL string `json:"DATABASE_URL"`
	ServerPort  string `json:"SERVER_PORT"`
}

// LoadConfig loads the configuration from the environment or `.env` file
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")

	// Automatic binding from environment variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// Allow missing config file if env vars are set
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
