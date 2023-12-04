package three

import (
	"2023/go/utils"
	"unicode"
)

const input = "../input/3.txt"

func parseInput() Schematic {
	ls := utils.ReadLines(input)
	var nbrs []PartNumber
	mx := make([][]int, len(ls))

	for i, l := range ls {
		mx[i] = make([]int, len(l))
		var previousDigit = false

		for j, r := range l {

			if unicode.IsDigit(r) {
				// If character is digit

				if !previousDigit {
					// If previous character was not digit, add to nbrs
					nbrs = append(nbrs, PartNumber{
						value:          int(r - '0'),
						adjacent:       false,
						symbol:         false,
						gear:           false,
						adjacentValues: map[int]struct{}{},
					})

				} else {
					// If previous was digit also, update value in nbrs
					nbrs[len(nbrs)-1].value = nbrs[len(nbrs)-1].value*10 + int(r-'0')
				}

				previousDigit = true

				// Add index reference to nbrs
				mx[i][j] = len(nbrs) - 1

			} else if r == 46 {
				// If character is '.' add "-1"
				mx[i][j] = -1
				previousDigit = false
			} else {
				// If character is "symbol"
				gear := false

				// If symbol is gear, i.e. "*"
				if r == 42 {
					gear = true
				}

				nbrs = append(nbrs, PartNumber{
					value:          int(r - '0'),
					adjacent:       false,
					symbol:         true,
					gear:           gear,
					adjacentValues: map[int]struct{}{},
				})

				// Add index reference to nbrs
				mx[i][j] = len(nbrs) - 1

				previousDigit = false
			}
		}
	}
	return Schematic{mx: mx, nbrs: nbrs}
}
