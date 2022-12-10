package ten

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/10.txt"

type CPU struct {
	cycle    int
	register int
	cs       []int
}

func NewCPU(cycles []int) CPU {
	return CPU{0, 1, cycles}
}

func (cpu *CPU) run(c int, r int) int {
	cpu.cycle += c
	cpu.register += r

	if utils.ContainsInt(cpu.cs, cpu.cycle) {
		return cpu.cycle * cpu.register
	}

	return 0
}

func (cpu *CPU) SignalStrengths(lines []string) int {
	signals := 0
	for _, l := range lines {

		if l == "noop" {
			signals += cpu.run(1, 0)
		} else {
			x, _ := strconv.Atoi(utils.SplitStringOnWhitespace(l)[1])
			signals += cpu.run(1, 0)
			signals += cpu.run(1, 0)
			cpu.run(0, x)
		}
	}

	return signals
}

func Run() {
	lines := utils.ReadLines(input)
	cpu := NewCPU([]int{20, 60, 100, 140, 180, 220})
	sum := cpu.SignalStrengths(lines)

	fmt.Println("Sums: ", sum)
}
