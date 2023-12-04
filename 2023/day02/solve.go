package dayone

import (
	"strconv"
	"strings"
)

const (
	maxGreen int64 = 13
	maxBlue  int64 = 14
	maxRed   int64 = 12
)

func checkGame(game string) (int, int) {
	badGame := false
	colourMap := map[string]int64{
		"green": 0,
		"red":   0,
		"blue":  0,
	}
	// In case of empty game.
	if game == "" {
		return 0, 0
	}

	// Clean up and parse input.
	game = strings.ReplaceAll(game, ",", "")
	parts := strings.Split(game, ": ")
	id, _ := strconv.ParseInt(strings.Split(parts[0], " ")[1], 10, 32)

	for _, set := range strings.Split(parts[1], "; ") {
		vals := strings.Split(set, " ")
		for i := 0; i < len(vals); i += 2 {
			num, _ := strconv.ParseInt(vals[i], 10, 32)
			// Part One.
			if vals[i+1] == "green" && num > maxGreen || vals[i+1] == "blue" && num > maxBlue || vals[i+1] == "red" && num > maxRed {
				badGame = true
			}
			// Part Two.
			if colourMap[vals[i+1]] < num {
				colourMap[vals[i+1]] = num
			}
		}
	}

	partOne := int(id)
	if badGame {
		partOne = 0
	}

	partTwo := 1
	for _, v := range colourMap {
		partTwo *= int(v)
	}
	return partOne, partTwo
}

func solve(input []string) (int, int) {
	partOne := 0
	partTwo := 0
	for _, line := range input {
		a, b := checkGame(line)
		partOne += a
		partTwo += b
	}
	return partOne, partTwo
}
