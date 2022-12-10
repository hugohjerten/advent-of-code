package ten

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/10.txt"

var strengthCycles = []int{20, 60, 100, 140, 180, 220}

type CPU struct {
	cycle    int
	x        int
	strength int
}

func (cpu *CPU) exec(c int, x int) {
	cpu.cycle += c
	cpu.x += x

	if c != 0 && utils.ContainsInt(strengthCycles, cpu.cycle) {
		cpu.strength += cpu.cycle * cpu.x
	}
}

type CRT struct {
	row    int
	pixels [6][40]string
}

func (crt *CRT) pixel(i int, x int) {
	p := "."
	i -= 1

	// Update row nbr (if not very first pixel)
	if i != 0 && i%40 == 0 {
		crt.row += 1
	}

	// Update row index
	i -= (i / 40) * 40

	// If x in sprite
	if x-1 <= i && i <= x+1 {
		p = "#"
	}

	crt.pixels[crt.row][i] = p
}

type VideoSystem struct {
	cpu CPU
	crt CRT
}

func NewVideoSystem() VideoSystem {
	return VideoSystem{CPU{0, 1, 0}, CRT{}}
}

func (vs *VideoSystem) add(x int) {
	vs.cpu.exec(0, x)
}

func (vs *VideoSystem) cycle(nbr int) {
	for i := 0; i < nbr; i++ {
		vs.cpu.exec(1, 0)
		vs.crt.pixel(vs.cpu.cycle, vs.cpu.x)
	}
}

func (vs *VideoSystem) Run(instructions []string) {
	for _, i := range instructions {

		if i == "noop" {
			vs.cycle(1)
		} else {
			x, _ := strconv.Atoi(utils.SplitStringOnWhitespace(i)[1])
			vs.cycle(2)
			vs.add(x)
		}
	}
}

func (vs *VideoSystem) SignalStrength() int {
	return vs.cpu.strength
}

func (vs *VideoSystem) Draw() {
	for _, r := range vs.crt.pixels {
		fmt.Println(r)
	}
}

func Run() {
	lines := utils.ReadLines(input)
	vs := NewVideoSystem()
	vs.Run(lines)
	fmt.Println("Sums: ", vs.SignalStrength())
	vs.Draw()
}
