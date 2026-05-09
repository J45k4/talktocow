package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func getEnvWithDefault(name string, defaultValue string) string {
	value := os.Getenv(name)

	if value == "" {
		return defaultValue
	}

	return value
}

func getCommaSeparatedEnvWithDefault(name string, defaultValue string) []string {
	value := getEnvWithDefault(name, defaultValue)
	parts := strings.Split(value, ",")
	result := []string{}

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)

		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

var PublicKeyPath = os.Getenv("PUBLIC_KEY_PATH")
var PrivateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
var DBName = os.Getenv("DB_NAME")
var DBUser = os.Getenv("DB_USER")
var DBPassword = os.Getenv("DB_PASSWORD")
var DbHost = os.Getenv("DB_HOST")
var DBPort = os.Getenv("DB_PORT")
var OpenAIApiKey = os.Getenv("OPENAI_API_KEY")
var WebAuthnRPID = getEnvWithDefault("WEBAUTHN_RP_ID", "localhost")
var WebAuthnRPOrigins = getCommaSeparatedEnvWithDefault("WEBAUTHN_RP_ORIGINS", "http://localhost:3080,http://localhost:12001")
var WebAuthnRPDisplayName = getEnvWithDefault("WEBAUTHN_RP_DISPLAY_NAME", "Talktocow")
var FileStoragePath = getEnvWithDefault("FILE_STORAGE_PATH", "./uploads")
