package dayone

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func solve(input []string) (int64, int64) {
	var left, right []int64
	for _, s := range input {
		parts := strings.Split(s, "   ")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		left = append(left, a)
		a, _ = strconv.ParseInt(parts[1], 10, 64)
		right = append(right, a)
	}
	slices.Sort(left)
	slices.Sort(right)

	total := int64(0)
	occurrences := map[int64]int64{}
	for i := range left {
		diff := abs(left[i] - right[i])
		total += diff
		if v, ok := occurrences[right[i]]; ok {
			occurrences[right[i]] = v + 1
		} else {
			occurrences[right[i]] = 1
		}
	}

	totalOccurrences := int64(0)
	for _, v := range left {
		if val, ok := occurrences[v]; ok {
			totalOccurrences += v * val
		}
	}

	return total, totalOccurrences
}
