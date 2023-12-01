package twentytwo

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/22.txt"

func ParseInput() Map {
	i := utils.SeparateSliceOnNewLine(utils.ReadLines(input))
	lines, path := i[0], i[1][0]
	max := 0
	for _, l := range lines {
		if len(l) > max {
			max = len(l)
		}
	}

	start := len(lines[0])
	jungle := make([][]string, len(lines))

	for y, l := range lines {
		jungle[y] = make([]string, max)

		for x := 0; x < max; x++ {
			if x < len(l) {
				jungle[y][x] = string(l[x])
				if y == 0 && string(l[x]) != " " && x < start {
					start = x
				}
				continue
			}
			jungle[y][x] = " "
		}
	}

	rowEdges := make([][]int, len(jungle))
	for y, row := range jungle {
		leftEdge, rightEdge := max, 0
		for x := 0; x < max; x++ {
			if row[x] != " " {
				if x < leftEdge {
					leftEdge = x
				}
				if x > rightEdge {
					rightEdge = x
				}
			}
		}
		rowEdges[y] = []int{leftEdge, rightEdge}
	}
	colEdges := make([][]int, max)
	for x := 0; x < max; x++ {
		topEdge, bottomEdge := len(jungle), 0
		for y := 0; y < len(jungle); y++ {
			if jungle[y][x] != " " {
				if y < topEdge {
					topEdge = y
				}
				if y > bottomEdge {
					bottomEdge = y
				}
			}
		}
		colEdges[x] = []int{topEdge, bottomEdge}
	}

	return Map{jungle, 0, start, ">", path, rowEdges, colEdges}
}

type Map struct {
	jungle   [][]string
	y        int
	x        int
	facing   string
	path     string
	rowEdges [][]int
	colEdges [][]int
}

func (m Map) Print() {
	for _, row := range m.jungle {
		fmt.Println(row)
	}
}

func (m *Map) rotate(r string) {
	if m.facing == ">" {
		if r == "R" {
			m.facing = "v"
		} else {
			m.facing = "^"
		}
	} else if m.facing == "v" {
		if r == "R" {
			m.facing = "<"
		} else {
			m.facing = ">"
		}
	} else if m.facing == "<" {
		if r == "R" {
			m.facing = "^"
		} else {
			m.facing = "v"
		}
	} else if m.facing == "^" {
		if r == "R" {
			m.facing = ">"
		} else {
			m.facing = "<"
		}
	}
}

func (m Map) isWall(x int, y int) bool {
	return m.jungle[y][x] == "#"
}

func (m *Map) moveVertically(s int) {
	topEdge, bottomEdge := m.colEdges[m.x][0], m.colEdges[m.x][1]

	for i := 0; i < s; i++ {
		y := m.y + 1

		if m.isWall(m.x, y) {
			return
		}
		if y > bottomEdge {
			y = topEdge
			if m.isWall(m.x, y) {
				return
			}
		} else if y < topEdge {
			y = bottomEdge
			if m.isWall(m.x, y) {
				return
			}
		}
		m.y = y
	}
}

func (m *Map) moveHorizontally(s int) {
	leftEdge, rightEdge := m.rowEdges[m.y][0], m.rowEdges[m.y][1]

	for i := 0; i < s; i++ {
		x := m.x + 1

		if m.isWall(x, m.y) {
			return
		}
		if x > rightEdge {
			x = leftEdge
			if m.isWall(x, m.y) {
				return
			}
		} else if x < leftEdge {
			x = rightEdge
			if m.isWall(x, m.y) {
				return
			}
		}
		m.x = x
	}
}

func (m *Map) move(s int) {
	if m.facing == ">" || m.facing == "<" {
		m.moveHorizontally(s)
	} else {
		m.moveVertically(s)
	}
}

func (m *Map) Walk() {
	for i := 0; i < len(m.path); i++ {

		str := string(m.path[i])
		cnt := 0
		if str == "R" || str == "L" {
			m.rotate(str)
			continue
		}

		for {
			if i+cnt == len(m.path) ||
				string(m.path[i+cnt]) == "R" ||
				string(m.path[i+cnt]) == "L" {
				str = m.path[i : i+cnt]
				break
			}
			cnt += 1
		}

		steps, _ := strconv.Atoi(str)
		m.move(steps)

		if cnt > 0 {
			i += cnt - 1
		}
	}
}

func (m Map) Password() {
	row, col := m.y+1, m.x+1
	facing := 0
	switch m.facing {
	case "v":
		facing = 1
	case "<":
		facing = 2
	case "^":
		facing = 3
	}
	fmt.Println("Password: ", 1000*row+4*col+facing)
}

func Run() {
	m := ParseInput()
	m.Print()
	m.Walk()
	m.Password()
}
