package one

import (
	"fmt"

	elf "2022/go/one/calorieElf"
)

func Run() {
	elves := elf.GetElves("one/input.txt")
	max, _ := elf.MaxCalories(elves)
	fmt.Println("Max calories: ", max)

	topThree := elf.TopThreeCalories(elves)
	fmt.Println("Top three calories: ", topThree)
}
