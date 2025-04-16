package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetNatsUrl() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	natsUrl := os.Getenv("NATS_URL")
	return natsUrl, nil
}
