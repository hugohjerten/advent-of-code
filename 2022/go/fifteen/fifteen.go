package fifteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
	"strings"
)

const input = "../input/15.txt"

type Coord struct {
	x int
	y int
}

type Pair = [2]Coord

type Tunnels struct {
	m       [][]rune
	w       int
	h       int
	xOffset int
	yOffset int
	pairs   []Pair
}

func (t Tunnels) Show() {
	for _, r := range t.m {
		fmt.Println(string(r))
	}
}

func (t *Tunnels) addSensorsAndBeacons(cs []Pair) {
	for _, p := range cs {
		sx := p[0].x - t.xOffset
		sy := p[0].y - t.yOffset
		bx := p[1].x - t.xOffset
		by := p[1].y - t.yOffset

		t.m[sy][sx] = 83 // "S"
		t.m[by][bx] = 66 // "B"
	}
}

func padding(nbr int) []rune {
	padding := make([]rune, nbr)
	for k := 0; k < nbr; k++ {
		padding[k] = 46 // "."
	}
	return padding
}

func (t *Tunnels) expandLeft(nbr int) {
	for i, row := range t.m {
		t.m[i] = append(padding(nbr), row...)
	}
	t.w += nbr
	t.xOffset -= nbr
}

func (t *Tunnels) expandRight(nbr int) {
	for i, row := range t.m {
		t.m[i] = append(row, padding(nbr)...)
	}
	t.w += nbr
}

func (t *Tunnels) expandUp(nbr int) {
	pad := make([][]rune, nbr)
	for i := 0; i < nbr; i++ {
		pad[i] = padding(t.w)
	}
	t.m = append(pad, t.m...)
	t.h += nbr
	t.yOffset -= nbr
}

func (t *Tunnels) expandDown(nbr int) {
	pad := make([][]rune, nbr)
	for i := 0; i < nbr; i++ {
		pad[i] = padding(t.w)
	}
	t.m = append(t.m, pad...)
	t.h += nbr
}

func (t *Tunnels) noBeacons(c Coord, d int) {
	if c.x-t.xOffset-d < 0 {
		t.expandLeft(utils.Abs(c.x - t.xOffset - d))
	}
	if c.x-t.xOffset+d >= t.w {
		t.expandRight((c.x - t.xOffset + d) - t.w + 1)
	}
	if c.y-t.yOffset-d < 0 {
		t.expandUp(utils.Abs(c.y - t.yOffset - d))
	}
	if c.y-t.yOffset+d >= t.h {
		t.expandDown((c.y - t.yOffset + d) - t.h + 1)
	}
	sx := c.x - t.xOffset
	sy := c.y - t.yOffset

	for y := sy - d; y <= sy+d; y++ {
		for x := sx - d; x <= sx+d; x++ {
			if (utils.ManhattanDistance(sx, sy, x, y) <= d) &&
				t.m[y][x] == 46 { // "."

				t.m[y][x] = 35 // "#"
			}
		}
	}
}

func (t *Tunnels) EmptyPositions(row int, show bool) {
	for _, p := range t.pairs {
		d := utils.ManhattanDistance(p[0].x, p[0].y, p[1].x, p[1].y)
		t.noBeacons(p[0], d)
	}

	row -= t.yOffset
	cnt := 0
	for _, r := range t.m[row] {
		if r == 35 { // "#"
			cnt += 1
		}
	}

	if show {
		t.Show()
	}

	fmt.Println("Empty positions: ", cnt)
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

func ReadInput(input string) Tunnels {
	lines := utils.ReadLines(input)
	pairs := make([][2]Coord, len(lines))

	xEdges := [2]int{0, 0}
	yEdges := [2]int{0, 0}

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
		pairs[i] = [2]Coord{{sensorX, sensorY}, {beaconX, beaconY}}

		xEdges = minMax(xEdges, sensorX)
		xEdges = minMax(xEdges, beaconX)
		yEdges = minMax(yEdges, sensorY)
		yEdges = minMax(yEdges, beaconY)
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

	tunnels := Tunnels{m, w, h, xEdges[0], yEdges[0], pairs}
	tunnels.addSensorsAndBeacons(pairs)

	return tunnels
}

func Run() {
	tunnels := ReadInput(input)
	tunnels.EmptyPositions(10, true)

}
