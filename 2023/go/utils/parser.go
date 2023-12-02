package utils

import (
	"strings"
)

// Cast byte to int
func IntifyByte(b byte) int {
	r := rune(b)
	return int(r - '0')
}

func SplitStringOnNewline(str string) []string {
	return strings.Split(string(str), "\n")
}
