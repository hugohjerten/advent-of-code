package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Cast string to int
func IntifyString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("string to convert to int: ", s)
		panic("Failed to convert string to int.")
	}
	return i
}

// Cast byte to int
func IntifyByte(b byte) int {
	r := rune(b)
	return int(r - '0')
}

func SplitStringOnNewline(str string) []string {
	return strings.Split(string(str), "\n")
}

// Split string on any character
func SplitStringOn(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Split string on whitespace
func SplitStringOnWhitespace(str string) []string {
	return strings.Fields(str)
}
