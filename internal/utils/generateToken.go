package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateToken generates a random token
func GenerateToken() string {
	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		panic(fmt.Sprintf("failed to generates random token: %v", err))
	}

	return hex.EncodeToString(bytes)
}
