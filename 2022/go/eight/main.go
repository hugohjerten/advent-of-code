package eight

import (
	"2022/go/utils"
	"fmt"
)

type Grid = [][]int

func containsTrue(s []bool) bool {
	for _, a := range s {
		if a {
			return true
		}
	}
	return false
}

func isVisible(x int, y int, g Grid) bool {
	h := g[x][y]

	vis := []bool{true, true, true, true}

	// Above
	for i := 0; i < x; i++ {
		if g[i][y] >= h {
			vis[0] = false
			break
		}
	}

	// Below
	for i := x + 1; i < len(g); i++ {
		if g[i][y] >= h {
			vis[1] = false
			break
		}
	}

	// Left
	for j := 0; j < y; j++ {
		if g[x][j] >= h {
			vis[2] = false
			break
		}
	}

	// Right
	for j := y + 1; j < len(g); j++ {
		if g[x][j] >= h {
			vis[3] = false
			break
		}
	}

	return containsTrue(vis)

}

func VisibleTrees(g Grid) {
	cnt := 0
	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[0])-1; j++ {

			if isVisible(i, j, g) {
				cnt += 1
			}
		}
	}

	// outer edge
	cnt += len(g)*2 + (len(g[0])-2)*2

	fmt.Println("COUNT: ", cnt)
}

func parseTrees(filePath string) Grid {
	ls := utils.ReadLines(input)
	ts := make([][]int, len(ls))

	for i, l := range ls {
		ts[i] = make([]int, len(l))

		for j, h := range l {
			ts[i][j] = int(h - '0')
		}
	}

	return ts
}

const input = "../input/8.txt"

func Run() {
	grid := parseTrees(input)
	VisibleTrees(grid)
}
