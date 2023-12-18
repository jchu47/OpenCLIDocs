# Documentation Generator

The `main` package contains the `main` function. This function initializes the root command of the documentation generator tool and adds two subcommands to it.

## Usage

```
docai [command]
```

## Commands

The following commands are available:

### generate

```
docai generate [flags]
```

The `generate` command generates the documentation.

### input

```
docai input [flags]
```

The `input` command specifies the input for the documentation generation.

## Exit Code

- `0`: The command executed successfully.
- `1`: An error occurred during command execution.

## Examples

- Generate the documentation:

```
docai generate --input=file.txt
```

- Specify input for documentation generation:

```
docai input --input=file.txt
```