package utils

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Distance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(float64(x2-x1)*float64(x2-x1) + float64(y2-y1)*float64(y2-y1))
}
