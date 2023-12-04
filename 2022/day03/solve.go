package daythree

import (
	s "github.com/grant93/advent-of-code/lib/set"
	"unicode"
)

func toInt(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 64 + 26
	} else {
		return int(char) - 96
	}
}

func findCommonThree(set1 s.Set[rune], set2 s.Set[rune], set3 s.Set[rune]) rune {
	intersection := set1.Intersection(set2).Intersection(set3)
	for key, _ := range intersection {
		return key
	}
	panic("uhoh")
	return rune('a')
}

func findCommonTwo(set1 s.Set[rune], set2 s.Set[rune]) rune {
	intersection := set1.Intersection(set2)
	for key, _ := range intersection {
		return key
	}
	panic("uhoh")
	return rune('a')
}

func solve(data []string) (int, int) {
	partOne := 0
	partTwo := 0
	for i, line := range data {
		size := len(line)
		partOne += toInt(findCommonTwo(s.CreateWithValues([]rune(line[0:(size/2)])), s.CreateWithValues([]rune(line[(size/2):size]))))
		if (i % 3) == 0 {
			partTwo += toInt(findCommonThree(s.CreateWithValues([]rune(data[i])), s.CreateWithValues([]rune(data[i+1])), s.CreateWithValues([]rune(data[i+2]))))
		}
	}
	return partOne, partTwo
}
