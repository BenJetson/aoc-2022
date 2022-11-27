package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BenJetson/aoc-2022/aoc"
)

var dayFlag = flag.Int("day", 0, "day of the advent calendar, 1-25")
var inputFileFlag = flag.String("input", "",
	"optional filename to use for input, other than the default")

func main() {
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		log.Fatal("invalid or missing AoC day number")
	}

	if len(*inputFileFlag) < 1 {
		*inputFileFlag = aoc.GetInputFilename(*dayFlag)
	}

	solution, err := runSolverForDay(*dayFlag, *inputFileFlag)
	if err != nil {
		log.Fatalf("error while solving puzzle: %v\n", err)
	}

	fmt.Print(solution.String())
}
