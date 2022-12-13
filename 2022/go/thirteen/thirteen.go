package thirteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/13.txt"

type Integer struct {
	val int
}

type Packet struct {
	val string
	idx int
}

func (p Packet) sliceFrom(start int) (string, bool) {
	nstd := 0
	for i := start; i < len(p.val); i++ {
		switch p.val[i] {
		case 91: // '['
			nstd += 1
		case 93:
			nstd -= 1

			if nstd == 0 {
				sub := p.val[start : i+1]
				if utils.ContainsOnly(p.val[i+1:], 93) { // ']'
					return sub, true
				}
				return sub, false
			}
		}
	}
	panic("Couldn't get slice.")
}

func (p *Packet) Next() (*Integer, *Packet) {
	if p.idx >= len(p.val) {
		return nil, nil
	}

	r := p.val[p.idx]

	switch r {
	case 91: // '['
		val, last := p.sliceFrom(p.idx)

		if last {
			return nil, &Packet{val, 1}
		}

		p.idx += 1
		return nil, p

	case 93: // ']'
		p.idx += 1
		return p.Next()

	case 44: // ','
		p.idx += 1
		return p.Next()

	default:
		var val int
		// i.e. if value is two digits, next digit is between 0<9
		if r >= 48 && r <= 57 {
			val, _ = strconv.Atoi(p.val[p.idx : p.idx+1])
			p.idx += 2
		} else {
			val, _ = strconv.Atoi(string(r))
			p.idx += 1
		}
		return &Integer{val}, nil
	}

}

type Pair struct {
	l Packet
	r Packet
}

func Listify(i int) string {
	return "[" + strconv.Itoa(i) + "]"
}

func IsEmpty(p *Packet) bool {
	return p == nil || p.val == "[]"
}

func compare(p Pair) bool {
	fmt.Println("Left: ", p.l.val)
	fmt.Println("Right: ", p.r.val)

	// cnt := 0

	for {
		leftInt, leftPacket := p.l.Next()
		rightInt, rightPacket := p.r.Next()

		// If both left and right run out of items
		if leftInt == nil && IsEmpty(leftPacket) &&
			rightInt == nil && IsEmpty(rightPacket) {
			fmt.Println("Both left and right have run out of items.")
			return true
		}

		// If left is list, but right is empty
		if leftPacket != nil && IsEmpty(rightPacket) && rightInt == nil {
			fmt.Println("Left list, right empty.")
			return false
		}

		// If left is empty, but right is list
		if IsEmpty(leftPacket) && leftInt == nil && rightPacket != nil {
			fmt.Println("Left empty, right list.")
			return true
		}

		// If both int
		if leftInt != nil && rightInt != nil {
			fmt.Println("Both are int.")
			if leftInt.val > rightInt.val {
				return false
			}
			continue
		}

		// If both lists
		if leftPacket != nil && rightPacket != nil {
			fmt.Println("Both are lists.")
			if !compare(Pair{*leftPacket, *rightPacket}) {
				return false
			}
			continue
		}

		// If left is int, but right is list
		if leftInt != nil && rightPacket != nil {
			packed := Packet{Listify(leftInt.val), 1}
			fmt.Println("Left int, right list.")
			fmt.Println("Packed: ", packed.val)
			fmt.Println("right", rightPacket.val)

			if !compare(Pair{packed, *rightPacket}) {
				return false
			}
		}

		// If left is list, but right is int
		if leftPacket != nil && rightInt != nil {
			packed := Packet{Listify(rightInt.val), 1}
			fmt.Println("Left list, right int.")
			fmt.Println("left", leftPacket.val)
			fmt.Println("Packed: ", packed.val)

			// os.Exit(2)
			if !compare(Pair{*leftPacket, packed}) {
				return false
			}
		}

	}
}

func ComparePairs(pairs []Pair) {
	sum := 0
	for i, p := range pairs {
		fmt.Println("Pair #", i)
		if compare(p) {
			fmt.Println("Is in right order.")
			sum += i
		}
		// os.Exit(1)
	}

	fmt.Println("Sum of indices: ", sum)
}

func ParsePairs(input string) []Pair {
	split := utils.SeparateSliceOnNewLine(utils.ReadLines(input))
	pairs := make([]Pair, len(split))
	for i, p := range split {
		pairs[i] = Pair{Packet{p[0], 1}, Packet{p[1], 1}}
	}

	return pairs
}

func Run() {
	pairs := ParsePairs(input)
	ComparePairs(pairs)
}
