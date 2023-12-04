package dayfour

import (
	"strconv"
	"strings"
)

var copies []int

func calculateScore(input string, pos int) (int, int) {
	if input == "" {
		return 0, 0
	}
	// Add 1 for the original.
	copies[pos] += 1
	// Parsing.
	tmp := strings.Split(input, ": ")
	tmp = strings.Split(tmp[1], " | ")

	// Use maps as a makeshift set.
	winners := map[int]bool{}
	for _, i := range strings.Split(tmp[0], " ") {
		a, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			continue
		}
		winners[int(a)] = true
	}

	hand := map[int]bool{}
	for _, i := range strings.Split(tmp[1], " ") {
		a, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			continue
		}
		hand[int(a)] = true
	}

	score := 0
	p2Score := 0
	for k, _ := range hand {
		if _, ok := winners[k]; ok {
			if score == 0 {
				p2Score = 1
				score = 1
				continue
			}
			score *= 2
			p2Score += 1
		}
	}

	max := pos + p2Score + 1
	if max >= len(copies) {
		max = len(copies) - 1
	}

	for i := pos + 1; i < max; i++ {
		size := copies[pos]
		copies[i] += size
	}

	return score, 0
}

func solve(input []string) (int, int) {
	partOne := 0
	partTwo := 0
	// Initialize slice of size len(input) to zero.
	copies = make([]int, len(input))
	for i, row := range input {
		a, _ := calculateScore(row, i)
		partOne += a
	}
	for _, num := range copies {
		partTwo += num
	}
	return partOne, partTwo
}
