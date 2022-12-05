package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func toInt(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 64 + 26
	} else {
		return int(char) - 96
	}
}

// TODO: create a generic module for sets and move this there.
func setCreate(str string) map[rune]bool {
	set := make(map[rune]bool)
	for _, char := range str {
		set[char] = true
	}
	return set
}

func setIntersection(s1 map[rune]bool, s2 map[rune]bool) map[rune]bool {
	s_intersection := map[rune]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] {
			s_intersection[k] = true
		}
	}
	return s_intersection
}

func findCommonThree(set1 map[rune]bool, set2 map[rune]bool, set3 map[rune]bool) rune {
	intersection := setIntersection(setIntersection(set1, set2), set3)
	for key, _ := range intersection {
		return key
	}
	return rune('a') // should never get here
}

func findCommonTwo(set1 map[rune]bool, set2 map[rune]bool) rune {
	intersection := setIntersection(set1, set2)
	for key, _ := range intersection {
		return key
	}
	return rune('a') // should never get here
}

func solve(data []string) []int {
	answers := make([]int, 2)
	for i, line := range data {
		size := len(line)
		answers[0] += toInt(findCommonTwo(setCreate(string(line[0:(size/2)])), setCreate(string(line[(size/2):size]))))
		if (i % 3) == 0 {
			answers[1] += toInt(findCommonThree(setCreate(data[i]), setCreate(data[i+1]), setCreate(data[i+2])))
		}
	}
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
