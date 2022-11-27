package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BenJetson/aoc-2022/days"
	"github.com/BenJetson/aoc-2022/utilities"
)

var dayFlag = flag.Int("day", 0, "day of the advent calendar, 1-25")

func main() {
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		log.Fatal("invalid or missing AoC day number")
	}

	if _, ok := days.Solvers[*dayFlag]; ok {
		log.Fatalln("this day has already been initialized")
	}

	dayName := fmt.Sprintf("day%02d", *dayFlag)

	err := os.Mkdir("days/"+dayName, 0755)
	if err != nil {
		log.Fatalf("could not make directory: %v", err)
	}

	solveSource := strings.Join([]string{
		"package " + dayName,
		"",
		`import "github.com/BenJetson/aoc-2022/aoc"`,
		"",
		`func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {`,
		"\treturn",
		"}",
	}, "\n") + "\n"

	err = os.WriteFile("days/"+dayName+"/solve.go", []byte(solveSource), 0644)
	if err != nil {
		log.Fatalf("could not write source file: %v", err)
	}

	daysSourceLines, err := utilities.ReadLinesFromFile("days/days.go")
	if err != nil {
		log.Fatalf("could not read days source file: %v", err)
	}

	for index, data := range daysSourceLines {
		if data == "\t// END DAY IMPORTS" {
			daysSourceLines = append(
				daysSourceLines[:index+1],
				daysSourceLines[index:]...,
			)
			daysSourceLines[index] = "\t" + fmt.Sprintf(
				`"github.com/BenJetson/aoc-2022/days/%s"`, dayName)
			break
		}
	}

	for index, data := range daysSourceLines {
		if data == "\t// END DAY SOLVERS" {
			daysSourceLines = append(
				daysSourceLines[:index+1],
				daysSourceLines[index:]...,
			)
			daysSourceLines[index] = "\t" + fmt.Sprintf(
				`%d: %s.SolvePuzzle,`, *dayFlag, dayName)
			break
		}
	}

	daysSource := strings.Join(daysSourceLines, "\n") + "\n"
	err = os.WriteFile("days/days.go", []byte(daysSource), 0644)
	if err != nil {
		log.Fatalf("could not write days source file: %v", err)
	}
}
