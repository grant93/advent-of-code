package daytwo

import (
	"strconv"
	"strings"

	slice "github.com/grant93/advent-of-code/lib"
)

func isSafe(report []int64) bool {
	dir := int64(0)
	for i := range len(report) - 1 {
		diff := report[i+1] - report[i]
		if dir == 0 {
			if diff == 0 {
				return false
			}
			dir = diff
		} else if (dir < 0 && diff >= 0) || (dir > 0 && diff <= 0) {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func solve(input []string) (int, int) {
	nums := [][]int64{}
	for _, s := range input {
		parts := strings.Fields(s)
		tmp := []int64{}
		for _, i := range parts {
			a, _ := strconv.ParseInt(i, 10, 64)
			tmp = append(tmp, a)
		}
		nums = append(nums, tmp)
	}

	partOne := 0
	partTwo := 0
	for _, report := range nums {
		if isSafe(report) {
			partOne += 1
		}
		for i := range len(report) {
			if isSafe(slice.RemoveIndex(report, i)) {
				partTwo += 1
				break
			}
		}
	}

	return partOne, partTwo
}
