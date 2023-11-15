package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/aether/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("AETHER")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// No file was found, no worries, we write default file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Oops, something else went wrong
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	viper.SetDefault("spotify.url", "http://localhost:24879")

	viper.SafeWriteConfig()
}