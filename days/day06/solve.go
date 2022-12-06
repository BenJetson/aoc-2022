package day06

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type CharSet map[rune]bool

func (cs CharSet) String() string {
	var charStrings []string
	for c := range cs {
		charStrings = append(charStrings, string(c))
	}
	return "[ " + strings.Join(charStrings, ", ") + " ]"
}

func CharSetFromString(s string, startIndex, endIndex int) CharSet {
	cs := make(CharSet)

	for i := startIndex; i < endIndex; i++ {
		c := rune(s[i])
		cs[c] = true
	}

	return cs
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	if len(input) > 1 {
		err = fmt.Errorf("expected one line of input; received %d", len(input))
		return
	}

	signal := input[0]

	if len(signal) < 4 {
		err = errors.New("signal must be at least 4 characters long")
		return
	}

	var endIndexOfMarker int
	for i := 4; i < len(signal); i++ {
		cs := CharSetFromString(signal, i-4, i)
		if len(cs) == 4 {
			endIndexOfMarker = i
			break
		}
	}

	if endIndexOfMarker == 0 {
		err = errors.New("no marker found")
		return
	}

	s.Part1.SaveIntAnswer(endIndexOfMarker)

	return
}
