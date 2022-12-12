package twelve

import (
	"2022/go/utils"
	"fmt"

	"github.com/gammazero/deque"
)

const input = "../input/12.txt"

type Coord struct {
	X int
	Y int
}

type Location struct {
	distance int
	c        Coord
}

type empty = struct{}

type HeightMap struct {
	m       [][]rune
	dest    Coord
	queue   deque.Deque[Location]
	visited map[Coord]empty // set
}

func (hm HeightMap) elevationOk(from Coord, to Coord) bool {
	return hm.m[to.X][to.Y] <= hm.m[from.X][from.Y]+1
}

func (hm HeightMap) coordinateOk(c Coord) bool {
	return (c.X >= 0 && c.Y >= 0) && (c.X < len(hm.m) && c.Y < len(hm.m[0]))
}

func (hm HeightMap) getNeighbours(c Coord) []Coord {
	ns := make([]Coord, 0)
	for _, n := range []Coord{
		{X: c.X + 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y},
		{X: c.X, Y: c.Y + 1},
		{X: c.X, Y: c.Y - 1},
	} {
		// If coordinates inside map, and ok elevation diff
		if hm.coordinateOk(n) && hm.elevationOk(c, n) {
			ns = append(ns, n)
		}
	}
	return ns
}

func (hm *HeightMap) Find(start Coord) int {
	// Add starting position to queue
	hm.queue.PushBack(Location{0, start})

	for {
		// If still locations left to visit
		if hm.queue.Len() == 0 {
			break
		}

		// If found destination
		loc := hm.queue.PopFront()
		if loc.c == hm.dest {
			return loc.distance
		}

		// If not visited this location yet
		_, visited := hm.visited[loc.c]
		if !visited {
			hm.visited[loc.c] = empty{}

			for _, n := range hm.getNeighbours(loc.c) {
				_, visited = hm.visited[n]

				// If not visited neighbour yet, add to queue
				if !visited {
					hm.queue.PushBack(Location{loc.distance + 1, n})
				}
			}
		}

	}

	return 10000000
}

func (hm *HeightMap) Scenic(starts []Coord) int {
	min := 100000
	for _, s := range starts {
		hm.queue.Clear()
		hm.visited = map[Coord]empty{}
		steps := hm.Find(s)

		// If shorter route
		if steps < min {
			min = steps
		}
	}

	return min
}

func ReadInput(input string) (HeightMap, Coord, []Coord) {
	lines := utils.ReadLines(input)
	m := make([][]rune, len(lines))
	starts := make([]Coord, 0)
	var start Coord
	var dest Coord

	for i, l := range lines {
		m[i] = make([]rune, len(l))
		for j, r := range l {

			switch r {
			case 83: // S
				r = 97 // a
				start = Coord{X: i, Y: j}
			case 69: // E
				r = 122 // z
				dest = Coord{X: i, Y: j}
			case 97:
				starts = append(starts, Coord{i, j})
			}

			m[i][j] = r
		}
	}

	return HeightMap{
		m:       m,
		dest:    dest,
		queue:   deque.Deque[Location]{},
		visited: map[Coord]empty{},
	}, start, starts
}

func Run() {
	hm, start, starts := ReadInput(input)
	shortest := hm.Find(start)
	fmt.Println("Shortest: ", shortest)
	scenic := hm.Scenic(starts)
	fmt.Println("Scenic: ", scenic)
}
