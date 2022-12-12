package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func breadthFirst(data []string, heights map[[2]int]int, start [2]int, end [2]int) int {
	parent := make(map[[2]int][2]int, 0)
	cost := make(map[[2]int]int, 0)
	queue := list.New()
	queue.PushBack(start)
	cost[start] = 0
	for queue.Len() > 0 {
		curr := queue.Front().Value.([2]int)
		queue.Remove(queue.Front())
		if curr == end {
			path := list.New()
			for _, ok := parent[curr]; ok; {
				path.PushBack(curr)
				curr = parent[curr]
				if curr == start {
					break
				}
			}
			return path.Len()
		}

		for _, dir := range dirs {
			xNew := curr[0] + dir[0]
			yNew := curr[1] + dir[1]
			next := [2]int{xNew, yNew}
			if exists := heights[next]; exists == 0 || heights[next] > heights[curr]+1 {
				continue
			}
			tmp := cost[curr] + 1
			a := cost[next]
			if a == 0 {
				a = math.MaxInt
			}
			if tmp < a {
				cost[next] = tmp
				parent[next] = curr
				queue.PushBack(next)
			}
		}
	}
	return math.MaxInt
}

func partOne(data []string) int {
	heights := make(map[[2]int]int, 0)
	start, end := [2]int{}, [2]int{}
	for y, line := range data {
		for x, value := range line {
			if value == 'S' {
				start = [2]int{x, y}
			} else if value == 'E' {
				end = [2]int{x, y}
			} else {
				heights[[2]int{x, y}] = int(value)
			}
		}
	}
	heights[start] = int('a') - 1
	heights[end] = int('z') + 1

	z := breadthFirst(data, heights, start, end)
	return z

}

func partTwo(data []string) int {
	heights := make(map[[2]int]int, 0)
	start, end := [][2]int{}, [2]int{}
	for y, line := range data {
		for x, value := range line {
			if value == 'S' {
				start = append(start, [2]int{x, y})
				heights[[2]int{x, y}] = int('a')
			} else if value == 'a' {
				start = append(start, [2]int{x, y})
				heights[[2]int{x, y}] = int('a')
			} else if value == 'E' {
				end = [2]int{x, y}
			} else {
				heights[[2]int{x, y}] = int(value)
			}
		}
	}
	heights[end] = int('z') + 1

	lowest := math.MaxInt
	for _, point := range start {
		val := breadthFirst(data, heights, point, end)
		if lowest > val {
			lowest = val
		}
	}
	return lowest

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
