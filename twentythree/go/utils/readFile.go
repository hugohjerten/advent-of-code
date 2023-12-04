package utils

import (
	"log"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func ReadLines(path string) []string {
	absolutePath, err := filepath.Abs(path)
	check(err)
	fileContent, err := os.ReadFile(absolutePath)
	check(err)

	content := string(fileContent)

	return SplitStringOnNewline(content)
}
