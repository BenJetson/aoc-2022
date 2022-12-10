package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type PositionHistory map[Position]struct{}

func (ph PositionHistory) Record(p Position) {
	ph[p] = struct{}{}
}

func (ph PositionHistory) Count() int {
	return len(ph)
}

type Position struct{ X, Y int }

func (p *Position) Equals(o Position) bool {
	return p.X == o.X && p.Y == o.Y
}

func Abs(value int) int {
	if value < 0 {
		return -1 * value
	}
	return value
}

func (p *Position) DistanceX(o Position) int {
	return Abs(p.X - o.X)
}

func (p *Position) DistanceY(o Position) int {
	return Abs(p.Y - o.Y)
}

func (p *Position) IsTouching(o Position) bool {

	return p.Equals(o) ||
		(p.DistanceY(o) < 2 && p.DistanceX(o) < 2)
}

type VectorFunc func(p *Position)

func (v VectorFunc) Apply(p *Position) { v(p) }

func VectorNorth(p *Position) { p.Y++ }
func VectorEast(p *Position)  { p.X++ }
func VectorSouth(p *Position) { p.Y-- }
func VectorWest(p *Position)  { p.X-- }

func VectorNorthEast(p *Position) { p.Y++; p.X++ }
func VectorSouthEast(p *Position) { p.Y--; p.X++ }
func VectorSouthWest(p *Position) { p.Y--; p.X-- }
func VectorNorthWest(p *Position) { p.Y++; p.X-- }

type RopeSimulator struct {
	Head Position
	Tail Position

	TailHistory PositionHistory
}

func NewRopeSimulator() *RopeSimulator {
	rs := &RopeSimulator{
		TailHistory: make(PositionHistory),
	}

	// Record initial tail position.
	rs.TailHistory.Record(rs.Tail)

	return rs
}

func Max(values ...int) (out int) {
	if len(values) < 1 {
		return 0
	}

	out = values[0]
	for _, v := range values {
		if v > out {
			out = v
		}
	}
	return out
}

func (rs *RopeSimulator) String() string {
	gridSize := Max(5, rs.Head.X, rs.Head.Y, rs.Tail.X, rs.Tail.Y) + 1

	grid := make([][]string, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]string, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = "."
		}
	}

	putAtPos := func(char string, p Position) {
		grid[gridSize-1-p.Y][p.X] = char
	}

	putAtPos("s", Position{X: 0, Y: 0})
	putAtPos("T", rs.Tail)
	putAtPos("H", rs.Head)

	var s strings.Builder

	// s.WriteString(" ")
	// for j := 0; j < gridSize; j++ {
	// 	s.WriteString(strconv.Itoa(j % 10))
	// }
	// s.WriteString("\n")

	// for i, line := range grid {
	for _, line := range grid {
		// y := gridSize - i - 1
		// s.WriteString(strconv.Itoa(y % 10))
		s.WriteString(strings.Join(line, "") + "\n")
	}

	return s.String()
}

type Instruction struct {
	Char      string
	Direction VectorFunc
	Steps     int
}

func (rs *RopeSimulator) TailStep() {
	if rs.Tail.Equals(rs.Head) {
		return
	}

	var v VectorFunc

	switch {
	case rs.Tail.X == rs.Head.X && rs.Tail.Y < rs.Head.Y:
		v = VectorNorth
	case rs.Tail.X == rs.Head.X && rs.Tail.Y > rs.Head.Y:
		v = VectorSouth
	case rs.Tail.X < rs.Head.X && rs.Tail.Y == rs.Head.Y:
		v = VectorEast
	case rs.Tail.X > rs.Head.X && rs.Tail.Y == rs.Head.Y:
		v = VectorWest
	case rs.Tail.X < rs.Head.X && rs.Tail.Y < rs.Head.Y:
		v = VectorNorthEast
	case rs.Tail.X < rs.Head.X && rs.Tail.Y > rs.Head.Y:
		v = VectorSouthEast
	case rs.Tail.X > rs.Head.X && rs.Tail.Y > rs.Head.Y:
		v = VectorSouthWest
	case rs.Tail.X > rs.Head.X && rs.Tail.Y < rs.Head.Y:
		v = VectorNorthWest
	}

	v.Apply(&rs.Tail)
	rs.TailHistory.Record(rs.Tail)
}

func (rs *RopeSimulator) Execute(instr Instruction) {
	// fmt.Printf("== %s %d ==\n\n", instr.Char, instr.Steps)
	for i := 0; i < instr.Steps; i++ {
		instr.Direction.Apply(&rs.Head)

		if !rs.Tail.IsTouching(rs.Head) {
			rs.TailStep()
		}

		// fmt.Println(rs.String())
	}
}

func ReadInstructionsFromLines(input aoc.Input) ([]Instruction, error) {
	var instrs []Instruction
	for index, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expect 2 parts on line %d but found %d ",
				index+1, len(parts))
		}

		var instr Instruction

		instr.Char = parts[0]

		switch instr.Char {
		case "U":
			instr.Direction = VectorNorth
		case "R":
			instr.Direction = VectorEast
		case "D":
			instr.Direction = VectorSouth
		case "L":
			instr.Direction = VectorWest
		default:
			return nil, fmt.Errorf("invalid direction '%s' on line %d",
				parts[0], index+1)
		}

		var err error
		if instr.Steps, err = strconv.Atoi(parts[1]); err != nil {
			return nil, fmt.Errorf("invalid step count '%s' on line %d: %w",
				parts[1], index+1, err)
		}

		instrs = append(instrs, instr)
	}

	return instrs, nil
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	instrs, err := ReadInstructionsFromLines(input)
	if err != nil {
		err = fmt.Errorf("could not read instructions input: %w", err)
		return
	}

	rs := NewRopeSimulator()
	for _, instr := range instrs {
		rs.Execute(instr)
	}

	s.Part1.SaveIntAnswer(rs.TailHistory.Count())

	return
}
