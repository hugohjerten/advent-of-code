package one

import (
	"fmt"

	elf "2022/go/one/calorieElf"
)

const input = "../input/1.txt"

func Run() {
	elves := elf.GetElves(input)
	max, _ := elf.MaxCalories(elves)
	fmt.Println("Max calories: ", max)

	topThree := elf.TopThreeCalories(elves)
	fmt.Println("Top three calories: ", topThree)
}
