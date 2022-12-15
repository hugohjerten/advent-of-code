package fourteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/14.txt"

type State int

const (
	Blocked State = iota
	Falling
	OutOfBounds
)

type Coord struct {
	x int
	y int
}

type Path = []Coord

type Cave struct {
	m      [][]rune
	w      int
	h      int
	offset int
	source Coord
}

func (c Cave) Show() {
	for _, r := range c.m {
		fmt.Println(string(r))
	}
}

func (s *Cave) addPath(p Path) {
	for _, c := range p {
		s.m[c.y][c.x] = 35 // "#"
	}
}

func (c *Cave) check(s Coord) State {
	if s.x < 0 || s.x > c.w || s.y < 0 || s.y > c.h {
		return OutOfBounds
	}
	if c.m[s.y][s.x] == 46 { // "."
		return Falling
	}
	return Blocked
}

func (c *Cave) fall(s Coord) (Coord, State) {
	falls := []Coord{{s.x, s.y + 1}, {s.x - 1, s.y + 1}, {s.x + 1, s.y + 1}}
	for _, f := range falls {
		switch c.check(f) {
		case Blocked:
			continue
		case Falling:
			return c.fall(f)
		case OutOfBounds:
			return s, OutOfBounds
		}
	}
	return s, Blocked
}

func (c *Cave) Simulate() {
	cnt := 0
	abyss := false

	// New unit of sand for each loop
	for {
		if abyss {
			break
		}

		n, s := c.fall(c.source)
		switch s {
		case Blocked:
			c.m[n.y][n.x] = 111 // "o"
			cnt += 1
		case OutOfBounds:
			abyss = true
		}
	}

	c.Show()
	fmt.Println("Units of sand: ", cnt)
}

func (c *Cave) addPaths(paths []Path) {
	// Draw source
	c.m[c.source.y][c.source.x] = 43 // "+"

	for _, cs := range paths {
		for i := 0; i < len(cs)-1; i++ {
			c.addPath(getRockPath(cs[i], cs[i+1], c.offset))
		}
	}
}

func getRockPath(c1 Coord, c2 Coord, offset int) []Coord {
	cs := make([]Coord, 0)

	if c1.x == c2.x {
		// Vertical rock path
		x := c1.x - offset

		if c1.y < c2.y {
			// Going down
			for y := c1.y; y <= c2.y; y++ {
				cs = append(cs, Coord{x, y})
			}
		} else {
			// Going up
			for y := c1.y; y >= c2.y; y-- {
				cs = append(cs, Coord{x, y})
			}
		}
	} else {
		// Horizontal rock path
		y := c1.y

		if c1.x < c2.x {
			// Going right
			for x := c1.x - offset; x <= c2.x-offset; x++ {
				cs = append(cs, Coord{x, y})
			}
		} else {
			// Going left
			for x := c1.x - offset; x >= c2.x-offset; x-- {
				cs = append(cs, Coord{x, y})
			}
		}
	}
	return cs
}

func minMax(old [2]int, n int) [2]int {
	if n < old[0] {
		old[0] = n
	}
	if n > old[1] {
		old[1] = n
	}
	return old
}

func FetchCave(input string, source Coord) Cave {
	lines := utils.ReadLines(input)
	paths := make([]Path, len(lines))

	xEdges := [2]int{source.x, source.x}
	yEdges := [2]int{source.y, source.y}

	for i, l := range lines {
		split := utils.SplitStringOn(l, " -> ")
		paths[i] = make(Path, len(split))

		for j, c := range split {
			cs := utils.SplitStringOn(c, ",")
			x, _ := strconv.Atoi(cs[0])
			y, _ := strconv.Atoi(cs[1])
			paths[i][j] = Coord{x, y}

			xEdges = minMax(xEdges, x)
			yEdges = minMax(yEdges, y)
		}
	}

	w := xEdges[1] - xEdges[0] + 1
	h := yEdges[1] - yEdges[0] + 1

	m := make([][]rune, h)
	for i := range m {
		m[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			m[i][j] = 46 // "."
		}
	}

	offset := xEdges[0]
	cave := Cave{m, w, h, offset, Coord{source.x - offset, source.y}}
	cave.addPaths(paths)
	return cave
}

func Run() {
	cave := FetchCave(input, Coord{500, 0})
	cave.Simulate()
}
