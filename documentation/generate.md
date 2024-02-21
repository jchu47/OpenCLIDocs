# Package Commands

The `commands` package contains functionality to generate documentation for a given source file.

## Functions

### GenerateCmd
- **Use**: generate
- **Short**: Generate documentation for a given source file
- **Args**: cobra.ExactArgs(1)
- **Flags**:
  - --output, -o: Output file for the documentation (Markdown format)

### findFileRecursive(dir, fileName string) (string, error)
Recursively searches for a file in all directories.

## Usage

To generate documentation for a source file, run the following command:
```
generate [source_file]
```

The command will recursively search for the source file in all directories, extract the base name without extension, create a "documentation" folder if it does not exist, and write the documentation in Markdown format to an output file within the "documentation" folder.

## Dependencies
- `github.com/jchu47/OpenCLIDocs/api`
- `github.com/sashabaranov/go-openai`
- `github.com/spf13/cobra`

## Example
```go
package main

import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{Use: "myapp"}
	rootCmd.AddCommand(commands.GenerateCmd)
	rootCmd.Execute()
}
```

This will add the `generate` command to your CLI application for generating documentation.