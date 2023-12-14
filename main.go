package main

import (
	"fmt"
	"log"
	"context"
	"os"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

func main(){
	var rootCmd = &cobra.Command {
	    Use: "docai",
			Short: "Documentation Generator",
	}

	var cmdGenerate = &cobra.Command {
		Use: "generate",
		Short: "Generate documentation for a given source file",
		Args: cobra.ExactArgs(1),
		Run:func(cmd *cobra.Command, args []string) {
			fileContent, err := os.ReadFile(args[0])
			if err != nil {
					log.Fatalf("Failed to read file: %v", err)
			}
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("You have not entered an API key yet.  Please enter one now: ")
			var inputAPIKey string
				fmt.Scanln(&inputAPIKey)
				err := os.Setenv("API_KEY", inputAPIKey)
				if err != nil {
					fmt.Println("Error setting API_KEY: ", err)
					return
				}
				apiKey = inputAPIKey
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

	var inputAPI = &cobra.Command {
	    Use: "input",
			Short: "User inputs OpenAPI Key",
			Run: func(cmd *cobra.Command, args []string) {
				var apiKey string
				fmt.Printf("Enter your API key: ")
				fmt.Scanln(&apiKey)
				err := os.Setenv("API_KEY", apiKey)
				if err != nil {
					fmt.Println("Error setting API_KEY: ", err)
				}
				test := os.Getenv("API_KEY")
			if test != "" {
				fmt.Println("API_KEY:", test)
			} else {
				fmt.Println("API_KEY is not set.  Please re-run with a valid API Key.")
			}
			},
	}

	rootCmd.AddCommand(cmdGenerate)
	rootCmd.AddCommand(inputAPI)

	if err := rootCmd.Execute(); err != nil {
	fmt.Println(err)
	defer os.Exit(1)
	}
}



