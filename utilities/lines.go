package utilities

import (
	"strings"
)

// MergeGapsBetweenLines takes a slice of lines (bricks) and shall merge
// lines together with the mortar, breaking upon a blank line.
func MergeGapsBetweenLines(bricks []string, mortar string) []string {
	if len(bricks) < 1 {
		return bricks
	}

	var pile []string
	var buffer strings.Builder

	for i := range bricks {
		if bricks[i] == "" {
			pile = append(pile, buffer.String())
			buffer.Reset()
			continue
		}

		if buffer.Len() > 0 {
			buffer.WriteString(mortar)
		}
		buffer.WriteString(bricks[i])
	}

	if buffer.Len() > 0 {
		pile = append(pile, buffer.String())
	}

	return pile
}
