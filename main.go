package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

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
			// filename, err := os.ReadFile(args[0])
			// if err != nil {
			// log.Printf("Failed to read file: %v", err)
			// }
			requestBody, err := json.Marshal(map[string]interface{}{
				"model": "gpt-3.5-turbo",
				"messages": []map[string]string{
					{
						"role":    "system",
						"content": "You are a helpful assistant.",
					},
					{
						"role":    "user",
						"content": "Please generate a short 5 sentance paragraph about dinorsaurs",
					},
				},
			})
			if err != nil {
				log.Fatalf("Failed to read json Marshil: %v", err)
			}
			api_key := ""
			req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
			if err != nil {
				log.Fatalf("Failed to create requestt: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+api_key)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatalf("Failed to read response body: %v", err)
				}
				log.Fatalf("Unexpected status code: %d, response: %s", resp.StatusCode, body)
			}
			paragraph, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Failed to read response: %v", err)
			}
			log.Printf("Response: %s", string(paragraph))
		},
	}

	rootCmd.AddCommand(cmdGenerate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
}
