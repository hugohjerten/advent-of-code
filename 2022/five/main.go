package five

import (
	"2022/utils"
	"fmt"
	"strings"
)

type Crate = string
type Stack []Crate
type Stacks []Stack

type Rearrangement struct {
	amount int
	from   int
	to     int
}

func emptyStacks(nbr int) Stacks {
	stacks := make([]Stack, nbr)
	for i := 0; i < nbr; i++ {
		var stack Stack
		stacks[i] = stack
	}
	return stacks
}

func trimCrate(crate Crate) Crate {
	return strings.Trim(strings.Trim(strings.TrimLeft(crate, " "), "["), "]")
}

func getStacks(filePath string) Stacks {
	lines := utils.ReadLines(filePath)
	nbr := 9
	stacks := emptyStacks(nbr)

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]

		for j := 0; j < nbr; j++ {
			start := j * 4
			end := start + 3

			var str string

			if j == 0 {
				str = line[:end]
			} else if end > len(line) {
				str = line[start:]
			} else {
				str = line[start:end]
			}

			crate := trimCrate(str)

			if crate != "" {
				stacks[j] = append(stacks[j], crate)
			}

		}
	}

	return stacks
}

func trimRearrangement(str string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(str, "move ", ""),
			"from ",
			""),
		"to ",
		"",
	)
}

func getRearrangements(filePath string) []Rearrangement {
	lines := utils.ReadLines(filePath)
	rearrangements := make([]Rearrangement, len(lines))

	for i, str := range lines {
		str = trimRearrangement(str)
		l := utils.Intify(utils.SplitStringOnWhitespace(str))

		// -1 for 0 index
		rearrangements[i] = Rearrangement{l[0], l[1] - 1, l[2] - 1}
	}

	return rearrangements
}

func rearrange9000(stacks Stacks, rearrangements []Rearrangement) Stacks {
	for _, r := range rearrangements {

		for i := 0; i < r.amount; i++ {
			n := len(stacks[r.from]) - 1
			pop := stacks[r.from][n]

			// Remove last element
			stacks[r.from] = stacks[r.from][:n]

			// Add popped element
			stacks[r.to] = append(stacks[r.to], pop)
		}
	}

	return stacks
}

func rearrange9001(stacks Stacks, rearrangements []Rearrangement) Stacks {
	for _, r := range rearrangements {
		n := len(stacks[r.from]) - r.amount
		pop := stacks[r.from][n:]

		// Remove elements
		stacks[r.from] = stacks[r.from][:n]

		// Add popped element
		stacks[r.to] = append(stacks[r.to], pop...)
	}

	return stacks
}

func (s Stacks) topOfStack() string {
	tops := make([]string, 9)
	for i, stack := range s {
		tops[i] = stack[len(stack)-1]
	}

	return strings.Join(tops, "")
}

func Run(stackFilePath string, rearrangementFilePath string) {
	stacks := getStacks(stackFilePath)
	rearrangements := getRearrangements(rearrangementFilePath)

	stacks = rearrange9000(stacks, rearrangements)
	fmt.Println("Top of stacks: ", stacks.topOfStack())

	stacks = getStacks(stackFilePath)
	stacks = rearrange9001(stacks, rearrangements)
	fmt.Println("Top of stacks: ", stacks.topOfStack())
}
