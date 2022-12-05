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

func FirstCommonItem(compartments ...Compartment) *Item {
	common := make(Compartment)

	for _, current := range compartments {
		for item := range current {
			common[item] += 1
		}
	}

	for item, count := range common {
		if count == len(compartments) {
			return &item
		}
	}
	return nil
}

type Rucksack struct {
	First, Second Compartment
}

func (s *Rucksack) String() string {
	return s.First.String() + " | " + s.Second.String()
}

func (s *Rucksack) CombinedContents() Compartment {
	c := make(Compartment)
	for item, count := range s.First {
		c[item] = count
	}
	for item, count := range s.Second {
		c[item] += count
	}
	return c
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

	for index, sack := range sacks {
		itemPtr := FirstCommonItem(sack.First, sack.Second)
		if itemPtr == nil {
			err = fmt.Errorf("no common item for sack on line %d", index+1)
			return
		}

		totalPriority += itemPtr.Priority()
	}

	s.Part1.SaveAnswer(totalPriority)

	if len(sacks)%3 != 0 {
		err = fmt.Errorf("sack count of %d not divisible by 3", len(sacks))
		return
	}

	totalPriority = 0

	for i := 0; i < len(sacks); i += 3 {
		c1 := sacks[i].CombinedContents()
		c2 := sacks[i+1].CombinedContents()
		c3 := sacks[i+2].CombinedContents()

		itemPtr := FirstCommonItem(c1, c2, c3)
		if itemPtr == nil {
			err = fmt.Errorf("no common item for group starting at line %d",
				(i*3)+1)
			return
		}

		totalPriority += itemPtr.Priority()
	}

	s.Part2.SaveAnswer(totalPriority)

	return
}
