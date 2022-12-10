package day10

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type CycleNumber int

type RegisterValue int

type Register struct {
	Value   RegisterValue
	History map[CycleNumber]RegisterValue
}

func NewRegister(initial RegisterValue) *Register {
	return &Register{
		Value: initial,
		History: map[CycleNumber]RegisterValue{
			1: initial,
		},
	}
}

func (r *Register) Set(c CycleNumber, v RegisterValue) {
	r.Value = v
	r.History[c] = v
}

func (r *Register) ValueAtCycle(target CycleNumber) (v RegisterValue) {
	var cycles []CycleNumber
	for c := range r.History {
		cycles = append(cycles, c)
	}
	sort.Slice(cycles, func(i, j int) bool { return cycles[i] < cycles[j] })

	for _, c := range cycles {
		if c > target {
			return
		}
		v = r.History[c]
	}
	return
}

func (r Register) HistoryString(end CycleNumber) string {
	var s strings.Builder

	var c CycleNumber
	for c = 1; c <= end; c++ {
		v := r.ValueAtCycle(c)
		s.WriteString(fmt.Sprintf("cycle %d: %d\n", c, v))
	}

	return s.String()
}

type RegisterHistory map[CycleNumber]Register

type Processor struct {
	Display *Display
	Cycle   CycleNumber
	X       *Register
}

func NewProcessor() *Processor {
	return &Processor{
		Display: NewDisplay(40, 6),
		Cycle:   1,
		X:       NewRegister(1),
	}
}

func (p *Processor) ClockTick() {
	p.Display.Render(p.Cycle, p.X.Value)
	p.Cycle++
}

func (p *Processor) ExecNoOp() {
	p.ClockTick()
}

func (p *Processor) ExecAddX(value int) {
	p.ClockTick()
	p.ClockTick()
	p.X.Set(p.Cycle, p.X.Value+RegisterValue(value))
}

func (p *Processor) Execute(instr Instruction) {
	switch instr.Operation {
	case OpCodeNoOp:
		p.ExecNoOp()
	case OpCodeAddX:
		p.ExecAddX(instr.Operand1)
	}
}

type Pixel bool

type Display struct {
	Width, Height int
	Output        [][]Pixel
}

func NewDisplay(width, height int) *Display {
	d := &Display{Width: width, Height: height}

	d.Output = make([][]Pixel, height)
	for y := 0; y < d.Height; y++ {
		d.Output[y] = make([]Pixel, width)
	}

	return d
}

func (d *Display) String() string {
	var s strings.Builder
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			if d.Output[y][x] {
				s.WriteString("#")
			} else {
				s.WriteString(".")
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (d *Display) RowString(y, maxCol int) string {
	var s strings.Builder
	for x := 0; x <= maxCol; x++ {
		if d.Output[y][x] {
			s.WriteString("#")
		} else {
			s.WriteString(".")
		}
	}
	s.WriteString("\n")
	return s.String()
}

func (d *Display) Render(c CycleNumber, r RegisterValue) {
	cycleIndex := int(c) - 1
	y := (cycleIndex / d.Width) % d.Height
	x := cycleIndex % d.Width
	p := Pixel(x > int(r)-2 && x < int(r)+2)

	d.Output[y][x] = p
}

type OpCode string

const (
	OpCodeNoOp OpCode = "noop"
	OpCodeAddX OpCode = "addx"
)

type Instruction struct {
	Operation OpCode
	Operand1  int
}

func ReadInstructionsFromLines(input aoc.Input) ([]Instruction, error) {
	var instrs []Instruction
	for index, line := range input {
		parts := strings.Split(line, " ")
		numParts := len(parts)

		if numParts < 1 {
			return nil, fmt.Errorf("instruction on line %d has no parts",
				index+1)
		}

		var instr Instruction

		switch OpCode(parts[0]) {
		case OpCodeNoOp:
			instr.Operation = OpCodeNoOp
		case OpCodeAddX:
			instr.Operation = OpCodeAddX
			if numParts != 2 {
				return nil, fmt.Errorf("addx on line %d takes one operand",
					index+1)
			}

			var err error
			if instr.Operand1, err = strconv.Atoi(parts[1]); err != nil {
				return nil, fmt.Errorf("addx operand on line %d invalid: %w",
					index+1, err)
			}
		default:
			return nil, fmt.Errorf("invalid opcode '%s' on line %d",
				parts[0], index+1)
		}

		instrs = append(instrs, instr)
	}

	return instrs, nil
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	instrs, err := ReadInstructionsFromLines(input)
	if err != nil {
		err = fmt.Errorf("could not read input instructions: %w", err)
		return
	}

	p := NewProcessor()
	for _, instr := range instrs {
		p.Execute(instr)
	}

	var sum int
	for _, c := range []CycleNumber{20, 60, 100, 140, 180, 220} {
		if c > p.Cycle {
			err = fmt.Errorf("processor did not execute cycle %d", c)
			return
		}
		sum += int(c) * int(p.X.ValueAtCycle(c))
	}

	s.Part1.SaveIntAnswer(sum)

	out := p.Display.String()

	// Using a lookup table of known outputs since it doesn't make sense to
	// develop a letter recognizer for this. This allows automated testing to
	// work while still relying on human interpreatation of the outputs.
	result, ok := stringToLetters[out]
	if !ok {
		err = fmt.Errorf(
			"output of display has not been seen before; "+
				"need human interpretation:\n%s", out)
		return
	}

	s.Part2.SaveAnswer(result)

	return
}
