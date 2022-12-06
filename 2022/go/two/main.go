package two

import (
	"2022/go/two/game"
)

const input = "../input/2.txt"

func Run() {
	g := game.GetGame(input, false)
	g.PrintTotalScores()

	g = game.GetGame(input, true)
	g.PrintTotalScores()

}
