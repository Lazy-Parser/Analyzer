package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDotenv(names ...string) ([]string, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return []string{}, fmt.Errorf("error trying to get .env var: %w", err)
	}

	var vars []string
	for idx, name := range names {
		vars[idx] = os.Getenv(name)
	}

	return vars, nil
}
