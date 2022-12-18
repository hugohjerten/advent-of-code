package sixteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
	"strings"
)

const input = "../input/16.txt"
const start = "AA"
const limit = 30

type Valve struct {
	l string
	r int
	t map[string]*Valve
}

type State struct {
	m int
	l string
	p int
	o []string
}

type MinuteValve struct {
	m int
	l string
}

type Network struct {
	vs map[string]*Valve
}

func (n Network) MaximumPressure() {
	// BFS
	queue := make([]State, 0)
	queue = append(queue, State{1, start, 0, make([]string, 0)})

	seens := map[MinuteValve]int{}
	max := 0

	for {
		if len(queue) == 0 {
			break
		}

		// pop
		state := queue[0]
		queue = queue[1:]

		// update seen states
		minuteValve := MinuteValve{state.m, state.l}
		pressure, seen := seens[minuteValve]

		// if already seen, and current pressure is worse than previous best, skip
		if seen && pressure >= state.p {
			continue
		}
		seens[minuteValve] = state.p

		// Limit reached
		if state.m == limit {
			if state.p > max {
				max = state.p
			}
			continue
		}

		// open current valve
		if n.vs[state.l].r > 0 && !utils.ContainsStr(state.o, state.l) {
			opened := utils.CopyStringList(state.o)
			opened = append(opened, state.l)
			pressure := state.p
			for _, l := range opened {
				pressure += n.vs[l].r
			}

			queue = append(queue, State{state.m + 1, state.l, pressure, opened})
		}

		// leave current valve closed
		pressure = state.p
		for _, l := range state.o {
			pressure += n.vs[l].r
		}
		for l := range n.vs[state.l].t {
			queue = append(queue, State{state.m + 1, l, pressure, state.o})
		}
		// queue = append(queue, n.keepValveClosed(state)...)
	}

	fmt.Println("Maximum pressure: ", max)
}

func Parse(input string) Network {
	lines := utils.ReadLines(input)

	vs := map[string]*Valve{}

	for _, l := range lines {
		t := map[string]*Valve{}

		l = strings.ReplaceAll(l, "Valve ", "")
		l = strings.ReplaceAll(l, " has flow rate=", ",")
		l = strings.ReplaceAll(l, "; tunnels lead to valves ", ",")
		l = strings.ReplaceAll(l, "; tunnel leads to valve ", ",")
		l = strings.ReplaceAll(l, " ", "")
		ls := utils.SplitStringOn(l, ",")

		for _, v := range ls[2:] {
			t[v] = nil
		}
		r, _ := strconv.Atoi(ls[1])
		vs[ls[0]] = &Valve{ls[0], r, t}
	}

	for _, v := range vs {
		for w := range v.t {
			v.t[w] = vs[w]
		}
	}

	return Network{vs}
}

func Run() {
	network := Parse(input)
	network.MaximumPressure()
}
