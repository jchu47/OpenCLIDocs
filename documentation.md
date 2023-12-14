# Documentation Generator

The `docai` package is a documentation generator that uses the OpenAI chat completion model to generate documentation for a given source file. It is written in Go.

## Commands

### Generate

The `generate` command is used to generate documentation for a given source file.

Usage: `docai generate [flags] <source-file>`

- `<source-file>`: The path to the source file for which documentation needs to be generated.
- Flags:
  - `-o, --output <file>`: Optional. The output file for the documentation in Markdown format.

Example: `docai generate -o output.md source.go`

## Implementation

The `main` function initializes the root command and the `generate` command using the `cobra` package.

The `generate` command reads the content of the source file specified in the command-line arguments. It then creates a client for the OpenAI chat completion API using the provided API key. The source file content is passed as a user message to the chat completion API.

The response from the API is printed to the console or written to the specified output file if provided.

## Dependencies

- `github.com/sashabaranov/go-openai`: Go client library for the OpenAI API.
- `github.com/spf13/cobra`: Command-line interface (CLI) library for Go.
- Other standard Go packages (`fmt`, `log`, `context`, `os`) are also used.
