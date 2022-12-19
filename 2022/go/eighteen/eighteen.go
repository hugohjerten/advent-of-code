package eighteen

import (
	"2022/go/utils"
	"fmt"
	"strconv"
)

const input = "../input/18.txt"

func ParseInput(input string) [][][]bool {
	coordinates := make([][]int, 0)
	lines := utils.ReadLines(input)

	max := []int{0, 0, 0}

	for _, l := range lines {
		split := utils.SplitStringOn(l, ",")
		cs := make([]int, 3)
		for i, s := range split {
			k, _ := strconv.Atoi(s)
			cs[i] = k

			if k > max[i] {
				max[i] = k
			}
		}
		coordinates = append(coordinates, cs)
	}

	axis := make([][][]bool, max[0]+1)
	for x := 0; x <= max[0]; x++ {
		axis[x] = make([][]bool, max[1]+1)
		for y := 0; y <= max[1]; y++ {
			axis[x][y] = make([]bool, max[2]+1)
			for z := 0; z <= max[2]; z++ {
				axis[x][y][z] = false
			}
		}
	}

	for _, c := range coordinates {
		axis[c[0]][c[1]][c[2]] = true
	}

	return axis
}

func countSurfaces(x int, y int, z int, axis [][][]bool) int {
	cnt := 0
	if x == 0 || (x-1 >= 0 && !axis[x-1][y][z]) {
		cnt += 1
	}
	if x == len(axis)-1 || (x+1 <= len(axis)-1 && !axis[x+1][y][z]) {
		cnt += 1
	}
	if y == 0 || (y-1 >= 0 && !axis[x][y-1][z]) {
		cnt += 1
	}
	if y == len(axis[0])-1 || (y+1 <= len(axis[0])-1 && !axis[x][y+1][z]) {
		cnt += 1
	}
	if z == 0 || (z-1 >= 0 && !axis[x][y][z-1]) {
		cnt += 1
	}
	if z == len(axis[0][0])-1 || (z+1 <= len(axis[0][0])-1 && !axis[x][y][z+1]) {
		cnt += 1
	}

	return cnt
}

func SurfaceArea(axis [][][]bool) {
	cnt := 0
	for x := 0; x < len(axis); x++ {
		for y := 0; y < len(axis[0]); y++ {
			for z := 0; z < len(axis[0][0]); z++ {
				if axis[x][y][z] {
					cnt += countSurfaces(x, y, z, axis)
				}
			}
		}
	}
	fmt.Println("Surfaces: ", cnt)
}

func Run() {
	axis := ParseInput(input)
	SurfaceArea(axis)
}
