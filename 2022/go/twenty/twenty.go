package twenty

import (
	"2022/go/utils"
	"container/ring"
	"fmt"
	"strconv"
)

const input = "../input/20.txt"

type IdxToRing = map[int]*ring.Ring

func Mix(idxs IdxToRing) {
	for i := 0; i < len(idxs); i++ {
		r := idxs[i].Prev()
		curr := r.Unlink(1)
		r.Move(curr.Value.(int)).Link(curr)
	}
}

func ParseInput() (IdxToRing, *ring.Ring) {
	lines := utils.ReadLines(input)
	r := ring.New(len(lines))
	idxs := map[int]*ring.Ring{}
	var zero *ring.Ring

	for i, l := range lines {
		nbr, _ := strconv.Atoi(l)
		r.Value = nbr
		idxs[i] = r
		r = r.Next()

		if nbr == 0 {
			zero = r.Prev()
		}
	}

	return idxs, zero
}

func Run() {
	idxs, zeroIdx := ParseInput()
	Mix(idxs)

	sum := zeroIdx.Move(1000).Value.(int) + zeroIdx.Move(2000).Value.(int) + zeroIdx.Move(3000).Value.(int)
	fmt.Println("Sum: ", sum)
}
