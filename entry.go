package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type NodeInfo struct {
	Name     string
	Path     string
	IsDir    bool
	Size     int64
	NumLines int
}

func getNodeInfo(path string, verbose bool) (NodeInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return NodeInfo{}, err
	}

	nodeInfo := NodeInfo{
		Name:  info.Name(),
		Path:  path,
		IsDir: info.IsDir(),
		Size:  info.Size(),
	}

	if verbose && !nodeInfo.IsDir {
		lines, err := countLines(path)
		if err == nil {
			nodeInfo.NumLines = lines
		}
	}

	return nodeInfo, nil
}

func printTree(path, prefix string, verbose bool) error {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range dirs {
		if entry.Name()[0] == '.' { // Skip hidden files and directories
			continue
		}

		isLast := i == len(dirs)-1
		var newPrefix string
		if isLast {
			fmt.Printf("%s└── %s\n", prefix, entry.Name())
			newPrefix = prefix + "    "
		} else {
			fmt.Printf("%s├── %s\n", prefix, entry.Name())
			newPrefix = prefix + "│   "
		}

		fullPath := filepath.Join(path, entry.Name())
		info, err := getNodeInfo(fullPath, verbose)
		if err != nil {
			log.Printf("Error getting info for %s: %v", fullPath, err)
			continue
		}

		if verbose && !info.IsDir {
			fmt.Printf("%s    Size: %d bytes, Lines: %d\n", newPrefix, info.Size, info.NumLines)
		}

		if info.IsDir {
			if err := printTree(fullPath, newPrefix, verbose); err != nil {
				log.Printf("Error reading directory %s: %v", fullPath, err)
			}
		}
	}

	return nil
}
