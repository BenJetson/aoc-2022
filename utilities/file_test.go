package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLinesFromFile(t *testing.T) {
	testCases := []struct {
		alias     string
		filename  string
		expect    []string
		expectErr bool
	}{
		{
			alias:     "nonexistent",
			filename:  "testdata/nope.txt",
			expectErr: true,
		},
		{
			alias:    "lorem",
			filename: "testdata/lorem.txt",
			expect: []string{
				"lorem ipsum",
				"dolor sit amet",
				"consectetur adipiscing elit",
				"sed do eiusmod",
				"tempor incididunt ut labore",
				"et dolore magna aliqua",
			},
		},
		{
			alias:    "double_linefeed",
			filename: "testdata/double_linefeed.txt",
			expect: []string{
				"this is text",
				"and another line",
				"",
				"and two lines further down",
				"we have more text",
			},
		},
		{
			alias:    "nums",
			filename: "testdata/nums.txt",
			expect: []string{
				"73", "54", "3", "41", "52", "63",
				"5", "8", "30", "50", "52", "33",
			},
		},
		{
			alias:    "nums_bad",
			filename: "testdata/nums_bad.txt",
			expect: []string{
				"14", "44", "9", "96", "6", "43", "24", "21", "7", "71", "31",
				"62", "70", "7", "99", "14", "14", "seventy four", "1", "54",
				"16", "42", "35", "18", "93", "23", "97", "36", "81", "70",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			lines, err := ReadLinesFromFile(tc.filename)

			if tc.expectErr {
				require.Error(t, err)
				assert.Nil(t, lines)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, tc.expect, lines)
			}
		})
	}
}

func TestReadIntegersFromFile(t *testing.T) {
	testCases := []struct {
		alias     string
		filename  string
		expect    []int
		expectErr bool
	}{
		{
			alias:     "nonexistent",
			filename:  "testdata/nope.txt",
			expectErr: true,
		},
		{
			alias:     "lorem",
			filename:  "testdata/lorem.txt",
			expectErr: true,
		},
		{
			alias:    "nums",
			filename: "testdata/nums.txt",
			expect: []int{
				73, 54, 3, 41, 52, 63,
				5, 8, 30, 50, 52, 33,
			},
		},
		{
			alias:     "nums_gap",
			filename:  "testdata/nums_gap.txt",
			expectErr: true,
		},
		{
			alias:     "nums_bad",
			filename:  "testdata/nums_bad.txt",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.alias, func(t *testing.T) {
			lines, err := ReadIntegersFromFile(tc.filename)

			if tc.expectErr {
				require.Error(t, err)
				assert.Nil(t, lines)
			} else {
				require.NoError(t, err)
				assert.EqualValues(t, tc.expect, lines)
			}
		})
	}
}
