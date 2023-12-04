package two

import (
	"strings"
	"twentythree/go/utils"
)

const input = "../input/2.txt"

type Game struct {
	id   int
	sets []map[string]int
}

func parseSet(s string) map[string]int {
	colours := map[string]int{"blue": 0, "red": 0, "green": 0}

	splitSet := utils.SplitStringOn(s, ",")
	for i := 0; i < len(splitSet); i++ {
		trimmed := strings.Trim(splitSet[i], " ")
		splitValue := utils.SplitStringOnWhitespace(trimmed)
		nbr := utils.IntifyString(splitValue[0])
		colour := splitValue[1]

		colours[colour] = nbr
	}
	return colours
}

func parseSets(s string) []map[string]int {
	split := utils.SplitStringOn(s, ";")
	sets := make([]map[string]int, len(split))
	for i := 0; i < len(split); i++ {
		sets[i] = parseSet(split[i])
	}

	return sets
}

func parseGameId(s string) int {
	return utils.IntifyString(utils.SplitStringOnWhitespace(s)[1])
}

func parseGame(line string) Game {
	split := utils.SplitStringOn(line, ":")
	id := parseGameId(split[0])
	sets := parseSets(split[1])

	return Game{id: id, sets: sets}
}

func parseGames() []Game {
	lines := utils.ReadLines(input)
	games := make([]Game, len(lines))

	for i := 0; i < len(lines); i++ {
		games[i] = parseGame(lines[i])
	}

	return games
}
