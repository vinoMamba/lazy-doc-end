package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	JwtSecret string
}

var _config config

func LoadConfig(path string) {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)   // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&_config); err != nil {
		panic(err)
	}
}

func GetJwtSecret() string {
	return _config.JwtSecret
}
