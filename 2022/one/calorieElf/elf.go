package calorieElf

import (
	"2022/utils"
)

type Elf struct {
	calories []int
	sum      int
}

func sum(calories []int) int {
	total := 0
	for _, v := range calories {
		total += v
	}
	return total
}

func GetElves(filePath string) []Elf {
	slice := utils.ReadLines(filePath)
	slices := utils.SeparateSliceOnNewLine(slice)

	elfs := make([]Elf, len(slices))
	for i, slice := range slices {
		calories := utils.Intify(slice)
		elfs[i] = Elf{calories, sum(calories)}
	}
	return elfs
}

func MaxCalories(elfs []Elf) (int, int) {
	max := 0
	index := 0
	for i, elf := range elfs {
		if elf.sum > max {
			max = elf.sum
			index = i
		}
	}

	return max, index
}

func removeElf(elfs []Elf, index int) []Elf {
	ret := make([]Elf, 0)
	ret = append(ret, elfs[:index]...)
	return append(ret, elfs[index+1:]...)
}

func TopThreeCalories(elfs []Elf) int {
	sum := 0
	for i := 0; i < 3; i++ {
		max, index := MaxCalories(elfs)
		sum = sum + max

		elfs = removeElf(elfs, index)
	}

	return sum
}
