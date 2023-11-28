package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)




func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data[0])
	answers[1] = partTwo(data[0])
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
