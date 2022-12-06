package two

import (
	"2022/go/two/game"
)

func Run() {
	g := game.GetGame("two/input.txt", false)
	g.PrintTotalScores()

	g = game.GetGame("two/input.txt", true)
	g.PrintTotalScores()

}
