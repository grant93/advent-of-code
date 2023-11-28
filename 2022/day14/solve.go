package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Cave struct {
	grid map[[2]int]rune
	lowerBoundary int
	leftBoundary int
	rightBoundary int
	floor int
}

func (this *Cave) printGrid() {
	for y := 0; y <= this.lowerBoundary; y++ {
		str := ""
		for x:= this.leftBoundary; x <=this.rightBoundary; x++ {
			switch this.grid[[2]int{x,y}] {
			case 0: 
				str += " "
			default:
				str += string(this.grid[[2]int{x,y}])
			}
		}
		fmt.Println(str)
	}
}

func (this *Cave) parse(data []string) {
	this.leftBoundary, this.rightBoundary, this.lowerBoundary = 999, 0, 0
	this.grid = make(map[[2]int]rune)
	for _, row := range data {
		coords := strings.Split(row, " -> ")
		for i := 0; i < len(coords) - 1; i++ {

			coord1 := strings.Split(coords[i], ",")
			x1, _ := strconv.Atoi(coord1[0])
			y1, _ := strconv.Atoi(coord1[1])

			coord2 := strings.Split(coords[i+1], ",")
			x2, _ := strconv.Atoi(coord2[0])
			y2, _ := strconv.Atoi(coord2[1])

			if x2 < x1 {
				x1, x2 = x2, x1
			}
			if y2 < y1 {
				y1, y2 = y2, y1
			}

			if y1 > this.lowerBoundary {
				this.lowerBoundary = y1
			} else if y2 > this.lowerBoundary {
				this.lowerBoundary = y2
			}
			if x1 > this.rightBoundary {
				this.rightBoundary = x1
			} else if x1 < this.leftBoundary {
				this.leftBoundary =  x1
			}
			if x2 > this.rightBoundary {
				this.rightBoundary = x2
			} else if x2 < this.leftBoundary {
				this.leftBoundary = x2
			}

			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					this.grid[[2]int{x, y}] = '#'
				}
			}
		}
	}
}

func (this *Cave) simulate() bool{
	x, y := 500, 0
	for true {
		if this.floor == 0 && y >= this.lowerBoundary {
			return false
		}
		if this.grid[[2]int{500, 0}] != 0 {
			return false
		}
		if this.grid[[2]int{x,y+1}] == 0  && (this.floor == 0 || y+1 < this.floor) {
			y++
		} else if this.grid[[2]int{x-1,y+1}] == 0 && (this.floor == 0 || y+1 < this.floor) {
			x--
			y++
		} else if this.grid[[2]int{x+1,y+1}] == 0 && (this.floor == 0 || y+1 < this.floor) {
			x++
			y++
		} else {
			this.grid[[2]int{x,y}] = 'o'
			return true
		}
	}
	return false
}

func partOne(data []string, debug bool) int {
	counter := 0
	cave := &Cave{}
	cave.parse(data)
	for cave.simulate() {
		counter ++
	}
	if debug {
		cave.printGrid()
	}
	return counter
}

func partTwo(data []string, debug bool) int {
	counter := 0
	cave := &Cave{}
	cave.parse(data)
	cave.floor = cave.lowerBoundary + 2
	for cave.simulate() {
		counter++
	}
	if debug {
		cave.printGrid()
	}
	return counter
}

func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data, false)
	answers[1] = partTwo(data, false)

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
