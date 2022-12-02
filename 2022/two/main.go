package two

import (
	"fmt"

	"2022/two/game"
)

func Run(filePath string) {
	game := game.GetGamePart("two/input.txt")
	total1, total2 := game.TotalScores()
	fmt.Println("Player 1: ", total1)
	fmt.Println("Player 2: ", total2)

}
