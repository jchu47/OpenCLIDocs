package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jchu47/OpenCLIDocs/api"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)
var outputFile string

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate documentation for a given source file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		// Attempt to find the file recursively in all directories
		filePath, err := findFileRecursive(".", fileName)
		if err != nil {
			log.Fatalf("Failed to find file: %v", err)
		}

		// Extract the base name of the input file without extension
		baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

		// Check if "documentation" folder exists, create it if not
		docFolder := "documentation"
		if _, err := os.Stat(docFolder); os.IsNotExist(err) {
			err := os.Mkdir(docFolder, 0755)
			if err != nil {
				log.Fatalf("Failed to create %s folder: %v", docFolder, err)
			}
		}

		// Construct the output file path within the "documentation" folder with a .md extension
		outputFile = filepath.Join(docFolder, baseName+".md")

		fileContent, err := os.ReadFile(filePath)
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

		// Write the documentation to the output file
		err = os.WriteFile(outputFile, []byte(resp.Choices[0].Message.Content), 0644)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		fmt.Printf("Documentation written to %s\n", outputFile)
	},

}

func init() {
	GenerateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the documentation (Markdown format)")
}

func findFileRecursive(dir, fileName string) (string, error) {
	filePath := filepath.Join(dir, fileName)
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	// Recursively search in subdirectories
	subdirs, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, subdir := range subdirs {
		if subdir.IsDir() {
			subdirPath := filepath.Join(dir, subdir.Name())
			foundPath, err := findFileRecursive(subdirPath, fileName)
			if err == nil {
				return foundPath, nil
			}
		}
	}

	return "", fmt.Errorf("file not found: %s", fileName)
}
