package three

import (
	sack "2022/go/three/rucksack"
	"fmt"
)

func Run() {
	rucksacks := sack.GetRucksacks("three/input.txt")
	sumPriority := sack.SumOfPriorityOfBadItems(rucksacks)
	fmt.Println("Sum of priority of item types: ", sumPriority)

	sumPriority = sack.SumOfPriorityOfGroupBadges(rucksacks)
	fmt.Println("Sum of priority of group badges: ", sumPriority)
}
