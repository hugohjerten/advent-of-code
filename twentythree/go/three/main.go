package three

import (
	"fmt"
)

type PartNumber struct {
	value          int
	adjacent       bool
	symbol         bool
	gear           bool
	adjacentValues map[int]struct{}
}

type Schematic struct {
	mx   [][]int
	nbrs []PartNumber
}

// Return true if is symbol at location
func (s *Schematic) isSymbol(i int, j int) bool {
	// If outside grid, skip
	if i < 0 || i == len(s.mx) || j < 0 || j >= len(s.mx[0]) {
		return false
	}

	// If symbol, return true
	if s.mx[i][j] != -1 && s.nbrs[s.mx[i][j]].symbol {
		return true
	}

	return false
}

// Return value if is number at location
func (s *Schematic) isNumber(i int, j int) (bool, int) {
	// If outside grid, skip
	if i < 0 || i == len(s.mx) || j < 0 || j >= len(s.mx[0]) {
		return false, 0
	}

	// If number, return true
	if s.mx[i][j] != -1 && !s.nbrs[s.mx[i][j]].symbol {
		return true, s.nbrs[s.mx[i][j]].value
	}

	return false, 0
}

// Check adjacent slots
func (s *Schematic) checkNeighbours(i int, j int, part2 bool) {
	// Iterate over 3x3 quadrant around point
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {

			// Skip itself
			if k != i || l != j {

				if !part2 {
					// Part 1 identify symbols
					if s.isSymbol(k, l) {
						s.nbrs[s.mx[i][j]].adjacent = true
					}

				} else {
					// Part 2 identify numbers
					isNumber, number := s.isNumber(k, l)
					if isNumber {
						s.nbrs[s.mx[i][j]].adjacentValues[number] = struct{}{}
					}
				}
			}
		}
	}
}

func (s *Schematic) identifyValidParts() {
	for i := 0; i < len(s.mx); i++ {
		for j := 0; j < len(s.mx[0]); j++ {

			// If number, check if number has symbol neighbour
			if s.mx[i][j] != -1 && !s.nbrs[s.mx[i][j]].symbol {
				s.checkNeighbours(i, j, false)
			}
		}
	}

	sum := 0
	for _, nbr := range s.nbrs {
		if nbr.adjacent {
			sum += nbr.value
		}
	}

	fmt.Println("Part 1: ", sum)
}

func (s *Schematic) identifyRelevantGears() {
	for i := 0; i < len(s.mx); i++ {
		for j := 0; j < len(s.mx[0]); j++ {

			// If gear
			if s.mx[i][j] != -1 && s.nbrs[s.mx[i][j]].gear {
				s.checkNeighbours(i, j, true)
			}
		}
	}

	sum := 0
	for _, nbr := range s.nbrs {
		if len(nbr.adjacentValues) == 2 {
			product := 1
			for value := range nbr.adjacentValues {
				product *= value
			}
			sum += product
		}
	}

	fmt.Println("Part 2: ", sum)
}

func Run() {
	schematic := parseInput()
	schematic.identifyValidParts()
	schematic.identifyRelevantGears()
}
