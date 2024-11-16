package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	verbose := flag.Bool("v", false, "Enable verbose output (show size and line count)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [-v] <directory_path>")
		return
	}

	root := flag.Arg(0)
	info, err := getNodeInfo(root, *verbose)
	if err != nil {
		log.Fatalf("Error getting info for %s: %v", root, err)
	}

	fmt.Println(info.Name)
	if info.IsDir {
		if err := printTree(root, "", *verbose); err != nil {
			log.Fatalf("Error printing tree: %v", err)
		}
	}
}
