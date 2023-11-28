package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	North int = iota
	South
	West
	East
)

var defaultQueue [4]int = [4]int{North, South, East, West}
var northCheck [3][2]int = [3][2]int{{0, -1}, {1, -1}, {-1, -1}}
var southCheck [3][2]int = [3][2]int{{0, 1}, {1, 1}, {-1, 1}}
var westCheck [3][2]int = [3][2]int{{-1, 0}, {-1, -1}, {-1, 1}}
var eastCheck [3][2]int = [3][2]int{{1, 0}, {1, -1}, {1, 1}}

var allDirs [8][2]int = [8][2]int{{-1, 0}, {1, 0}, {-1, 1}, {-1, -1},  {1, 1}, {1, -1}, {0, 1}, {0, -1}}

type Elf struct {
	prevPos [2]int
	currPos [2]int
	dirQueue [4]int
}

func parse(data []string) map[[2]int]Elf {
	retval := map[[2]int]Elf{}

	for y, row := range data {
		for x, char := range row {
			if char == '#' {
				retval[[2]int{x, y}] = Elf{[2]int{x, y}, [2]int{x,y}, defaultQueue}
			}
		}
	}
	return retval
}

func checkSurroundings(state map[[2]int]Elf, newState map[[2]int]Elf,  elf Elf) bool {
	for _, dir := range allDirs {
		xNew, yNew := elf.currPos[0] + dir[0], elf.currPos[1] + dir[1] 
		if _, ok := state[[2]int{xNew, yNew}]; ok {
			break
		}
		return false
	}

	for _, val := range elf.dirQueue {
		switch val {
		case North:
			for _, dir := range northCheck {
				xNew, yNew := elf.currPos[0] + dir[0], elf.currPos[1] + dir[1] 
				if _, ok := state[[2]int{xNew, yNew}]; ok {
					break
				}
				return true 
			}
		case South:
			for _, dir := range southCheck {
				xNew, yNew := elf.currPos[0] + dir[0], elf.currPos[1] + dir[1] 
				if _, ok := state[[2]int{xNew, yNew}]; ok {
					break
				}
				return true
			}
		case West:
			for _, dir := range westCheck {
				xNew, yNew := elf.currPos[0] + dir[0], elf.currPos[1] + dir[1] 
				if _, ok := state[[2]int{xNew, yNew}]; ok {
					break
				}
				return true
			}
		case East:
			for _, dir := range eastCheck {
				xNew, yNew := elf.currPos[0] + dir[0], elf.currPos[1] + dir[1] 
				if _, ok := state[[2]int{xNew, yNew}]; ok {
					break
				}
				return true 
			}
		}
		if val, ok := newState[[2]int{xNew, yNew}]; ok {
			// already exists, lets move it back
			delete(newState, [2]int{xNew, yNew})
			newState[val.prevPos] = val
			newState[elf.currPos] = elf
		} else {
			elf.prevPos = currPos
			elf.currPos = [2]int{xNew, yNew}
		}
	}
	return false
}

func partOne(state map[[2]int]Elf) int {
	var newState map[[2]int]Elf
	for i := 0; i < 10; i++ {
		newState := map[[2]int]Elf{}
		for key, value := range state {
			if !checkSurroundings(state, newState, value) {
				return -1
			}
		}
		state = newState
	}
	fmt.Println(state)
	return 0
	
}

func solve(data []string) []int {
	answers := make([]int, 2)
	state := parse(data)
	fmt.Println(state)
	answers[0] = partOne(state)

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
