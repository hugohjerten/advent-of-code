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
	ns := []Coord{
		{X: c.X + 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y},
		{X: c.X, Y: c.Y + 1},
		{X: c.X, Y: c.Y - 1},
	}

	valid := make([]Coord, 0)
	for _, n := range ns {
		if hm.coordinateOk(n) && hm.elevationOk(c, n) {
			valid = append(valid, n)
		}
	}
	return valid
}

func (hm *HeightMap) Find() int {
	for {
		if hm.queue.Len() == 0 {
			break
		}

		loc := hm.queue.PopFront()
		if loc.c == hm.dest {
			return loc.distance
		}

		_, visited := hm.visited[loc.c]

		if !visited {
			hm.visited[loc.c] = empty{}

			for _, n := range hm.getNeighbours(loc.c) {
				_, visited = hm.visited[n]

				if !visited {
					hm.queue.PushBack(Location{loc.distance + 1, n})
				}
			}
		}

	}

	panic("No path could be found!")
}

func NewHeightMap(lines []string) HeightMap {
	m := make([][]rune, len(lines))
	var s Coord
	var e Coord

	for i, l := range lines {
		m[i] = make([]rune, len(l))
		for j, r := range l {

			switch r {
			case 83: // S
				r = 97 // a
				s = Coord{X: i, Y: j}
			case 69: // E
				r = 122 // z
				e = Coord{X: i, Y: j}
			}

			m[i][j] = r
		}
	}

	var queue deque.Deque[Location]
	queue.PushBack(Location{0, s})

	return HeightMap{
		m:       m,
		dest:    e,
		queue:   queue,
		visited: map[Coord]struct{}{},
	}
}

func Run() {
	lines := utils.ReadLines(input)
	hm := NewHeightMap(lines)
	shortest := hm.Find()
	fmt.Println("Shortest: ", shortest)

}
