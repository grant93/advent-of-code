package dayone

import (
	"sort"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func solve(data []string) (int, int) {
	var a []int
	total := 0

	for _, line := range data {
		if line == "" {
			a = append(a, total)
			total = 0
		} else {
			z, _ := strconv.Atoi(line)
			total += z
		}
	}
	sort.Ints(a)
	return a[len(a)-1], sum(a[len(a)-3 : len(a)])
}
