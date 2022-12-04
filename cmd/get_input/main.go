package main

import (
	"flag"
	"log"
	"os"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/client"
	"github.com/BenJetson/aoc-2022/days"
)

var dayFlag = flag.Int("day", 0, "day of the advent calendar, 1-25")

func main() {
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		log.Fatal("invalid or missing AoC day number")
	}

	if _, ok := days.Solvers[*dayFlag]; !ok {
		log.Fatalln("this day is not initialized")
	}

	client, err := client.New()
	if err != nil {
		log.Fatalf("failed to initailize client: %v\n", err)
	}

	input, err := client.GetPuzzleInput(*dayFlag)
	if err != nil {
		log.Fatalf("failed to get puzzle input: %v\n", err)
	}

	inputFilename := aoc.GetInputFilename(*dayFlag)

	err = os.WriteFile(inputFilename, []byte(input), 0644)
	if err != nil {
		log.Fatalf("could not write input file: %v\n", err)
	}
}
