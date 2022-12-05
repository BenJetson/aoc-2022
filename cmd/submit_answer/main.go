package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BenJetson/aoc-2022/client"
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

	client, err := client.New()
	if err != nil {
		log.Fatalf("failed to initailize client: %v\n", err)
	}

	solution, err := solver.RunForDay(*dayFlag)
	if err != nil {
		log.Fatalf("error while solving puzzle: %v\n", err)
	}

	answer := map[int]string{
		1: solution.Part1.Value,
		2: solution.Part2.Value,
	}[*partFlag]

	fmt.Printf("Submitting for day %d, part %d.\n", *dayFlag, *partFlag)
	fmt.Printf("Your answer is: %s.\n", answer)

	fmt.Println("---")
	fmt.Println("Result:")

	result, err := client.SubmitAnswer(*dayFlag, *partFlag, answer)
	if err != nil {
		log.Fatalf("error while submitting answer: %v\n", err)
	}

	fmt.Println(result)
}
