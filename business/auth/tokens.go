package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const APITokenPrefix = "app-"

func (b *Business) hash(token string) string {
	token_hash := sha256.Sum256([]byte(token))
	return base64.StdEncoding.EncodeToString(token_hash[:])
}

func (b *Business) generateToken(length int) (string, error) {
	// Generate a cryptographically secure random hex string
	token := make([]byte, length)
	n, err := rand.Read(token)
	if err != nil || n != length {
		return "", err
	}

	return fmt.Sprintf("%s%X", APITokenPrefix, token), nil
}
