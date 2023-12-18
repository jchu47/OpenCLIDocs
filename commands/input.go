package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jchu47/OpenCLIDocs/api"
)

var InputCmd = &cobra.Command{
	Use:   "input",
	Short: "User inputs OpenAPI Key",
	Run: func(cmd *cobra.Command, args []string) {
		var apiKey string
		fmt.Printf("Enter your API key: ")
		fmt.Scanln(&apiKey)
		err := api.SetAPIKeyToFile(apiKey)
		if err != nil {
			fmt.Println("Error setting API key: ", err)
			return
		}
		fmt.Println("API key set successfully.")
	},
}