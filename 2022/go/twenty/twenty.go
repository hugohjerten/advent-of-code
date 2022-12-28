package twenty

import (
	"2022/go/utils"
	"container/ring"
	"fmt"
	"strconv"
)

const input = "../input/20.txt"

type IdxToRing = map[int]*ring.Ring

func Mix(idxs IdxToRing, zeroIdx *ring.Ring, times int) {
	for k := 0; k < times; k++ {
		for i := 0; i < len(idxs); i++ {
			r := idxs[i].Prev()
			curr := r.Unlink(1)
			r.Move(curr.Value.(int) % (len(idxs) - 1)).Link(curr)
		}
	}

	sum := zeroIdx.Move(1000).Value.(int) + zeroIdx.Move(2000).Value.(int) + zeroIdx.Move(3000).Value.(int)
	fmt.Println("Sum: ", sum)
}

func ParseInput(key int) (IdxToRing, *ring.Ring) {
	lines := utils.ReadLines(input)
	r := ring.New(len(lines))
	idxs := map[int]*ring.Ring{}
	var zero *ring.Ring

	for i, l := range lines {
		nbr, _ := strconv.Atoi(l)
		r.Value = nbr * key
		idxs[i] = r
		r = r.Next()

		if nbr == 0 {
			zero = r.Prev()
		}
	}

	return idxs, zero
}

func Run() {
	// Part 1
	idxs, zeroIdx := ParseInput(1)
	Mix(idxs, zeroIdx, 1)

	// Part 2
	idxs, zeroIdx = ParseInput(811589153)
	Mix(idxs, zeroIdx, 10)
}
