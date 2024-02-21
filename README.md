# OpenCLIDocs

OpenCLIDocs is a command-line interface (CLI) tool for generating documentation using the OpenAI GPT-3.5 Turbo model.

## Installation

To use OpenCLIDocs, you need to have [Go](https://golang.org/) installed on your machine. Once Go is installed, you can get the OpenCLIDocs tool using:

```bash
go get github.com/your-username/OpenCLIDocs


## Usage
Generate Documentation

The generate command is used to generate documentation for a given source file. By default, it looks for the source file in the current directory and its subdirectories.

bash

./OpenCLIDocs generate path/to/your/source/file.go

You can specify the output file using the -o or --output flag. The generated documentation will be written to the specified Markdown file.

bash

./OpenCLIDocs generate path/to/your/source/file.go --output path/to/output/file.md

Input OpenAI API Key

The input command allows you to input your OpenAI API key. If the API key is not set, it prompts you to enter the key interactively.

bash

./OpenCLIDocs input

## Examples
Generate Documentation

Generate documentation for a Go source file named example.go:

bash

./OpenCLIDocs generate example.go

Generate documentation and save it to a specific Markdown file:

bash

./OpenCLIDocs generate example.go --output example.md

Input OpenAI API Key

Input your OpenAI API key interactively:

bash

./OpenCLIDocs input



Contributions are welcome! If you find any issues or have suggestions, please open an issue or submit a pull request.
License

This project is licensed under the MIT License - see the LICENSE file for details.

vbnet
