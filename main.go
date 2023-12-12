package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
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
		client := openai.NewClient("")
		resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Please write a 5 sentace paragraph about ducks.",
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
// https://api.openai.com/v1/chat/completions

//VARIABLE IS API_KEY
// Please generate a short 5 sentance paragraph about dinorsaurs.

	rootCmd.AddCommand(cmdGenerate)
	if err := rootCmd.Execute(); err != nil {
	fmt.Println(err)
	defer os.Exit(1)
	}
}

