package utils

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig loads defaults, optional .env file, then binds OS environment variables.
func InitConfig() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config: could not read .env: %v (using defaults and environment)", err)
	}

	viper.AutomaticEnv()
	return nil
}
