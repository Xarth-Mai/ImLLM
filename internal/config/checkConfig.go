package config

import (
	"fmt"
	"strconv"
	"strings"
)

// CheckConfig checks if the configuration parameters in the userConfig are valid.
// 1. Port cannot be empty and must be a valid number between 1 and 65535.
// 2. Addr cannot be empty.
// 3. UserPasswd cannot be empty, and each username and password cannot be empty and cannot contain the illegal character ';'.
func CheckConfig(userConfig Config) error {
	// Check if Port is empty
	if strings.TrimSpace(userConfig.Port) == "" {
		return fmt.Errorf("port cannot be empty")
	}
	// Check if Port is a valid number
	portNum, err := strconv.Atoi(userConfig.Port)
	if err != nil {
		return fmt.Errorf("port must be a valid number, current value: %s", userConfig.Port)
	}
	// Check if the port number is within the reasonable range
	if portNum < 1 || portNum > 65535 {
		return fmt.Errorf("port must be within 1-65535, current value: %d", portNum)
	}

	// Check if Addr is empty
	if strings.TrimSpace(userConfig.Addr) == "" {
		return fmt.Errorf("addr cannot be empty")
	}

	// Check if UserPasswd is empty
	if userConfig.UserPasswd == nil || len(userConfig.UserPasswd) == 0 {
		return fmt.Errorf("UserPasswd cannot be empty")
	}

	// Check if each username and password is empty and if they contain illegal character ';'
	for k, v := range userConfig.UserPasswd {
		// Trim any leading or trailing spaces
		trimmedKey := strings.TrimSpace(k)
		trimmedValue := strings.TrimSpace(v)

		if trimmedKey == "" {
			return fmt.Errorf("username cannot be empty")
		}
		if trimmedValue == "" {
			return fmt.Errorf("password cannot be empty, for user '%s'", k)
		}
		if strings.Contains(trimmedKey, ";") || strings.Contains(trimmedValue, ";") {
			return fmt.Errorf("username or password cannot contain illegal character ';'")
		}
	}
	return nil
}
