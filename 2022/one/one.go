package one

import (
	"fmt"

	elf "2022/one/calorieElf"
)

func Run(filePath string) {
	elves := elf.GetElves(filePath)
	max, _ := elf.MaxCalories(elves)
	fmt.Println("Max calories: ", max)

	topThree := elf.TopThreeCalories(elves)
	fmt.Println("Top three calories: ", topThree)
}
