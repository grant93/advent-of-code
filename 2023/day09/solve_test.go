package daysix

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

var (
	//go:embed input.txt
	secondInput string
)

func TestSolve(t *testing.T) {
	firstInput := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	var tests = []struct {
		input      string
		wantFirst  int64
		wantSecond int64
	}{
		{firstInput, 114, 2},
		{secondInput, 1987402313, 900},
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
