package main

import (
	"fmt"
	"os"

	"github.com/jchu47/OpenCLIDocs/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "docai",
		Short: "Documentation Generator",
	}

	rootCmd.AddCommand(commands.GenerateCmd)
	rootCmd.AddCommand(commands.InputCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		fmt.Println(err)
		defer os.Exit(1)
	}
}
