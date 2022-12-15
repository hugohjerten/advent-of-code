package fourteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/14.txt"

var source = Coord{500, 0}

type State int

const (
	Blocked State = iota
	Falling
	OutOfBounds
	ExpandLeft
	ExpandRight
)

type Coord struct {
	x int
	y int
}

type Path = []Coord

type Cave struct {
	m        [][]rune
	w        int
	h        int
	offset   int
	source   Coord
	voidless bool
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
	if s.y >= c.h {
		return OutOfBounds
	}
	if c.voidless && (s.x < 0 || s.x > c.w) {
		return OutOfBounds
	}
	if !c.voidless && s.y == c.h-1 {
		return Blocked
	}
	if !c.voidless && s.x < 0 && s.y < c.h {
		return ExpandLeft
	}
	if !c.voidless && s.x >= c.w && s.y < c.h {
		return ExpandRight
	}
	if c.m[s.y][s.x] == 32 { // " "
		return Falling
	}
	return Blocked
}

func (c *Cave) expandLeft() {
	for i, row := range c.m {
		var r rune
		if i < c.h-1 {
			r = 32 // " "
		} else {
			r = 35 // "#"
		}
		c.m[i] = append([]rune{r}, row...)
	}
	c.w += 1
	c.source.x += 1
}

func (c *Cave) expandRight() {
	for i, row := range c.m {
		var r rune
		if i < c.h-1 {
			r = 32 // " "
		} else {
			r = 35 // "#"
		}
		c.m[i] = append(row, r)
	}
	c.w += 1
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
		case ExpandLeft:
			c.expandLeft()
			return c.fall(Coord{f.x + 1, f.y})
		case ExpandRight:
			c.expandRight()
			return c.fall(f)
		}
	}
	return s, Blocked
}

func (c *Cave) Simulate(show bool) {
	cnt := 0
	stop := false

	// New unit of sand for each loop
	for {
		if stop {
			break
		}

		n, s := c.fall(c.source)
		switch s {
		case Blocked:
			c.m[n.y][n.x] = 111 // "o"
			cnt += 1
			if n == c.source {
				stop = true
			}
		case OutOfBounds:
			stop = true
		}
	}

	if show {
		c.Show()
	}
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

	// Add floor
	if !c.voidless {
		for x := 0; x < c.w; x++ {
			c.m[c.h-1][x] = 35 // "#"
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

func FetchCave(input string, voidless bool) Cave {
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

	if !voidless {
		h += 2
	}

	m := make([][]rune, h)
	for i := range m {
		m[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			m[i][j] = 32 // " "
		}
	}

	offset := xEdges[0]
	cave := Cave{m, w, h, offset, Coord{source.x - offset, source.y}, voidless}
	cave.addPaths(paths)
	return cave
}

func Run() {
	cave := FetchCave(input, true)
	cave.Simulate(false)

	cave = FetchCave(input, false)
	cave.Simulate(false)
}
