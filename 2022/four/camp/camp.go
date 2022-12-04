package camp

import (
	"2022/utils"
	"strconv"
)

type Pair struct {
	first  []int
	second []int
}

// Expect a string of format X-Y, eg "4-7", which will return [4,5,6,7]
func getInterval(str string) []int {
	split := utils.SplitStringOn(str, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	if start == end {
		interval := make([]int, 1)
		interval[0] = start
		return interval
	}

	var interval []int
	for i := start; i < end+1; i++ {
		interval = append(interval, i)
	}

	return interval
}

func GetPairs(filePath string) []Pair {
	lines := utils.ReadLines(filePath)
	pairs := make([]Pair, len(lines))

	for i, str := range lines {
		split := utils.SplitStringOn(str, ",")
		pairs[i] = Pair{getInterval(split[0]), getInterval(split[1])}
	}

	return pairs
}

func (p Pair) OneRangeFullyContainOther() bool {
	fullyContains := true
	for _, id := range p.first {
		if !(utils.ContainsInt(p.second, id)) {
			fullyContains = false
			break
		}
	}

	if fullyContains {
		return true
	}

	for _, id := range p.second {
		if !(utils.ContainsInt(p.first, id)) {
			return false
		}
	}

	return true
}

func NumberAssignmentsWithRangeFullyContainOther(pairs []Pair) int {
	nbr := 0
	for _, p := range pairs {
		if p.OneRangeFullyContainOther() {
			nbr += 1
		}
	}
	return nbr
}
