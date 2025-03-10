package config

import (
	"fmt"
	"strings"
)

// Init initializes the configuration from the yaml file.
func Init(addr *string, port *string, userPasswd *map[string]string) error {
	*port = "4000"
	*userPasswd = map[string]string{
		"admin":  "passwd",
		"admin1": "passwd1",
	}
	*addr = "localhost"

	for k, v := range *userPasswd {
		if strings.Contains(k, ";") || strings.Contains(v, ";") {
			return fmt.Errorf("username or password contains illegal character ';'")
		}
	}

	return nil
}
