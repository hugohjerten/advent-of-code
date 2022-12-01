package utils

import (
	"strconv"
)

func Intify(strs []string) []int {
	ints := make([]int, len(strs))
	for i := range ints {
		ints[i], _ = strconv.Atoi(strs[i])
	}

	return ints
}
