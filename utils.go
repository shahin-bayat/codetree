package main

import (
	"os"
	"strings"
)

func countLines(path string) (int, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return len(strings.Split(string(bs), "\n")), nil

}
