package dayseven

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var valMap = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12,
	'J': 11, 'T': 10, '9': 9,
	'8': 8, '7': 7, '6': 6,
	'5': 5, '4': 4, '3': 3,
	'2': 2, '1': 1, 'X': 0,
}

type hand struct {
	cards          string
	bet            int
	occurrences    map[rune]int
	maxOccurrences int
	twoPair        bool
}

func parse(input []string) []hand {
	m := []hand{}
	for _, line := range input {
		t := strings.Fields(line)
		n, _ := strconv.ParseInt(t[1], 10, 32)
		o := map[rune]int{}
		max := 1
		for _, c := range t[0] {
			if v, ok := o[c]; ok {
				o[c] = v + 1
				if m := o[c]; m > max {
					max = m
				}
			} else {
				o[c] = 1
			}
		}
		c := 0
		if max == 2 || max == 3 {
			for _, v := range o {
				if v == 2 {
					c += 1
				}
			}
		}
		m = append(m, hand{t[0], int(n), o, max, (c == 2 || (max == 3 && c == 1))})
	}
	return m
}

func compareHands(a, b string) int {
	for i := range a {
		first := valMap[rune(a[i])]
		second := valMap[rune(b[i])]
		if first > second {
			return 1
		} else if second > first {
			return -1
		}
	}
	return 0
}

func calculateScore(hands []hand) int {
	total := 0
	for i, h := range hands {
		total = total + (i+1)*h.bet
	}
	return total
}

func sortCards(a, b hand) int {
	switch {
	case a.maxOccurrences > b.maxOccurrences:
		return 1
	case a.maxOccurrences < b.maxOccurrences:
		return -1
	case a.maxOccurrences == b.maxOccurrences:
		if a.maxOccurrences < 4 && a.twoPair && !b.twoPair {
			return 1
		} else if a.maxOccurrences < 4 && !a.twoPair && b.twoPair {
			return -1
		} else {
			return compareHands(a.cards, b.cards)
		}
	}
	return 0
}

func newMax(a hand) int {
	newMax := a.maxOccurrences
	if a.maxOccurrences == a.occurrences['J'] {
		for k, v := range a.occurrences {
			if k != 'J' && v+a.maxOccurrences >= newMax {
				newMax = v + a.maxOccurrences
			}
		}
	} else if jokerCount := a.occurrences['J']; jokerCount > 0 {
		newMax = a.maxOccurrences + jokerCount
	}
	return newMax
}

func updateHands(hands []hand) []hand {
	newHands := []hand{}
	for _, h := range hands {
		h.maxOccurrences = newMax(h)
		h.cards = strings.ReplaceAll(h.cards, "J", "X")
		newHands = append(newHands, h)
	}
	return newHands
}

func solve(input []string) (int, int) {
	h := parse(input)
	slices.SortFunc(h, sortCards)
	partOne := calculateScore(h)
	fmt.Println(partOne)

	h = updateHands(h)
	slices.SortFunc(h, sortCards)
	partTwo := calculateScore(h)

	return partOne, partTwo

}
