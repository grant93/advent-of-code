package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

var neighbours = [][3]int{
	{-1, 0, 0},
	{1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func compare(max int, min int, current int) (int, int) {
	if (current > max) {
		max = current
	}
	if (current < min) {
		min = current
	}
	return max, min
}

func parse(data []string) (map[[3]int]int, [3]int, [3]int){
	cache := make(map[[3]int]int)
	max := [3]int{0, 0, 0}
	min := [3]int{math.MaxInt, math.MaxInt, math.MaxInt}
	for _, line := range data {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		max[0], min[0] = compare(max[0], min[0], x)
		y, _ := strconv.Atoi(coords[1])
		max[1], min[1] = compare(max[1], min[1], y)
		z, _ := strconv.Atoi(coords[2])
		max[2], min[2] = compare(max[2], min[2], z)
		cache[[3]int{int(x),int(y),int(z)}] = 6
	}
	return cache, max, min
}

func partOne(cache map[[3]int]int) int {
	total := 0
	for key, _ := range cache {
		for _, diff := range neighbours {
			x, y, z := key[0] + diff[0], key[1] + diff[1], key[2] + diff[2]
			_, ok := cache[[3]int{x,y,z}]
			if !ok {
				total++
			}
		}
	}
	return total 
}

func partTwo(cache map[[3]int]int, max [3]int, min [3]int) int {
	// FIXME: is using the empty struct hack a better idea here?
	visited := map[[3]int]int{min: 0}
	q := [][3]int{min}

	total := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, diff := range neighbours {
			x, y, z := curr[0] + diff[0], curr[1] + diff[1], curr[2] + diff[2]
			if curr[0] <= max[0] || curr[0] >= min[0] || curr[1] <= max[1] || curr[1] >= min[1] || curr[2] <= max[2] || curr[2] >= min[2] {
				total++
			} else if _, seen := visited[[3]int{x,y,z}]; !seen {
				visited[[3]int{x,y,z}] = 0
				q = append(q, [3]int{x,y,z})
			}
		}
	}
	return total

}


func solve(data []string) []int {
	cache, max, min := parse(data)
	answers := make([]int, 2)
	answers[0] = partOne(cache)
	answers[1] = partTwo(cache, max, min)
	fmt.Println(max, min)

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
