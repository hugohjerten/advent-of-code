package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
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
	fileContent, err := ioutil.ReadFile(absolutePath)
	check(err)

	content := string(fileContent)

	return strings.Split(string(content), "\n")
}

func SeparateSlices(baseList []string) [][]string {
	slices := make([][]string, 0, len(baseList))

	sliceStart := 0
	for i := 0; i < len(baseList); i++ {

		// When empty string, have reached end of slice
		if baseList[i] == "" {
			length := i - sliceStart
			slices = append(slices, baseList[sliceStart:sliceStart+length])
			sliceStart = i + 1
		}
	}

	// Add last slice which is missed in above for-loop
	slices = append(slices, baseList[sliceStart:])

	return slices
}
