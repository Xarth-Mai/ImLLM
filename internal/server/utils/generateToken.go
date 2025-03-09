package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateToken generates a token for the given user.
func GenerateToken(userTokenSeed *map[string]string, username string) string {
	seedBytes := make([]byte, 16)
	_, err := rand.Read(seedBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random seed: %v", err))
	}

	seedHex := hex.EncodeToString(seedBytes)

	(*userTokenSeed)[username] = seedHex

	data := username + seedHex

	hashBytes := sha256.Sum256([]byte(data))

	token := hex.EncodeToString(hashBytes[:])
	return token
}
