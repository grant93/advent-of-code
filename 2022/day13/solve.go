package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

/* treat the input as JSON, initially I wanted to use python to eval it - but
* this seemed like a way to do it in go without too much messing around. */

func testEquality(first any, second any) int {
	a, okA := first.(float64)
	b, okB := second.(float64)

	// both ints
	if okA && okB {
		return int(a - b)
	}

	// at least 1 list is involved
	var firstEntry []any
	switch first.(type) {
	case float64:
		firstEntry = []any{first}
	case []any, []float64:
		firstEntry = first.([]any)
	}

	var secondEntry []any
	switch second.(type) {
	case float64:
		secondEntry = []any{second}
	case []any, []float64:
		secondEntry = second.([]any)
	}

	// compare lists
	for i := range firstEntry {
		if len(secondEntry) <= i {
			return 1
		}
		if result := testEquality(firstEntry[i], secondEntry[i]); result != 0 {
			return result
		}
	}

	if len(firstEntry) == len(secondEntry) {
		return 0
	}
	return -1
}

func partOne(data []string, packets *[]any) int {
	count := 0
	for i := 0; i < len(data); i += 3 {
		var a, b any
		if err := json.Unmarshal([]byte(data[i]), &a); err == nil {
			if err := json.Unmarshal([]byte(data[i+1]), &b); err == nil {
				*packets = append(*packets, a, b)
				if testEquality(a, b) <= 0 {
					count += (i / 3) + 1
				}
			}
		}
	}
	return count
}

func partTwo(packets *[]any) int {
	var extra1, extra2 any
	json.Unmarshal([]byte("[[2]]"), &extra1)
	json.Unmarshal([]byte("[[6]]"), &extra2)
	*packets = append(*packets, extra1, extra2)

	allPackets := *packets

	sort.Slice(allPackets, func(x, y int) bool {
		return testEquality(allPackets[x], allPackets[y]) < 0
	})

	decoder_key := 1
	for key, value := range allPackets {
		str, _ := json.Marshal(value)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			decoder_key *= key + 1
		}
	}
	return decoder_key

}

func solve(data []string) []int {
	packets := []any{}
	answers := make([]int, 2)
	answers[0] = partOne(data, &packets)
	answers[1] = partTwo(&packets)

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
