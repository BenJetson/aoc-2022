package day01

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/BenJetson/aoc-2022/aoc"
)

type Food struct {
	Calories int
}

type Elf struct {
	Inventory []Food
}

func (e *Elf) TotalCaloriesCarried() (total int) {
	for _, f := range e.Inventory {
		total += f.Calories
	}
	return
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var elves []Elf

	var current Elf
	for index, line := range input {
		if line == "" {
			elves = append(elves, current)
			current = Elf{}
			continue
		}

		var calories int

		if calories, err = strconv.Atoi(line); err != nil {
			err = fmt.Errorf("cannot convert calories on line %d: %w",
				index+1, err)
			return
		}

		current.Inventory = append(current.Inventory, Food{Calories: calories})
	}

	if len(elves) < 1 {
		err = errors.New("no elves found")
		return
	}

	// Sort so the elf with the most calories will be at index zero.
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCaloriesCarried() > elves[j].TotalCaloriesCarried()
	})

	s.Part1.SaveAnswer(elves[0].TotalCaloriesCarried())

	return
}
