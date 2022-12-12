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
	knots []Knot
	// set of historic knot locations
	hist map[Knot]void
}

func NewRope(knots int) Rope {
	ks := make([]Knot, knots)
	hist := map[Knot]void{}

	for i := 0; i < knots; i++ {
		knot := Knot{0, 0}
		ks[i] = knot
	}
	hist[ks[0]] = member

	return Rope{ks, hist}
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
	ks := make([]Knot, len(r.knots))

	switch d {
	case Up:
		ks[0] = Knot{r.knots[0].x, r.knots[0].y + 1}
	case Down:
		ks[0] = Knot{r.knots[0].x, r.knots[0].y - 1}
	case Right:
		ks[0] = Knot{r.knots[0].x + 1, r.knots[0].y}
	case Left:
		ks[0] = Knot{r.knots[0].x - 1, r.knots[0].y}
	}

	for i := 1; i < len(r.knots); i++ {
		ks[i] = ks[i-1].drag(r.knots[i])
	}

	r.knots = ks
	r.hist[ks[len(ks)-1]] = member
}

func (r *Rope) GetPositions(ds []Direction) int {
	for _, d := range ds {
		r.move(d)
	}

	nbrPositions := len(r.hist)

	return nbrPositions
}

func ParseDirections(lines []string) []Direction {
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
	directions := ParseDirections(utils.ReadLines(input))
	r := NewRope(2)
	nbr := r.GetPositions(directions)
	fmt.Println("Number of positions: ", nbr)

	r = NewRope(10)
	nbr = r.GetPositions(directions)
	fmt.Println("Number of positions: ", nbr)
}
