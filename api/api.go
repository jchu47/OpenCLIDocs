package api

import (
	"os"
	"strings"
	"fmt"
	"bufio"
)

var apiKeyFile = ".env"

func GetAPIKeyFromFile() (string, error) {
	file, err := os.Open(apiKeyFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "OPENAI_KEY=") {
			return strings.TrimPrefix(line, "OPENAI_KEY="), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("OPENAI_KEY not found in %s", apiKeyFile)
}

func SetAPIKeyToFile(apiKey string) error {
	file, err := os.OpenFile(apiKeyFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var updatedContent strings.Builder
	scanner := bufio.NewScanner(file)
	keyFound := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "OPENAI_KEY=") {
			keyFound = true
			line = fmt.Sprintf("OPENAI_KEY=%s", apiKey)
		}
		updatedContent.WriteString(line + "\n")
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if !keyFound {
		updatedContent.WriteString(fmt.Sprintf("OPENAI_KEY=%s\n", apiKey))
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(updatedContent.String())
	if err != nil {
		return err
	}

	return nil
}