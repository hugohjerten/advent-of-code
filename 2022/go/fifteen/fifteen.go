package fifteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
	"strings"
)

const input = "../input/15.txt"
const row = 2000000
const min = 0
const max = 4000000

type Coord struct {
	x int
	y int
}

type Coefficients = map[int]int

type Pair struct {
	s Coord
	b Coord
	d int
}

func (p Pair) noBeacons(set map[Coord]struct{}) map[Coord]struct{} {
	for x := p.s.x - p.d; x <= p.s.x+p.d; x++ {
		if !(x == p.b.x && row == p.b.y) &&
			!(x == p.s.x && row == p.s.y) &&
			(utils.ManhattanDistance(p.s.x, p.s.y, x, row) <= p.d) {
			set[Coord{x, row}] = struct{}{}
		}
	}
	return set
}

// The boundary of a scanner is four line segments.
// If a scanner is in position (sx,sy) and has distance d, then we
// want the line segments just outside, i.e. of distance d+1.

// There will be two line segments of gradient 1 (y=x+a):
//   y = x + sy-sx+d+1
//   y = x + sy-sx-d-1
// and two line segments of gradient -1 (y=-x+b):
//   y = -x + sx+sy+d+1
//   y = -x + sx+sy-d-1

// Line y=x+a and line y=-x+b intersect at the point ( (b-a)/2 , (a+b)/2 ).
// One of these intersection points will be the missing scanner location.
// So, we assemble a set of all the 'a' coefficients (lines of gradient 1)
// and all the 'b' coefficients (lines of gradient -1), then look at their
// intersections to see if they are the point we need.

func (p Pair) coefficients(a Coefficients, b Coefficients) (Coefficients, Coefficients) {
	a = add(a, p.s.y-p.s.x+p.d+1)
	a = add(a, p.s.y-p.s.x-p.d-1)
	b = add(b, p.s.x+p.s.y+p.d+1)
	b = add(b, p.s.x+p.s.y-p.d-1)
	return a, b
}

func add(cs Coefficients, c int) Coefficients {
	_, ok := cs[c]
	if ok {
		cs[c] += 1
	} else {
		cs[c] = 1
	}
	return cs
}

func ReadInput(input string) []Pair {
	lines := utils.ReadLines(input)
	pairs := make([]Pair, len(lines))

	for i, l := range lines {
		l = strings.ReplaceAll(l, "Sensor at x=", "")
		l = strings.ReplaceAll(l, " y=", "")
		l = strings.ReplaceAll(l, " closest beacon is at x=", "")
		split := utils.SplitStringOn(l, ":")
		sensor := utils.SplitStringOn(split[0], ",")
		beacon := utils.SplitStringOn(split[1], ",")
		sensorX, _ := strconv.Atoi(sensor[0])
		sensorY, _ := strconv.Atoi(sensor[1])
		beaconX, _ := strconv.Atoi(beacon[0])
		beaconY, _ := strconv.Atoi(beacon[1])
		pairs[i] = Pair{
			Coord{sensorX, sensorY},
			Coord{beaconX, beaconY},
			utils.ManhattanDistance(sensorX, sensorY, beaconX, beaconY),
		}
	}

	return pairs
}

func EmptyPositions(ps []Pair) {
	set := map[Coord]struct{}{}
	for _, p := range ps {
		set = p.noBeacons(set)
	}
	fmt.Println("Number positions: ", len(set))
}

func intersect(a int, b int) Coord {
	x := (b - a) / 2
	y := (a + b) / 2
	return Coord{x, y}
}

func missingLocation(c Coord, ps []Pair) bool {
	if !(c.x >= min && c.x <= max && c.y >= min && c.y <= max) {
		return false
	}
	for _, p := range ps {
		d := utils.ManhattanDistance(p.s.x, p.s.y, c.x, c.y)
		if d <= p.d {
			return false
		}
	}
	return true
}

func TuningFrequency(ps []Pair) {
	as := Coefficients{}
	bs := Coefficients{}
	for _, p := range ps {
		as, bs = p.coefficients(as, bs)
	}

	for a := range as {
		for b := range bs {
			i := intersect(a, b)
			if missingLocation(i, ps) {
				fmt.Println("Tuning Frequency: ", i.x*max+i.y)
				return
			}
		}
	}
}

func Run() {
	pairs := ReadInput(input)
	EmptyPositions(pairs)
	TuningFrequency(pairs)
}
