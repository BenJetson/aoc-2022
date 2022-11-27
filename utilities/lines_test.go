package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeGapsBetweenLines(t *testing.T) {
	testCases := []struct {
		alias  string
		bricks []string
		mortar string
		expect []string
	}{
		{
			alias:  "no_gaps",
			bricks: []string{"one two", "three four", "five six", "seven"},
			mortar: "-x-",
			expect: []string{"one two-x-three four-x-five six-x-seven"},
		},
		{
			alias:  "nil",
			bricks: nil,
			mortar: "asdf",
			expect: nil,
		},
		{
			alias:  "empty",
			bricks: []string{},
			mortar: "jkl;",
			expect: []string{},
		},
		{
			alias:  "blank",
			bricks: []string{""},
			mortar: "!!",
			expect: []string{""},
		},
		{
			alias: "a_few",
			bricks: []string{
				"one two",
				"three four",
				"",
				"purple blue",
				"green yellow",
				"orange red",
			},
			mortar: " ",
			expect: []string{
				"one two three four",
				"purple blue green yellow orange red",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			actual := MergeGapsBetweenLines(tc.bricks, tc.mortar)
			assert.EqualValues(t, tc.expect, actual)
		})
	}
}
