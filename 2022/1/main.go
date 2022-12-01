package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type Elf struct {
	calories []int
	sum      int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func readLines(path string) []string {
	absolutePath, err := filepath.Abs(path)
	check(err)
	fileContent, err := ioutil.ReadFile(absolutePath)
	check(err)

	content := string(fileContent)

	return strings.Split(string(content), "\n")
}

func separateSlices(baseList []string) [][]string {
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

func intify(strs []string) []int {
	ints := make([]int, len(strs))
	for i := range ints {
		ints[i], _ = strconv.Atoi(strs[i])
	}

	return ints
}

func sum(calories []int) int {
	total := 0
	for _, v := range calories {
		total += v
	}
	return total
}

func getElfs(filePath string) []Elf {
	slice := readLines(filePath)
	slices := separateSlices(slice)

	elfs := make([]Elf, len(slices))
	for i, slice := range slices {
		calories := intify(slice)
		elfs[i] = Elf{calories, sum(calories)}
	}
	return elfs
}

func maxCalories(elfs []Elf) int {
	max := 0
	for _, elf := range elfs {
		if elf.sum > max {
			max = elf.sum
		}
	}

	return max
}

func main() {
	elfs := getElfs("input.txt")
	max := maxCalories(elfs)
	fmt.Println("Max calories: ", max)
}
