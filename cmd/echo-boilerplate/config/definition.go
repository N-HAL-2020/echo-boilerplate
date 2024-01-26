package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

const (
	configName = "config"
	configPath = "./cmd/echo-boilerplate/config"
	configType = "yaml"
)

func Load() (*Config, error) {
	setup()

	err := readConfigFile()
	if err != nil {
		return nil, err
	}

	config, err := unmarshalConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func setup() {
	viper.SetDefault("port", "8080")

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func readConfigFile() error {
	err := viper.ReadInConfig()
	if err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	return nil
}

func unmarshalConfig() (*Config, error) {
	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
