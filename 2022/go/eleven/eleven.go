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
	ms       []Monkey
	divisors []int
	worried  bool
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
	// Lowest Common Multiple
	lcm := utils.LCM(mim.divisors[0], mim.divisors[1], mim.divisors[2:]...)

	// For each round
	for r := 0; r < rounds; r++ {

		// For each monkey
		for mi, monkey := range mim.ms {
			m := monkey

			// For each item
			for _, item := range m.items {
				mim.ms[mi].inspect += 1
				i := item

				i = m.update(i)
				if !mim.worried {
					i /= 3
				} else {
					i = i % lcm
				}

				idx := m.throwTo(i)
				mim.ms[idx].items = append(mim.ms[idx].items, i)
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

func parseTestAndDivisor(one string, two string, three string) (func(int) int, int) {
	div, _ := strconv.Atoi(strings.ReplaceAll(one, "  Test: divisible by ", ""))
	happy, _ := strconv.Atoi(strings.ReplaceAll(two, "    If true: throw to monkey ", ""))
	sad, _ := strconv.Atoi(strings.ReplaceAll(three, "    If false: throw to monkey ", ""))

	return func(x int) int {
		if x%div == 0 {
			return happy
		}
		return sad
	}, div
}

func GetMonkeys(lines []string, worried bool) MonkeyInTheMiddle {
	slices := utils.SeparateSliceOnNewLine(lines)
	ms := make([]Monkey, len(slices))
	divisors := make([]int, len(slices))

	for i, s := range slices {
		test, div := parseTestAndDivisor(s[3], s[4], s[5])
		divisors[i] = div
		ms[i] = Monkey{0, parseItems(s[1]), parseOperation(s[2]), test}
	}

	return MonkeyInTheMiddle{ms, divisors, worried}
}

func Run() {
	lines := utils.ReadLines(input)
	monkeys := GetMonkeys(lines, false)
	monkeys.Rounds(20)
	fmt.Println("Monkey Business :", monkeys.MonkeyBusiness())

	monkeys = GetMonkeys(lines, true)
	monkeys.Rounds(10000)
	fmt.Println("Monkey Business :", monkeys.MonkeyBusiness())
}
