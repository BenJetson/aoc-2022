package day02

import (
	"fmt"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type Points int

type Shape Points

const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

func ShapeFromChar(char string) (s Shape, err error) {
	switch char {
	case "A", "X":
		s = Rock
	case "B", "Y":
		s = Paper
	case "C", "Z":
		s = Scissors
	default:
		err = fmt.Errorf("unknown selection character '%s'", char)
	}

	return
}

func ShapeThatLosesAgainst(s Shape) Shape {
	if s == Rock {
		return Scissors
	} else if s == Paper {
		return Rock
	}

	// s == Scissors
	return Paper
}

func ShapeThatWinsAgainst(s Shape) Shape {
	if s == Rock {
		return Paper
	} else if s == Paper {
		return Scissors
	}

	// s == Scissors
	return Rock
}

type Result Points

const (
	Win  Result = 6
	Draw Result = 3
	Lose Result = 0
)

func ResultFromChar(char string) (r Result, err error) {
	switch char {
	case "X":
		r = Lose
	case "Y":
		r = Draw
	case "Z":
		r = Win
	default:
		err = fmt.Errorf("unknown selection character '%s'", char)
	}

	return
}

type Match struct {
	OpponentChoice, MyChoice Shape
	DesiredResult            Result
}

func (m *Match) Result() Result {
	if m.OpponentChoice == m.MyChoice {
		return Draw
	}

	if m.OpponentChoice == ShapeThatLosesAgainst(m.MyChoice) {
		return Win
	}

	return Lose
}

func (m *Match) MyChoiceForDesiredResult() Shape {
	if m.DesiredResult == Win {
		return ShapeThatWinsAgainst(m.OpponentChoice)
	} else if m.DesiredResult == Lose {
		return ShapeThatLosesAgainst(m.OpponentChoice)
	}

	// Draw
	return m.OpponentChoice
}

func (m *Match) Score() Points {
	return Points(m.Result()) + Points(m.MyChoice)
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var strategy []Match
	for index, line := range input {
		choices := strings.Split(line, " ")
		if len(choices) != 2 {
			err = fmt.Errorf("invalid match on line %d", index+1)
			return
		}

		var match Match

		match.OpponentChoice, err = ShapeFromChar(choices[0])
		if err != nil {
			err = fmt.Errorf("opponent choice is invalid on line %d", index+1)
			return
		}

		match.MyChoice, err = ShapeFromChar(choices[1])
		if err != nil {
			err = fmt.Errorf("my choice is invalid on line %d", index+1)
			return
		}

		strategy = append(strategy, match)
	}

	var totalScore Points
	for _, match := range strategy {
		totalScore += match.Score()
	}

	s.Part1.SaveAnswer(int(totalScore))

	strategy = []Match{}
	for index, line := range input {
		chars := strings.Split(line, " ")
		if len(chars) != 2 {
			err = fmt.Errorf("invalid match on line %d", index+1)
			return
		}

		var match Match

		match.OpponentChoice, err = ShapeFromChar(chars[0])
		if err != nil {
			err = fmt.Errorf("opponent choice is invalid on line %d", index+1)
			return
		}

		match.DesiredResult, err = ResultFromChar(chars[1])
		if err != nil {
			err = fmt.Errorf("desired result is invalid on line %d", index+1)
			return
		}

		match.MyChoice = match.MyChoiceForDesiredResult()

		strategy = append(strategy, match)
	}

	totalScore = 0
	for _, match := range strategy {
		totalScore += match.Score()
	}

	s.Part2.SaveAnswer(int(totalScore))

	return
}
