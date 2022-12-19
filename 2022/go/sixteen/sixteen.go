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

type EState struct {
	m  int
	l1 string
	l2 string
	p  int
	o  []string
}

type MinuteValve struct {
	m int
	l string
}

type MinuteValve2 struct {
	m  int
	l1 string
	l2 string
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

		// elf opens valve
		if n.vs[state.l].r > 0 && !utils.ContainsStr(state.o, state.l) {
			opened := utils.CopyStringList(state.o)
			opened = append(opened, state.l)
			pressure := state.p
			for _, l := range opened {
				pressure += n.vs[l].r
			}

			queue = append(queue, State{state.m + 1, state.l, pressure, opened})
		}

		// elf moves through tunnel
		pressure = state.p
		for _, l := range state.o {
			pressure += n.vs[l].r
		}
		for l := range n.vs[state.l].t {
			queue = append(queue, State{state.m + 1, l, pressure, state.o})
		}
	}

	fmt.Println("Maximum pressure: ", max)
}

func (n Network) Elephants() {
	// BFS
	queue := make([]EState, 0)
	queue = append(queue, EState{1, start, start, 0, make([]string, 0)})

	allOpen := 0
	for _, v := range n.vs {
		allOpen += v.r
	}

	seens := map[MinuteValve2]int{}
	max := 0

	for {
		if len(queue) == 0 {
			break
		}

		// pop
		state := queue[0]
		queue = queue[1:]

		// update seen states
		minuteValve := MinuteValve2{state.m, state.l1, state.l2}
		pressure, seen := seens[minuteValve]

		// if already seen, and current pressure is worse than previous best, skip
		if seen && pressure >= state.p {
			continue
		}
		seens[minuteValve] = state.p

		// Limit reached
		if state.m == limit-4 {
			if state.p > max {
				max = state.p
			}
			continue
		}

		// Case 0: All valves are open
		pressure = 0
		for _, l := range state.o {
			pressure += n.vs[l].r
		}
		if pressure >= allOpen {
			cumPress := state.p + pressure
			m := state.m
			for {
				if m >= limit-4 {
					break
				}
				cumPress += pressure
				m += 1
			}
			queue = append(queue, EState{m, state.l1, state.l2, cumPress, state.o})
			continue
		}

		// Case 1: elf open valve
		if n.vs[state.l1].r > 0 && !utils.ContainsStr(state.o, state.l1) {

			// Case 1A: and elephant opens valve
			if n.vs[state.l2].r > 0 && !utils.ContainsStr(state.o, state.l2) {
				opened := utils.CopyStringList(state.o)
				opened = append(opened, state.l1)
				opened = append(opened, state.l2)
				pressure := state.p
				for _, l := range opened {
					pressure += n.vs[l].r
				}
				queue = append(queue, EState{state.m + 1, state.l1, state.l2, pressure, opened})
			}

			// Case 1B: and elephant moves through tunnel
			opened := utils.CopyStringList(state.o)
			opened = append(opened, state.l1)
			pressure := state.p
			for _, l := range opened {
				pressure += n.vs[l].r
			}

			for l2 := range n.vs[state.l2].t {
				queue = append(queue, EState{state.m + 1, state.l1, l2, pressure, state.o})
			}
		}

		// Case 2: elf moves through tunnel
		for l1 := range n.vs[state.l1].t {

			// Case 2A: and elephant opens valve
			if n.vs[state.l2].r > 0 && !utils.ContainsStr(state.o, state.l2) {
				opened := utils.CopyStringList(state.o)
				opened = append(opened, state.l2)
				pressure := state.p
				for _, l := range opened {
					pressure += n.vs[l].r
				}
				queue = append(queue, EState{state.m + 1, l1, state.l2, pressure, opened})
			}

			// Case 2B: and elephant moves through tunnel
			opened := utils.CopyStringList(state.o)
			pressure := state.p
			for _, l := range opened {
				pressure += n.vs[l].r
			}

			for l2 := range n.vs[state.l2].t {
				queue = append(queue, EState{state.m + 1, l1, l2, pressure, opened})
			}

		}
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
	network.Elephants()
}
