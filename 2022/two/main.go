package two

import (
	"2022/two/game"
)

func Run(filePath string) {
	g := game.GetGame("two/input.txt", false)
	g.PrintTotalScores()

	g = game.GetGame("two/input.txt", true)
	g.PrintTotalScores()

}
