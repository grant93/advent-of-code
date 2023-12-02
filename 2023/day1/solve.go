package dayone

import (
	"strconv"
	"strings"
	"unicode"
)

func partOne(input []string) int {
	total := 0
	for _, line := range input {
		if line == "" {
			break
		}
		var first, last string
		for _, a := range line {
			if unicode.IsNumber(a) {
				first = string(a)
				break
			}
		}
		for i := range line {
			i = len(line) - 1 - i
			if unicode.IsNumber(rune(line[i])) {
				last = string(line[i])
				break
			}
		}
		num, err := strconv.ParseInt(first+last, 10, 64)
		if err != nil {
			panic("failed to convert string")
		}
		total += int(num)
	}
	return total
}

func partTwo(input []string) int {

	valueMap := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
	var output []string
	for _, line := range input {
		for k, v := range valueMap {
			line = strings.ReplaceAll(line, k, v)
		}
		output = append(output, line)
	}
	return partOne(output)
}
