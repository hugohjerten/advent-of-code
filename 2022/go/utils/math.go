package utils

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func EuclidianDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(float64(x2-x1)*float64(x2-x1) + float64(y2-y1)*float64(y2-y1))
}

func ManhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
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

// Return arithmetic operation from string
func ArithmeticOperation(op string) func(int, int) int {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	return ops[op]
}
