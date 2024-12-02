package dayfifteen

import (
	"strings"
	"strconv"
)

type lense struct {
	label string
	fl int
}

func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current %= 256
	}
	return current
}

func solvePartOne(input []string) int {
	total := 0
	for _, s := range input {
		total += hash(s)
	}
	return total
}

func checkBoxes(s string, boxes [][]lense) {
	if strings.HasSuffix(s, "-") {
		l := s[0:len(s)-1]
		boxNum := hash(l)
		box := boxes[boxNum]
		for i, v := range box {
			if v.label == l {
				if len(box) == 1 {
					boxes[boxNum] = []lense{}
				} else if i == 0 {
					boxes[boxNum] = boxes[boxNum][1:]
				} else if i == len(boxes[boxNum]) - 1 {
					boxes[boxNum] = boxes[boxNum][0:len(box)-1]
				} else {
					boxes[boxNum] = append(boxes[boxNum][0:i], boxes[boxNum][i+1:len(box)]...)
				}
			}
		}
	} else {
		v := strings.Split(s, "=")
		boxNum := hash(v[0])
		l := v[0]
		n, _ := strconv.ParseInt(v[1], 10, 32)
		for i, v := range boxes[boxNum] {
			if v.label == l {
				boxes[boxNum][i].fl = int(n)
				return
			}
		}
		boxes[boxNum] = append(boxes[boxNum], lense{l, int(n)})
	}
}

func solvePartTwo(input []string) int {
	boxes := make([][]lense, 256)
	for i := range boxes {
		boxes[i] = []lense{}
	}
	for _, s := range input {
		checkBoxes(s, boxes)
	}
	total := 0
	for boxNum, box := range boxes {
		for i, v := range box {
			total += (boxNum + 1) * (i + 1) * v.fl
		}
	}
	return total
}

func solve(input []string) (int, int) {
	i := strings.Split(input[0], ",")
	partOne := solvePartOne(i)
	partTwo := solvePartTwo(i)
	return partOne, partTwo
}
