package two

func setsOk(red int, green int, blue int, sets []map[string]int) bool {
	for _, s := range sets {
		if !(s["red"] <= red && s["green"] <= green && s["blue"] <= blue) {
			return false
		}
	}
	return true
}

func sumOfGames(red int, green int, blue int, games []Game) {
	sum := 0
	for _, g := range games {
		if setsOk(red, green, blue, g.sets) {
			sum += g.id
		}
	}
	println("Part 1: ", sum)
}

func minimumPower(sets []map[string]int) int {
	var red, green, blue int = 0, 0, 0
	for _, s := range sets {
		if s["red"] > red {
			red = s["red"]
		}
		if s["green"] > green {
			green = s["green"]
		}
		if s["blue"] > blue {
			blue = s["blue"]
		}
	}
	return red * green * blue
}

func sumOfPowerOfSets(games []Game) {
	sum := 0
	for _, g := range games {
		sum += minimumPower(g.sets)
	}
	println("Part 2: ", sum)
}

func Run() {
	games := parseGames()
	sumOfGames(12, 13, 14, games)
	sumOfPowerOfSets(games)
}
