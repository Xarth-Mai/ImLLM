package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// CreateConfig creates a config template file at the given path.
func CreateConfig(configPath string) (Config, error) {
	config := Config{
		Port: "0",
		Addr: "localhost",
		UserPasswd: map[string]string{
			"ChangeMe":  "ChangeMe",
			"ChangeMe!": "ChangeMe!",
		},
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return Config{}, err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return config, err
	}

	return config, nil
}
