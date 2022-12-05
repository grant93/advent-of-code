package main

import (
	"bufio"
	"fmt"
	"os"
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

func solve(data []string) []int {
	var a []int
	answers := make([]int, 2)
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
	answers[0] = a[len(a)-1]
	answers[1] = sum(a[len(a)-3 : len(a)])
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
