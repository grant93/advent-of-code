package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var dirMap = map[string][2]int{
	"R" : [2]int{1, 0},
	"L" : [2]int{-1, 0},
	"U" : [2]int{0, 1},
	"D" : [2]int{0, -1},
}

func setIntersection(s1 map[[2]int]int, s2 map[[2]int]int) map[[2]int]int {
	s_intersection := map[[2]int]int{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] > 0 {
			val := s1[k] + s2[k]
			if s_intersection[k] == 0 || s_intersection[k] > val {
				s_intersection[k] = val 
			}
		}
	}
	return s_intersection
}

func manhattanDistance(x int, y int) int {
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}

	return x + y
}

func solve(data []string) []int {
	answers := make([]int, 2)
	visits := [2]map[[2]int]int{}
	for num, line := range data {
		visits[num] = make(map[[2]int]int, 0)
		current := [2]int{0, 0}
		moves := strings.Split(line, ",")
		steps := 0
		for _, move := range moves {
			direction := string(move[0])
			scalar, _ := strconv.Atoi(move[1:])
			for i:=0; i < scalar; i++ {
				steps ++
				current[0] += dirMap[direction][0]
				current[1] += dirMap[direction][1]
				visits[num][current] = steps 
			}
		}
	}

	intersection := setIntersection(visits[0], visits[1])
	shortest := 10000
	minSteps := 1000000
	for key, value := range intersection{
		fmt.Println(key, value)
		distance := manhattanDistance(key[0], key[1])
		if distance < shortest {
			shortest = distance
		}
		if value < minSteps {
			minSteps = value
		}
	}
	answers[0] = shortest
	answers[1] = minSteps
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
