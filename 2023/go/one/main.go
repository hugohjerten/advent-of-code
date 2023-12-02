package one

import (
	"2023/go/utils"
	"fmt"
)

const input = "../input/1.txt"

func getCalibrationValue(line string) int {
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

func Run() {
	lines := utils.ReadLines(input)
	sum := 0
	for _, line := range lines {
		sum += getCalibrationValue(line)
	}
	fmt.Println("Sum of all calibration values is: ", sum)
}
