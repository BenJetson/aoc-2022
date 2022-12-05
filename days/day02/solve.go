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

type Result Points

const (
	Win  Result = 6
	Draw Result = 3
	Lose Result = 0
)

type Match struct {
	OpponentChoice, MyChoice Shape
}

func (m *Match) Result() Result {
	if m.OpponentChoice == m.MyChoice {
		return Draw
	}

	if m.OpponentChoice == Rock && m.MyChoice == Paper ||
		m.OpponentChoice == Paper && m.MyChoice == Scissors ||
		m.OpponentChoice == Scissors && m.MyChoice == Rock {
		return Win
	}

	return Lose
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

	return
}
