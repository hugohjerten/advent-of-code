package twentyone

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/21.txt"

type Monkey struct {
	val     int                                  // value
	waiting bool                                 // waiting for other monkeys
	job     func(map[string]*Monkey) (int, bool) // math operation.
}

func ParseInput() (map[string]*Monkey, []*Monkey) {
	lines := utils.ReadLines(input)
	monkeys := map[string]*Monkey{}
	var queue []*Monkey

	for _, l := range lines {
		split := utils.SplitStringOn(l, ": ")
		name := split[0]

		rightSplit := utils.SplitStringOnWhitespace(split[1])

		if len(rightSplit) == 1 {
			nbr, _ := strconv.Atoi(rightSplit[0])
			monkeys[name] = &Monkey{nbr, false, nil}

		} else {
			l := rightSplit[0]
			op := utils.ArithmeticOperation(rightSplit[1])
			r := rightSplit[2]

			// Math operation for given monkey. Return false if dependent monkey aren't ready
			monkeys[name] = &Monkey{0, true, func(ms map[string]*Monkey) (int, bool) {
				left, right := ms[l], ms[r]
				if left.waiting || right.waiting {
					return 0, false
				}
				return op(left.val, right.val), true
			}}

			queue = append(queue, monkeys[name])
		}
	}

	return monkeys, queue
}

func MonkeyShout(final string, ms map[string]*Monkey, queue []*Monkey) {
	for {
		if len(queue) == 0 {
			// All monkeys are calculated
			break
		}

		// pop monkey
		m := queue[0]
		queue = queue[1:]

		// This monkey is already calculated, skip
		if !m.waiting {
			continue
		}

		val, ok := m.job(ms)
		if ok {
			m.val = val
			m.waiting = false
		} else {
			// put monkey back in queue
			queue = append(queue, m)
		}
	}

	fmt.Println("Shouts: ", ms[final].val)
}

func Run() {
	monkeys, queue := ParseInput()
	MonkeyShout("root", monkeys, queue)
}
