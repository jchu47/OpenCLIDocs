package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/jchu47/OpenCLIDocs/api"
)
var outputFile string

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate documentation for a given source file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		apiKey, err := api.GetAPIKeyFromFile()
		if err != nil || apiKey == "" {
			fmt.Printf("API key is not set, please enter: ")
			fmt.Scanln(&apiKey)
			err := api.SetAPIKeyToFile(apiKey)
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
						Content: "Generate documentation for this code in markdown format: " + string(fileContent),
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		if outputFile != "" {
			err = os.WriteFile(outputFile, []byte(resp.Choices[0].Message.Content), 0644)
			if err != nil {
				log.Fatalf("Failed to write to file: %v", err)
			}
			fmt.Printf("Documentation written to %s\n", outputFile)
		} else {
			fmt.Println(resp.Choices[0].Message.Content)
		}

	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the documentation (Markdown format)")
}