package rucksack

import (
	"2022/utils"
	"strings"
)

type Rucksack struct {
	first  string
	second string
}

func GetRucksacks(filePath string) []Rucksack {
	lines := utils.ReadLines(filePath)

	rucksacks := make([]Rucksack, len(lines))
	for i, str := range lines {
		first, second := utils.SplitStringInMiddle(str)
		rucksacks[i] = Rucksack{first, second}
	}

	return rucksacks

}

func (r Rucksack) findBadItem() rune {
	for _, item := range r.first {
		if strings.ContainsRune(r.second, item) {
			return item
		}
	}
	panic("Couldn't find any bad items in rucksack!")
}

func itemPriority(item rune) int {
	if (96 < item) && (item < 123) {
		// a - z
		return int(item) - 96

	} else if (64 < item) && (item < 91) {
		// A - Z
		return int(item) - 64 + 26

	} else {
		panic("Bad Item!")
	}
}

func SumOfPriorityOfBadItems(rucksacks []Rucksack) int {
	sumPriority := 0
	for _, r := range rucksacks {
		item := r.findBadItem()
		sumPriority += itemPriority(item)
	}

	return sumPriority
}
