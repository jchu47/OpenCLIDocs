# OpenCLIDocs

OpenCLIDocs is a command-line interface (CLI) tool for generating documentation using the OpenAI GPT-3.5 Turbo model.


## Installation

### Binary Download

You can download the latest release of the OpenCLIDocs binary from the [Releases](https://github.com/jchu47/OpenCLIDocs/releases) page. Choose the appropriate version for your operating system and architecture.

### Add to PATH

1. **Locate the OpenCLIDocs Binary:**
   After downloading the binary, locate the `OpenCLIDocs` executable in the extracted folder.

2. **Add to PATH:**
   Add the directory containing the `OpenCLIDocs` binary to your system's `PATH`. You can do this by adding the following line to your shell profile file (e.g., `~/.bashrc`, `~/.zshrc`, etc.):

   ```bash
   export PATH=$PATH:/path/to/folder/containing/OpenCLIDocs

Replace /path/to/folder/containing/OpenCLIDocs with the actual path to the directory containing the OpenCLIDocs binary.

After updating your shell profile, restart your terminal or run source ~/.bashrc (or the corresponding command for your shell) to apply the changes.

Now, you should be able to run OpenCLIDocs generate example.go directly from any directory.

## Usage

### Generate Documentation

The `generate` command is used to generate documentation for a given source file. By default, it looks for the source file in the current directory and its subdirectories.

```bash
./OpenCLIDocs generate path/to/your/source/file.go
```
You can specify the output file using the -o or --output flag. The generated documentation will be written to the specified Markdown file.
```bash
./OpenCLIDocs generate path/to/your/source/file.go --output path/to/output/file.md
```

### Input OpenAI API Key

The input command allows you to input your OpenAI API key. If the API key is not set, it prompts you to enter the key interactively.
```bash
./OpenCLIDocs input
```

### Examples
Generate Documentation

Generate documentation for a Go source file named example.go:

Input OpenAI API Key (optional)

Skip to input your OpenAI API key interactively:

```bash

./OpenCLIDocs input
```

```bash
./OpenCLIDocs generate example.go
```

Generate documentation and save it to a specific Markdown file:

``` bash

./OpenCLIDocs generate example.go -o different_example.md
```


## Contributions
Contributions are welcome! If you find any issues or have suggestions, please open an issue or submit a pull request.
License

This project is licensed under the MIT License - see the LICENSE file for details.

vbnet
