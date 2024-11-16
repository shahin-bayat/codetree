package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	verbose := flag.Bool("v", false, "Enable verbose output (show size and line count)")
	ignore := flag.String("ignore", "", "Comma-separated list of folders to ignore")
	output := flag.String("o", "", "File to save output (optional)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [-v] [--ignore=folder1,folder2] [-o output_file] <directory_path>")
		return
	}

	root := flag.Arg(0)
	ignoreList := strings.Split(*ignore, ",")
	if *ignore == "" {
		ignoreList = []string{}
	}

	var outputFile *os.File
	var err error
	if *output != "" {
		outputFile, err = os.Create(*output)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer outputFile.Close()
	}

	writer := os.Stdout
	if outputFile != nil {
		writer = outputFile
	}

	info, err := getNodeInfo(root, *verbose)
	if err != nil {
		log.Fatalf("Error getting info for %s: %v", root, err)
	}

	fmt.Fprintln(writer, info.Name)
	if info.IsDir {
		if err := printTree(root, "", *verbose, ignoreList, writer); err != nil {
			log.Fatalf("Error printing tree: %v", err)
		}
	}
}
