package eleven

import (
	"2022/go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const input = "../input/11.txt"

type Monkey struct {
	inspect int
	items   []int
	update  func(int) int
	throwTo func(int) int
}

type MonkeyInTheMiddle struct {
	ms []Monkey
}

func (mim MonkeyInTheMiddle) MonkeyBusiness() int {
	insp := make([]int, len(mim.ms))
	for i, m := range mim.ms {
		insp[i] = m.inspect
	}

	sort.Slice(insp, func(i, j int) bool {
		return insp[i] > insp[j]
	})

	return insp[0] * insp[1]
}

func (mim *MonkeyInTheMiddle) Rounds(rounds int) {
	// For each round
	for r := 0; r < rounds; r++ {

		// For each monkey
		for mi, monkey := range mim.ms {
			m := monkey

			// For each item
			for _, item := range m.items {
				i := item

				i = m.update(i) / 3
				idx := m.throwTo(i)
				mim.ms[idx].items = append(mim.ms[idx].items, i)

				mim.ms[mi].inspect += 1
			}
			mim.ms[mi].items = nil
		}
	}
}

func parseItems(str string) []int {
	str = strings.ReplaceAll(str, "  Starting items: ", "")
	split := utils.SplitStringOn(str, ", ")

	items := make([]int, 0)
	for _, i := range split {
		item, _ := strconv.Atoi(i)
		items = append(items, item)
	}

	return items
}

func parseOperation(str string) func(int) int {
	str = strings.ReplaceAll(str, "  Operation: new = old ", "")
	split := utils.SplitStringOnWhitespace(str)

	if split[1] == "old" {
		return func(x int) int {
			return x * x
		}
	}

	val, _ := strconv.Atoi(split[1])

	if split[0] == "+" {
		return func(x int) int {
			return x + val
		}
	}
	return func(x int) int {
		return x * val
	}
}

func parseTest(one string, two string, three string) func(int) int {
	div, _ := strconv.Atoi(strings.ReplaceAll(one, "  Test: divisible by ", ""))
	happy, _ := strconv.Atoi(strings.ReplaceAll(two, "    If true: throw to monkey ", ""))
	sad, _ := strconv.Atoi(strings.ReplaceAll(three, "    If false: throw to monkey ", ""))

	return func(x int) int {
		if x%div == 0 {
			return happy
		}
		return sad
	}
}

func GetMonkeys(lines []string) MonkeyInTheMiddle {
	slices := utils.SeparateSliceOnNewLine(lines)
	ms := make([]Monkey, len(slices))

	for i, s := range slices {
		ms[i] = Monkey{0, parseItems(s[1]), parseOperation(s[2]), parseTest(s[3], s[4], s[5])}
	}

	return MonkeyInTheMiddle{ms}
}

func Run() {
	lines := utils.ReadLines(input)
	monkeys := GetMonkeys(lines)
	monkeys.Rounds(20)
	fmt.Println("Monkey Business :", monkeys.MonkeyBusiness())
}
