package daysix

import (
	"strconv"
	"strings"
)

func recurseNums(input []int64, backwards bool) int64 {
	diffs := []int64{}
	total := int64(0)
	pos := len(input)-1
	if backwards {
		pos = 0
		for i := 0; i < len(input)-1; i++ {
			diff := input[i] - input[i+1]
			diffs = append(diffs, diff)
			total += diff
		}
	} else {
		for i := 0; i < len(input)-1; i++ {
			diff := input[i+1] - input[i]
			diffs = append(diffs, diff)
			total += diff
		}
	}
	if last := input[pos]; last == 0 && total == 0 {
		return last 
	}
	z := recurseNums(diffs, backwards)
	return z + input[pos]
}

func solve(input []string) (int64, int64) {
	partOne := int64(0) 
	partTwo := int64(0)
	for _, line := range input {
		nums := []int64{}
		for _, n := range strings.Split(line, " ") {
			num, _ := strconv.ParseInt(n, 10, 64)
			nums = append(nums, num)
		}
		partOne += recurseNums(nums, false)
		partTwo += recurseNums(nums, true)

	}
	return partOne, partTwo
}
