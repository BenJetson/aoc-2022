package day03

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/BenJetson/aoc-2022/aoc"
)

type Item rune

func (i Item) Priority() int {
	if unicode.IsUpper(rune(i)) {
		return int(rune(i)-'A') + 1 + 26
	}
	return int(rune(i)-'a') + 1
}

type Compartment map[Item]int

func (c Compartment) String() string {
	var out []string
	for item, count := range c {
		for i := 0; i < count; i++ {
			out = append(out, string(item))
		}
	}
	sort.Strings(out)
	return strings.Join(out, "")
}

type Rucksack struct {
	First, Second Compartment
}

func (s *Rucksack) String() string {
	return s.First.String() + " | " + s.Second.String()
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var sacks []Rucksack
	for index, line := range input {
		if len(line)%2 != 0 {
			err = fmt.Errorf("input line %d has odd length", index+1)
		}

		var items []Item
		for _, r := range line {
			items = append(items, Item(r))
		}

		var sack Rucksack

		sack.First = make(Compartment)
		for _, item := range items[:len(items)/2] {
			sack.First[item] += 1
		}

		sack.Second = make(Compartment)
		for _, item := range items[len(items)/2:] {
			sack.Second[item] += 1
		}

		sacks = append(sacks, sack)
	}

	var totalPriority int

	for _, sack := range sacks {
		for item := range sack.First {
			if _, ok := sack.Second[item]; ok {
				totalPriority += item.Priority()
				break
			}
		}
	}

	s.Part1.SaveAnswer(totalPriority)

	return
}
