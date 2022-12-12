package twelve

import (
	"2022/go/utils"
	"fmt"
	"sort"
)

const input = "../input/12.txt"

type Path struct {
	p     [][]string
	steps int
}

func (p Path) AddStep(direction string, c utils.Coordinates) *Path {
	new := p
	new.p = nil
	new.p = make([][]string, len(p.p))
	for i, r := range p.p {
		new.p[i] = make([]string, len(r))
		copy(new.p[i], r)
	}
	new.p[c.X][c.Y] = direction
	new.steps += 1

	return &new
}

func (p Path) NewSpot(c utils.Coordinates) bool {
	return p.p[c.X][c.Y] == "."
}

func NewPath(x int, y int) *Path {
	p := make([][]string, x)
	for i := 0; i < x; i++ {
		p[i] = make([]string, y)
		for j := 0; j < y; j++ {
			p[i][j] = "."
		}
	}

	return &Path{p, 0}
}

type HeightMap struct {
	m [][]rune
	e utils.Coordinates
}

func (hm HeightMap) elevationOk(from utils.Coordinates, to utils.Coordinates) bool {
	return hm.m[to.X][to.Y] <= hm.m[from.X][from.Y]+1
}

func (hm HeightMap) coordinateOk(c utils.Coordinates) bool {
	return (c.X >= 0 && c.Y >= 0) && (c.X < len(hm.m) && c.Y < len(hm.m[0]))
}

func (hm HeightMap) try(c utils.Coordinates, direction string, p *Path) *Path {
	var n utils.Coordinates
	switch direction {
	case ">":
		n = utils.Coordinates{X: c.X + 1, Y: c.Y}
	case "<":
		n = utils.Coordinates{X: c.X - 1, Y: c.Y}
	case "v":
		n = utils.Coordinates{X: c.X, Y: c.Y + 1}
	case "^":
		n = utils.Coordinates{X: c.X, Y: c.Y - 1}
	}

	if !hm.coordinateOk(n) || !hm.elevationOk(c, n) || !p.NewSpot(n) {
		return nil
	}

	new := p.AddStep(direction, n)
	return hm.Find(n, new)
}

func (hm HeightMap) Find(c utils.Coordinates, p *Path) *Path {
	if c == hm.e {
		p.p[c.X][c.Y] = "E"
		return p
	}

	paths := make([]*Path, 4)
	paths[0] = hm.try(c, ">", p)
	paths[1] = hm.try(c, "<", p)
	paths[2] = hm.try(c, "v", p)
	paths[3] = hm.try(c, "^", p)

	sort.Slice(paths, func(i, j int) bool {
		return (paths[i] == nil && paths[j] == nil) || (paths[i] == nil) || ((paths[j] != nil) && paths[i].steps < paths[j].steps)
	})

	return paths[0]
}

func NewHeightMap(lines []string) (HeightMap, utils.Coordinates) {
	m := make([][]rune, len(lines))
	var c utils.Coordinates
	var e utils.Coordinates

	for i, l := range lines {
		m[i] = make([]rune, len(l))
		for j, r := range l {

			switch r {
			case 83: // S
				r = 97 // a
				c = utils.Coordinates{X: i, Y: j}
			case 69: // E
				r = 122 // z
				e = utils.Coordinates{X: i, Y: j}
			}

			m[i][j] = r
		}
	}

	return HeightMap{m, e}, c
}

func Run() {
	lines := utils.ReadLines(input)
	hm, start := NewHeightMap(lines)
	p := NewPath(len(hm.m), len(hm.m[0]))
	path := hm.Find(start, p)
	for _, r := range path.p {
		fmt.Println(r)
	}

}
