package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	testRange := func(t *testing.T, startExpect int, startChar, endChar rune) {
		expect := startExpect
		for c := startChar; c <= endChar; c++ {
			item := Item(c)
			priority := item.Priority()

			assert.Equal(t, expect, priority,
				fmt.Sprintf("priority of %s ought to be %d", string(c), expect))

			expect++
		}
	}

	t.Run("lowercase", func(t *testing.T) {
		testRange(t, 1, 'a', 'z')
	})
	t.Run("uppercase", func(t *testing.T) {
		testRange(t, 27, 'A', 'Z')
	})
}
