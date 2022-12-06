package day06

import (
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

func FindUniqueStreak(signal string, quantity int) (endIndex int, err error) {
	if len(signal) < quantity {
		err = fmt.Errorf("signal must be at least %d characters long", quantity)
		return
	}

	for i := quantity; i < len(signal); i++ {
		cs := CharSetFromString(signal, i-quantity, i)
		if len(cs) == quantity {
			endIndex = i
			return
		}
	}

	err = fmt.Errorf("no unique streak of %d characters found", quantity)
	return
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	if len(input) > 1 {
		err = fmt.Errorf("expected one line of input; received %d", len(input))
		return
	}

	signal := input[0]

	startOfPacketIndex, err := FindUniqueStreak(signal, 4)
	if err != nil {
		err = fmt.Errorf("could not find start of packet: %w", err)
		return
	}

	s.Part1.SaveIntAnswer(startOfPacketIndex)

	startOfMessageIndex, err := FindUniqueStreak(signal, 14)
	if err != nil {
		err = fmt.Errorf("could not find start of message: %w", err)
		return
	}

	s.Part2.SaveIntAnswer(startOfMessageIndex)

	return
}
