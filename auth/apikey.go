package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const APIKeyPrefix = "ttc_"

func GenerateAPIKey() (string, string, string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", "", err
	}

	token := APIKeyPrefix + base64.RawURLEncoding.EncodeToString(bytes)
	prefix := APIKeyDisplayPrefix(token)
	hash := HashAPIKey(token)

	return token, prefix, hash, nil
}

func HashAPIKey(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func APIKeyDisplayPrefix(token string) string {
	if len(token) <= 12 {
		return token
	}

	return token[:12]
}

func LooksLikeAPIKey(token string) bool {
	return strings.HasPrefix(token, APIKeyPrefix)
}

func ValidateAPIKeyName(name string) error {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return fmt.Errorf("name is required")
	}
	if len(trimmed) > 100 {
		return fmt.Errorf("name must be at most 100 characters")
	}
	return nil
}
