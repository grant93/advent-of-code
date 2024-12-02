package dayseven

import (
	"fmt"
	"strings"
	"testing"

	_ "embed"
)

var (
	//go:embed input.txt
	secondInput string
)

func TestSolve(t *testing.T) {
	firstInput := `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`

	var tests = []struct {
		input      string
		wantFirst  int
		wantSecond int
	}{
		{firstInput, 6440, 5905},
		{secondInput, 251136060, 249400220},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			p1, p2 := solve(strings.Split(tt.input, "\n"))
			if p1 != tt.wantFirst {
				t.Errorf("part one: got %d, want %d", p1, tt.wantFirst)
			}
			if p2 != tt.wantSecond {
				t.Errorf("part two: got %d, want %d", p2, tt.wantSecond)
			}
		})
	}
}
