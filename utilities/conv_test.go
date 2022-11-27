package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceStringsToInts(t *testing.T) {
	testCases := []struct {
		alias     string
		input     []string
		expect    []int
		expectErr bool
	}{
		{
			alias:  "nil",
			input:  nil,
			expect: nil,
		},
		{
			alias:  "empty",
			input:  []string{},
			expect: nil,
		},
		{
			alias: "numbers",
			input: []string{
				"14", "44", "9", "96", "6", "43", "24", "21", "7", "71", "31",
				"62", "70", "7", "99", "14", "14", "1", "54", "16", "42", "35",
				"18", "93", "23", "97", "36", "81", "70",
			},
			expect: []int{
				14, 44, 9, 96, 6, 43, 24, 21, 7, 71, 31,
				62, 70, 7, 99, 14, 14, 1, 54, 16, 42, 35,
				18, 93, 23, 97, 36, 81, 70,
			},
		},
		{
			alias: "not_really_number",
			input: []string{
				"14", "44", "9", "96", "6", "43", "24", "21", "7", "71", "31",
				"62", "70", "7", "99", "14", "14", "seventy four", "1", "54",
				"16", "42", "35", "18", "93", "23", "97", "36", "81", "70",
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			nums, err := SliceStringsToInts(tc.input)

			if tc.expectErr {
				require.Error(t, err)
				assert.Nil(t, nums)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, tc.expect, nums)
			}
		})
	}
}
