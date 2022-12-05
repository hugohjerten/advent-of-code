package game

import (
	"2022/go/utils"
	"fmt"
)

type Shape int
type Strategy int

const (
	Rock Shape = iota
	Paper
	Scissors
)

const (
	Lose Strategy = iota
	Draw
	Win
)

type Round struct {
	shape1 Shape
	shape2 Shape
	score1 int
	score2 int
}

type Game struct {
	rounds       []Round
	withStrategy bool
}

func (g Game) PrintTotalScores() {
	total1 := 0
	total2 := 0
	for _, round := range g.rounds {
		total1 += round.score1
		total2 += round.score2
	}

	fmt.Println("WithStrategy: ", g.withStrategy)
	fmt.Println("Player 1: ", total1)
	fmt.Println("Player 2: ", total2)

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
	strategyMap = map[string]Strategy{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}
)

var (
	scoreMap = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
)

func parseShape(str string) Shape {
	s := shapesMap[str]
	return s
}

func parseStrategy(str string) Strategy {
	s := strategyMap[str]
	return s
}

// Return outcome of round, 0 = loss, 3 = draw, 6 = win
func roundOutcome(one Shape, two Shape) (int, int) {
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

func drawAgainst(shape Shape) Shape {
	return shape
}

func loseAgainst(shape Shape) Shape {
	if shape == Rock {
		return Scissors
	}
	if shape == Paper {
		return Rock
	}
	return Paper
}

func winAgainst(shape Shape) Shape {
	if shape == Rock {
		return Paper
	}
	if shape == Paper {
		return Scissors
	}
	return Rock
}

var (
	determineShape = map[Strategy]func(Shape) Shape{
		Lose: loseAgainst,
		Draw: drawAgainst,
		Win:  winAgainst,
	}
)

func evaluateRound(round []string, part2 bool) Round {
	shape1 := parseShape(round[0])
	var shape2 Shape

	if part2 {
		strategy := parseStrategy(round[1])
		shape2 = determineShape[strategy](shape1)
	} else {
		shape2 = parseShape(round[1])
	}

	outcome1, outcome2 := roundOutcome(shape1, shape2)

	return Round{
		shape1,
		shape2,
		scoreMap[shape1] + outcome1,
		scoreMap[shape2] + outcome2,
	}

}

func GetGame(filePath string, withStrategy bool) Game {
	lines := utils.ReadLines(filePath)
	split := utils.SplitStringsOnWhitespace(lines)

	rounds := make([]Round, len(split))
	for i, round := range split {

		rounds[i] = evaluateRound(round, withStrategy)
	}

	return Game{rounds, withStrategy}
}
