package aoc

import (
	"fmt"
	"strconv"
)

type Input []string

type Answer struct {
	Value string
	Valid bool
}

func (a *Answer) SaveAnswer(ans string) {
	if a.Valid {
		panic("attempt to overwrite answer")
	}

	a.Value = ans
	a.Valid = true
}

func (a *Answer) SaveIntAnswer(ans int) {
	a.SaveAnswer(strconv.Itoa(ans))
}

func (a *Answer) String() string {
	if !a.Valid {
		return "blank"
	}
	return a.Value
}

type Solution struct {
	Part1, Part2 Answer
}

func (s *Solution) String() string {
	return fmt.Sprintf("Part one answer is: %s\n", s.Part1.String()) +
		fmt.Sprintf("Part two answer is: %s\n", s.Part2.String())
}

type Solver func(input Input) (Solution, error)

type PuzzleText struct {
	Title string
	Body  []string
}
