package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/BenJetson/aoc-2022/client"
	"github.com/BenJetson/aoc-2022/days"
	"github.com/BenJetson/aoc-2022/utilities"
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

	puzzle, err := client.GetPuzzleMarkdown(*dayFlag, *partFlag)
	if err != nil {
		log.Fatalf("failed to get puzzle markdown: %v\n", err)
	}

	readmeFilename := fmt.Sprintf("days/day%02d/README.md", *dayFlag)
	readmeLines, err := utilities.ReadLinesFromFile(readmeFilename)
	if err != nil {
		log.Fatalf("failed to get read README file from disk: %v\n", err)
	}

	if *partFlag == 1 {
		readmeLines[0] = strings.Replace(
			readmeLines[0],
			`<!-- PUZZLE TITLE PLACEHOLDER -->`,
			puzzle.Title,
			1,
		)
	}

	targetPlaceholder := map[int]string{
		1: `<!-- PART ONE PLACEHOLDER -->`,
		2: `<!-- PART TWO PLACEHOLDER -->`,
	}[*partFlag]

	found := false
	for index, lineText := range readmeLines {
		if lineText != targetPlaceholder {
			continue
		}

		found = true

		newReadmeLines := append(readmeLines[:index], puzzle.Body...)
		newReadmeLines = append(newReadmeLines, readmeLines[index+1:]...)

		readmeLines = newReadmeLines
		break
	}

	if !found {
		log.Fatalln("could not find target placeholder for this part")
	}

	readmeTxt := strings.Join(readmeLines, "\n") + "\n"
	err = os.WriteFile(readmeFilename, []byte(readmeTxt), 0644)
	if err != nil {
		log.Fatalf("could not write README file: %v\n", err)
	}

	err = exec.
		Command("npx", "prettier", "--write", readmeFilename).
		Run()
	if err != nil {
		log.Fatalf("failed to format README file: %v\n", err)
	}
}
