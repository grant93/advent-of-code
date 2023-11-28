package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	index int
	value int
}

func toInt(data []string) []Node {
	converted := []Node{}
	for i, item := range data {
		tmp, _ := strconv.Atoi(item)
		converted = append(converted, Node{i, tmp})
	}
	return converted
}

func abs(num int) int {
	if num < 0 {
		return num*-1
	}
	return num
}

func findPos(haystack []Node, needle int) int {
	for index, value := range haystack {
		if value.index == needle {
			return index
		}
	}
	return -1
}

func findValPos(haystack []Node, needle int) int {
	for index, value := range haystack {
		if value.value == needle {
			return index
		}
	}
	return -1
}

func partOne(data []Node) int {
	workingCopy := data

	for i, item := range data {
		pos := findPos(workingCopy, i)
		newPos := pos + item.value
		if newPos < 0 {
			newPos = len(data) -1 + (newPos % (len(data)-1))
		} else {
			newPos %= len(data)
		}

		if newPos == len(data)-1 {
			newPos--
		} else if newPos == 0 {
			newPos = len(data)-1
		}
		if item.value != 0 {
			tmp := []Node{}
			for j, obj := range workingCopy {
				if j == pos {
					continue
				}
				tmp = append(tmp, obj)
				if j == newPos {
					tmp = append(tmp, item)
				}
			}
			workingCopy = tmp
		}
	}
	index := findValPos(workingCopy, 0)
	fmt.Println(workingCopy)
	fmt.Println(index, workingCopy[(index + 1000) % len(workingCopy)].value, workingCopy[(index + 2000) % len(workingCopy)].value, workingCopy[(index + 3000) % len(workingCopy)].value)
	return workingCopy[(index + 1000) % len(workingCopy)].value + workingCopy[(index + 2000) % len(workingCopy)].value + workingCopy[(index + 3000) % len(workingCopy)].value 
}

/*
func partOne(data []Node) int {
	workingCopy := data
	for _, item := range data {
		pos := findPos(workingCopy, item)
		newPos := pos + *item 
		fmt.Println(newPos)
		if newPos < 0 {
			newPos = len(data) - 1 + (newPos % (len(data)-1))
		} else {
			newPos %= len(data)

		}
		fmt.Println(newPos)
		if newPos == len(data)-1 {
			newPos--
		}
		if *item != 0 {
			tmp := []*int{}
			for i, obj := range workingCopy {
				if i == pos {
					continue
				} 
				tmp = append(tmp, obj)
				if i==newPos {
					tmp = append(tmp, item)
				}
			}
			workingCopy = tmp
			fmt.Println(workingCopy)
		}
	}

	index := findValPos(workingCopy, 0)
	fmt.Println(index, *(workingCopy[(index + 1000) % len(workingCopy)]), *(workingCopy[(index + 2000) % len(workingCopy)]), *(workingCopy[(index + 3000) % len(workingCopy)]))
	return *(workingCopy[(index + 1000) % len(workingCopy)]) + *(workingCopy[(index + 2000) % len(workingCopy)]) + *(workingCopy[(index + 3000) % len(workingCopy)])

}
*/


func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(toInt(data))

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
