package main

import (
	"bufio"
	"fmt"
	"os"
)

func findRepeats(data string) bool {
	for i, char := range data[0 : len(data)-1] {
		for _, char2 := range data[i+1 : len(data)] {
			if char == char2 {
				return true
			}
		}
	}
	return false
}

func solve(data []string) []int {
	answers := make([]int, 2)
	stream := data[0]

	for i := 0; i < len(stream); i++ {
		if !findRepeats(stream[i : i+4]) {
			if answers[0] == 0 {
				answers[0] = i + 4
			}
		}
		if !findRepeats(stream[i : i+14]) {
			answers[1] = i + 14
		}
		if answers[0] != 0 && answers[1] != 0 {
			break
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
