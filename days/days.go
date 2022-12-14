package days

import (
	"github.com/BenJetson/aoc-2022/aoc"
	// BEGIN DAY IMPORTS
	"github.com/BenJetson/aoc-2022/days/day01"
	"github.com/BenJetson/aoc-2022/days/day02"
	"github.com/BenJetson/aoc-2022/days/day03"
	"github.com/BenJetson/aoc-2022/days/day04"
	"github.com/BenJetson/aoc-2022/days/day05"
	"github.com/BenJetson/aoc-2022/days/day06"
	"github.com/BenJetson/aoc-2022/days/day07"
	"github.com/BenJetson/aoc-2022/days/day08"
	"github.com/BenJetson/aoc-2022/days/day09"
	"github.com/BenJetson/aoc-2022/days/day10"
	"github.com/BenJetson/aoc-2022/days/day11"
	// END DAY IMPORTS
)

var Solvers = map[int]aoc.Solver{
	// BEGIN DAY SOLVERS
	1: day01.SolvePuzzle,
	2: day02.SolvePuzzle,
	3: day03.SolvePuzzle,
	4: day04.SolvePuzzle,
	5: day05.SolvePuzzle,
	6: day06.SolvePuzzle,
	7: day07.SolvePuzzle,
	8: day08.SolvePuzzle,
	9: day09.SolvePuzzle,
	10: day10.SolvePuzzle,
	11: day11.SolvePuzzle,
	// END DAY SOLVERS
}
