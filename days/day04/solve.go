package day04

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type SectionID int
type Assignment struct {
	Min, Max SectionID
}

func (a *Assignment) IsWithinRange(s SectionID) bool {
	return s >= a.Min && s <= a.Max
}

type Pair struct {
	A, B Assignment
}

func (p *Pair) DoesPartiallyOverlap() bool {
	return p.A.IsWithinRange(p.B.Min) || p.A.IsWithinRange(p.B.Max) ||
		p.B.IsWithinRange(p.A.Min) || p.B.IsWithinRange(p.A.Max)
}

func (p *Pair) DoesFullyOverlap() bool {
	return (p.A.IsWithinRange(p.B.Min) && p.A.IsWithinRange(p.B.Max)) ||
		(p.B.IsWithinRange(p.A.Min) && p.B.IsWithinRange(p.A.Max))
}

func AssignmentFromString(str string) (a Assignment, err error) {
	idStrings := strings.Split(str, "-")
	if len(idStrings) != 2 {
		err = errors.New("expected 2 section ID strings")
		return
	}

	minID, err := strconv.Atoi(idStrings[0])
	if err != nil {
		err = fmt.Errorf("invalid min ID: %w", err)
	}

	maxID, err := strconv.Atoi(idStrings[1])
	if err != nil {
		err = fmt.Errorf("invalid max ID: %w", err)
	}

	a.Min = SectionID(minID)
	a.Max = SectionID(maxID)

	return
}

func PairFromLine(line string) (p Pair, err error) {
	assignmentStrings := strings.Split(line, ",")
	if len(assignmentStrings) != 2 {
		err = errors.New("expected 2 assignment strings")
		return
	}

	if p.A, err = AssignmentFromString(assignmentStrings[0]); err != nil {
		err = fmt.Errorf("assignment A parse error: %w", err)
		return
	}
	if p.B, err = AssignmentFromString(assignmentStrings[1]); err != nil {
		err = fmt.Errorf("assignment B parse error: %w", err)
		return
	}

	return
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var pairs []Pair

	for index, line := range input {
		var p Pair
		p, err = PairFromLine(line)
		if err != nil {
			err = fmt.Errorf("problem on line %d: %w", index+1, err)
			return
		}
		pairs = append(pairs, p)
	}

	var fullOverlapCount, partialOverlapCount int
	for _, p := range pairs {
		if p.DoesFullyOverlap() {
			fullOverlapCount++
			partialOverlapCount++
		} else if p.DoesPartiallyOverlap() {
			partialOverlapCount++
		}
	}

	s.Part1.SaveAnswer(fullOverlapCount)
	s.Part2.SaveAnswer(partialOverlapCount)

	return
}
