package twentyone

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/21.txt"

const root = "root"

type Monkey struct {
	name    string
	val     int  // value
	waiting bool // waiting for other monkeys
	left    string
	right   string
	op      string
}

// Math operation for given monkey. Return false if dependent monkey aren't ready
func (m Monkey) job(ms map[string]*Monkey) (int, bool) {
	l, r := ms[m.left], ms[m.right]
	if l.waiting || r.waiting {
		return 0, false
	}
	return utils.ArithmeticOperation(m.op)(l.val, r.val), true
}

func ParseInput() (map[string]*Monkey, []*Monkey) {
	lines := utils.ReadLines(input)
	monkeys := map[string]*Monkey{}
	var queue []*Monkey

	for _, l := range lines {
		split := utils.SplitStringOn(l, ": ")
		name, rightSplit := split[0], utils.SplitStringOnWhitespace(split[1])

		if len(rightSplit) == 1 {
			nbr, _ := strconv.Atoi(rightSplit[0])
			monkeys[name] = &Monkey{name, nbr, false, "", "", ""}

		} else {
			left, op, right := rightSplit[0], rightSplit[1], rightSplit[2]
			monkeys[name] = &Monkey{name, 0, true, left, right, op}
			queue = append(queue, monkeys[name])
		}
	}

	return monkeys, queue
}

func MonkeyShout(ms map[string]*Monkey, queue []*Monkey) {
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

	fmt.Println("Shouts: ", ms[root].val)
}

func Run() {
	monkeys, queue := ParseInput()
	MonkeyShout(monkeys, queue)
}
