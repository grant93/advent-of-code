package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x     int
	y     int
	moved bool
}

var dirMap = map[string][2]int{
	"R": [2]int{1, 0},
	"L": [2]int{-1, 0},
	"U": [2]int{0, 1},
	"D": [2]int{0, -1}}

func moveTail(head Part, tail Part) Part {
	xDiff := head.x - tail.x
	yDiff := head.y - tail.y

	if (math.Abs(float64(xDiff)) <= 1) && (math.Abs(float64(yDiff)) <= 1) {
		return tail
	} else if xDiff != 0 && yDiff != 0 {
		tail.x += xDiff / int(math.Abs(float64(xDiff)))
		tail.y += yDiff / int(math.Abs(float64(yDiff)))
	} else {
		if xDiff != 0 {
			tail.x += xDiff / int(math.Abs(float64(xDiff)))
		} else if yDiff != 0 {
			tail.y += yDiff / int(math.Abs(float64(yDiff)))
		}
	}

	return tail
}

func partTwo(data []string) int {
	visited := make(map[[2]int]bool)
	knots := [10]Part{}
	for i := 0; i < 10; i++ {
		knots[i] = Part{0, 0, false}
	}
	for _, line := range data {
		tmp := strings.Fields(line)
		direction := tmp[0]
		distance, _ := strconv.Atoi(tmp[1])
		head := &knots[0]
		for i := 0; i < distance; i++ {
			head.x += dirMap[direction][0]
			head.y += dirMap[direction][1]
			for j := 1; j < 10; j++ {
				knots[j] = moveTail(knots[j-1], knots[j])
				if j == 9 {
					coords := [2]int{knots[j].x, knots[j].y}
					visited[coords] = true
				}
			}
		}
	}
	return len(visited)
}

func partOne(data []string) int {
	visited := make(map[[2]int]bool)
	head := Part{0, 0, false}
	tail := Part{0, 0, false}
	for _, line := range data {
		tmp := strings.Fields(line)
		direction := tmp[0]
		distance, _ := strconv.Atoi(tmp[1])
		for i := 0; i < distance; i++ {
			head.x += dirMap[direction][0]
			head.y += dirMap[direction][1]
			tail = moveTail(head, tail)
			coords := [2]int{tail.x, tail.y}
			visited[coords] = true
		}
	}
	return len(visited)
}

func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data)
	answers[1] = partTwo(data)
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
