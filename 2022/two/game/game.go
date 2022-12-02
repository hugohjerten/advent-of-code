package game

import (
	"2022/utils"
)

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

type Round struct {
	shape1 Shape
	shape2 Shape
	score1 int
	score2 int
}

type Game struct {
	rounds []Round
}

func (g Game) TotalScores() (int, int) {
	total1 := 0
	total2 := 0
	for _, round := range g.rounds {
		total1 += round.score1
		total2 += round.score2
	}

	return total1, total2
}

var (
	shapesMap = map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
)

var (
	scoreMap = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
)

func check(ok bool, str string) {
	if !ok {
		panic("Could not parse Shape!")
	}
}

func ParseShape(str string) Shape {
	c, ok := shapesMap[str]
	check(ok, str)
	return c
}

// Return outcome of round, 0 = loss, 3 = draw, 6 = win
func roundOutcome1(one Shape, two Shape) (int, int) {
	if one == two {
		return 3, 3
	}
	if one == Rock && two == Scissors {
		return 6, 0
	}
	if one == Scissors && two == Paper {
		return 6, 0
	}
	if one == Paper && two == Rock {
		return 6, 0
	}
	return 0, 6
}

func evaluateRound1(round []string) Round {
	if len(round) != 2 {
		panic("Bad number of shapes!")
	}

	shape1 := ParseShape(round[0])
	shape2 := ParseShape(round[1])
	outcome1, outcome2 := roundOutcome1(shape1, shape2)

	return Round{
		shape1,
		shape2,
		scoreMap[shape1] + outcome1,
		scoreMap[shape2] + outcome2,
	}

}

func GetGamePart(filePath string) Game {
	lines := utils.ReadLines(filePath)
	split := utils.SplitStringsOnWhitespace(lines)

	rounds := make([]Round, len(split))
	for i, round := range split {
		rounds[i] = evaluateRound1(round)
	}

	return Game{rounds}
}
