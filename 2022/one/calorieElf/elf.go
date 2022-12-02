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

	elves := make([]Elf, len(slices))
	for i, slice := range slices {
		calories := utils.Intify(slice)
		elves[i] = Elf{calories, sum(calories)}
	}
	return elves
}

func MaxCalories(elves []Elf) (int, int) {
	max := 0
	index := 0
	for i, elf := range elves {
		if elf.sum > max {
			max = elf.sum
			index = i
		}
	}

	return max, index
}

func removeElf(elves []Elf, index int) []Elf {
	ret := make([]Elf, 0)
	ret = append(ret, elves[:index]...)
	return append(ret, elves[index+1:]...)
}

func TopThreeCalories(elves []Elf) int {
	sum := 0
	for i := 0; i < 3; i++ {
		max, index := MaxCalories(elves)
		sum = sum + max

		elves = removeElf(elves, index)
	}

	return sum
}
