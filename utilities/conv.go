package utilities

import "strconv"

// SliceStringsToInts takes a slice of strings and attempts to convert to a
// slice of integers.
func SliceStringsToInts(numStrs []string) ([]int, error) {
	var nums []int

	var n int
	var err error

	for _, s := range numStrs {
		if n, err = strconv.Atoi(s); err != nil {
			return nil, err
		}

		nums = append(nums, n)
	}

	return nums, nil
}
