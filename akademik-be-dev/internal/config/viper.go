package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	// read config
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
