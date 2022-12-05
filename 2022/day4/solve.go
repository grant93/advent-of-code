package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve(data []string) []int {
	answers := make([]int, 2)

	for _, line := range data {
		/* parse */
		q := regexp.MustCompile(`[-,]+`)
		var tmp []int
		for _, result := range q.Split(line, 4) {
			a, _ := strconv.Atoi(result)
			tmp = append(tmp, a)
		}
		a, b, c, d := tmp[0], tmp[1], tmp[2], tmp[3]

		/* part one */
		if (a <= c && b >= d) || (c <= a && d >= b) {
			answers[0]++
		}

		/* part two */
		if (a >= c && a <= d) || (b >= c && b <= d) || (c >= a && c <= b) || (d >= a && d <= b) {
			answers[1]++
		}
	}
	return answers
}

/* standard boilerplate */
func main() {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	fmt.Println(solve(data))
}
