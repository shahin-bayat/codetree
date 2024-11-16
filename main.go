package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	verbose := flag.Bool("v", false, "Enable verbose output (show size and line count)")
	ignore := flag.String("ignore", "", "Comma-separated list of folders to ignore")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [-v] [--ignore=folder1,folder2] <directory_path>")
		return
	}

	root := flag.Arg(0)
	ignoreList := strings.Split(*ignore, ",")
	if *ignore == "" {
		ignoreList = []string{} // Ensure empty list if no folders specified
	}

	info, err := getNodeInfo(root, *verbose)
	if err != nil {
		log.Fatalf("Error getting info for %s: %v", root, err)
	}

	fmt.Println(info.Name)
	if info.IsDir {
		if err := printTree(root, "", *verbose, ignoreList); err != nil {
			log.Fatalf("Error printing tree: %v", err)
		}
	}
}
