```markdown
# Code Tree Generator

This tool generates a tree structure of a directory, similar to the `tree` command. It provides an option to display additional information about files, such as their size and the number of lines.

## Usage

To run the tool, use the following command:

```sh
go run main.go [-v] <directory_path>
```

- `<directory_path>`: The path to the directory you want to generate the tree for.
- `-v`: (Optional) Enable verbose output to show the size and line count of each file.

## Example

To generate a simple tree structure:

```sh
go run main.go /path/to/directory
```

To generate a tree structure with verbose output:

```sh
go run main.go -v /path/to/directory
```

## Output

The output will display the directory structure in a tree format. If the `-v` flag is used, it will also show the size and number of lines for each file.

## Installation

To install the tool, clone the repository and build the Go project:

```sh
git clone <repository_url>
cd <repository_directory>
go build -o codetree
```

## Using as a Script

To build the tool and use it as a script, run the following command:

```sh
go build -o codetree
```

This will create an executable named `codetree` in the current directory. You can then run the tool using:

```sh
./codetree [-v] <directory_path>
```

## Adding an Alias

To make it easier to run the tool, you can add an alias to your shell configuration file.

For Linux and macOS, add the following line to your `~/.bashrc` or `~/.zshrc` file:

```sh
alias codetree='/path/to/codetree'
```

Replace `/path/to/codetree` with the actual path to the `codetree` executable. After adding the alias, reload your shell configuration:

```sh
source ~/.bashrc  # For bash
source ~/.zshrc   # For zsh
```

Now you can run the tool using the `codetree` command:

```sh
codetree [-v] <directory_path>
```

## License

This project is licensed under the MIT License.
```