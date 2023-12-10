package dayten

import (
	"strings"
)

type Direction int

const (
	MOVE_LEFT Direction = iota
	MOVE_UP
	MOVE_RIGHT
	MOVE_DOWN
)

type coordinates struct {
	x, y int
}

var moveMap = map[Direction]coordinates{ 
	MOVE_UP: coordinates{0, -1},
	MOVE_LEFT: coordinates{-1, 0},
	MOVE_RIGHT: coordinates{1, 0},
	MOVE_DOWN: coordinates{0, 1},
}

var valMap = map[byte][]Direction{
	'|': []Direction{MOVE_UP, MOVE_DOWN},
	'-': []Direction{MOVE_LEFT, MOVE_RIGHT},
	'L': []Direction{MOVE_RIGHT, MOVE_UP},
	'J': []Direction{MOVE_LEFT, MOVE_UP},
	'7': []Direction{MOVE_LEFT, MOVE_DOWN},
	'F': []Direction{MOVE_RIGHT, MOVE_DOWN},
}

var oppositeMap = map[Direction]Direction{
	MOVE_UP: MOVE_DOWN,
	MOVE_DOWN: MOVE_UP,
	MOVE_RIGHT: MOVE_LEFT,
	MOVE_LEFT: MOVE_RIGHT,
}

var loopCoords = map[coordinates]bool{}

func contains(a Direction, b []Direction) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func determineStart(input []string, start coordinates) ([]coordinates, byte) {
	var retval []coordinates
	var pipeDirs []Direction
	maxX, maxY := len(input[0]), len(input)
	// Find the neighbour that works with the starting position.
	for k,v := range moveMap {
		newX := start.x + v.x
		if newX < 0 || newX >= maxX {
			continue
		}
		newY := start.y + v.y
		if newY < 0 || newY >= maxY {
			continue
		}
		val := input[newY][newX]
		oppo := oppositeMap[k]

		for _, v := range valMap[val] {
			if v == oppo {
				retval = append(retval, coordinates{newX, newY})
				pipeDirs = append(pipeDirs, k)
				break
			}
		}
	}

	// Determine the starting pipe value.
	for k, v := range valMap {
		bad := false
		for _, p := range pipeDirs {
			if contains(p, v) {
				continue
			}
			bad = true
		}
		if !bad {
			return retval, k
		}
	}
	panic("this should never happen!")
}

// Recursion is fun.
func move(input []string, cur, prev []coordinates) int {
	var newPos []coordinates 
	loopCoords[cur[0]] = true
	loopCoords[cur[1]] = true

	if cur[0].x == cur[1].x && cur[0].y == cur[1].y {
		return 1
	}

	for i:=0; i<2; i++ {
		val := input[cur[i].y][cur[i].x]
		diff := coordinates{prev[i].x - cur[i].x, prev[i].y - cur[i].y}
		for _, n := range valMap[val] {
			m := moveMap[n]
			if diff.x == m.x && diff.y == m.y {
				continue
			}
			newPos = append(newPos, coordinates{cur[i].x + m.x, cur[i].y + m.y})
		}
	}
	return move(input, newPos, cur) + 1
}

func run(input []string) int {
	for y, line := range input {
		if x := strings.Index(line, "S"); x != -1 {
			start := coordinates{x, y}
			firstStep, val := determineStart(input, start) 
			input[y] = strings.Replace(input[y], "S", string(val), -1)
			loopCoords[coordinates{x, y}] = true
			return move(input, firstStep, []coordinates{start, start})
		}
	}
	return 0
}

func solve(input []string) (int, int) {
	partTwo := int64(0)
	partOne := run(input)
	return int(partOne), int(partTwo)
}
