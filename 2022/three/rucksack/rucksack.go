package rucksack

import (
	"2022/utils"
	"strings"
)

type Rucksack struct {
	first    string
	second   string
	allItems string
}

type Group struct {
	rucksacks []Rucksack
}

func GetRucksacks(filePath string) []Rucksack {
	lines := utils.ReadLines(filePath)

	rucksacks := make([]Rucksack, len(lines))
	for i, str := range lines {
		first, second := utils.SplitStringInMiddle(str)
		rucksacks[i] = Rucksack{first, second, str}
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

func groupsOfThrees(rucksacks []Rucksack) []Group {
	nbr := len(rucksacks) / 3
	groups := make([]Group, nbr)
	j := 0
	k := 3
	for i := 0; i < nbr; i++ {
		groups[i] = Group{rucksacks[j:k]}
		j = k
		k += 3
	}

	return groups
}

func (g Group) groupBadge() rune {
	for _, item := range g.rucksacks[0].allItems {
		if strings.ContainsRune(g.rucksacks[1].allItems, item) &&
			strings.ContainsRune(g.rucksacks[2].allItems, item) {
			return item
		}
	}
	panic("Couldn't find badge in group!")
}

func SumOfPriorityOfGroupBadges(rucksacks []Rucksack) int {
	groups := groupsOfThrees(rucksacks)
	sumPriority := 0

	for _, g := range groups {
		sumPriority += itemPriority(g.groupBadge())
	}

	return sumPriority
}
