package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the environment or `.env` file
func LoadConfig() {

	// Dynamically find the root directory
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Error finding root path: %s", err)
	}

	// Specify the file name and type
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Add the root directory as a search path
	viper.AddConfigPath(rootPath)

	// Read the config file
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}

}

// GetConfig retrieves the value for the given key
func GetConfig(key string, baseValue string) string {
	envValue := viper.GetString(key)

	if envValue == "" {
		return baseValue
	}
	return envValue
}
