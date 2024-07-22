package config

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func InitConfigFile() {
	// Set the config file name and path
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	fmt.Println("Successfully loaded config.yaml")

}
