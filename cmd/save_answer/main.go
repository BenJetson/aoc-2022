package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/days"
	"github.com/BenJetson/aoc-2022/solver"
)

var dayFlag = flag.Int("day", 0, "day of the advent calendar, 1-25")
var partFlag = flag.Int("part", 0, "part of the puzzle, 1 or 2")

func main() {
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		log.Fatal("invalid or missing AoC day number")
	}
	if *partFlag != 1 && *partFlag != 2 {
		log.Fatal("invalid or missing AoC puzzle part number")
	}

	if _, ok := days.Solvers[*dayFlag]; !ok {
		log.Fatalln("this day is not initialized")
	}

	foundSolution, err := solver.RunForDay(*dayFlag)
	if err != nil {
		log.Fatalf("error while solving puzzle: %v\n", err)
	}

	knownSolution, err := aoc.GetSolution(*dayFlag)
	if errors.Is(err, os.ErrNotExist) {
		knownSolution = aoc.Solution{}
	} else if err != nil {
		log.Fatalf("error while reading known answer: %v\n", err)
	}

	switch *partFlag {
	case 1:
		knownSolution.Part1 = foundSolution.Part1
	case 2:
		knownSolution.Part2 = foundSolution.Part2
	}

	fmt.Printf("Saving solution for day %d, part %d.\n\n", *dayFlag, *partFlag)
	fmt.Println(knownSolution.String())

	err = os.WriteFile(
		aoc.GetSolutionFilename(*dayFlag),
		[]byte(knownSolution.String()),
		0644,
	)
	if err != nil {
		log.Fatalf("could not write input file: %v\n", err)
	}
}
