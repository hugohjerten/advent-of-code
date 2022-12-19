package seventeen

import (
	"2022/go/utils"
	"fmt"
)

const input = "../input/17.txt"
const nbrRocks = 2022
const caveWidth = 7

type Rock = [][]string

func getShapes() []Rock {
	return []Rock{
		{{"@", "@", "@", "@"}},
		{{".", "@", "."}, {"@", "@", "@"}, {".", "@", "."}},
		{{"@", "@", "@"}, {"@", ".", "."}, {"@", ".", "."}},
		{{"@"}, {"@"}, {"@"}, {"@"}},
		{{"@", "@"}, {"@", "@"}},
	}
}

type Cave struct {
	c       [][]string
	jet     []string
	rocks   []Rock
	nbr     int
	jetIdx  int
	rockIdx int
	height  int
}

func (c Cave) Show() {
	for _, s := range c.c {
		fmt.Println(s)
	}
}

func (c *Cave) elongateCave(nbr int) {
	for i := 0; i < nbr; i++ {
		row := make([]string, caveWidth)
		for j := 0; j < caveWidth; j++ {
			row[j] = "."
		}
		c.c = append(c.c, row)
	}
}

func (c *Cave) calcHeight() int {
	for y := len(c.c) - 1; y >= 0; y-- {
		if utils.ContainsStr(c.c[y], "#") {
			return y + 1
		}
	}
	return 0
}

func (c *Cave) rockRest(x int, y int, rock Rock) {
	for rY := 0; rY < len(rock); rY++ {
		for rX := 0; rX < len(rock[0]); rX++ {
			if rock[rY][rX] == "@" {
				c.c[rY+y][rX+x] = "#"
			}
		}
	}
	c.height = c.calcHeight()
}

func (c Cave) check(x int, y int, rock Rock) bool {
	// If bottom
	if y < 0 {
		return false
	}

	// If beyond cave walls
	if x < 0 || x+len(rock[0])-1 >= caveWidth {
		return false
	}

	// If hit another rock
	for rY := 0; rY < len(rock); rY++ {
		for rX := 0; rX < len(rock[0]); rX++ {
			if rock[rY][rX] == "@" &&
				c.c[rY+y][rX+x] == "#" {
				return false
			}
		}
	}
	return true
}

func (c *Cave) rockFall(x int, y int, rock Rock) (int, int) {
	d := c.jet[c.jetIdx]
	c.jetIdx += 1
	if c.jetIdx == len(c.jet) {
		c.jetIdx = 0
	}

	switch d {
	case ">":
		if c.check(x-1, y, rock) {
			x -= 1
		}
	case "<":
		if c.check(x+1, y, rock) {
			x += 1
		}
	}

	if c.check(x, y-1, rock) {
		y -= 1

	} else {
		// Bottom or other rock reached
		return x, y
	}

	return c.rockFall(x, y, rock)
}

func (c *Cave) nextRock() {
	rock := c.rocks[c.rockIdx]
	c.rockIdx += 1
	if c.rockIdx == len(c.rocks) {
		c.rockIdx = 0
	}

	// left edge is 2 units from the left wall
	// bottom edge is 3 units above the highest rock in the room
	x := caveWidth - (len(rock[0]) + 2)
	y := c.height + 3

	// Increase matrix if necessary
	if len(c.c) < y+len(rock) {
		c.elongateCave((y + len(rock)) - len(c.c))
	}

	x, y = c.rockFall(x, y, rock)
	c.rockRest(x, y, rock)
	c.nbr += 1
}

func (c *Cave) Simulate() {
	for {
		if c.nbr == nbrRocks {
			break
		}
		c.nextRock()
	}

	fmt.Println("Height: ", c.height)
}

func ParseInput(input string) Cave {
	jets := make([]string, 0)
	for _, d := range utils.ReadLines(input)[0] {
		jets = append(jets, string(d))
	}

	rs := make([][]string, 4)
	for i := range rs {
		rs[i] = make([]string, caveWidth)
		for j := 0; j < caveWidth; j++ {
			rs[i][j] = "."
		}
	}
	return Cave{rs, jets, getShapes(), 0, 0, 0, 0}
}

func Run() {
	cave := ParseInput(input)
	cave.Simulate()
}
