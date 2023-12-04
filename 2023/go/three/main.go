package three

import (
	"fmt"
)

type PartNumber struct {
	value    int
	adjacent bool
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

	// If "-2", i.e. symbol, return true
	if s.mx[i][j] == -2 {
		return true
	}

	return false
}

// Check adjacent slots
func (s *Schematic) checkNeighbours(i int, j int) {
	// Iterate over 3x3 quadrant around point
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {

			// Skip itself
			if k != i || l != j {
				if s.isSymbol(k, l) {
					s.nbrs[s.mx[i][j]].adjacent = true
				}
			}
		}
	}
}

func (s *Schematic) identifyValidParts() {
	for i := 0; i < len(s.mx); i++ {
		for j := 0; j < len(s.mx[0]); j++ {

			// If number, check if number has symbol neighbour
			if s.mx[i][j] != -1 && s.mx[i][j] != -2 {
				s.checkNeighbours(i, j)
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

func Run() {
	schematic := parseInput()
	schematic.identifyValidParts()
}
