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
	Cycle CycleNumber
	X     *Register
}

func NewProcessor() *Processor {
	return &Processor{
		Cycle: 1,
		X:     NewRegister(1),
	}
}

func (p *Processor) ExecNoOp() {
	p.Cycle++
}

func (p *Processor) ExecAddX(value int) {
	p.Cycle++
	p.Cycle++
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

	return
}
