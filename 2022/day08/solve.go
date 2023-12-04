package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* there's likely more efficient solutions than this hack, but it will do */

var visibilityMap [][]rune

func printMap() {
	fmt.Println(strings.Repeat("-", len(visibilityMap[0])))
	for _, a := range visibilityMap {
		fmt.Println(string(a))
	}
	fmt.Println(strings.Repeat("-", len(visibilityMap[0])))
}

const (
	Left int = iota
	Right
	Top
	Bottom
)

func partOne(data []string, direction int) int {
	visible := 0
	/*						 (x,y)
	calculate from left: 0,0 -> len, len 	(y=0; y++ outer, x=0; x++ inner
	calculate from right len, 0 -> 0, len 	(y=0; y++ outer, x=len-1; x-- inner)
	calculate from top   0, 0, -> len, len 	(x=0; x++ outer, y=0; y++ inner
	calculate from bottom 0, len -> len, 0 	(x=0; x++ outer, y=len-1; y-- inner)

	*/
	counterOneMax := 0
	counterTwoMax := 0

	if direction == Top || direction == Bottom {
		counterOneMax = len(data[0])
		counterTwoMax = len(data)
	} else {
		counterTwoMax = len(data[0])
		counterOneMax = len(data)
	}

	for counter1 := 0; counter1 < counterOneMax; counter1++ {
		prevHeight := -1
		y := 0
		x := 0

		if direction == Top || direction == Bottom {
			x = counter1
		} else {
			y = counter1
		}

		for counter2 := 0; counter2 < counterTwoMax; counter2++ {
			if direction == Left {
				x = counter2
			} else if direction == Right {
				x = counterTwoMax - counter2 - 1
			} else if direction == Top {
				y = counter2
			} else {
				y = counterTwoMax - counter2 - 1
			}

			height := int(data[y][x] - '0')

			if height > prevHeight {
				if visibilityMap[y][x] == ' ' {
					visibilityMap[y][x] = rune(data[y][x])
					visible++
				}
				prevHeight = height
			}

		}
	}
	//printMap()
	return visible
}

func calculateScore(data []string, value int, x int, y int) int {
	score := 1

	count := 0
	/* check right */
	for _x := x + 1; _x < len(data[0]); _x++ {
		count++
		if int(data[y][_x]-'0') >= value {
			break
		}
	}

	if count != 0 {
		score *= count
		count = 0
	}

	/* check left */
	for _x := x - 1; _x >= 0; _x-- {
		count++
		if int(data[y][_x]-'0') >= value {
			break
		}
	}

	if count != 0 {
		score *= count
		count = 0
	}

	/* check below */
	for _y := y + 1; _y < len(data); _y++ {
		count++
		if int(data[_y][x]-'0') >= value {
			break
		}
	}

	if count != 0 {
		score *= count
		count = 0
	}

	/* check above */
	for _y := y - 1; _y >= 0; _y-- {
		count++
		if int(data[_y][x]-'0') >= value {
			break
		}
	}

	if count != 0 {
		score *= count
	}

	return score
}

/* run through every element and check surrounds */
func partTwo(data []string) int {
	highScore := 0
	for y, valY := range data {
		for x, valX := range valY {
			score := calculateScore(data, int(valX-'0'), x, y)
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func solve(data []string) []int {
	answers := make([]int, 2)
	visibilityMap = make([][]rune, len(data))
	transformations := []int{Left, Right, Top, Bottom}
	for i := range visibilityMap {
		visibilityMap[i] = []rune(strings.Repeat(" ", len(data[0])))
	}
	visible := 0

	/* "rotate" and check from each direction */
	for _, dydx := range transformations {
		visible += partOne(data, dydx)
	}

	answers[0] = visible
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
