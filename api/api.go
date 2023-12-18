package api

import (
	"os"
)

var apiKeyFile = ".env"

func GetAPIKeyFromFile() (string, error) {
	content, err := os.ReadFile(apiKeyFile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SetAPIKeyToFile(apiKey string) error {
	return os.WriteFile(apiKeyFile, []byte(apiKey), 0600)
}