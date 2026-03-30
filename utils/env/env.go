package env

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig loads defaults, optional .env file, then binds OS environment variables.
func InitConfig(environment map[string]any) error {
	for key, value := range environment {
		viper.SetDefault(strings.ToUpper(key), value)
	}
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config: could not read .env: %v (using defaults and environment)", err)
	}

	viper.AutomaticEnv()
	return nil
}
