package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/* cheating and using a string as a stack */
var (
	stacks  = make([]string, 9)
	stacks2 = make([]string, 9)
)

func parseStacks(data []string) []string {
	for i, line := range data {
		counter := 0
		if len(line) == 0 {
			return data[i+1 : len(data)]
		}
		/* every 4th character is a potential value */
		for j := 1; j <= len(line); j += 4 {
			/* dirty hack to ignore the column numbers */
			if unicode.IsDigit(rune(line[j])) {
				break
			} else if line[j] != ' ' {
				/* only process actual values */
				stacks[counter] = string(line[j]) + stacks[counter]
			}
			counter++
		}
	}

	return data // should never happen

}

func move(src int, dst int, size int) {
	/* part one */
	for i := 0; i < size; i++ {
		length := len(stacks[src])
		if length > 0 {
			stacks[dst] += string(stacks[src][length-1])
			stacks[src] = stacks[src][0 : length-1]

		}
	}

	/* part two */
	length := len(stacks2[src])
	stacks2[dst] += stacks2[src][length-size : length]
	stacks2[src] = stacks2[src][0 : length-size]
}

func solve(data []string) []string {
	answers := make([]string, 2)
	data = parseStacks(data)
	copy(stacks2, stacks)

	for _, line := range data {
		tmp := strings.Fields(line)
		size, _ := strconv.Atoi(tmp[1])
		src, _ := strconv.Atoi(tmp[3])
		dst, _ := strconv.Atoi(tmp[5])
		move(src-1, dst-1, size)
	}

	for i := range stacks {
		if len(stacks[i]) > 0 {
			answers[0] += string(stacks[i][len(stacks[i])-1])
		}
		if len(stacks2[i]) > 0 {
			answers[1] += string(stacks2[i][len(stacks2[i])-1])
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
