package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var PublicKeyPath = os.Getenv("PUBLIC_KEY_PATH")
var PrivateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
var DBName = os.Getenv("DB_NAME")
var DBUser = os.Getenv("DB_USER")
var DBPassword = os.Getenv("DB_PASSWORD")
var DbHost = os.Getenv("DB_HOST")
var DBPort = os.Getenv("DB_PORT")
var OpenAIApiKey = os.Getenv("OPENAI_API_KEY")
