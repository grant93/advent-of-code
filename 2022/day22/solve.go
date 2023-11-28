package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var monkeyMap []string
var monkeyMap2 [][]byte

var charMap = map[int]byte{0: '>', 1:'v', 2:'<', 3:'^'}

type Coords struct {
	x, y int
}

var directions []Coords = []Coords{Coords{1, 0}, Coords{0, 1}, Coords{-1, 0}, Coords{0, -1}};
var currentDirs = 0

func printMap() {
	for _, row := range monkeyMap2 {
		fmt.Println(string(row))
	}
}

func changeDirection(dir rune) Coords {
	if dir == 'R' {
		currentDirs++
	} else {
		currentDirs--
	}
	if currentDirs == 4 {
		currentDirs =  0
	}
	if currentDirs == -1 {
		currentDirs = 3
	}
	return directions[currentDirs]
}

func step(pos Coords) Coords {
	direction := directions[currentDirs]
	pos.x += direction.x
	pos.y += direction.y
	if pos.x >= len(monkeyMap[0]) {
		pos.x = 0
	} else if pos.x < 0 {
		pos.x = len(monkeyMap[0]) - 1
	}
	if pos.y >= len(monkeyMap) {
		pos.y = 0
	} else if pos.y < 0 {
		pos.y = len(monkeyMap) - 1
	}
	return pos
}

func findWrapAround(pos Coords) Coords {
	for true {
		pos = step(pos)
		if monkeyMap[pos.y][pos.x] != ' ' {
			return pos
		}
	}
	return Coords{0,0}
}

// there's a bug in here somewhere...
func findWrapAround2(pos Coords) Coords {
	if currentDirs == 0 { //right
		if pos.x == 149 && pos.y >=0 && pos.y <= 49 { // right 1
			currentDirs = 2
			pos.y = 149 - pos.y
			pos.x = 99
		} else if pos.x == 99 && pos.y >= 50 && pos.y <= 99 { // bottom 3
			pos.y, pos.x = pos.x + 50, 49 
			currentDirs = 3
		} else if pos.x == 99 && pos.y >= 100 && pos.y <= 149 { // back 4
			pos.y = 49 - (pos.y - 100) 
			pos.x = 149
			currentDirs = 2
		} else if pos.x == 49 && pos.y >= 150 && pos.y <= 199 { // top 6
			pos.x = pos.y - 100
			pos.y = 149
			currentDirs = 3
		}
	} else if currentDirs == 1 { // down
		if pos.y == 49 && pos.x >= 100 && pos.x <= 149 { // right 1
			pos.y = pos.x - 50
			pos.x = 99 
			currentDirs = 3
		} else if pos.y == 149 && pos.x >= 50 && pos.x <= 99 { // back 4
			pos.y = 150 + (pos.x - 50)
			pos.x = 49
			currentDirs = 3
		} else if pos.y == 199 && pos.x >= 0 && pos.x <= 49 { // top 6
			pos.x += 100
			pos.y = 0
		}
	} else if currentDirs == 2 { // left
		if pos.x == 50 && pos.y >= 0 && pos.y <= 49 { // front 2
			pos.x = 0
			pos.y += 100
			currentDirs = 0
		} else if pos.x == 50 && pos.y >= 50 && pos.y <= 99 { // bottom 3
			pos.x = pos.y - 50
			pos.y = 100
			currentDirs = 1
		} else if pos.x == 0 && pos.y >= 100 && pos.y <= 149 { // left 5
			pos.x = 50
			pos.y = 49 - (pos.y - 100)
			currentDirs = 0
		} else if pos.x == 0 && pos.y >= 150 && pos.y <= 199 { // top 6
			pos.x = pos.y - 100
			pos.y = 0
			currentDirs = 1
		}
	} else if currentDirs == 3 { // up
		if pos.y == 0 && pos.x >= 50 && pos.x <= 99 { // front 2
			pos.y = pos.x + 100
			pos.x = 0
			currentDirs = 0
		} else if pos.y == 0 && pos.x >=100 && pos.x <= 149 { // right 1
			pos.x = pos.x - 100 
			pos.y = 199
			currentDirs = 2
		} else if pos.y == 100 && pos.x >= 0 && pos.x <= 49 { // left 5
			pos.y = 50 + pos.x
			pos.x = 50
			currentDirs = 0
		}
	}
	return pos 
}

func move(pos Coords, scalar int, partTwo bool) Coords {
	for i:=0; i < scalar; i++ {
		tmp := step(pos)
		switch monkeyMap[tmp.y][tmp.x] {
		case '#':
			return pos
		case '.':
			pos = tmp
			fmt.Println(pos)
			monkeyMap2[tmp.y][tmp.x] = charMap[currentDirs]
		case ' ':
			var tmp2 Coords
			if partTwo {
				tmp2 = findWrapAround2(pos)
			} else {
				tmp2 = findWrapAround(pos)
			}
			if monkeyMap[tmp2.y][tmp2.x] == '#' {
				return pos
			}
			if monkeyMap[tmp2.y][tmp2.x] == ' '{
				fmt.Println("WTF")
			}
			pos = tmp2
			fmt.Println(pos)
			monkeyMap2[pos.y][pos.x] = charMap[currentDirs]
		}
	}
	return pos
}

func partOne(path string, partTwo bool) {
	instr := ""
	currentDirs = 0
	pos := Coords{strings.Index(monkeyMap[0], "."), 0}
	fmt.Println("starting pos: ", pos)
	for i, char := range path {
		dir := directions[currentDirs]
		if unicode.IsNumber(char) {
			instr += string(char)
			if (i == len(path)-1) || !unicode.IsNumber(rune(path[i+1])) {
				tmp, _ := strconv.Atoi(instr)
				fmt.Println("Scalar: ", tmp)
				pos = move(pos, tmp, partTwo)
				instr = ""
			}
		} else {
			dir = changeDirection(char)
			fmt.Println("Changing direction: ", dir)
		}
	}
	printMap()
	fmt.Println((1000 * (pos.y + 1)) + (4 * (pos.x + 1)) + currentDirs )
}


func solve(data []string) []int {
	monkeyMap = data[:len(data)-2]
	for _, row := range monkeyMap {
		monkeyMap2 = append(monkeyMap2, []byte(row))
	}
	path := data[len(data)-1]
	answers := make([]int, 2)
	//partOne(path, false)
	monkeyMap2 = [][]byte{}
	for _, row := range monkeyMap {
		monkeyMap2 = append(monkeyMap2, []byte(row))
	}
	partOne(path, true)

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
