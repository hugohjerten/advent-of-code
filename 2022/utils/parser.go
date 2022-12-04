package utils

import (
	"strconv"
	"strings"
)

// Cast slice of string to slice of int
func Intify(strs []string) []int {
	ints := make([]int, len(strs))
	for i := range ints {
		ints[i], _ = strconv.Atoi(strs[i])
	}

	return ints
}

// Split slice of strings, on empty elements; returning multiple slices.
// Eg ["A", "B", "C", "", "D", "E"] -> [["A", "B", "C"], ["D", "E"]]
func SeparateSliceOnNewLine(baseList []string) [][]string {
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

// Split string on whitespace
func SplitStringOnWhitespace(str string) []string {
	return strings.Fields(str)
}

func SplitStringsOnWhitespace(strs []string) [][]string {
	split := make([][]string, len(strs))
	for i, str := range strs {
		split[i] = SplitStringOnWhitespace(str)
	}
	return split
}

func SplitStringOn(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Split string in middle, raise error if not even number
func SplitStringInMiddle(str string) (string, string) {
	if len(str)%2 != 0 {
		panic("Not even number for splitting.")
	}

	middle := len(str) / 2
	return str[:middle], str[middle:]
}
