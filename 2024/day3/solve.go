package daythree

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulRE = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var mulCondRE = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)

func parseMuls(input []string) int {
	total := 0
	for _, line := range input {
		matches := mulRE.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)
		for _, i := range matches {
			first, _ := strconv.ParseInt(i[1], 10, 64)
			second, _ := strconv.ParseInt(i[2], 10, 64)
			total += int(first * second)
		}
	}
	return total
}

func parseConditionalMuls(input []string) int {
	total := 0
	state := true
	for _, line := range input {
		matches := mulCondRE.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)
		for _, i := range matches {
			switch {
			case strings.HasPrefix(i[0], "don't"):
				state = false
			case strings.HasPrefix(i[0], "do"):
				state = true
			default:
				if !state {
					continue
				}
				first, _ := strconv.ParseInt(i[1], 10, 64)
				second, _ := strconv.ParseInt(i[2], 10, 64)
				total += int(first * second)
			}
		}
	}
	return total
}

func solve(input []string) (int, int) {

	return parseMuls(input), parseConditionalMuls(input)
}
