package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type ScenicScoreSheet struct {
	North, East, South, West int
}

func (s *ScenicScoreSheet) Total() int {
	return s.North * s.East * s.South * s.West
}

type Tree struct {
	Height            int
	ExternallyVisible bool
	ScenicScore       ScenicScoreSheet
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

func (f Forest) ExternallyVisibleString() string {
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

// ScenicScoreCSVString creates a comma separated value string with groups that
// look like a cross:
//
//	  | N | T
//	W | H | E
//	  | S |
//
// Where N/E/S/W are the cardinal direction scenic score values, H is the
// height of the tree, and T is the total scenic score.
func (f Forest) ScenicScoreCSVString() string {
	var s strings.Builder
	for r := 0; r < f.RowCount(); r++ {
		var line1, line2, line3 strings.Builder
		for c := 0; c < f.ColCount(); c++ {
			t := f.Row(r).Col(c)

			line1.WriteString(fmt.Sprintf(",,%d,%d,", t.ScenicScore.North,
				t.ScenicScore.Total()))
			line2.WriteString(fmt.Sprintf(",%d,%d,%d,",
				t.ScenicScore.West, t.Height, t.ScenicScore.East))
			line3.WriteString(fmt.Sprintf(",,%d,,", t.ScenicScore.South))
		}
		s.WriteString(line1.String() + "\n")
		s.WriteString(line2.String() + "\n")
		s.WriteString(line3.String() + "\n")
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

type VectorFunc func(p *Position)

func (v VectorFunc) Apply(p *Position) { v(p) }

func VectorNorth(p *Position) { p.r-- }
func VectorEast(p *Position)  { p.c++ }
func VectorSouth(p *Position) { p.r++ }
func VectorWest(p *Position)  { p.c-- }

func (f Forest) ScanExternalVisibility(start Position, v VectorFunc) {
	p := start
	t := f.Row(p.r).Col(p.c)
	highestSoFar := t.Height
	v.Apply(&p)
	t = f.Row(p.r).Col(p.c)

	for t != nil {
		isTallerThanHighest := t.Height > highestSoFar
		if isTallerThanHighest {
			t.ExternallyVisible = true
			highestSoFar = t.Height
		}

		v.Apply(&p)
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
		f.ScanExternalVisibility(Position{r: r, c: 0}, VectorEast)
		f.ScanExternalVisibility(Position{r: r, c: colCount - 1}, VectorWest)
	}

	for c := 1; c < colCount-1; c++ {
		f.ScanExternalVisibility(Position{r: 0, c: c}, VectorSouth)
		f.ScanExternalVisibility(Position{r: rowCount - 1, c: c}, VectorNorth)
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

func (f Forest) ScanScenicScore(start Position, v VectorFunc) (score int) {
	p := start
	t := f.Row(p.r).Col(p.c)

	myHeight := t.Height

	v.Apply(&p)
	t = f.Row(p.r).Col(p.c)

	for t != nil {
		score++
		if t.Height >= myHeight {
			break
		}

		v.Apply(&p)
		t = f.Row(p.r).Col(p.c)
	}

	return
}

func (f Forest) CalculateScenicScores() {
	for r := 0; r < f.RowCount(); r++ {
		for c := 0; c < f.ColCount(); c++ {
			t := f.Row(r).Col(c)
			p := Position{r: r, c: c}

			t.ScenicScore.North = f.ScanScenicScore(p, VectorNorth)
			t.ScenicScore.East = f.ScanScenicScore(p, VectorEast)
			t.ScenicScore.South = f.ScanScenicScore(p, VectorSouth)
			t.ScenicScore.West = f.ScanScenicScore(p, VectorWest)
		}
	}
}

func (f Forest) HighestScenicScore() (highestScore int) {
	for r := 0; r < f.RowCount(); r++ {
		for c := 0; c < f.ColCount(); c++ {
			score := f.Row(r).Col(c).ScenicScore.Total()
			if score > highestScore {
				highestScore = score
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

	f.CalculateScenicScores()
	highest := f.HighestScenicScore()
	s.Part2.SaveIntAnswer(highest)

	return
}
