package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var apiKeyFile = ".env"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "docai",
		Short: "Documentation Generator",
	}

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "Generate documentation for a given source file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileContent, err := os.ReadFile(args[0])
			if err != nil {
				log.Fatalf("Failed to read file: %v", err)
			}

			apiKey, err := getAPIKeyFromFile() // Attempt to retrieve the API key from the file
			if err != nil || apiKey == "" {
				fmt.Printf("API key is not set, please enter: ")
				fmt.Scanln(&apiKey)
				err := setAPIKeyToFile(apiKey) // Save the API key to the file
				if err != nil {
					fmt.Println("Error setting API key: ", err)
					return
			}
			}

			client := openai.NewClient(apiKey)
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: "Generate documentation for this code: " + string(fileContent),
						},
					},
				},
			)

			if err != nil {
				fmt.Printf("ChatCompletion error: %v\n", err)
				return
			}

			fmt.Println(resp.Choices[0].Message.Content)
		},
	}

	var inputAPI = &cobra.Command{
		Use:   "input",
		Short: "User inputs OpenAPI Key",
		Run: func(cmd *cobra.Command, args []string) {
			var apiKey string
			fmt.Printf("Enter your API key: ")
			fmt.Scanln(&apiKey)
			err := setAPIKeyToFile(apiKey) // Save the API key to the file
			if err != nil {
				fmt.Println("Error setting API key: ", err)
				return
			}
			fmt.Println("API key set successfully.")
		},
	}

	rootCmd.AddCommand(cmdGenerate)
	rootCmd.AddCommand(inputAPI)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getAPIKeyFromFile() (string, error) {
	content, err := os.ReadFile(apiKeyFile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func setAPIKeyToFile(apiKey string) error {
	return os.WriteFile(apiKeyFile, []byte(apiKey), 0600)
}