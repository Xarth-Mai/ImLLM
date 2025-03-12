package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config defines the configuration data structure.
type Config struct {
	Port       string            `yaml:"Port"`
	Addr       string            `yaml:"Addr"`
	UserPasswd map[string]string `yaml:"UserPasswd"`
}

// InitConfig initializes the configuration from the yaml file.
func InitConfig(configPath string) (Config, error) {

	// Read the configuration file content
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	// Parse YAML data into the config struct
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	// Validate if the configuration is legal
	if err := CheckConfig(config); err != nil {
		return Config{}, err
	}

	return config, nil
}
