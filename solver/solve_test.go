package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/days"
)

func TestSolvers(t *testing.T) {
	require.NoError(t, os.Chdir(".."), "must get to repo root for test")

	for day := 1; day <= 25; day++ {
		if _, ok := days.Solvers[day]; !ok {
			continue
		}

		dayStr := fmt.Sprintf("day%02d", day)

		t.Run(dayStr, func(t *testing.T) {
			knownSolution, err := aoc.GetSolution(day)
			require.NoError(t, err, "require solution to check")
			require.True(t, knownSolution.Part1.Valid,
				"known solution part 1 answer must be valid to check")
			require.True(t, knownSolution.Part2.Valid,
				"known solution part 2 answer must be valid to check")

			solution, err := RunForDay(day)
			require.NoError(t, err, "expect no error when solving")

			assert.True(t, solution.Part1.Valid,
				"part 1 answer ought to be valid")
			if solution.Part1.Valid {
				assert.Equal(t,
					knownSolution.Part1.Value,
					solution.Part1.Value,
					"part 1 answer ought to match known answer",
				)
			}
			assert.True(t, solution.Part2.Valid,
				"part 2 answer ought to be valid")
			if solution.Part2.Valid {
				assert.Equal(t,
					knownSolution.Part2.Value,
					solution.Part2.Value,
					"part 2 answer ought to match known answer",
				)
			}
		})
	}
}
