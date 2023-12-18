# Documentation Generator

A Go program that generates documentation for a given source file in markdown format.

## Usage

```
docai generate [flags] [source file]
```

## Flags

- `--output, -o`:  Output file for the documentation (Markdown format)

## Example

```
docai generate -o output.md main.go
```

This command will generate Markdown documentation for the `main.go` file and save it to `output.md`.