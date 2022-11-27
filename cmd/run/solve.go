package main

import (
	"log"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/days"
	"github.com/BenJetson/aoc-2022/utilities"
)

func runSolverForDay(day int, inputFile string) (aoc.Solution, error) {
	solver, ok := days.Solvers[day]
	if !ok {
		log.Fatalf("no solver registered for day %d\n", day)
	}

	lines, err := utilities.ReadLinesFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not get lines from input file '%s': %v\n",
			inputFile, err)
	}

	return solver(lines)
}
