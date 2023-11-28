package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)


func partOne(data []string) int {
}

func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data)

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
