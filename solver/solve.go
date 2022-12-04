package solver

import (
	"fmt"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/days"
	"github.com/BenJetson/aoc-2022/utilities"
)

func RunForDayWithInput(day int, inputFile string) (sol aoc.Solution, err error) {
	solver, ok := days.Solvers[day]
	if !ok {
		err = fmt.Errorf("no solver registered for day %d\n", day)
		return
	}

	lines, err := utilities.ReadLinesFromFile(inputFile)
	if err != nil {
		err = fmt.Errorf("could not get lines from input file '%s': %w\n",
			inputFile, err)
		return
	}

	return solver(lines)
}

func RunForDay(day int) (aoc.Solution, error) {
	inputFile := aoc.GetInputFilename(day)
	return RunForDayWithInput(day, inputFile)
}
