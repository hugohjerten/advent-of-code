package fifteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
	"strings"
)

const input = "../input/15.txt"
const row = 2000000

type Coord struct {
	x int
	y int
}

type Pair = [2]Coord

func ReadInput(input string) []Pair {
	lines := utils.ReadLines(input)
	pairs := make([][2]Coord, len(lines))

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
	}

	return pairs
}

func noBeacons(p Pair, set map[Coord]struct{}) map[Coord]struct{} {
	sx := p[0].x
	sy := p[0].y
	bx := p[1].x
	by := p[1].y
	d := utils.ManhattanDistance(sx, sy, bx, by)

	for x := sx - d; x <= sx+d; x++ {
		if !(x == bx && row == by) &&
			!(x == sx && row == sy) &&
			(utils.ManhattanDistance(sx, sy, x, row) <= d) {
			set[Coord{x, row}] = struct{}{}
		}
	}
	return set
}

func EmptyPositions(ps []Pair) {
	set := map[Coord]struct{}{}
	for _, p := range ps {
		set = noBeacons(p, set)
	}
	fmt.Println("Number positions: ", len(set))
}

func Run() {
	pairs := ReadInput(input)
	EmptyPositions(pairs)
}
