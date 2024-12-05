package dayfive

import (
	"slices"
	"strconv"
	"strings"
)

type challengeData struct {
	rules map[int64][]int64
	data  [][]int64
}

// Parsing is the biggest challenge.
func parse(input []string) challengeData {
	a := make(map[int64][]int64)
	b := [][]int64{}
	rules := true
	for _, line := range input {
		if line == "" {
			rules = false
			continue
		}
		if rules {
			parts := strings.Split(line, "|")
			key, _ := strconv.ParseInt(parts[1], 10, 64)
			val, _ := strconv.ParseInt(parts[0], 10, 64)
			if n, ok := a[key]; ok {
				a[key] = append(n, val)
			} else {
				a[key] = []int64{val}
			}
		} else {
			parts := strings.Split(line, ",")
			c := []int64{}
			for _, x := range parts {
				val, _ := strconv.ParseInt(x, 10, 64)
				c = append(c, val)
			}
			b = append(b, c)
		}
	}
	return challengeData{rules: a, data: b}
}

// Sort the update by the page ordering rules.
func sortedByRules(line []int64, data *challengeData) []int64 {
	slices.SortFunc(line, func(a, b int64) int {
		for _, v := range data.rules[b] {
			if v == a {
				return -1
			}
		}
		return 1
	})
	return line
}

// Sort the ordering rules and check if they're different to validate.
func validate(data challengeData) (int64, int64) {
	middle := int64(0)
	incorrect := int64(0)
	for _, line := range data.data {
		if sorted := sortedByRules(slices.Clone(line), &data); slices.Compare(sorted, line) != 0 {
			incorrect += sorted[(len(sorted)-1)/2]
		} else {
			middle += line[(len(line)-1)/2]
		}
	}
	return middle, incorrect
}

func solve(input []string) (int64, int64) {
	data := parse(input)
	return validate(data)
}
