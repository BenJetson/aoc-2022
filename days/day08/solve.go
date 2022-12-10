package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type Tree struct {
	Height            int
	ExternallyVisible bool
}

type ForestRow []Tree

func (fr ForestRow) Col(c int) *Tree {
	if fr == nil || c < 0 || c >= len(fr) {
		return nil
	}
	return &fr[c]
}

// Forest represents a grid of trees. Coordinates are like the PyGame grid
// system, where (0,0) is the upper left and y values increase as you go down.
type Forest []ForestRow

func (f Forest) Row(r int) ForestRow {
	if f == nil || r < 0 || r >= len(f) {
		return nil
	}
	return f[r]
}

func (f Forest) RowCount() int {
	if len(f) == 0 {
		return 0
	}
	return len(f[0])
}

func (f Forest) ColCount() int {
	return len(f)
}

func (f Forest) String() string {
	var s strings.Builder
	for r := 0; r < f.RowCount(); r++ {
		for c := 0; c < f.ColCount(); c++ {
			t := f.Row(r).Col(c)
			s.WriteString(strconv.Itoa(t.Height))
			if t.ExternallyVisible {
				s.WriteString("e")
			} else {
				s.WriteString("_")
			}
			s.WriteString(" ")
		}
		s.WriteString("\n")
	}
	return s.String()
}

func ReadForestFromLines(lines aoc.Input) (f Forest, err error) {
	f = make(Forest, len(lines))
	for r, line := range lines {
		f[r] = make(ForestRow, len(line))
	}

	for r, line := range lines {
		for c, char := range line {
			t := f.Row(r).Col(c)
			t.Height, err = strconv.Atoi(string(char))
			if err != nil {
				err = fmt.Errorf("bad height for tree at pos (%d, %d): %w",
					r, c, err)
				return
			}
		}
	}

	return
}

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

type Position struct {
	r, c int
}

func (p *Position) Move(v Vector) {
	p.r += v.r
	p.c += v.c
}

type Vector struct {
	r, c int
}

func (f Forest) Scan(start Position, v Vector) {
	p := start
	t := f.Row(p.r).Col(p.c)
	highestSoFar := t.Height
	p.Move(v)
	t = f.Row(p.r).Col(p.c)

	for t != nil {
		isTallerThanHighest := t.Height > highestSoFar
		if isTallerThanHighest {
			t.ExternallyVisible = true
			highestSoFar = t.Height
		}

		p.Move(v)
		t = f.Row(p.r).Col(p.c)
	}
}

func (f Forest) MarkExternallyVisibleTrees() {
	rowCount := f.RowCount()
	colCount := f.ColCount()

	for r := 0; r < rowCount; r++ {
		row := f.Row(r)

		// Leftmost and rightmost tree are always perimiter.
		row.Col(0).ExternallyVisible = true
		row.Col(colCount - 1).ExternallyVisible = true

		// On first and last rows, all trees from 1 to colCount-1 are perimiter.
		if r == 0 || r == rowCount-1 {
			for c := 1; c < colCount-1; c++ {
				row.Col(c).ExternallyVisible = true
			}
		}
	}

	for r := 1; r < rowCount-1; r++ {
		f.Scan(Position{r: r, c: 0}, Vector{r: 0, c: 1})
		f.Scan(Position{r: r, c: colCount - 1}, Vector{r: 0, c: -1})
	}

	for c := 1; c < colCount-1; c++ {
		f.Scan(Position{r: 0, c: c}, Vector{r: 1, c: 0})
		f.Scan(Position{r: rowCount - 1, c: c}, Vector{r: -1, c: 0})
	}

}

func (f Forest) CountExternallyVisibleTrees() (count int) {
	for r := 0; r < f.RowCount(); r++ {
		for c := 0; c < f.ColCount(); c++ {
			if f.Row(r).Col(c).ExternallyVisible {
				count++
			}
		}
	}
	return
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	f, err := ReadForestFromLines(input)
	if err != nil {
		err = fmt.Errorf("could not read forest: %w", err)
		return
	}

	f.MarkExternallyVisibleTrees()
	count := f.CountExternallyVisibleTrees()
	s.Part1.SaveIntAnswer(count)

	return
}
