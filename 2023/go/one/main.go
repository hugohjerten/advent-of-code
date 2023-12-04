package one

import (
	"2023/go/utils"
	"fmt"
	"strings"
)

const input = "../input/1.txt"

var numbers = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var forward = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var backward = [10]string{"orez", "eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func getCalibrationValuePart1(line string) int {
	var first, last int
	for i := 0; i < len(line); i++ {
		if utils.IsByteInt(line[i]) {
			first = utils.IntifyByte(line[i])
			break
		}

	}
	for i := len(line) - 1; i >= 0; i-- {
		if utils.IsByteInt(line[i]) {
			last = utils.IntifyByte(line[i])
			break
		}
	}

	return first*10 + last
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += getCalibrationValuePart1(line)
	}
	fmt.Println("Part 1: ", sum)
}

func matchNumber(s string) (int, bool) {
	for i, sub := range numbers {
		if strings.Contains(s, sub) {
			return i, true
		}
	}
	for i, sub := range forward {
		if strings.Contains(s, sub) {
			return i, true
		}
	}
	for i, sub := range backward {
		if strings.Contains(s, sub) {
			return i, true
		}
	}
	return -1, false
}

func getCalibrationValuePart2(line string) int {
	var first, last int
	var found bool
	for i := 1; i < len(line)+1; i++ {
		if first, found = matchNumber(line[:i]); found {
			break
		}

	}

	for i := len(line) - 1; i >= 0; i-- {
		if last, found = matchNumber(line[i:]); found {
			break
		}
	}

	return first*10 + last
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += getCalibrationValuePart2(line)
	}
	fmt.Println("Part 2: ", sum)
}

func Run() {
	lines := utils.ReadLines(input)
	part1(lines)
	part2(lines)
}
