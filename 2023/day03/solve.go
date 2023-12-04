package daythree

import (
	"fmt"
	"strconv"
	"unicode"
)

type coords struct {
	x, y int
}

var (
	neighbours = []coords{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	cache = make(map[string]int)
)

func getNeighbours(input []string, x, y int) (int, int) {
	partOne := 0
	partTwo := map[int]int{}
	for _, diff := range neighbours {
		newX := x + diff.x
		newY := y + diff.y
		// Out of bounds.
		if newX < 0 || newX >= len(input[0]) || newY < 0 || newY >= len(input) {
			continue
		}
		if unicode.IsNumber(rune(input[newY][newX])) {
			leftBound := newX
			rightBound := newX
			// Check left.
			for xx := newX; xx >= 0; xx-- {
				if unicode.IsNumber(rune(input[newY][xx])) {
					leftBound = xx
					continue
				}
				break
			}
			// Check right.
			for xx := newX; xx < len(input[0]); xx++ {
				if unicode.IsNumber(rune(input[newY][xx])) {
					rightBound = xx
					continue
				}
				break
			}
			num, _ := strconv.ParseInt(input[newY][leftBound:rightBound+1], 10, 64)
			a := fmt.Sprintf("%d-%d->%d", leftBound, newY, num)
			if _, ok := cache[a]; !ok {
				partOne += int(num)
				cache[a] = 1
			}
			if input[y][x] == byte('*') {
				partTwo[int(num)] = 1
			}
		}
	}
	total := 0
	if len(partTwo) == 2 {
		total = 1
		for k, _ := range partTwo {
			total *= k
		}
	}
	return partOne, total
}

func checkNeighbours(input []string, x, y int) (int, int) {
	if unicode.IsNumber(rune(input[y][x])) || input[y][x] == byte('.') {
		return 0, 0
	}
	// Is a symbol, check the neighbours.
	return getNeighbours(input, x, y)
}

func solve(input []string) (int, int) {
	partOne := 0
	partTwo := 0
	for y, row := range input {
		for x := range row {
			a, b := checkNeighbours(input, x, y)
			partOne += a
			partTwo += b
		}
	}
	return partOne, partTwo
}
