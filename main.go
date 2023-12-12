package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io"
	"log"
	// "net/http"
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
		client := openai.NewClient("") // api key goes here
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

	rootCmd.AddCommand(cmdGenerate)
	if err := rootCmd.Execute(); err != nil {
	fmt.Println(err)
	defer os.Exit(1)
	}
}

