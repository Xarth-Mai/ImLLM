package config

import (
	"gopkg.in/yaml.v3"
	"log"
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
	var config Config
	// Read the configuration file content
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Config file does not exist, creating a new one...")
			config, err = CreateConfig(configPath)
			if err != nil {
				return Config{}, err
			}
		} else {
			return Config{}, err
		}
	}

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
