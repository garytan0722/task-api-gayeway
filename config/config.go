package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	LogLevel string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("port", 8080)
	viper.SetDefault("log-level", "debug")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return &Config{
		Port:     viper.GetInt("port"),
		LogLevel: viper.GetString("log-level"),
	}, nil
}
