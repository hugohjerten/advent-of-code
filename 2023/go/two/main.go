package two

func setsOk(red int, green int, blue int, sets []map[string]int) bool {
	for _, s := range sets {
		if !(s["red"] <= red && s["green"] <= green && s["blue"] <= blue) {
			return false
		}
	}
	return true
}

func determineSumOfGames(red int, green int, blue int, games []Game) {
	sum := 0
	for _, g := range games {
		if setsOk(red, green, blue, g.sets) {
			sum += g.id
		}
	}
	println("Part 1: ", sum)
}

func Run() {
	games := parseGames()
	determineSumOfGames(12, 13, 14, games)
}
