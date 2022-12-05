package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type Crate rune

type Stack []Crate

func (s *Stack) Push(c Crate) {
	*s = append(*s, c)
}

func (s *Stack) Pop() (c Crate, ok bool) {
	if len(*s) > 0 {
		c = (*s)[len(*s)-1]
		ok = true

		*s = (*s)[:len(*s)-1]
	}
	return
}

func (s *Stack) Peek() (c Crate, ok bool) {
	if len(*s) > 0 {
		c = (*s)[len(*s)-1]
		ok = true
	}
	return
}

type CargoShip []Stack

func (cs *CargoShip) String() string {
	var lines []string

	for index, stack := range *cs {
		var cargoStrings []string
		for _, cargo := range stack {
			cargoStrings = append(cargoStrings, string(cargo))
		}

		lines = append(lines, fmt.Sprintf("%d: %s",
			index, strings.Join(cargoStrings, ", ")))
	}

	return strings.Join(lines, "\n")
}

func (cs *CargoShip) HasStack(index int) bool {
	return index < len(*cs)
}

func (cs *CargoShip) TopOfEach() (out string, err error) {
	for index, stack := range *cs {
		topCrate, ok := stack.Peek()
		if !ok {
			err = fmt.Errorf("no top crate on stack at index %d", index)
			return
		}

		out += string(topCrate)
	}
	return
}

type CraneInstruction struct {
	Quantity               int
	SourceStack, DestStack int
}

func (cs *CargoShip) CheckBounds(instruction CraneInstruction) error {
	if !cs.HasStack(instruction.SourceStack) {
		return fmt.Errorf("invalid source stack %d",
			instruction.SourceStack)
	} else if !cs.HasStack(instruction.DestStack) {
		return fmt.Errorf("invalid destination stack %d",
			instruction.DestStack)
	}
	return nil
}

func (cs *CargoShip) UseCrane9000(instruction CraneInstruction) error {
	err := cs.CheckBounds(instruction)
	if err != nil {
		return fmt.Errorf("bounds error: %w", err)
	}

	for i := 0; i < instruction.Quantity; i++ {
		crate, ok := (*cs)[instruction.SourceStack].Pop()
		if !ok {
			return fmt.Errorf("stack %d does not have enough crates to move",
				instruction.SourceStack)
		}

		(*cs)[instruction.DestStack].Push(crate)
	}

	return nil
}

func (cs *CargoShip) UseCrane9001(instruction CraneInstruction) error {
	err := cs.CheckBounds(instruction)
	if err != nil {
		return fmt.Errorf("bounds error: %w", err)
	}

	var movingStack Stack
	for i := 0; i < instruction.Quantity; i++ {
		crate, ok := (*cs)[instruction.SourceStack].Pop()
		if !ok {
			return fmt.Errorf("stack %d does not have enough crates to move",
				instruction.SourceStack)
		}

		movingStack.Push(crate)
	}

	for len(movingStack) > 0 {
		crate, _ := movingStack.Pop()
		(*cs)[instruction.DestStack].Push(crate)
	}

	return nil
}

func ReadCrateCountFromLine(line string) (count int, err error) {
	line = strings.TrimSpace(line)
	lastSpaceIndex := strings.LastIndex(line, " ")
	countStr := line[lastSpaceIndex+1:]

	count, err = strconv.Atoi(countStr)
	if err != nil {
		err = fmt.Errorf("failed to parse create count string: %w", err)
		return
	}

	return
}

func ReadCratesFromLine(
	line string, crateCount int,
) (crates []Crate, err error) {

	if len(line) < 1+(crateCount-1)*4 {
		err = fmt.Errorf("not enough characters to read %d crates", crateCount)
		return
	}

	crates = make([]Crate, crateCount)
	for i := 0; i < crateCount; i++ {
		crateChar := line[i*4+1]

		if crateChar == ' ' {
			continue
		}

		crates[i] = Crate(crateChar)
	}

	return crates, nil
}

func ReadCargoShip(
	cargoLines []string, crateCount int,
) (cs CargoShip, err error) {

	cs = make(CargoShip, crateCount)
	for lineIndex := len(cargoLines) - 1; lineIndex >= 0; lineIndex-- {
		var crates []Crate
		crates, err = ReadCratesFromLine(cargoLines[lineIndex], crateCount)
		if err != nil {
			err = fmt.Errorf("failed to read crate line at index %d: %w",
				lineIndex, err)
			return
		}

		for crateIndex, crate := range crates {
			if crate != 0 {
				cs[crateIndex].Push(crate)
			}
		}
	}

	return cs, nil
}

func ReadInstructions(lines []string) ([]CraneInstruction, error) {
	var instructions []CraneInstruction

	for _, line := range lines {
		var i CraneInstruction
		n, err := fmt.Sscanf(line, "move %d from %d to %d",
			&i.Quantity, &i.SourceStack, &i.DestStack)
		if err != nil {
			return nil, fmt.Errorf("failed to scan instruction line: %w", err)
		} else if n != 3 {
			return nil, fmt.Errorf("expect three scanned values, found: %d", n)
		}

		// Decrement one since we use index.
		i.SourceStack--
		i.DestStack--

		instructions = append(instructions, i)
	}

	return instructions, nil
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	var cargoLines []string
	var crateCountLine string
	var instructionLines []string

	for _, line := range input {
		if line == "" {
			continue
		}

		if strings.Contains(line, "[") {
			cargoLines = append(cargoLines, line)
		} else if len(line) > 2 && line[:2] == " 1" {
			crateCountLine = line
		} else if strings.Contains(line, "move") {
			instructionLines = append(instructionLines, line)
		}
	}

	crateCount, err := ReadCrateCountFromLine(crateCountLine)
	if err != nil {
		err = fmt.Errorf("could not determine crate count: %w", err)
		return
	}

	cs, err := ReadCargoShip(cargoLines, crateCount)
	if err != nil {
		err = fmt.Errorf("could not parse cargo ship: %w", err)
		return
	}

	instructions, err := ReadInstructions(instructionLines)
	if err != nil {
		err = fmt.Errorf("could not parse instructions list: %w", err)
		return
	}

	for index, instruction := range instructions {
		err = cs.UseCrane9000(instruction)
		if err != nil {
			err = fmt.Errorf("crane 9000 instruction at index %d failed: %w",
				index, err)
			return
		}

	}

	topOfEach, err := cs.TopOfEach()
	if err != nil {
		err = fmt.Errorf("could not determine top of each stack: %w", err)
		return
	}

	s.Part1.SaveAnswer(topOfEach)

	cs, err = ReadCargoShip(cargoLines, crateCount)
	if err != nil {
		err = fmt.Errorf("could not parse cargo ship again: %w", err)
		return
	}

	for index, instruction := range instructions {
		err = cs.UseCrane9001(instruction)
		if err != nil {
			err = fmt.Errorf("crane 9001 instruction at index %d failed: %w",
				index, err)
			return
		}

	}

	topOfEach, err = cs.TopOfEach()
	if err != nil {
		err = fmt.Errorf("could not determine top of each stack: %w", err)
		return
	}

	s.Part2.SaveAnswer(topOfEach)

	return
}
