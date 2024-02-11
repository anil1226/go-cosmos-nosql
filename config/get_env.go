package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../config/.env")
}

func GetEnvKey(s string) string {
	return os.Getenv(s)
}
