package nine

import (
	"2022/go/utils"
	"fmt"
	"math"
	"strconv"
)

type void struct{}

var member void

type Direction int

const (
	Up Direction = iota
	Down
	Right
	Left
)

var (
	directionMap = map[string]Direction{
		"U": Up,
		"D": Down,
		"R": Right,
		"L": Left,
	}
)

type Knot struct {
	x int
	y int
}

type Rope struct {
	head Knot
	tail Knot
	// set of historic locations
	hist map[Knot]void
}

func NewRope() Rope {
	loc := Knot{0, 0}
	hist := make(map[Knot]void, 0)
	hist[loc] = member
	return Rope{loc, loc, hist}
}

func (head Knot) drag(tail Knot) Knot {
	x := tail.x
	y := tail.y

	dist := utils.Distance(x, y, head.x, head.y)

	if dist > 1 {
		if utils.Abs(head.x-x) > 1 || dist > math.Sqrt(2) {
			if head.x > x {
				x += 1
			} else if head.x < x {
				x -= 1
			}
		}

		if utils.Abs(head.y-y) > 1 || dist > math.Sqrt(2) {
			if head.y > y {
				y += 1
			} else if head.y < y {
				y -= 1
			}
		}
	}

	return Knot{x, y}
}

func (r *Rope) move(d Direction) {
	var head Knot
	if d == Up {
		head = Knot{r.head.x, r.head.y + 1}
	} else if d == Down {
		head = Knot{r.head.x, r.head.y - 1}
	} else if d == Right {
		head = Knot{r.head.x + 1, r.head.y}
	} else if d == Left {
		head = Knot{r.head.x - 1, r.head.y}
	}

	tail := head.drag(r.tail)

	r.head = head
	r.tail = tail
	r.hist[tail] = member

}

func (r *Rope) GetPositions(ds []Direction) {
	for _, d := range ds {
		r.move(d)
	}

	nbrPositions := len(r.hist)

	fmt.Println("Number of positions: ", nbrPositions)
}

func ReadDirections(path string) []Direction {
	lines := utils.ReadLines(path)
	ds := make([]Direction, 0)

	for _, line := range lines {
		split := utils.SplitStringOnWhitespace(line)
		d := directionMap[split[0]]
		nbr, _ := strconv.Atoi(split[1])

		for i := 0; i < nbr; i++ {
			ds = append(ds, d)
		}
	}

	return ds
}

const input = "../input/9.txt"

func Run() {
	directions := ReadDirections(input)
	r := NewRope()
	r.GetPositions(directions)
}
