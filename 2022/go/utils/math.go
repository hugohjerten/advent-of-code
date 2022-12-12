package utils

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coordinates struct {
	X int
	Y int
}

func Distance(p1 Coordinates, p2 Coordinates) float64 {
	return math.Sqrt(float64(p2.X-p1.X)*float64(p2.X-p1.X) + float64(p2.Y-p1.Y)*float64(p2.Y-p1.Y))
}

// Greatest Common Divisor (GCD) via Euclidean algorithm
func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Find Least Common Multiple (LCM) via GCD
func LCM(a int, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
