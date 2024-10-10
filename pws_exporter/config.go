package pws_exporter

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Pws     string   `mapstructure:"pws"`
	Sensors []string `mapstructure:"sensors"`
}

func Configure() *Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return &c
}
